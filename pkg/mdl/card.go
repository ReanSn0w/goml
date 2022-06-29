package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

func Card(elements ...view.View) view.View {
	return dom.Attributed(
		dom.Div(
			view.For(uint(len(elements)), func(i int) view.View {
				return elements[i]
			}),
		),
		dom.Class("mdl-card mdl-shadow--2dp"),
	)
}

// Тайтл для карточки в дизайне MDL
func CardTitle(expand bool, value string) view.View {
	return dom.Attributed(
		dom.Div(
			dom.Attributed(
				dom.Div(
					dom.Text(value),
				),
				dom.Class(func() string {
					if expand {
						return "mdl-card__title-text mdl-card--expand"
					} else {
						return "mdl-card__title-text"
					}
				}()),
			),
		),
		dom.Class("mdl-card__title"),
	)
}

// Основной текст в карточке
func CardSupportingText(text string) view.View {
	return dom.Attributed(
		dom.Div(
			dom.Text(text),
		),
		dom.Class("mdl-card__supporting-text"),
	)
}

// Действия отображаемые на карточке
func CardActions(content ...view.View) view.View {
	return dom.Attributed(
		dom.Div(content...),
		dom.Class("mdl-card__actions mdl-card--border"),
	)
}

// Меню отображаемое на карточке
func CardMenu(content ...view.View) view.View {
	return dom.Attributed(
		dom.Div(content...),
		dom.Class("mdl-card__menu"),
	)
}
