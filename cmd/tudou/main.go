package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/iochen/tudou"
)

func help() {
	fmt.Println("Usage: tudou [encode/decode]")
}

func main() {
	if len(os.Args) !=2 {
		help()
		return
	}

	switch os.Args[1] {
	case "encode":
		fmt.Println("tudou: Go ahead and type your message ...")
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
			return
		}
		s, err := tudou.Encode(b)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(s)
	case "decode":
		fmt.Println("tudou: Go ahead and type your code ...")
		b, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
			return
		}
		s := strings.TrimSpace(string(b))
		msg, err := tudou.Decode(s)
		if err != nil {
			panic(err)
			return
		}
		fmt.Println(string(msg))
	default:
		help()
		return
	}

}