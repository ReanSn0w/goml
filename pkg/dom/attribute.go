package dom

import (
	"bytes"
	"context"
	"fmt"

	"github.com/ReanSn0w/gew/pkg/view"
)

var (
	attributeStorageCtxKey = &ctxKey{}
)

type ctxKey struct{}

type AttributeStorage map[string]string
type AttributeEditor func(AttributeStorage)

func (as AttributeStorage) add(key string, value string) {
	as[key] = value
}

func (as AttributeStorage) update(key string, update func(string) string) {
	val := as[key]
	as[key] = update(val)
}

func (as AttributeStorage) remove(key string) {
	delete(as, key)
}

func (as AttributeStorage) build() string {
	buffer := new(bytes.Buffer)

	for k, v := range as {
		if v == "" {
			buffer.WriteString(fmt.Sprintf(" %v", k))
		} else {
			buffer.WriteString(fmt.Sprintf(" %v=\"%v\"", k, v))
		}
	}

	return buffer.String()
}

func getAttributeStorageFromCtx(ctx context.Context) (AttributeStorage, context.Context) {
	value := ctx.Value(attributeStorageCtxKey)
	storage, even := value.(AttributeStorage)

	if !even {
		newStorage := AttributeStorage{}
		newCtx := context.WithValue(ctx, attributeStorageCtxKey, newStorage)
		return newStorage, newCtx
	}

	return storage, ctx
}

func Attributed(tag view.View, editors ...AttributeEditor) view.View {
	return view.CustomContextPreparation(tag, func(ctx context.Context) context.Context {
		storage, ctx := getAttributeStorageFromCtx(ctx)

		for _, editor := range editors {
			editor(storage)
		}

		return ctx
	})
}

func SetAttribute(name string, value string) AttributeEditor {
	return func(as AttributeStorage) {
		as.add(name, value)
	}
}

func UnsetAttribute(name string) AttributeEditor {
	return func(as AttributeStorage) {
		as.remove(name)
	}
}

func UpdateAttribute(name string, update func(string) string) AttributeEditor {
	return func(as AttributeStorage) {
		as.update(name, update)
	}
}

func AccessKey(value string) AttributeEditor {
	return SetAttribute("accesskey", value)
}

func Class(value string) AttributeEditor {
	return SetAttribute("class", value)
}

func ContentEditable(value string) AttributeEditor {
	return SetAttribute("contenteditable", value)
}

func ContextMenu(value string) AttributeEditor {
	return SetAttribute("contextmenu", value)
}

func Dir(value string) AttributeEditor {
	return SetAttribute("dir", value)
}

func Hidden(value string) AttributeEditor {
	return SetAttribute("hidden", value)
}

func Id(value string) AttributeEditor {
	return SetAttribute("id", value)
}

func Lang(value string) AttributeEditor {
	return SetAttribute("lang", value)
}

func Name(value string) AttributeEditor {
	return SetAttribute("name", value)
}

func Spellcheck(value string) AttributeEditor {
	return SetAttribute("spellcheck", value)
}

func TabIndex(value string) AttributeEditor {
	return SetAttribute("tabindex", value)
}

func AttrTitle(value string) AttributeEditor {
	return SetAttribute("title", value)
}
