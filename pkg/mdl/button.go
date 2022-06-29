package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

const (
	ButtonFab     ButtonStyle = " mdl-button--fab"
	ButtonMiniFab ButtonStyle = " mdl-button--fab mdl-button--mini-fab"
	ButtonRaised  ButtonStyle = " mdl-button--raised"
	ButtonIcon    ButtonStyle = " mdl-button--icon"

	ButtonColored ButtonColor = " mdl-button--colored"
	ButtonPrimary ButtonColor = " mdl-button--primary"
	ButtonAccent  ButtonColor = " mdl-button--accent"
)

type ButtonStyle string
type ButtonColor string

type ButtonPreferences struct {
	Style    ButtonStyle
	Color    ButtonColor
	Raised   bool
	Ripple   bool
	Disabled bool
	Id       string // Идентификатор таргета, для случаев когда кнопка открывает меню
}

func (bp *ButtonPreferences) attributes(as dom.AttributeStorage) {
	{
		class := as["class"]
		class += "mdl-button mdl-js-button"
		class += string(bp.Style)
		class += string(bp.Color)

		if bp.Raised {
			class += " mdl-js-ripple-effect"
		}

		if bp.Ripple {
			class += " mdl-js-ripple-effect"
		}

		as["class"] = class
	}

	{
		if bp.Disabled {
			as["disabled"] = ""
		}
	}

	{
		if bp.Id != "" {
			as["id"] = bp.Id
		}
	}
}

// Button - метод для создания кнопки в стиле material design
func Button(pref ButtonPreferences, content ...view.View) view.View {
	return dom.Attributed(
		dom.Button(content...),
		pref.attributes,
	)
}

func ChipButton(content ...view.View) view.View {
	return dom.Attributed(
		dom.Button(content...),
		dom.SetAttribute("class", "mdl-chip"),
		dom.SetAttribute("type", "button"),
	)
}
