package routes

import (
	"fmt"
	"myapp/controller"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func InitializeRoutes() {
	router := mux.NewRouter()

	// router.HandleFunc("/home", controller.AddContent).Methods("POST")
	// router.HandleFunc("/home/{myname}", controller.ParameterHandler)
	router.HandleFunc("/content", controller.AddContent).Methods("POST")
	router.HandleFunc("/content/{cid}", controller.GetContent).Methods("GET")
	router.HandleFunc("/content/{cid}", controller.UpdateContent).Methods("PUT")
	router.HandleFunc("/content/{cid}", controller.DeleteContent).Methods("DELETE")
	router.HandleFunc("/contents", controller.GetAllContent)

	// router.HandleFunc("/course", controller.AddCourse).Methods("POST")
	// router.HandleFunc("/course/{cid}", controller.GetCourse).Methods("GET")
	// router.HandleFunc("/course/{cid}", controller.UpdateCourse).Methods("PUT")
	// router.HandleFunc("/course/{cid}", controller.DeleteCourse).Methods("DELETE")
	// router.HandleFunc("/courses", controller.GetAllCourse)

	fhandler := http.FileServer(http.Dir("./view"))
	router.PathPrefix("/").Handler(fhandler)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
