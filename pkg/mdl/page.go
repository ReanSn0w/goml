package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

type PagePreferences struct {
	FixedHeader bool
	FixedDrawer bool
}

func (pp *PagePreferences) Class() string {
	class := "mdl-layout mdl-js-layout"

	if pp.FixedDrawer {
		class += " mdl-layout--fixed-drawer"
	}

	if pp.FixedHeader {
		class += " mdl-layout--fixed-header"
	}

	return class
}

func Page(pref PagePreferences, content ...view.View) view.View {
	return classifiedDiv(
		pref.Class(),
		content...,
	)
}

func Main(content ...view.View) view.View {
	return dom.Attributed(
		dom.Main(content...),
		dom.SetAttribute("class", "mdl-layout__content"),
	)
}

func Content(content ...view.View) view.View {
	return classifiedDiv("page-content", content...)
}

func ContentSection(id string, isActive bool, content ...view.View) view.View {
	return dom.Attributed(
		dom.Section(Content(content...)),
		dom.SetAttribute("class", func() string {
			if isActive {
				return "mdl-layout__tab-panel is-active"
			} else {
				return "mdl-layout__tab-panel"
			}
		}()),
		dom.SetAttribute("id", id),
	)
}
