# Write simple cache proxy

The main purpose of the service is to proxy requests and save the responses into a cache.

## Requirements
* Use the Host Header and path to build an URL of a target service Host Header syntax.
Doc - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Host

* Use "Effective Request URI" RFC section as a guideline on how to build a target URL:
Doc - https://tools.ietf.org/html/rfc7230#section-5.5

## Examples:
* Request to the proxy (8080 is for non-TLS connections)
    curl -H "godoc.org" <http://0.0.0.0:8080/github.com/stretchr/testify/assert>
     	 
  Target URL:
    <http://godoc.org/github.com/stretchr/testify/assert>
     	 
* Request to the proxy (8080 is for non-TLS connections)
    curl -H "godoc.org:80" <http://0.0.0.0:8080/github.com/stretchr/testify/assert>
     	 
  Target URL:
    <http://godoc.org:80/github.com/stretchr/testify/assert>
     	 
* Request to the proxy (9443 is for TLS connections)
    curl -H "godoc.org" <http://0.0.0.0:9443/github.com/stretchr/testify/assert>
     	 
  Target URL:
    <https://godoc.org/github.com/stretchr/testify/assert>
     	 
* Request to the proxy (9443 is for TLS connections)
    curl -H "godoc.org:443" <http://0.0.0.0:9443/github.com/stretchr/testify/assert>
     	 
  Target URL:
    <https://godoc.org:443/github.com/stretchr/testify/assert>

** This implies the proxy server should be capable of receiving both HTTP & HTTPS requests. **

Use Cache-Control Header to control the lifetime of cached requests or to bypass caching.
Doc - https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Cache-Control

Service should only cache safe HTTP methods, see the link to Wikipedia to find out what is the safe method, in case the request isn't cacheable just proxy request
The cache is purely in-memory, i.e. while a service is running. There's no persistence of cached data in any persistent storage.

