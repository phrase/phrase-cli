#!/bin/bash
set -euo pipefail
umask 077

CERTIFICATE_BASE64="${SIGNING_CERTIFICATE}"
P12_PASSWORD="${CERTIFICATE_PASSWORD}"
SIGNING_IDENTITY="${SIGNING_IDENTITY}"
KEYCHAIN_PASSWORD="${KEYCHAIN_PASSWORD}"
DIST_DIR="${DIST_DIR:-dist}"

CERTIFICATE_PATH="./build_certificate.p12"
KEYCHAIN_PATH="./my-signing.keychain-db"

# Basic env validation to fail fast
require_env() {
  local name="$1" value="$2"
  if [[ -z "$value" ]]; then
    echo "❌ Missing required environment variable: $name" >&2
    exit 1
  fi
}

require_env "SIGNING_CERTIFICATE" "${CERTIFICATE_BASE64}"
require_env "CERTIFICATE_PASSWORD" "${P12_PASSWORD}"
require_env "SIGNING_IDENTITY" "${SIGNING_IDENTITY}"
require_env "KEYCHAIN_PASSWORD" "${KEYCHAIN_PASSWORD}"
require_env "NOTARIZATION_APPLE_ID" "${NOTARIZATION_APPLE_ID:-}"
require_env "NOTARIZATION_APP_PASSWORD" "${NOTARIZATION_APP_PASSWORD:-}"
require_env "NOTARIZATION_TEAM_ID" "${NOTARIZATION_TEAM_ID:-}"


cleanup() {
  echo "🧹 Cleaning up keychain and certificate..."
  # Attempt to delete the temporary keychain
  security delete-keychain "$KEYCHAIN_PATH" || true
  # Remove certificate file
  rm -f "$CERTIFICATE_PATH" || true
}
trap cleanup EXIT

echo "🔐 Setting up certificate and keychain..."

# Decode the certificate (macOS-only)
echo "$CERTIFICATE_BASE64" | /usr/bin/base64 -D > "$CERTIFICATE_PATH"
# Restrict permissions on sensitive certificate material
chmod 600 "$CERTIFICATE_PATH"

# Create temporary keychain
security create-keychain -p "$KEYCHAIN_PASSWORD" "$KEYCHAIN_PATH"
security set-keychain-settings -lut 21600 "$KEYCHAIN_PATH"
security unlock-keychain -p "$KEYCHAIN_PASSWORD" "$KEYCHAIN_PATH"

# Import certificate into keychain
security import "$CERTIFICATE_PATH" -P "$P12_PASSWORD" -A -t cert -f pkcs12 -k "$KEYCHAIN_PATH"
security set-key-partition-list -S apple-tool:,apple: -s -k "$KEYCHAIN_PASSWORD" "$KEYCHAIN_PATH"

# Show available signing identities for visibility
echo "🔎 Available signing identities (codesigning):"
security find-identity -v -p codesigning "$KEYCHAIN_PATH" || true

# Find and sign all macOS binaries dynamically
echo "🔎 Searching for macOS binaries in $DIST_DIR..."

find "$DIST_DIR" -type f \( -name "phrase_macosx_*" ! -name "*.tar.gz" \) -print0 | while IFS= read -r -d '' binary; do
  echo "🔏 Signing $binary..."
  codesign --timestamp --options runtime --keychain "$KEYCHAIN_PATH" --sign "$SIGNING_IDENTITY" "$binary"
  codesign --verify --verbose=2 --keychain "$KEYCHAIN_PATH" "$binary"
done

echo "✅ All macOS binaries signed successfully."

# --- Zip artifacts for notarization ---
echo "📦 Zipping macOS binaries for notarization..."
shopt -s nullglob
for bin in "$DIST_DIR"/phrase_macosx_*; do
  [[ "$bin" == *.tar.gz ]] && continue
  relbin="${bin#${DIST_DIR}/}"
  echo "Creating $DIST_DIR/${relbin}.zip"
  (
    cd "$DIST_DIR" && /usr/bin/zip -o "${relbin}.zip" "${relbin}"
  )
done

# --- Notarization via Apple notarytool (Apple ID + app-specific password) ---
echo "📝 Notarizing zipped binaries with Apple Notary (Apple ID)..."
for zip in "$DIST_DIR"/phrase_macosx_*.zip; do
  [[ -e "$zip" ]] || continue
  echo "Submitting $zip to Apple Notary..."
  xcrun notarytool submit "$zip" \
    --apple-id "$NOTARIZATION_APPLE_ID" \
    --password "$NOTARIZATION_APP_PASSWORD" \
    --team-id "$NOTARIZATION_TEAM_ID" \
    --wait
  echo "ℹ️ Notarization complete for $zip."
done

echo "🎉 Signing and notarization finished."