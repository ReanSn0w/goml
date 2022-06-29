package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

// TODO: - Необходимо доделать вариации
func List(elements ...view.View) view.View {
	return dom.Attributed(
		dom.Ul(elements...),
		dom.SetAttribute("class", "demo-list-icon mdl-list"),
	)
}

func ListItem(lines uint, content ...view.View) view.View {
	return dom.Attributed(
		dom.Li(content...),
		dom.SetAttribute("class", func() string {
			switch lines {
			case 2:
				return "mdl-list__item mdl-list__item--two-line"
			case 3:
				return "mdl-list__item mdl-list__item--three-line"
			default:
				return "mdl-list__item"
			}
		}()),
	)
}

func ListItemTitle(text string) view.View {
	return dom.Span(dom.Text(text))
}

func ListItemTextBody(text string) view.View {
	return dom.Attributed(
		dom.Span(dom.Text(text)),
		dom.SetAttribute("class", "mdl-list__item-text-body"),
	)
}

func ListItemSubtitle(text string) view.View {
	return dom.Attributed(
		dom.Span(dom.Text(text)),
		dom.SetAttribute("class", "mdl-list__item-sub-title"),
	)
}

func ListItemAction(icon string, link string) view.View {
	return dom.Attributed(
		dom.Span(
			dom.Attributed(
				dom.A(
					dom.Attributed(
						dom.I(dom.Text(icon)),
						dom.SetAttribute("class", "material-icons"),
					),
				),
				dom.SetAttribute("class", "mdl-list__item-secondary-action"),
				dom.SetAttribute("href", link),
			),
		),
		dom.SetAttribute("class", "mdl-list__item-secondary-content"),
	)
}
