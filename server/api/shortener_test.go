package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chinmayweb3/urlshortner/model"
	"github.com/stretchr/testify/require"
)

type shortenerType struct {
	lurl         string
	isStatusCode int
	isBroken     bool
}

var Shortenerargs = []shortenerType{
	{lurl: "https://youtube.com/", isStatusCode: 200, isBroken: true},
	{lurl: "https://udemy.com/", isStatusCode: 200, isBroken: true},
	{lurl: "https://quickref.me/", isStatusCode: 200, isBroken: true},
	{lurl: "https://primevideo.com/", isStatusCode: 200, isBroken: true},
	{lurl: "https://asuratoon.com/", isStatusCode: 200, isBroken: true},
}

func TestShortener(t *testing.T) {
	r := SetUpRouter()
	const urlreq = "/api/v1/shortener"
	r.POST(urlreq, Shortener)
	var now = time.Now()

	for _, arg := range Shortenerargs {
		var s model.Url

		by, err := json.Marshal(map[string]string{"url": arg.lurl})
		require.NoError(t, err)

		req, _ := http.NewRequest("POST", urlreq, bytes.NewBuffer(by))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		body, _ := io.ReadAll(w.Body)
		json.Unmarshal(body, &s)

		if w.Code == 200 {
			require.Equal(t, arg, s.LUrl)
		} else if w.Code == 403 {

		}

		require.WithinDuration(t, now.Local(), s.LastViewed.Local(), 2*time.Second)

	}

}
