package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

type getUrlType struct {
	param  string
	status int
}

var geturlargs = []getUrlType{
	{param: "/test", status: 404},
	{param: "/H57EuVixKg0v", status: 301},
}

func TestGetUrl(t *testing.T) {
	r := SetUpRouter()
	r.GET("/:shortUrl", GetHandler)

	for _, arg := range geturlargs {
		var s map[string]string

		req, _ := http.NewRequest("GET", arg.param, nil)
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
	fmt.Println("getUrl test pass")

}
