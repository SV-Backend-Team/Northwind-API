package suppliers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GET
func GetSuppliers(res http.ResponseWriter, req *http.Request) {
	log.Println("select all")
	var supplierModel SupplierModel
	suppliers, err := supplierModel.GetSuppliers()
	if err != nil {
		log.Panicln(err)
		respondWithError(res, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, suppliers)
	}
}

//GET{ID}
func GetSuppliersByID(res http.ResponseWriter, req *http.Request) {
	log.Println("select by id")
	vars := mux.Vars(req)
	id := vars["id"]
	supplierid, err := strconv.Atoi(id)
	if err != nil {
		log.Panicln(err)
		respondWithError(res, http.StatusBadRequest, err.Error())
	}
	var supplierModel SupplierModel
	suppliers, err := supplierModel.GetSupplierByID(supplierid)
	if err != nil {
		log.Panicln(err)
		respondWithError(res, http.StatusInternalServerError, err.Error())
	} else {
		respondWithJson(res, http.StatusOK, suppliers)
	}
}

//POST
func CreateSuppliers(res http.ResponseWriter, req *http.Request) {
	log.Println("create")
	var supplier Supplier
	err := json.NewDecoder(req.Body).Decode(&supplier)
	if err != nil {
		log.Panicln(err)
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var supplierModel SupplierModel
		err2 := supplierModel.CreateSuppliers(&supplier)
		if err2 != nil {
			log.Panicln(err)
			respondWithError(res, http.StatusInternalServerError, err2.Error())
		} else {
			respondWithJson(res, http.StatusOK, supplier)
		}
	}
}

//Update
func UpdateSupplier(res http.ResponseWriter, req *http.Request) {
	log.Println("update")
	var supplier Supplier
	err := json.NewDecoder(req.Body).Decode(&supplier)
	if err != nil {
		log.Panicln(err)
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		var supplierModel SupplierModel
		err2 := supplierModel.UpdateSupplier(&supplier)
		if err2 != nil {
			log.Panicln(err)
			respondWithError(res, http.StatusInternalServerError, err2.Error())
		} else {
			respondWithJson(res, http.StatusOK, supplier)
		}
	}
}

//Delete
func DeleteSupplierByID(res http.ResponseWriter, req *http.Request) {
	log.Println("delete")
	vars := mux.Vars(req)
	id := vars["id"]
	supplierid, _ := strconv.Atoi(id)
	var supplierModel SupplierModel
	suppliers, err := supplierModel.GetSupplierByID(supplierid)
	if err != nil {
		log.Panicln(err)
		respondWithError(res, http.StatusBadRequest, err.Error())
	} else {
		err2 := supplierModel.DeleteSupplier(suppliers)
		if err2 != nil {
			log.Panicln(err)
			respondWithError(res, http.StatusInternalServerError, err.Error())
		} else {
			respondWithJson(res, http.StatusOK, suppliers)
		}
	}
}

//RespondWith
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "applicaction/json")
	w.WriteHeader(code)
	w.Write(response)
}
