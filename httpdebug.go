package httpdebugtransport

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

type DebugAuthTransport struct {
	Username string
	Password string
	Token    string
	Debug    bool

	Transport http.RoundTripper
}

// New DebugAuthTransport
func New() *DebugAuthTransport {

	return &DebugAuthTransport{
		Debug: false,
	}

}

// SetCredentails - Set username/password for basic authentication
func (t *DebugAuthTransport) SetCredentails(username string, password string) {
	t.Username = username
	t.Password = password
}

// SetToken - Set Authorization header in request.
func (t *DebugAuthTransport) SetToken(token string) {
	t.Token = token
}

// SetDebug - Enable debug output
func (t *DebugAuthTransport) SetDebug(debug bool) {
	t.Debug = debug
}

// Returns http client with transport
func (t *DebugAuthTransport) Client() *http.Client {
	return &http.Client{Transport: t}
}

func (t *DebugAuthTransport) transport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

func (t *DebugAuthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req2 := cloneRequest(req) // per RoundTripper contract

	if t.Token != "" {
		req2.Header.Add("Authorization", fmt.Sprintf("Bearer %s", t.Token))
	} else {
		req2.SetBasicAuth(t.Username, t.Password)
	}

	if t.Debug {
		dump, _ := httputil.DumpRequestOut(req, true)
		fmt.Printf("****REQUEST****\n%q\n", dump)
	}
	resp, err := t.transport().RoundTrip(req2)

	if t.Debug {
		dump, _ := httputil.DumpResponse(resp, true)
		fmt.Printf("****RESPONSE****\n%q\n****************\n\n", dump)
	}

	return resp, err
}

// cloneRequest returns a clone of the provided *http.Request.
// The clone is a shallow copy of the struct and its Header map.
func cloneRequest(r *http.Request) *http.Request {
	// shallow copy of the struct
	r2 := new(http.Request)
	*r2 = *r
	// deep copy of the Header
	r2.Header = make(http.Header, len(r.Header))
	for k, s := range r.Header {
		r2.Header[k] = append([]string(nil), s...)
	}
	return r2
}
