# Echo

This server displays the request information in text form:

```
POST /sfsdfsdf/dsfsdfsdf HTTP/1.1
Host: localhost:3000
User-Agent: Apache-HttpClient/4.5.6 (Java/1.8.0_201)
Content-Length: 43
Accept-Encoding: gzip,deflate
Connection: Keep-Alive
Content-Type: application/json; charset=UTF-8

{
	"usr-id": 1,
	"name": "Subhash Chandran"
}

```

This is useful when configuring reverse proxies where you need to understand how the reverse proxy forwards the requests.
