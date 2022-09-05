package ac

import (
	"regexp"
	"strings"
	"testing"
)

var mockTexts = []string{
	"he",
	"him",
	"she",
	"her",
	"shy",
	"say",
	"aye",
	"err",
}

func TestACMatch(t *testing.T) {
	ac := NewAC(mockTexts)
	text := "ashimsahery"

	matches := ac.Match(text)
	for i, m := range matches {
		t.Logf("[%d] r=[%s]", i, m)
	}
}

func TestAnalyse(t *testing.T) {
	entries := AnalyseTextList(mockTexts)
	for i, r := range entries {
		t.Logf("[%d] r=[%q]", i, r)
	}
}

func TestRegexpMatch(t *testing.T) {
	re, _ := regexp.Compile(strings.Join(mockTexts, "|"))
	text := "ashimsahery"

	matches := re.FindAllStringSubmatch(text, -1)
	for i, m := range matches {
		t.Logf("[%d] r=[%#v]", i, m)
	}
}
