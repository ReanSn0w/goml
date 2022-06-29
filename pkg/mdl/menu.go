package mdl

import (
	"fmt"

	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

const (
	MenuPositionTopLeft     = "mdl-menu--top-left"
	MenuPositionTopRight    = "mdl-menu--top-right"
	MenuPositionBottomLeft  = "mdl-menu--bottom-left"
	MenuPositionBotoomRight = "mdl-menu--bottom-right"
)

type MenuPosition string

func Menu(forID string, position MenuPosition, content ...view.View) view.View {
	return dom.Attributed(
		dom.Ul(content...),
		dom.SetAttribute("for", forID),
		dom.Class(fmt.Sprintf("mdl-menu %v mdl-js-menu mdl-js-ripple-effect", position)),
	)
}

func MenuItem(divider, disabled bool, content ...view.View) view.View {
	return dom.Attributed(
		dom.Li(content...),
		dom.Class(func() string {
			class := "mdl-menu__item"
			if disabled {
				class += " mdl-menu__item--full-bleed-divider"
			}
			return class
		}()),
		func(as dom.AttributeStorage) {
			if disabled {
				as["disabled"] = ""
			}
		},
	)
}
