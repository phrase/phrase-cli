require "json"

$mock = {
  routes: {},
  requests: [] # array of hashes
}

def json(status, obj, headers: {})
  [status, { "content-type" => "application/json" }.merge(headers), [JSON.dump(obj)]]
end

def read_body(env)
  io = env["rack.input"]
  return "" unless io
  body = io.read
  io.rewind if io.respond_to?(:rewind)
  body
end

def record_request!(env)
  method = env["REQUEST_METHOD"]
  path   = env["PATH_INFO"]
  query  = env["QUERY_STRING"]
  body   = read_body(env)

  # Record a safe subset of headers (HTTP_... keys + content-type/length)
  headers = env.each_with_object({}) do |(k, v), h|
    next unless k.start_with?("HTTP_") || k == "CONTENT_TYPE" || k == "CONTENT_LENGTH"
    h[k] = v
  end

  $mock[:requests] << {
    method: method,
    path: path,
    query: query,
    headers: headers,
    body: body
  }
end

run lambda { |env|
  method = env["REQUEST_METHOD"]
  path   = env["PATH_INFO"]

  # ---- Control API ----
  if method == "POST" && path == "/__control__/reset"
    $mock[:routes].clear
    $mock[:requests].clear
    return json(200, { ok: true })
  end

  if method == "POST" && path == "/__control__/set"
    payload = JSON.parse(read_body(env))
    m = payload.fetch("method")
    p = payload.fetch("path")

    $mock[:routes][[m, p]] = {
      status:  payload.fetch("status"),
      headers: payload["headers"] || {},
      body:    payload["body"] || ""
    }

    return json(200, { ok: true })
  end

  if method == "GET" && path == "/__control__/requests"
    return json(200, { requests: $mock[:requests] })
  end

  if method == "POST" && path == "/__control__/requests/clear"
    $mock[:requests].clear
    return json(200, { ok: true })
  end

  # ---- Record every non-control request ----
  record_request!(env)

  # ---- Serve mocked responses ----
  route = $mock[:routes][[method, path]]
  return json(404, { error: "not mocked", method: method, path: path }) unless route

  status  = route[:status]
  headers = route[:headers] || {}
  body    = route[:body]

  body_str =
    case body
    when String then body
    else JSON.dump(body)
    end

  headers = { "content-type" => "application/json" }.merge(headers) if body.is_a?(Hash) || body.is_a?(Array)

  [status, headers, [body_str]]
}
