package web

import ()

func loadIndexContents() (*indexPage, error) {
	return &indexPage{Title: "CafeLivro", Body: []byte("First GO Language Project")}, nil
}

type indexPage struct {
	Title string
	Body  []byte
}
