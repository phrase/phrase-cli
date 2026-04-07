require "spec_helper"
require "fileutils"
require "tmpdir"

RSpec.describe "phrase pull" do
  let(:token) { "test-token-pull" }
  let(:project_id) { "test-project-123" }
  let(:locale_en_id) { "locale-en-id" }
  let(:locale_de_id) { "locale-de-id" }

  let(:en_locale) do
    {
      id: locale_en_id,
      name: "English",
      code: "en",
      default: true,
      main: true,
      rtl: false,
      plural_forms: ["one", "other"],
      created_at: "2024-01-01T00:00:00Z",
      updated_at: "2024-01-01T00:00:00Z"
    }
  end

  let(:de_locale) do
    {
      id: locale_de_id,
      name: "German",
      code: "de",
      default: false,
      main: false,
      rtl: false,
      plural_forms: ["one", "other"],
      created_at: "2024-01-01T00:00:00Z",
      updated_at: "2024-01-01T00:00:00Z"
    }
  end

  let(:en_content) do
    <<~YAML
      en:
        hello: "Hello"
        world: "World"
    YAML
  end

  let(:de_content) do
    <<~YAML
      de:
        hello: "Hallo"
        world: "Welt"
    YAML
  end

  around do |example|
    Dir.mktmpdir do |tmpdir|
      @tmpdir = tmpdir
      example.run
    ensure
      @tmpdir = nil
    end
  end

  before do
    # Clear previous mock requests
    mock_clear_requests!

    # Mock the locales list endpoint
    mock_set!("GET", "/projects/#{project_id}/locales",
      status: 200,
      body: [en_locale, de_locale]
    )

    # Mock the download endpoints for each locale
    mock_set!("GET", "/projects/#{project_id}/locales/#{locale_en_id}/download",
      status: 200,
      body: en_content,
      headers: { "content-type" => "application/x-yaml" }
    )

    mock_set!("GET", "/projects/#{project_id}/locales/#{locale_de_id}/download",
      status: 200,
      body: de_content,
      headers: { "content-type" => "application/x-yaml" }
    )
  end

  describe "basic pull operation" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          pull:
            targets:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    it "downloads locale files successfully" do
      r = run_cli("pull", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify files were created
      en_file_path = File.join(@tmpdir, "locales", "en.yml")
      de_file_path = File.join(@tmpdir, "locales", "de.yml")

      expect(File.exist?(en_file_path)).to be true
      expect(File.exist?(de_file_path)).to be true

      # Verify file contents
      expect(File.read(en_file_path)).to eq(en_content)
      expect(File.read(de_file_path)).to eq(de_content)
    end

    it "makes authenticated requests to the API" do
      run_cli("pull", config: config)

      requests_made = mock_requests

      # Should have made: 1 request to list locales + 2 download requests
      expect(requests_made.length).to be >= 3

      # Check locales list request
      locales_request = requests_made.find { |r| r["path"] == "/projects/#{project_id}/locales" }
      expect(locales_request).not_to be_nil
      expect(locales_request["method"]).to eq("GET")
      expect(locales_request["headers"]["HTTP_AUTHORIZATION"]).to eq("token #{token}")

      # Check download requests
      en_download = requests_made.find { |r| r["path"] == "/projects/#{project_id}/locales/#{locale_en_id}/download" }
      de_download = requests_made.find { |r| r["path"] == "/projects/#{project_id}/locales/#{locale_de_id}/download" }

      expect(en_download).not_to be_nil
      expect(de_download).not_to be_nil
      expect(en_download["headers"]["HTTP_AUTHORIZATION"]).to eq("token #{token}")
      expect(de_download["headers"]["HTTP_AUTHORIZATION"]).to eq("token #{token}")
    end
  end

  describe "pull with locale_name placeholder" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          pull:
            targets:
            - file: "#{@tmpdir}/i18n/<locale_name>.yml"
              params:
                file_format: yml
      YAML
    end

    it "uses locale name in file path" do
      r = run_cli("pull", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify files were created with locale names
      en_file_path = File.join(@tmpdir, "i18n", "English.yml")
      de_file_path = File.join(@tmpdir, "i18n", "German.yml")

      expect(File.exist?(en_file_path)).to be true
      expect(File.exist?(de_file_path)).to be true

      # Verify file contents
      expect(File.read(en_file_path)).to eq(en_content)
      expect(File.read(de_file_path)).to eq(de_content)
    end
  end

  describe "pull with specific locale filter" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          pull:
            targets:
            - file: "#{@tmpdir}/locales/en.yml"
              params:
                file_format: yml
                locale_id: "#{locale_en_id}"
      YAML
    end

    it "downloads only the specified locale" do
      r = run_cli("pull", config: config)

      expect(r[:exit_code]).to eq(0)

      # Only English file should be created
      en_file_path = File.join(@tmpdir, "locales", "en.yml")

      expect(File.exist?(en_file_path)).to be true
      expect(File.read(en_file_path)).to eq(en_content)

      # Verify only the English locale was requested
      requests_made = mock_requests
      download_requests = requests_made.select { |r| r["path"].include?("/download") }

      # Should only have one download request for the en locale
      expect(download_requests.length).to eq(1)
      expect(download_requests.first["path"]).to include(locale_en_id)
    end
  end

  describe "pull with tag filter" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          pull:
            targets:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
                tags: "frontend"
      YAML
    end

    it "successfully pulls with tag filter configured" do
      r = run_cli("pull", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify files were created
      en_file_path = File.join(@tmpdir, "locales", "en.yml")
      de_file_path = File.join(@tmpdir, "locales", "de.yml")

      expect(File.exist?(en_file_path)).to be true
      expect(File.exist?(de_file_path)).to be true

      # Verify download requests were made
      requests_made = mock_requests
      download_requests = requests_made.select { |r| r["path"].include?("/download") }

      expect(download_requests).not_to be_empty
    end
  end

  describe "pull with locale_mapping" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          locale_mapping:
            English: "eng"
            German: "ger"
          pull:
            targets:
            - file: "#{@tmpdir}/locales/<locale_name>.yml"
              params:
                file_format: yml
      YAML
    end

    it "applies locale mapping to file names" do
      r = run_cli("pull", config: config)
      expect(r[:exit_code]).to eq(0)

      # Verify files were created with mapped locale names
      en_file_path = File.join(@tmpdir, "locales", "eng.yml")
      de_file_path = File.join(@tmpdir, "locales", "ger.yml")
      expect(File.exist?(en_file_path)).to be true
      expect(File.exist?(de_file_path)).to be true

      # Verify file contents
      expect(File.read(en_file_path)).to eq(en_content)
      expect(File.read(de_file_path)).to eq(de_content)
    end
  end

  describe "error handling" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          pull:
            targets:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    it "handles authentication errors" do
      # Override the mock to return 401
      mock_set!("GET", "/projects/#{project_id}/locales",
        status: 401,
        body: { message: "Unauthorized" }
      )

      r = run_cli("pull", config: config)

      expect(r[:exit_code]).not_to eq(0)
      expect(r[:stderr]).to include("401")
    end

    it "handles missing project errors" do
      # Override the mock to return 404
      mock_set!("GET", "/projects/#{project_id}/locales",
        status: 404,
        body: { message: "Project not found" }
      )

      r = run_cli("pull", config: config)

      expect(r[:exit_code]).not_to eq(0)
    end
  end

  describe "pull without configuration" do
    it "returns error when no targets are specified" do
      config = <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
      YAML

      r = run_cli("pull", config: config)

      expect(r[:exit_code]).not_to eq(0)
      expect(r[:stderr]).to match(/no targets|pull.*not.*specified/i)
    end
  end

  describe "pull with multiple targets" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          pull:
            targets:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
            - file: "#{@tmpdir}/translations/<locale_name>.yaml"
              params:
                file_format: yml
      YAML
    end

    it "downloads files to multiple target locations" do
      r = run_cli("pull", config: config)

      expect(r[:exit_code]).to eq(0)

      # Check first target (locale_code placeholder)
      expect(File.exist?(File.join(@tmpdir, "locales", "en.yml"))).to be true
      expect(File.exist?(File.join(@tmpdir, "locales", "de.yml"))).to be true

      # Check second target (locale_name placeholder)
      expect(File.exist?(File.join(@tmpdir, "translations", "English.yaml"))).to be true
      expect(File.exist?(File.join(@tmpdir, "translations", "German.yaml"))).to be true
    end
  end

  describe "directory creation" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          pull:
            targets:
            - file: "#{@tmpdir}/deeply/nested/path/locales/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    it "creates nested directories automatically" do
      r = run_cli("pull", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify nested directory was created
      nested_path = File.join(@tmpdir, "deeply", "nested", "path", "locales")
      expect(Dir.exist?(nested_path)).to be true

      # Verify files exist in nested directory
      expect(File.exist?(File.join(nested_path, "en.yml"))).to be true
      expect(File.exist?(File.join(nested_path, "de.yml"))).to be true
    end
  end
end
