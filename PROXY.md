New ReverseProxy Rewrite hook
The httputil.ReverseProxy forwarding proxy includes a new Rewrite hook function, superseding the previous Director hook.

The Rewrite hook accepts a ProxyRequest parameter, which includes both the inbound request received by the proxy and the outbound request that it will send. Unlike Director hooks, which only operate on the outbound request, this permits Rewrite hooks to avoid certain scenarios where a malicious inbound request may cause headers added by the hook to be removed before forwarding. See issue #50580.

The ProxyRequest.SetURL method routes the outbound request to a provided destination and supersedes the NewSingleHostReverseProxy function. Unlike NewSingleHostReverseProxy, SetURL also sets the Host header of the outbound request.

The ProxyRequest.SetXForwarded method sets the X-Forwarded-For, X-Forwarded-Host, and X-Forwarded-Proto headers of the outbound request. When using a Rewrite, these headers are not added by default.

An example of a Rewrite hook using these features is:

proxyHandler := &httputil.ReverseProxy{
  Rewrite: func(r *httputil.ProxyRequest) {
    r.SetURL(outboundURL) // Forward request to outboundURL.
    r.SetXForwarded()     // Set X-Forwarded-* headers.
    r.Out.Header.Set("X-Additional-Header", "header set by the proxy")
  },
}
ReverseProxy no longer adds a User-Agent header to forwarded requests when the incoming request does not have one.