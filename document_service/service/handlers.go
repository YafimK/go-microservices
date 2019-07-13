package service

import (
	"encoding/json"
	"github.com/Yafimk/go-microservices/common"
	"github.com/Yafimk/go-microservices/document_service/model"
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
		result, err := sdh.client.QueryValue(bucket, id, func(value []byte) (interface{}, error) {
			document := model.Document{}
			json.Unmarshal(value, &document)
			return document, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		document := result.(model.Document)
		data, _ := json.Marshal(document)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Content-Length", strconv.Itoa(len(data)))
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	}

}
