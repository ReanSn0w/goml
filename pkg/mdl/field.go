package mdl

import (
	"fmt"

	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

type FieldPreferences struct {
	Placeholder  string
	Pattern      string
	ErrorMessage string

	Rows          int // 0, 1 - использует input; > 1 - использует textarea
	FloatingLabel bool
}

func TextField(id string, preferences FieldPreferences) view.View {
	return dom.Attributed(
		dom.Div(
			view.If(preferences.Rows <= 1, dom.Attributed(
				dom.Input(),
				dom.SetAttribute("class", "mdl-textfield__input"),
				dom.SetAttribute("type", "text"),
				dom.SetAttribute("id", id),
				func(as dom.AttributeStorage) {
					if preferences.Pattern != "" {
						as["pattern"] = preferences.Pattern
					}
				},
			)),
			view.If(preferences.Rows > 1, dom.Attributed(
				dom.Textarea(),
				dom.SetAttribute("class", "mdl-textfield__input"),
				dom.SetAttribute("type", "text"),
				dom.SetAttribute("rows", fmt.Sprintf("%v", preferences.Rows)),
				dom.SetAttribute("id", id),
			)),
			dom.Attributed(
				dom.Label(dom.Text(preferences.Placeholder)),
				dom.SetAttribute("class", "mdl-textfield__label"),
				dom.SetAttribute("for", id),
			),
			view.If(preferences.ErrorMessage != "", dom.Attributed(
				dom.Span(dom.Text(preferences.ErrorMessage)),
				dom.SetAttribute("class", "mdl-textfield__error"),
			)),
		),
		dom.UpdateAttribute("class", func(s string) string {
			s += "mdl-textfield mdl-js-textfield"

			if preferences.FloatingLabel {
				s += " mdl-textfield--floating-label"
			}

			return s
		}),
	)
}

func ExpandableTextField(id, icon, placeholder string) view.View {
	return dom.Attributed(
		dom.Div(
			dom.Attributed(
				dom.Label(
					dom.Attributed(
						dom.I(dom.Text(icon)),
						dom.SetAttribute("class", ",aterial-icons"),
					),
				),
				dom.SetAttribute("class", "mdl-button mdl-js-button mdl-button--icon"),
				dom.SetAttribute("for", id),
			),
			dom.Attributed(
				dom.Div(
					dom.Attributed(
						dom.Input(),
						dom.SetAttribute("class", "mdl-textfield__input"),
						dom.SetAttribute("type", "text"),
						dom.SetAttribute("id", id),
					),
					dom.Attributed(
						dom.Label(dom.Text(placeholder)),
						dom.SetAttribute("class", "mdl-textfield__label"),
						dom.SetAttribute("for", fmt.Sprintf("%s-expandable", id)),
					),
				),
				dom.SetAttribute("class", "mdl-textfield__expandable-holder"),
			),
		),
		dom.SetAttribute("class", "mdl-textfield mdl-js-textfield mdl-textfield--expandable"),
	)
}
