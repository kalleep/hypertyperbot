package util

import (
	"log"
)

//Just becaouse im lazy
func PanicIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
