package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"practiceRestServices/service"
	"strings"
	"time"

	// "github.com/gorilla/mux"
	"net/http"
	"practiceRestServices/Dtos"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var myData []Dtos.Customer
	var foundedCostumer Dtos.Customer
	log.Print("customer id request", vars)
	resp, err := http.Get("http://localhost:8000/customers")
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()

	errores := json.NewDecoder(resp.Body).Decode(&myData)

	if errores != nil {
		panic(errores)
	}

	for i := range myData {
		if myData[i].Id == vars["customer_id"] {
			foundedCostumer = myData[i]
		}
	}

	json.NewEncoder(w).Encode(foundedCostumer)
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {

	customers := []Dtos.Customer{
		{Id: "1", Name: "New York", City: "10050ZX"},
		{Id: "2", Name: "Chicago", City: "2313DF"},
		{Id: "3", Name: "Portland", City: "8273847F"},
		{Id: "4", Name: "New Jersey", City: "14432HF"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "POST REQUEST")
}

func getTimeArg(w http.ResponseWriter, r *http.Request) {

	vars := r.URL.Query()
	x := make(map[string][]string)
	tz := vars.Get("tz")
	if len(tz) > 0 {
		splitedCities := strings.Split(tz, ",")
		for i := 0; i < len(splitedCities); i++ {
			fmt.Print(splitedCities[i])
			fmt.Print(len(splitedCities))
			loc, err := time.LoadLocation(splitedCities[i])
			if err != nil {
				splitedCities[i] = "No LOCATION"
			}
			x[splitedCities[i]] = append(x[splitedCities[i]], time.Now().In(loc).String())
		}
		json.NewEncoder(w).Encode(x)
	} else {
		loc, _ := time.LoadLocation("America/Buenos_Aires")
		var currentTime Dtos.CurrentFile
		currentTime.SetCurrentFileString(time.Now().In(loc).String())
		json.NewEncoder(w).Encode(currentTime)
	}

}
