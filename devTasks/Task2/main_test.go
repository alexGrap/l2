package main

import (
	"fmt"
	"testing"
)

func TestUnpackingFunc(t *testing.T) {
	testingStrings := []struct {
		inputStr    string
		expectedStr string
		err         error
	}{
		{
			"a4bc2d5e",
			"aaaabccddddde",
			nil,
		},
		{
			"abcd",
			"abcd",
			nil,
		},
		{
			"45",
			"",
			fmt.Errorf("incorrent string"),
		},
		{
			"",
			"",
			nil,
		},
		{
			`qwe\4\5`,
			"qwe45",
			nil,
		},
		{
			`qwe\45`,
			"qwe44444",
			nil,
		},
		{
			`qwe\\5`,
			`qwe\\\\\`,
			nil,
		},

		{
			`a4b5r8\9\56\\7`,
			`aaaabbbbbrrrrrrrr9555555\\\\\\\`,
			nil,
		},
	}
	for _, testItem := range testingStrings {
		s, err := UnpackingFunc(testItem.inputStr)
		if s != testItem.expectedStr {
			t.Errorf("fail test with string: %v", testItem.inputStr)
		}
		if !((err != nil && testItem.err != nil) || (err == nil && testItem.err == nil)) {
			t.Errorf("fail test with string: %v err: %v", testItem.inputStr, testItem.err)
		}
	}
}
