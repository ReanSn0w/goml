package mdl

import (
	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

type TableItemValue struct {
	Value   string
	Numeric bool
}

func (item *TableItemValue) View() view.View {
	return dom.Attributed(
		dom.Th(dom.Text(item.Value)),
		func(as dom.AttributeStorage) {
			if !item.Numeric {
				as["class"] = "mdl-data-table__cell--non-numeric"
			}
		},
	)
}

func Table(selectable bool, header view.View, content view.View) view.View {
	class := "mdl-data-table mdl-js-data-table"

	if selectable {
		class += " mdl-data-table--selectable"
	}

	return dom.Attributed(
		dom.Table(
			view.If(header != nil, dom.Thead(header)),
			dom.Tbody(content),
		),
		dom.Class(class),
	)
}

func TableItem(items []TableItemValue) view.View {
	return dom.Tr(
		view.For(uint(len(items)), func(i int) view.View {
			return items[i].View()
		}),
	)
}
