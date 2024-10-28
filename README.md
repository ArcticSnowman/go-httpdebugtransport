# go-httpdebugtransport

[![Go Reference](https://pkg.go.dev/badge/github.com/ArcticSnowman/go-httpdebugtransport.svg)](https://pkg.go.dev/github.com/ArcticSnowman/go-httpdebugtransport)

Go utility to debug HTTP requests/response, with support for authenication

## Usage

```go

    import (
        github.com/ArcticSnowman/go-httpdebugtransport
    )

    ct := httpdebugtransport.New()  
    ct.SetToken("some-token")
    ct.SetDebug(true)
  
  	client := http.Client{
		Transport: ct.Client(),
		Timeout:   4 * time.Second,
	}

    client.Get(url)
```


