package httpRequest

import (
	"log"
	"testing"
)

func TestRequest(t *testing.T) {
	r := ApiRequest{
		BaseUri: "",
		Url:     "https://studygolang.com/object/comments?objid=12784",
		Params: map[string]interface{}{
			// "objid": 12784,
			"objtype": 1,
			"p":       0,
		},
		Method: "get",
	}

	res := r.Do()
	if res.Err != nil {
		log.Println("err: ", res.Err)
		t.Error(res.Err)
		return
	}

	log.Println("data: ", string(res.Body))

}

/**
$ go test -v
--- PASS: TestRequest (0.68s)
PASS
ok      github.com/daheige/thinkgo/httpRequest  0.681s
*/
