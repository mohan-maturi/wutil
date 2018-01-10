package w

import (
	"strings"
	"testing"
)

func TestPrintWorld(t *testing.T) {
	arg := "Anika"
	ret := PrintWorld(arg)
	expected_ret := "Hello World!!! This is " + arg
	if strings.Compare(ret, expected_ret) != 0 {
		t.Error("Expected \"%s\", instead got \"%s\"", expected_ret, ret)
	}
}
