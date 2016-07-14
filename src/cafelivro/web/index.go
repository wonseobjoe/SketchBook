<<<<<<< HEAD
package web

import ()

func loadIndexContents() (*indexPage, error) {
	return &indexPage{Title: "CafeLivro", Body: []byte("First GO Language Project")}, nil
}

type indexPage struct {
	Title string
	Body  []byte
}
=======
package web

import ()

func loadIndexContents() (*indexPage, error) {
	return &indexPage{Title: "CafeLivro", Body: []byte("First GO Language Project")}, nil
}

type indexPage struct {
	Title string
	Body  []byte
}
>>>>>>> 2b3e8bf30f00c1359cdb82cc02321e53c21373f3
