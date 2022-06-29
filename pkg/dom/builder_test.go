package dom_test

import (
	"testing"

	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

func Test_Text(t *testing.T) {
	br := dom.NewBuilder()
	res := br.Bytes(dom.Text("some text"))
	resStr := string(res)

	if resStr != "some text" {
		t.Log("error: res $ne some text")
		t.FailNow()
	}
}

func Test_Tag(t *testing.T) {
	br := dom.NewBuilder()
	res := br.Bytes(dom.Attributed(
		dom.P(dom.Text("some text")),
		dom.SetAttribute("class", "test"),
	))
	resStr := string(res)

	if resStr != "<p class=\"test\">some text</p>" {
		t.Log("error: res $ne some text")
		t.FailNow()
	}
}

func Test_Page(t *testing.T) {
	br := dom.NewBuilder()
	res := br.Bytes(dom.Html(view.Group(
		dom.Head(dom.Title(dom.Text("simple text"))),
	)))

	if string(res) != "<html><head><title>simple text</title></head></html>" {
		t.Logf("error: res $ne some text <html><head><title>simple text</title></head></html> is %v/\n", string(res))
		t.FailNow()
	}
}
