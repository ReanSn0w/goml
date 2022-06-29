package mdl

import (
	"fmt"

	"github.com/ReanSn0w/gew/pkg/view"
	"github.com/ReanSn0w/goml/pkg/dom"
)

type CellPosition string

const (
	CellPositionScretch CellPosition = "mdl-cell--stretch"
	CellPositionTop     CellPosition = "mdl-cell--top"
	CellPositionMiddle  CellPosition = "mdl-cell--middle"
	CellPositionBottom  CellPosition = "mdl-cell--bottom"
)

type CellPreferences struct {
	Size        uint // Default: 1, Общий размер ячейки
	MobileSize  uint // Размер ячейки на телефоне
	TabletSize  uint // Размер ячейки на планшете
	DesktopSize uint // Размер ячейки на компьютере

	Offset        uint // общий отступ для ячейки
	MobileOffset  uint // отступ на телефоне
	TabletOffset  uint // отступ на планшете
	DesktopOffset uint // отступ на компьютере

	Order        uint // позиция ячейки
	MobileOrder  uint // позиция элемента на мобильном телефоне
	TabletOrder  uint // позиция элемента на планшете
	DesktopOrder uint // позиция элемента на компьютере

	MobileHide  bool // сокрытие ячейки на телефоне
	TabletHide  bool // сокрытие ячейки на планшете
	DesktopHide bool // сокрытие элемента на десктопе

	Position CellPosition // Default: - scretch
}

func (cp *CellPreferences) ClassValue() string {
	class := ""

	if cp.Size == 0 {
		class += fmt.Sprintf(" mdl-cell--%v-col", 1)
	} else {
		class += fmt.Sprintf(" mdl-cell--%v-col", cp.Size)
	}

	if cp.MobileSize != 0 {
		class += fmt.Sprintf(" mdl-cell--%v-col-phone", cp.MobileSize)
	}

	if cp.TabletSize != 0 {
		class += fmt.Sprintf(" mdl-cell--%v-col-tablet", cp.TabletSize)
	}

	if cp.DesktopSize != 0 {
		class += fmt.Sprintf(" mdl-cell--%v-col-desktop", cp.DesktopSize)
	}

	if cp.Offset != 0 {
		class += fmt.Sprintf(" mdl-cell--%v-offset", cp.Offset)
	}

	if cp.MobileOffset != 0 {
		class += fmt.Sprintf(" mdl-cell--%v-offset-phone", cp.MobileOffset)
	}

	if cp.TabletOffset != 0 {
		class += fmt.Sprintf(" mdl-cell--%v-offset-tablet", cp.TabletOffset)
	}

	if cp.DesktopOffset != 0 {
		class += fmt.Sprintf(" mdl-cell--%v-offset-desktop", cp.DesktopOffset)
	}

	if cp.Order != 0 {
		class += fmt.Sprintf(" mdl-cell--order-%v", cp.Order)
	}

	if cp.MobileOrder != 0 {
		class += fmt.Sprintf(" mdl-cell--order-%v-phone", cp.MobileOrder)
	}

	if cp.TabletOrder != 0 {
		class += fmt.Sprintf(" mdl-cell--order-%v-tablet", cp.TabletOrder)
	}

	if cp.DesktopOrder != 0 {
		class += fmt.Sprintf(" mdl-cell--order-%v-desktop", cp.DesktopOrder)
	}

	if cp.MobileHide {
		class += " mdl-cell--hide-phone"
	}

	if cp.TabletHide {
		class += " mdl-cell--hide-tablet"
	}

	if cp.DesktopHide {
		class += " mdl-cell--hide-desktop"
	}

	if cp.Position != CellPositionScretch {
		class += " "
		class += string(cp.Position)
	}

	return class
}

// Создание сетки для размещения элементов
func Grid(spacing bool, content ...view.View) view.View {
	return dom.Attributed(
		dom.Div(content...),
		dom.Class(func() string {
			class := "mdl-grid"

			if spacing {
				class += "mdl-grid--no-spacing"
			}

			return class
		}()),
	)
}

// Создание ячейки для сетки
func Cell(preferences CellPreferences, content ...view.View) view.View {
	return dom.Attributed(
		dom.Div(content...),
		dom.Class(preferences.ClassValue()),
	)
}
