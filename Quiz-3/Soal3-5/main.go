package main

import (
	"fmt"
	"log"
	"net/http"
	"quiz3/functions"

	"github.com/julienschmidt/httprouter"
)

// Fungi Log yang berguna sebagai middleware
func Auth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "password" {
			h(w, r, ps)
			return
		}
		if uname == "editor" && pwd == "secret" {
			h(w, r, ps)
			return
		}
		if uname == "trainer" && pwd == "rahasia" {
			h(w, r, ps)
			return
		}

		w.Write([]byte("Username atau Password tidak sesuai"))
		// return
	}
}

func main() {
	router := httprouter.New()

	router.GET("/categories", functions.GetAllCategory)
	router.POST("/categories", Auth(functions.PostCategory))
	router.PUT("/categories/:id", Auth(functions.UpdateCategory))
	router.DELETE("/categories/:id", Auth(functions.DeleteCategory))
	router.GET("/categories/:id/books", functions.GetBookByCategory)

	router.GET("/books", functions.GetAllBook)
	router.POST("/books", Auth(functions.PostBook))
	router.PUT("/books/:id", Auth(functions.UpdateBook))
	router.DELETE("/books/:id", Auth(functions.DeleteBook))

	fmt.Println("Server Running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
