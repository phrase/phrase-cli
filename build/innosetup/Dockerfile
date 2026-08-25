FROM ubuntu:22.04

RUN dpkg --add-architecture i386
RUN apt-get update && apt-get install -y --no-install-recommends curl ca-certificates wine-stable wine32 xvfb

ENV WINEDEBUG fixme-all
ENV WINEARCH win32
ENV FORCE_COLOR 1

# # Install Inno Setup binaries
# jrsoftware.org's old download.php shortcut no longer serves the binary
# directly (it now returns an HTML landing page); pin to a specific release
# from GitHub instead. Stick to the last unified (non x86/x64-split) build,
# since WINEARCH is win32.
RUN curl -SL "https://github.com/jrsoftware/issrc/releases/download/is-6_7_3/innosetup-6.7.3.exe" -o is.exe
RUN xvfb-run wine is.exe /VERYSILENT /SUPPRESSMSGBOXES /NORESTART
