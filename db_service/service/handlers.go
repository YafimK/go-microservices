package service

import (
	"encoding/json"
	"fmt"
	"github.com/Yafimk/go-microservices/common"
	"github.com/Yafimk/go-microservices/common/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DbHandler struct {
	client *common.DbDriver
}

func (sdh DbHandler) GetDocument(bucket string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var id = mux.Vars(r)["Id"]
		rawResult, err := sdh.client.QueryValue(bucket, id, func(value []byte) (interface{}, error) {
			document := model.Document{}
			json.Unmarshal(value, &document)
			return document, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, fmt.Sprintf("Didn't find the document with id - %v\n", id))
			return
		}
		document := rawResult.(model.Document)
		data, _ := json.Marshal(document)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}
}

func (sdh DbHandler) AddDocument(bucket string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var document model.Document
		err := decoder.Decode(&document)
		if err != nil {
			http.Error(w, fmt.Sprintf("error during addition of value to db: %v", err), http.StatusInternalServerError)
			return
		}
		jsonBytes, _ := json.Marshal(document)

		if err := sdh.client.AddValue(bucket, []byte(document.Id), jsonBytes); err != nil {
			http.Error(w, fmt.Sprintf("error during addition of value to db: %v", err), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte{})
	}
}

func (sdh DbHandler) CheckDocumentServiceHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result := sdh.client.IsAlive()
		isBucketExists := sdh.client.IsBucketExists("documents")
		var status string
		if result && isBucketExists {
			status = "up"
		} else {
			status = "down"
		}
		data, _ := json.Marshal(model.HealthCheck{Status: status})
		writeJsonResponse(w, http.StatusOK, data)
	}
}

func writeJsonResponse(w http.ResponseWriter, status int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.WriteHeader(status)
	w.Write(data)
}
