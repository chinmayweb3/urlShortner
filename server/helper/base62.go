package helper

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/deatil/go-encoding/base62"
)

func Base62Encoding(t time.Time) string {
	return string(base62.StdEncoding.Encode([]byte(t.Format("2006-01-02 15:04:05" + strconv.Itoa(rand.Intn(10000))))))[0:7]
}
