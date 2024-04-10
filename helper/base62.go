package helper

import (
	"fmt"

	"github.com/chinmayweb3/urlshortner/model"
	"github.com/deatil/go-encoding/base62"
)

func Base62Encoding() string {
	countNum := model.GetCountNumber()
	encode := string(base62.StdEncoding.Encode([]byte(fmt.Sprintf("%09d", countNum))))
	return encode
}
