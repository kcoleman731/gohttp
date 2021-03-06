# gohhtp

This project contains the source code for the `gohttp` REST client. `gohttp` provides syntactic sugar built on top of the `net/http` package, designed to make interacting with REST APIs from `golang` applications simple and easy.

## Install

```
$ go get github.com/mercambia/gohttp
```

## Getting started

### Client

`gohttp` provides a `Client` struct which can be initialized with a base URL and optionally, global headers.

```
headers := map[string]string{"Accept" : "application/json"}
gohttp.NewClient("www.some-url.com", headers)
```

### Request

`gohttp` provides a `Request` struct which makes building HTTP request simple and readable.

```
request := &gohttp.Request{
	Method: gohttp.GET,
	URL:    "www.test.com/users",
}
body, err := client.Execute(request)
if err != nil {
        return err
}
fmt.Printf("Response Body: %v", body)
```

Responses from a call to `Execute` on the `Client` struct will return an `error` object and a `map[string]interface{}` object containing the deserialized response body from the request.

// TODO - Build a `Response` struct.
