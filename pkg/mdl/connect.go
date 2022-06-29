package mdl

import (
	"fmt"

	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

const (
	ColorCyan       Color = "cyan"
	ColorTeal       Color = "teal"
	ColorGreen      Color = "green"
	ColorLightGreen Color = "light_green"
	ColorLime       Color = "lime"
	ColorYellow     Color = "yellow"
	ColorAmber      Color = "amber"
	ColorOrange     Color = "orange"
	ColorBrown      Color = "brown"
	ColorBlueGray   Color = "blue_gray"
	ColorGray       Color = "gray"
	ColorDeepOrange Color = "deep_orange"
	ColorRed        Color = "red"
	ColorPink       Color = "pink"
	ColorPurple     Color = "purple"
	ColorDeepPurple Color = "deep_purple"
	ColorIndigo     Color = "indigo"
	ColorBlue       Color = "blue"
	ColorLightBlue  Color = "light_blue"
)

type Color string

func ConnectIcons() view.View {
	return dom.Attributed(
		dom.Link(),
		dom.SetAttribute("rel", "stylesheet"),
		dom.SetAttribute("href", "https://fonts.googleapis.com/icon?family=Material+Icons"),
	)
}

func ConnectStyles(primary, accent Color) view.View {
	return dom.Attributed(
		dom.Link(),
		dom.SetAttribute("rel", "stylesheet"),
		dom.SetAttribute("href", fmt.Sprintf("https://code.getmdl.io/1.3.0/material.%s-%s.min.css", primary, accent)),
	)
}

func ConnectScripts() view.View {
	return dom.Attributed(
		dom.Script(),
		dom.SetAttribute("defer", ""),
		dom.SetAttribute("src", "https://code.getmdl.io/1.3.0/material.min.js"),
	)
}
