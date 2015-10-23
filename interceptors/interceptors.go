package interceptors

import (
	"fmt"
	"net/http"
)

//Logger is a dummy example of possible interceptors for this service
type Logger struct{}

//Intercept is the Interceptor interface implementation
func (l Logger) Intercept(rw http.ResponseWriter, r *http.Request) error {
	fmt.Println("[LOG] Request from: ", r.RemoteAddr, " on: ", r.URL.String())
	return nil
}
