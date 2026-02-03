require "net/http"
require "json"
require "uri"

module MockControl
  def mock_reset!
    post_json("/__control__/reset", {})
  end

  def mock_set!(method, path, status:, body:, headers: {})
    post_json("/__control__/set", {
      method: method,
      path: path,
      status: status,
      headers: headers,
      body: body
    })
  end

  def mock_requests
    get_json("/__control__/requests").fetch("requests")
  end

  def mock_clear_requests!
    post_json("/__control__/requests/clear", {})
  end

  private

  def base_uri
    URI(ENV.fetch("BASE_URL"))
  end

  def post_json(path, payload)
    uri = URI.join(base_uri.to_s, path)
    req = Net::HTTP::Post.new(uri)
    req["content-type"] = "application/json"
    req.body = JSON.dump(payload)

    Net::HTTP.start(uri.host, uri.port) do |http|
      res = http.request(req)
      raise "Mock control failed: #{res.code} #{res.body}" unless res.is_a?(Net::HTTPSuccess)
    end
  end

  def get_json(path)
    uri = URI.join(base_uri.to_s, path)
    res = Net::HTTP.get_response(uri)
    raise "Mock control failed: #{res.code} #{res.body}" unless res.is_a?(Net::HTTPSuccess)
    JSON.parse(res.body)
  end
end
