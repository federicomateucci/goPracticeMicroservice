package app

import (
	"github.com/ashishjuyal/banking/domain"
	"log"
	"net/http"
	"practiceRestServices/service"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id}", GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/api/time", getTimeArg).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))

}
