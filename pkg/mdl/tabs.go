package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

// <div class="mdl-layout__tab-bar mdl-js-ripple-effect">
//       <a href="#scroll-tab-1" class="mdl-layout__tab is-active">Tab 1</a>
//       <a href="#scroll-tab-2" class="mdl-layout__tab">Tab 2</a>
//       <a href="#scroll-tab-3" class="mdl-layout__tab">Tab 3</a>
//       <a href="#scroll-tab-4" class="mdl-layout__tab">Tab 4</a>
//       <a href="#scroll-tab-5" class="mdl-layout__tab">Tab 5</a>
//       <a href="#scroll-tab-6" class="mdl-layout__tab">Tab 6</a>
//     </div>

// <div class="mdl-layout__tab-bar mdl-js-ripple-effect">
//       <a href="#fixed-tab-1" class="mdl-layout__tab is-active">Tab 1</a>
//       <a href="#fixed-tab-2" class="mdl-layout__tab">Tab 2</a>
//       <a href="#fixed-tab-3" class="mdl-layout__tab">Tab 3</a>
//     </div>

func TabView(content ...view.View) view.View {
	return classifiedDiv("mdl-tabs mdl-js-tabs mdl-js-ripple-effect", content...)
}

func TabBar(content ...view.View) view.View {
	return classifiedDiv("mdl-tabs__tab-bar", content...)
}

func TabItem(link string, isActive bool, content ...view.View) view.View {
	return dom.Attributed(
		dom.A(content...),
		dom.SetAttribute("class", func() string {
			if isActive {
				return "mdl-tabs__tab is-active"
			} else {
				return "mdl-tabs__tab"
			}
		}()),
		dom.SetAttribute("href", link),
	)
}

func TabPanel(id string, isActive bool, content ...view.View) view.View {
	return classifiedDiv(
		func() string {
			if isActive {
				return "mdl-tabs__panel is-active"
			} else {
				return "mdl-tabs__panel"
			}
		}(),
		content...,
	)
}
