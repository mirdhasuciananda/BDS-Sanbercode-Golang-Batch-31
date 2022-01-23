package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"quiz3/models"
	"quiz3/query"
	"quiz3/utils"

	"github.com/julienschmidt/httprouter"
)

func getThickness(total_page uint) string {
	switch {
	case total_page <= 100:
		return "tipis"
	case total_page >= 101 && total_page < 201:
		return "sedang"
	default:
		return "tebal"
	}
}

func GetAllBook(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := query.GetAllBook(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, nilai, http.StatusOK)
}

func GetBookByCategory(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var idCat = ps.ByName("id")

	book, err := query.GetBookByCategory(ctx, idCat)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, book, http.StatusOK)
}

// func FilterBookCategory(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	fmt.Println("Menjalankan FilterBookCategory")
// 	ctx, cancel := context.WithCancel(context.Background())

// 	defer cancel()

// 	var title = ps.ByName("title")
// 	var minYear = ps.ByName("minYear")
// 	var maxYear = ps.ByName("maxYear")
// 	var minPage = ps.ByName("minPage")
// 	var maxPage = ps.ByName("maxPage")
// 	var sortByTitle = ps.ByName("sortByTitle")

// 	fmt.Println("title = " + title)
// 	fmt.Println("minYear = " + minYear)

// 	book, err := query.FilterBookCategory(ctx, title, minYear, maxYear, minPage, maxPage, sortByTitle)

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	utils.ResponseJSON(rw, book, http.StatusOK)
// }

func PostBook(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	flag_out := 0

	// Pengecekan URL
	_, err := url.ParseRequestURI(book.ImageUrl)
	if err != nil {
		http.Error(rw, "image_url tidak valid", http.StatusBadRequest)
		flag_out += 1
	}

	// Pengecekan release_year
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		http.Error(rw, "release_year minimal 1980 dan maksimal 2021 ", http.StatusBadRequest)
		flag_out += 1
	}

	if flag_out > 0 {
		return
	}

	// Konversi total_page ke thickness
	book.Thickness = getThickness(book.TotalPage)

	if err := query.InsertBook(ctx, book); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(rw, res, http.StatusCreated)

}

func UpdateBook(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idBook = ps.ByName("id")

	flag_out := 0

	// Pengecekan URL
	_, err := url.ParseRequestURI(book.ImageUrl)
	if err != nil {
		http.Error(rw, "image_url tidak valid", http.StatusBadRequest)
		flag_out += 1
	}

	// Pengecekan release_year
	if book.ReleaseYear < 1980 || book.ReleaseYear > 2021 {
		http.Error(rw, "release_year minimal 1980 dan maksimal 2021 ", http.StatusBadRequest)
		flag_out += 1
	}

	if flag_out > 0 {
		return
	}

	// Konversi total_page ke thickness
	book.Thickness = getThickness(book.TotalPage)

	if err := query.UpdateBook(ctx, book, idBook); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(rw, res, http.StatusCreated)

}

func DeleteBook(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var idBook = ps.ByName("id")

	if err := query.DeleteBook(ctx, idBook); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(rw, res, http.StatusCreated)

}
