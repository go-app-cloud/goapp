package goapp

import (
	"fmt"
	"testing"
)

func TestBuildPassword(t *testing.T) {
	fmt.Println(BuildPassword(32, Number))
	fmt.Println(BuildPassword(32, Char))
	fmt.Println(BuildPassword(32, NumberAndChar))
	fmt.Println(BuildPassword(32, Advance))
}
