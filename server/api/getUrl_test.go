package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type getType struct {
	test   string
	status int
}

var args = []getType{
	{test: "/test", status: 404},
	{test: "/H57EuVixKg0v", status: 301},
}

func TestGetUrl(t *testing.T) {
	r := SetUpRouter()
	r.GET("/:shortUrl", GetHandler)

	for _, arg := range args {
		var s map[string]string

		req, _ := http.NewRequest("GET", arg.test, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		body, _ := io.ReadAll(w.Body)
		json.Unmarshal(body, &s)

		if arg.status == 404 {
			require.Equal(t, arg.status, w.Result().StatusCode)
		} else if arg.status == 301 {
			require.Empty(t, s["error"])
			require.Equal(t, arg.status, w.Result().StatusCode)
		} else {
			log.Fatalln("error in test case")
		}

	}

}
