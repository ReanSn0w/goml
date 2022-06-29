package dom

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/ReanSn0w/gew/pkg/view"
)

func NewBuilder() *Builder {
	return &Builder{}
}

type Builder struct{}

func (b *Builder) Bytes(view view.View) []byte {
	buffer := new(bytes.Buffer)
	b.build(context.Background(), buffer, view)
	return buffer.Bytes()
}

func (b *Builder) Write(wr io.Writer, view view.View) {
	b.build(context.Background(), wr, view)
}

func (b *Builder) build(ext context.Context, wr io.Writer, node view.View) {
	view.Build(node, ext, func(i interface{}, ctx context.Context) {
		switch v := i.(type) {
		case tagValue:
			attributes, _ := getAttributeStorageFromCtx(ctx)

			wr.Write([]byte(fmt.Sprintf("<%v%v>", v.name, attributes.build())))
			b.build(ext, wr, v.content)
			wr.Write([]byte(fmt.Sprintf("</%v>", v.name)))
		case textValue:
			wr.Write([]byte(v.content))
		}
	})
}
