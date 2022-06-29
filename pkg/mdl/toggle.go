package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

func Checkbox(id string, checked bool, text string) view.View {
	return dom.Attributed(
		dom.Label(
			dom.Attributed(
				dom.Input(),
				dom.SetAttribute("type", "checkbox"),
				dom.SetAttribute("id", id),
				dom.SetAttribute("class", "mdl-checkbox__input"),
			),
			dom.Attributed(
				dom.Span(dom.Text(text)),
				dom.Class("mdl-checkbox__label"),
			),
		),
		dom.Class("mdl-checkbox mdl-js-checkbox mdl-js-ripple-effect"),
		dom.SetAttribute("for", id),
	)
}

func RadioButton(id, name, value, text string) view.View {
	return dom.Attributed(
		dom.Label(
			dom.Attributed(
				dom.Input(),
				dom.SetAttribute("type", "radio"),
				dom.SetAttribute("id", id),
				dom.SetAttribute("class", "mdl-radio__button"),
				dom.SetAttribute("name", name),
				dom.SetAttribute("value", value),
			),
			dom.Attributed(
				dom.Span(dom.Text(text)),
				dom.SetAttribute("class", "mdl-radio__label"),
			),
		),
		dom.SetAttribute("class", "mdl-radio mdl-js-radio mdl-js-ripple-effect"),
		dom.SetAttribute("for", id),
	)
}

func IconToggle(id string, checked bool, icon string) view.View {
	return dom.Attributed(
		dom.Label(
			dom.Attributed(
				dom.Input(),
				dom.SetAttribute("type", "checkbox"),
				dom.SetAttribute("id", id),
				dom.SetAttribute("class", "mdl-icon-toggle__input"),
				func(as dom.AttributeStorage) {
					if checked {
						as["checked"] = ""
					}
				},
			),
			dom.Attributed(
				dom.I(dom.Text(icon)),
				dom.SetAttribute("class", "mdl-icon-toggle__label material-icons"),
			),
		),
		dom.SetAttribute("class", "mdl-icon-toggle mdl-js-icon-toggle mdl-js-ripple-effect"),
		dom.SetAttribute("for", id),
	)
}

func Switch(id string, value bool) view.View {
	return dom.Attributed(
		dom.Label(
			dom.Attributed(
				dom.Input(),
				dom.SetAttribute("type", "checkbox"),
				dom.SetAttribute("id", id),
				dom.Class("mdl-switch__input"),
				func(as dom.AttributeStorage) {
					if value {
						as["checked"] = ""
					}
				},
			),
			dom.Attributed(
				dom.Span(),
				dom.Class("mdl-switch__label"),
			),
		),
		dom.Class("mdl-switch mdl-js-switch mdl-js-ripple-effect"),
		dom.SetAttribute("for", id),
	)
}
