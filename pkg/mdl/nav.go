package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

type NavigationPreferences struct {
	Transparent bool
	Waterfall   bool
}

func (np *NavigationPreferences) class() string {
	class := "mdl-layout__header"

	if np.Waterfall {
		class += " mdl-layout__header--waterfall"
	}

	if np.Transparent {
		class += " mdl-layout__header--transparent"
	}

	return class
}

func NavigationHeader(preferences NavigationPreferences, rows ...view.View) view.View {
	return dom.Attributed(
		dom.Header(rows...),
		dom.SetAttribute("class", preferences.class()),
	)
}

func WatefallNavigation(topRow, bottomRow view.View) view.View {
	return dom.Attributed(
		dom.Header(topRow, bottomRow),
		dom.SetAttribute("class", "mdl-layout__header mdl-layout__header--transparent"),
	)
}

func NavigationRow(content ...view.View) view.View {
	return classifiedDiv("mdl-layout__header-row", content...)
}

// Меню навигации расположенное справа от основного контента страницы
func SideNavigation(title string, elements ...view.View) view.View {
	return dom.Attributed(
		dom.Div(elements...),
		dom.SetAttribute("class", "mdl-layout__drawer"),
	)
}

func NavigationList(content ...view.View) view.View {
	return dom.Attributed(
		dom.Nav(content...),
		dom.SetAttribute("class", "mdl-navigation"),
	)
}

func NavigationLink(title, link string) view.View {
	return dom.Attributed(
		dom.A(dom.Text(title)),
		dom.SetAttribute("class", "mdl-navigation__link"),
		dom.SetAttribute("href", link),
	)
}
