package common_test

import (
	"testing"

	"github.com/patos-ufscar/http-web-server-example-go/common"
)

func TestExtractRegExpFromHostStr(t *testing.T) {
	r := common.ExtractRegExpFromHostStr("r`.+`")
	if r != ".+" {
		t.Errorf("common.ExtractRegExpFromHostStr(\"r`.+`\") = \"%s\"; want \".+\"", r)
	}

	r = common.ExtractRegExpFromHostStr("*")
	if r != ".+" {
		t.Errorf("common.ExtractRegExpFromHostStr(\"*\") = \"%s\"; want \".+\"", r)
	}

	r = common.ExtractRegExpFromHostStr("abc.example.com")
	if r != "^abc\\.example\\.com$" {
		t.Errorf("common.ExtractRegExpFromHostStr(\"abc.example.com\") = \"%s\"; want \"^abc\\.example\\.com\"$", r)
	}
}