package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"log"
	"net/http"
)

type hello struct {
	app.Compo
}

func (h *hello) Render() app.UI {
	return app.H1().Text("Hello World!")
}

func main() {
	// パスに対してどのコンポーネントをレンダリングすれば良いのか定義する
	app.Route("/", &hello{})
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
