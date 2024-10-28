# go-httpdebugtransport

Go utility to debug HTTP requests/response, with support for authenication

## Simple Usage

```go
    ct := httpdebugtransport.New()  
    ct.SetToken("some-token")
    ct.SetDebug(true)
  
  	client := http.Client{
		Transport: ct.Client(),
		Timeout:   4 * time.Second,
	}

    client.Get(url)
```


