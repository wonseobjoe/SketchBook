package web

import ()

func loadIndexContents() (*indexPage, error) {
	return &indexPage{Title: "SketchBook", Body: []byte("First GO Language Project")}, nil
}

type indexPage struct {
	Title string
	Body  []byte
}
