package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

func Footer(content ...view.View) view.View {
	return dom.Attributed(
		dom.Footer(content...),
		dom.SetAttribute("class", "mdl-mega-footer"),
	)
}

func LeftFooterSection(content ...view.View) view.View {
	return classifiedDiv("mdl-mega-footer__left-section", content...)
}

func RightFooterSection(content ...view.View) view.View {
	return classifiedDiv("mdl-mega-footer__right-section", content...)
}

func DropdownFooterSection(title string, elements uint, builder func(int) view.View) view.View {
	return classifiedDiv(
		"mdl-mega-footer__drop-down-section",
		dom.Attributed(
			dom.H1(dom.Text(title)),
			dom.SetAttribute("class", "mdl-mega-footer__heading"),
		),
		dom.Attributed(
			dom.Ul(
				view.For(elements, builder),
			),
			dom.SetAttribute("class", "mdl-mega-footer__link-list"),
		),
	)
}

func FooterList(content ...view.View) view.View {
	return dom.Attributed(
		dom.Ul(
			view.For(uint(len(content)), func(i int) view.View {
				return dom.Li(content[i])
			}),
		),
		dom.SetAttribute("class", "mdl-mega-footer__link-list"),
	)
}

func SmallFooter(content ...view.View) view.View {
	return dom.Attributed(
		dom.Footer(
			dom.Attributed(
				dom.Div(content...),
				dom.SetAttribute("class", "mdl-mini-footer__left-section"),
			),
		),
		dom.SetAttribute("class", "mdl-mini-footer"),
	)
}

func SmallFooterList(content ...view.View) view.View {
	return dom.Attributed(
		dom.Ul(
			view.For(uint(len(content)), func(i int) view.View {
				return dom.Li(content[i])
			}),
		),
		dom.SetAttribute("class", "mdl-mega-footer__link-list"),
	)
}
