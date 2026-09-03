require "spec_helper"

RSpec.describe "phrase locales download" do
  let(:token) { "test-token-locales-download" }
  let(:project_id) { "test-project-456" }
  let(:locale_id) { "en-US" }

  before do
    mock_clear_requests!
  end

  # The API responds with 200 and an empty body when a --tags filter
  # matches zero translations. This used to crash the generated
  # "locales download" command with a nil pointer dereference, because
  # it unconditionally called Close()/Name() on the (nil) result file.
  it "does not crash when a tag filter matches no translations" do
    mock_set!("GET", "/projects/#{project_id}/locales/#{locale_id}/download",
      status: 200,
      body: "",
      headers: { "content-type" => "application/x-properties" }
    )

    r = run_cli(
      "locales", "download",
      "--id", locale_id,
      "--project_id", project_id,
      "--file_format", "properties",
      "--tags", "nonexistent-tag",
      "-t", token,
      "--host", ENV.fetch("BASE_URL")
    )

    expect(r[:stderr]).not_to include("panic")
    expect(r[:exit_code]).to eq(0)
    expect(r[:stdout]).to eq("")
  end

  it "still prints the downloaded content for a normal response" do
    mock_set!("GET", "/projects/#{project_id}/locales/#{locale_id}/download",
      status: 200,
      body: "hello=world",
      headers: { "content-type" => "application/x-properties" }
    )

    r = run_cli(
      "locales", "download",
      "--id", locale_id,
      "--project_id", project_id,
      "--file_format", "properties",
      "-t", token,
      "--host", ENV.fetch("BASE_URL")
    )

    expect(r[:exit_code]).to eq(0)
    expect(r[:stdout]).to include("hello=world")
  end
end
