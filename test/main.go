package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("error one")
	err2 := errors.New("error two")
	er3 := fmt.Errorf("error three")
	err := errors.Join(err1, err2, er3)
	fmt.Println(err)
}
