package tudou_test

import (
	"fmt"
	"testing"

	"github.com/iochen/tudou"
)

func TestEncodeAndDecode(t *testing.T) {
	s, err := tudou.Encode([]byte("iochen.com "))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(s)

	b, err := tudou.Decode(s)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
}