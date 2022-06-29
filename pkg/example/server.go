package main

import (
	"net/http"

	"github.com/ReanSn0w/goml/pkg/dom"
	"github.com/ReanSn0w/goml/pkg/mdl"
)

func main() {
	hf := http.HandlerFunc(root)
	http.ListenAndServe(":8080", hf)
}

func root(w http.ResponseWriter, r *http.Request) {
	builder := dom.NewBuilder()

	builder.Write(w, dom.Html(
		dom.Head(
			dom.Title(
				dom.Text("Тестовая страница"),
			),
			mdl.ConnectStyles(mdl.ColorIndigo, mdl.ColorPink),
			mdl.ConnectScripts(),
		),
		dom.Body(
			dom.H1(dom.Text("Тестовая страница")),
			dom.P(dom.Text("Привет, мир!")),
			mdl.Button(
				mdl.ButtonPreferences{
					Raised: true,
				},
				dom.Text("Привет, мир!"),
			),
			dom.Attributed(
				mdl.Card(
					mdl.CardTitle(true, "Hello, world"),
					mdl.CardSupportingText("some supported text"),
					mdl.CardActions(
						mdl.Button(
							mdl.ButtonPreferences{
								Raised: true,
								Color:  mdl.ButtonAccent,
								Style:  mdl.ButtonFab,
							},
							dom.Text("Hello, world!"),
						),
					),
				),
				dom.SetAttribute("style", "height: 300px; width: 300px; padding: 10px"),
			),
		),
	))
}
