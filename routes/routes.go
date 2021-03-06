package routes

import "github.com/gorilla/mux"

func setPersonaRoutes(router *mux.Router) {
	subRoute := router.PathPrix("/persona/api").subRoute()
	subRoute.HandleFunc("/all", controllers.GetAll).Methods("GET")
	subRoute.HandleFunc("/save", controllers.Save).Methods("POST")
	subRoute.HandleFunc("/delete", controllers.Delete).Methods("POST")
	subRoute.HandleFunc("/find/{id}", controllers.Get).Methods("GET")

}
