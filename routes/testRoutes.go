package routes


import (
	http "net/http"
        controllers "meenu_be/controllers"
)

func addTestRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/test", controllers.TestFunc) 
}
