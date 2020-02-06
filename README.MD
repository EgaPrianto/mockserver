# mockserver

This is for mock an server and logs any request to the server.

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
curl -X POST \
>   'http://localhost:1234/test/path?test=query' \
>   -H 'Content-Length: 15' \
>   -H 'Test: Header' \
>   -d '{"test":"body"}'

{
  "message": "test",
  "data": {},
  "meta": {
      "http_status": 201
  }
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
## Contributing

Please read [CONTRIBUTING.md](https://gist.github.com/PurpleBooth/b24679402957c63ec426) for details on our code of conduct, and the process for submitting pull requests to us.

## Authors

* **Ega Prianto**