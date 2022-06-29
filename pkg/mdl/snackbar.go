package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

func Snackbar(id string) view.View {
	return dom.Attributed(
		dom.Div(
			dom.Attributed(
				dom.Div(),
				dom.Class("mdl-snackbar__text"),
			),
			dom.Attributed(
				dom.Button(),
				dom.SetAttribute("class", "mdl-snackbar__action"),
				dom.SetAttribute("type", "button"),
			),
		),
		dom.SetAttribute("id", id),
		dom.SetAttribute("class", "mdl-js-snackbar mdl-snackbar"),
	)
}
