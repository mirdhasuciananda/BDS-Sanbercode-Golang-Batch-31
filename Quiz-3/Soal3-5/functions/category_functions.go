package functions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"quiz3/models"
	"quiz3/query"
	"quiz3/utils"

	"github.com/julienschmidt/httprouter"
)

func GetAllCategory(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	nilai, err := query.GetAllCategory(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, nilai, http.StatusOK)
}

func PostCategory(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var cat models.Category

	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if err := query.InsertCategory(ctx, cat); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(rw, res, http.StatusCreated)
}

func UpdateCategory(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var cat models.Category

	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idCat = ps.ByName("id")

	if err := query.UpdateCategory(ctx, cat, idCat); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(rw, res, http.StatusCreated)

}

func DeleteCategory(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var idCat = ps.ByName("id")

	if err := query.DeleteCategory(ctx, idCat); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Succesfully",
	}

	utils.ResponseJSON(rw, res, http.StatusCreated)

}
