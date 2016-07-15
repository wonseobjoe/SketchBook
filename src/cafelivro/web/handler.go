<<<<<<< HEAD
package web

import (
	"net/http"
)

func IndexHandler(writer http.ResponseWriter, req *http.Request) {
	indexContents, _ := loadIndexContents()
	renderTemplate(writer, "index", indexContents)
}
=======
package web

import (
	"net/http"
)

func IndexHandler(writer http.ResponseWriter, req *http.Request) {
	indexContents, _ := loadIndexContents()
	renderTemplate(writer, "index", indexContents)
}
>>>>>>> 2b3e8bf30f00c1359cdb82cc02321e53c21373f3
