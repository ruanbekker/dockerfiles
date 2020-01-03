package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/xid"
)

type Beer struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Category  string `json:"category,omitempty"`
  Brewer    string `json:"brewer,omitempty"`
	Country  `json:"country,omitempty"`
}

type Location struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}

var beer []Beer

func CreateBeerEndpoint(w http.ResponseWriter, r *http.Request) {
	var beer Beer
	_ = json.NewDecoder(r.Body).Decode(&beer)
	beer = append(beer, beer)
	json.NewEncoder(w).Encode(beer)
}

func GetBeerEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(beer)
}

func GetBeerEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, p := range beer {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Beer Not Found")
}

func UpdateBeerEndpoint(w http.ResponseWriter, r *http.Request) {
	var beer Beer
	_ = json.NewDecoder(r.Body).Decode(&beer)
	params := mux.Vars(r)
	for i, p := range beer {
		if p.ID == params["id"] {
			beer[i] = beer
			json.NewEncoder(w).Encode(beer)
			break
		}
	}
}

func DeleteBeerEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, p := range beer {
		if p.ID == params["id"] {
			copy(beer[i:], beer[i+1:])
			beer = beer[:len(beer)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(beer)
}

func main() {

	router := mux.NewRouter()
	beer = append(beer, Beer{ID: xid.New().String(), Name: "CBC Pale Ale", Category: "IPA", Brewer: "CBC", Location: Location{Country: "South Africa", City: "Paarl"}})
	beer = append(beer, Beer{ID: xid.New().String(), Name: "Heineken", Category: "Lager", Brewer: "Heineken", Location: Location{Country: "Netherlands", City: "Amsterdam"}})
	beer = append(beer, Beer{ID: xid.New().String(), Name: "Jack Black Butchers Block", Category: "IPA", Brewer: "Devils Peak", Location: Location{Country: "South Africa", City: "Woodstock"}})
	beer = append(beer, Beer{ID: xid.New().String(), Name: "CBC Amber Weiss", Category: "Weiss", Brewer: "CBC", Location: Location{Country: "South Africa", City: "Paarl"}})
	
  router.HandleFunc("/beer", GetBeerEndpoint).Methods("GET")
	router.HandleFunc("/beer/{id}", GetBeerEndpoint).Methods("GET")
	router.HandleFunc("/beer", CreateBeerEndpoint).Methods("POST")
	router.HandleFunc("/beer/{id}", DeleteBeerEndpoint).Methods("DELETE")
	router.HandleFunc("/beer/{id}", UpdateBeerEndpoint).Methods("PUT")
	fmt.Println("Starting Server on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
  
}
