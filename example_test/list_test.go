package example_test

import (
	"fmt"
	"github.com/T-c33/gostl/list"
	"testing"
)

func TestList(t *testing.T) {
	l := list.NewList()
	l.Append(3)
	l.Append(4)
	fmt.Printf("%+v",l)
	fmt.Printf("%+v",l.GetListHead())
	fmt.Printf("%+v",l.GetListTail())
}