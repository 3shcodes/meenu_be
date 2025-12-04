package routes


import (
	http "net/http"
        controllers "meenu_be/controllers"
	sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func addClientRoutes(mux *http.ServeMux, db *sql.DB) {

	mux.HandleFunc("/clients/new", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateNewClient(w, r, db);
	}) 
}
