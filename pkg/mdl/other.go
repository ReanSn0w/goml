package mdl

import (
	"fmt"

	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

func Tooltip(forID string, large bool, value string) view.View {
	return dom.Attributed(
		dom.Div(dom.Text(value)),
		dom.UpdateAttribute("class", func(s string) string {
			s += "mdl-tooltip"

			if large {
				s += " mdl-tooltip--large"
			}

			return s
		}),
		dom.SetAttribute("data-mdl-for", forID),
	)
}

func Slider(min, max, value int) view.View {
	return dom.Attributed(
		dom.Input(),
		dom.Class("mdl-slider mdl-js-slider"),
		dom.SetAttribute("type", "range"),
		dom.SetAttribute("min", fmt.Sprintf("%v", min)),
		dom.SetAttribute("max", fmt.Sprintf("%v", max)),
		dom.SetAttribute("value", fmt.Sprintf("%v", value)),
		dom.SetAttribute("tabindex", fmt.Sprintf("%v", 0)),
	)
}

// SetBadge - устанавливает бейдж для элемента
//
// позволяет добавить наклейку к элементу,
// имеется возможность добавить строку или число
func SetBadge(overlap bool, value interface{}) dom.AttributeEditor {
	return func(as dom.AttributeStorage) {
		s := as["class"]

		switch v := value.(type) {
		case int:
			if v == 0 {
				return
			}
		case string:
			if len(s) == 0 {
				return
			}
		default:
			return
		}

		s += " mdl-badge"

		if overlap {
			s += " mdl-badge--overlap"
		}

		as["class"] = s
		as["data-badge"] = fmt.Sprintf("%v", value)
	}
}

func Progress(id string, intermediate bool) view.View {
	return dom.Attributed(
		dom.Div(),
		dom.UpdateAttribute("class", func(s string) string {
			s += "mdl-progress mdl-js-progress"

			if intermediate {
				s += " mdl-progress__indeterminate"
			}

			return s
		}),
	)
}

func Spinner(singleColor bool) view.View {
	return dom.Attributed(
		dom.Div(),
		dom.SetAttribute("class", func() string {
			if singleColor {
				return "mdl-spinner mdl-spinner--single-color mdl-js-spinner is-active"
			} else {
				return "mdl-spinner mdl-js-spinner is-active"
			}
		}()),
	)
}

func Logo(content ...view.View) view.View {
	return classifiedDiv("mdl-logo", content...)
}

func Title(title string) view.View {
	return dom.Attributed(
		dom.Span(dom.Text(title)),
		dom.SetAttribute("class", "mdl-layout-title"),
	)
}

func Spacer() view.View {
	return classifiedDiv("mdl-layout-spacer")
}

func classifiedDiv(class string, content ...view.View) view.View {
	return dom.Attributed(
		dom.Div(content...),
		dom.SetAttribute("class", class),
	)
}

func MaterialIcon(text string) view.View {
	return dom.Attributed(
		dom.I(dom.Text(text)),
		dom.Class("material-icons"),
	)
}

func ChipText(text string) view.View {
	return dom.Attributed(
		dom.Span(dom.Text("Basic Chip")),
		dom.Class("mdl-chip__text"),
	)
}

func ChipAction(icon string) view.View {
	return dom.Attributed(
		dom.Button(
			MaterialIcon(icon),
		),
		dom.SetAttribute("class", "mdl-chip__action"),
		dom.SetAttribute("type", "button"),
	)
}

func ChipPlaceholder(text string) view.View {
	return dom.Attributed(
		dom.Span(dom.Text(text)),
		dom.SetAttribute("class", "mdl-chip__contact mdl-color--teal mdl-color-text--white"),
	)
}

func ChipImage(src string) view.View {
	return dom.Attributed(
		dom.Img(),
		dom.SetAttribute("class", "mdl-chip__contact"),
		dom.SetAttribute("src", src),
	)
}

func Chip(contact, deletable bool, content ...view.View) view.View {
	return dom.Attributed(
		dom.Span(content...),
		dom.UpdateAttribute("class", func(s string) string {
			s += "mdl-chip"

			if contact {
				s += " mdl-chip--contact"
			}

			if deletable {
				s += " mdl-chip--deletable"
			}

			return s
		}),
	)
}
