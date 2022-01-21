package routing

import (
	core "api-test/core/controllers"
	database "api-test/interface/database/mongo"

	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()
	uc := core.NewUserController(database.CreateMongoClient())
	r.HandleFunc("/user/:id", uc.GetUser).Methods("GET")
	r.HandleFunc("/user", uc.CreateUser).Methods("POST")
	r.HandleFunc("/user/:id", uc.DeleteUser).Methods("DELETE")
	return r
}
