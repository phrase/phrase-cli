require "spec_helper"
require "fileutils"
require "tmpdir"

RSpec.describe "phrase push" do
  let(:token) { "test-token-push" }
  let(:project_id) { "test-project-123" }
  let(:locale_en_id) { "locale-en-id" }
  let(:locale_de_id) { "locale-de-id" }
  let(:upload_en_id) { "upload-en-id" }
  let(:upload_de_id) { "upload-de-id" }

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

  let(:upload_response) do
    lambda do |upload_id, locale_id, filename|
      {
        id: upload_id,
        filename: filename,
        format: "yml",
        state: "success",
        summary: {
          locales_created: 0,
          translation_keys_created: 2,
          translation_keys_updated: 0,
          translation_keys_unmentioned: 0,
          translations_created: 2,
          translations_updated: 0
        },
        created_at: "2024-01-01T00:00:00Z",
        updated_at: "2024-01-01T00:00:00Z",
        url: "https://phrase.com/projects/#{project_id}/uploads/#{upload_id}"
      }
    end
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

    # Mock the formats list endpoint
    mock_set!("GET", "/formats",
      status: 200,
      body: [
        {
          api_name: "yml",
          name: "YAML",
          description: "YAML format",
          extension: ".yml",
          default_encoding: "UTF-8",
          importable: true,
          exportable: true,
          includes_locale_information: true
        },
        {
          api_name: "json",
          name: "JSON",
          description: "JSON format",
          extension: ".json",
          default_encoding: "UTF-8",
          importable: true,
          exportable: true,
          includes_locale_information: false
        }
      ]
    )

    # Mock the upload create endpoint
    mock_set!("POST", "/projects/#{project_id}/uploads",
      status: 201,
      body: upload_response.call(upload_en_id, locale_en_id, "en.yml")
    )

    # Mock the upload show endpoint (for --wait flag)
    mock_set!("GET", "/projects/#{project_id}/uploads/#{upload_en_id}",
      status: 200,
      body: upload_response.call(upload_en_id, locale_en_id, "en.yml")
    )

    mock_set!("GET", "/projects/#{project_id}/uploads/#{upload_de_id}",
      status: 200,
      body: upload_response.call(upload_de_id, locale_de_id, "de.yml")
    )
  end

  def create_locale_file(path, content)
    FileUtils.mkdir_p(File.dirname(path))
    File.write(path, content)
  end

  describe "basic push operation" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    it "uploads locale files successfully" do
      # Create locale files
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "de.yml"), de_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify upload requests were made
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests.length).to be >= 2
    end

    it "makes authenticated requests to the API" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "de.yml"), de_content)

      run_cli("push", config: config)

      requests_made = mock_requests

      # Check locales list request
      locales_request = requests_made.find { |r| r["path"] == "/projects/#{project_id}/locales" }
      expect(locales_request).not_to be_nil
      expect(locales_request["headers"]["HTTP_AUTHORIZATION"]).to eq("token #{token}")

      # Check upload requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }
      expect(upload_requests).not_to be_empty

      upload_requests.each do |upload_request|
        expect(upload_request["headers"]["HTTP_AUTHORIZATION"]).to eq("token #{token}")
      end
    end
  end

  describe "push with locale_name placeholder" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/i18n/<locale_name>.yml"
              params:
                file_format: yml
      YAML
    end

    it "reads locale name from file path" do
      # Create files using locale names
      create_locale_file(File.join(@tmpdir, "i18n", "English.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "i18n", "German.yml"), de_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify upload requests were made
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests.length).to be >= 2
    end
  end

  describe "push with specific locale" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/en.yml"
              params:
                file_format: yml
                locale_id: "#{locale_en_id}"
      YAML
    end

    it "uploads only the specified locale" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify only one upload request was made
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests.length).to eq(1)
    end
  end

  describe "push with tags" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
                tags: "frontend,v2"
      YAML
    end

    it "successfully pushes with tags configured" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "de.yml"), de_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify upload requests were made
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests).not_to be_empty

      # Verify that tags are included in the request
      upload_requests.each do |request|
        body = request["body"]
        # Check that the tags parameter contains both tags
        expect(body).to match(/name="tags"\r?\n\r?\n[^\r\n]*frontend[^\r\n]*v2/m)
      end
    end
  end

  describe "push with locale_mapping" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          locale_mapping:
            English: "eng"
            German: "ger"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_name>.yml"
              params:
                file_format: yml
      YAML
    end

    it "applies reverse locale mapping when uploading" do
      # Create files with mapped locale names (local names)
      create_locale_file(File.join(@tmpdir, "locales", "eng.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "ger.yml"), de_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify upload requests were made
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      # Should upload both files, mapping back to remote locale names
      expect(upload_requests.length).to be >= 2

      # Verify that the reverse mapping worked by checking that:
      # - Files with local names "eng" and "ger" were successfully uploaded
      # - They matched the correct remote locales (English -> locale-en-id, German -> locale-de-id)
      locale_ids_used = upload_requests.map { |req|
        # Extract locale_id from multipart form data
        req["body"].match(/name="locale_id"\r?\n\r?\n([^\r\n]+)/m)&.captures&.first&.strip
      }.compact

      # Should contain the IDs of the English and German locales
      expect(locale_ids_used).to include(locale_en_id)
      expect(locale_ids_used).to include(locale_de_id)

      # Verify files were uploaded with their local names in the filename
      filenames = upload_requests.map { |req|
        req["body"].match(/filename="([^"]+)"/m)&.captures&.first
      }.compact

      expect(filenames).to include("eng.yml")
      expect(filenames).to include("ger.yml")
    end

    context "two locales map to the same local name" do
      let(:config) do
        <<~YAML
          phrase:
            host: #{ENV.fetch("BASE_URL")}
            project_id: "#{project_id}"
            access_token: "#{token}"
            locale_mapping:
              en-US: "english"
              en-GB: "english"
            push:
              sources:
              - file: "#{@tmpdir}/locales/<locale_code>.yml"
                params:
                  file_format: yml
        YAML
      end

      it "returns an error indicating the conflict" do
        r = run_cli("push", config: config)

        expect(r[:exit_code]).not_to eq(0)
        expect(r[:stderr]).to match(/locale_mapping error.*both 'en-..' and 'en-..' map to the same local name 'english'/i)
      end
    end
  end

  describe "push with --wait flag" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    it "waits for upload processing to complete" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)

      r = run_cli("push", "--wait", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify upload show requests were made (polling for status)
      requests_made = mock_requests
      upload_show_requests = requests_made.select { |r| r["method"] == "GET" && r["path"].include?("/uploads/") }

      expect(upload_show_requests).not_to be_empty
    end
  end

  describe "push with update_translations parameter" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
                update_translations: true
      YAML
    end

    it "successfully pushes with update_translations flag" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "de.yml"), de_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests).not_to be_empty

      # Verify that update_translations parameter is set to true
      upload_requests.each do |request|
        body = request["body"]
        expect(body).to match(/name="update_translations"\r?\n\r?\n(true|1)/m)
      end
    end
  end

  describe "error handling" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    it "handles authentication errors" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)

      # Override the mock to return 401
      mock_set!("POST", "/projects/#{project_id}/uploads",
        status: 401,
        body: { message: "Unauthorized" }
      )

      r = run_cli("push", config: config)

      expect(r[:exit_code]).not_to eq(0)
      expect(r[:stderr]).to include("401")
    end

    it "handles missing project errors" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)

      # Override the mock to return 404
      mock_set!("GET", "/projects/#{project_id}/locales",
        status: 404,
        body: { message: "Project not found" }
      )

      r = run_cli("push", config: config)

      expect(r[:exit_code]).not_to eq(0)
    end

    it "returns error when source files don't exist" do
      # Don't create any files

      r = run_cli("push", config: config)

      expect(r[:exit_code]).not_to eq(0)
      expect(r[:stderr]).to match(/could not find|no files/i)
    end

    it "handles upload processing errors" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)

      # Mock upload to return error state
      error_upload = upload_response.call(upload_en_id, locale_en_id, "en.yml")
      error_upload[:state] = "error"

      mock_set!("POST", "/projects/#{project_id}/uploads",
        status: 201,
        body: error_upload
      )

      mock_set!("GET", "/projects/#{project_id}/uploads/#{upload_en_id}",
        status: 200,
        body: error_upload
      )

      r = run_cli("push", "--wait", config: config)

      expect(r[:exit_code]).not_to eq(0)
    end
  end

  describe "push without configuration" do
    it "returns error when no sources are specified" do
      config = <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
      YAML

      r = run_cli("push", config: config)

      expect(r[:exit_code]).not_to eq(0)
      expect(r[:stderr]).to match(/no sources|push.*not.*specified/i)
    end
  end

  describe "push with multiple sources" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
            - file: "#{@tmpdir}/translations/<locale_name>.yaml"
              params:
                file_format: yml
      YAML
    end

    it "uploads files from multiple source locations" do
      # Create files for first source
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "de.yml"), de_content)

      # Create files for second source
      create_locale_file(File.join(@tmpdir, "translations", "English.yaml"), en_content)
      create_locale_file(File.join(@tmpdir, "translations", "German.yaml"), de_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      # Should have uploaded 4 files total (2 from each source)
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests.length).to be >= 4
    end
  end

  describe "push with wildcard pattern" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/*.yml"
              params:
                file_format: yml
                locale_id: "#{locale_en_id}"
      YAML
    end

    it "uploads all matching files" do
      create_locale_file(File.join(@tmpdir, "locales", "translations.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "messages.yml"), en_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      # Should upload both files
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests.length).to be >= 2
    end
  end

  describe "push with recursive wildcard pattern" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/**/locales/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    it "uploads files from nested directories" do
      # Create files in multiple nested directories
      create_locale_file(File.join(@tmpdir, "app1", "locales", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "app1", "locales", "de.yml"), de_content)
      create_locale_file(File.join(@tmpdir, "app2", "locales", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "app2", "locales", "de.yml"), de_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      # Should upload 4 files total
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests.length).to be >= 4
    end
  end

  describe "push with --branch flag" do
    let(:branch_name) { "feature-branch" }
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    before do
      # Mock branch show endpoint
      mock_set!("GET", "/projects/#{project_id}/branches/#{branch_name}",
        status: 200,
        body: {
          name: branch_name,
          state: "success",
          created_at: "2024-01-01T00:00:00Z"
        }
      )
    end

    it "uploads to specified branch" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "de.yml"), de_content)

      r = run_cli("push", "--branch", branch_name, config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify upload requests were made
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests).not_to be_empty

      # Verify that branch parameter is included in the requests
      upload_requests.each do |request|
        body = request["body"]
        expect(body).to match(/name="branch"\r?\n\r?\n#{branch_name}/m)
      end
    end
  end

  describe "push with --tag flag" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    it "applies tag to uploaded files" do
      create_locale_file(File.join(@tmpdir, "locales", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "de.yml"), de_content)

      r = run_cli("push", "--tag", "release-v1.0", config: config)

      expect(r[:exit_code]).to eq(0)

      # Verify upload requests were made
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests).not_to be_empty

      # Verify that the tag is included in the requests
      upload_requests.each do |request|
        body = request["body"]
        expect(body).to match(/name="tags"\r?\n\r?\n[^\r\n]*release-v1\.0/m)
      end
    end
  end

  describe "push with tag placeholder" do
    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<tag>/<locale_code>.yml"
              params:
                file_format: yml
      YAML
    end

    it "extracts tag from file path" do
      create_locale_file(File.join(@tmpdir, "locales", "frontend", "en.yml"), en_content)
      create_locale_file(File.join(@tmpdir, "locales", "frontend", "de.yml"), de_content)
      create_locale_file(File.join(@tmpdir, "locales", "backend", "en.yml"), en_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      # Should upload 3 files with tags extracted from path
      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests.length).to be >= 3

      # Verify that tags were extracted from the file paths
      tags_used = upload_requests.map { |req|
        req["body"].match(/name="tags"\r?\n\r?\n([^\r\n]+)/m)&.captures&.first&.strip
      }.compact

      # Should have both "frontend" and "backend" tags
      expect(tags_used).to include("frontend")
      expect(tags_used).to include("backend")
    end
  end

  describe "push with JSON format" do
    let(:json_content) do
      <<~JSON
        {
          "hello": "Hello",
          "world": "World"
        }
      JSON
    end

    let(:config) do
      <<~YAML
        phrase:
          host: #{ENV.fetch("BASE_URL")}
          project_id: "#{project_id}"
          access_token: "#{token}"
          push:
            sources:
            - file: "#{@tmpdir}/locales/<locale_code>.json"
              params:
                file_format: json
                locale_id: "#{locale_en_id}"
      YAML
    end

    it "uploads JSON files successfully" do
      create_locale_file(File.join(@tmpdir, "locales", "en.json"), json_content)

      r = run_cli("push", config: config)

      expect(r[:exit_code]).to eq(0)

      requests_made = mock_requests
      upload_requests = requests_made.select { |r| r["method"] == "POST" && r["path"].include?("/uploads") }

      expect(upload_requests).not_to be_empty
    end
  end
end
