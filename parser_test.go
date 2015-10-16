package clier_test

import "testing"

func TestOptionalsSingleDash(t *testing.T) {
	cases := []struct {
		Input  string
		Output interface{}
	}{
		{
			Input:  "",
			Output: NameSpace{},
		},
	}
}
