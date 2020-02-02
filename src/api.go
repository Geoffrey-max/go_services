package main
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/gorilla/mux"
)

import "models"
import "dao"

var dao = RatingsDAO{}

func AllRatingsEndPoint(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
func FindRatingEndpoint(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
func CreateRatingEndPoint(w http.ResponseWriter, r * http.Request) {
	defer r.Body.Close()
	var rating Rating
	json.NewDecoder(r.Body).Decode(&rating)
	rating.id = bson.NewObjectId()
	dao.Insert(rating)
	respondWithJson(w, http.StatusCreated, rating)
}
func DeleteRatingEndPoint(w http.ResponseWriter, r * http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ratings", AllRatingsEndPoint).Methods("GET")
	r.HandleFunc("/rating", CreateRatingEndPoint).Methods("POST")
	r.HandleFunc("/rating", DeleteRatingEndPoint).Methods("DELETE")
	r.HandleFunc("/ratings/{id}", FindRatingEndpoint).Methods("GET")
	err := http.ListenAndServe(":8002", r)
	if  err != nil {
		log.Fatal(err)
	}
}