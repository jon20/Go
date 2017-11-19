package get_wes

import (
	"reflect"
	"testing"
)

func TestWes(t *testing.T) {
	s := reflect.TypeOf(Getwes_a())
	if s.Kind() != reflect.String {
		t.Fatal("not a string")
	}
}
