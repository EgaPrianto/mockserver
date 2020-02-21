# mockserver

This is for mock a server and logs any request to the server.

### Prerequisites

- Golang 1.13

### Installing

Simply

```
go get github.com/EgaPrianto/mockserver
```

### Usage

Commands format
```
mockserver <port> <response_body>
```
Parameter port and response_body are optional

You need to add response.yml file in you current working directory.
Below are the example of response.yml.
```
Status: 200
Body: >-
  {
    "message": "test",
    "data": {},
    "meta": {
        "http_status": 201
    }
  }
Header:
  Content-Type:
    - application/json
  X-Example-Type:
    - Example
```

Run the program, with default port. Default port is 1234

```
/$ ls
response.yml
/$ mockserver
Listening on Port :1234
```

Then try do any request to port 1234, for example:
```
curl -v POST \
>   'http://localhost:1234/test/path?test=query' \
>   -H 'Content-Length: 15' \
>   -H 'Test: Header' \
>   -d '{"test":"body"}'

* Rebuilt URL to: POST/
* Could not resolve host: POST
* Closing connection 0
curl: (6) Could not resolve host: POST
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 1234 (#1)
> POST /test/path?test=query HTTP/1.1
> Host: localhost:1234
> User-Agent: curl/7.58.0
> Accept: */*
> Content-Length: 15
> Test: Header
> Content-Type: application/x-www-form-urlencoded
> 
* upload completely sent off: 15 out of 15 bytes
< HTTP/1.1 200 OK
< Content-Type: application/json
< X-Example-Type: Example
< Date: Thu, 06 Feb 2020 05:21:09 GMT
< Content-Length: 79
< 
{
  "message": "test",
  "data": {},
  "meta": {
      "http_status": 201
  }
* Connection #1 to host localhost left intact
}
```

And mockserver will log the reques as shown below:

```
Listening on Port :1234
METHOD = POST
======URL================================
/test/path?test=query
======HEADERS================================
{
  "Accept": [
    "*/*"
  ],
  "Content-Length": [
    "15"
  ],
  "Content-Type": [
    "application/x-www-form-urlencoded"
  ],
  "Test": [
    "Header"
  ],
  "User-Agent": [
    "curl/7.58.0"
  ]
}
======QUERY PARAMS================================
{
  "test": [
    "query"
  ]
}
======BODY================================
{
  "test": "body"
}

```

## Authors

* **Ega Prianto**