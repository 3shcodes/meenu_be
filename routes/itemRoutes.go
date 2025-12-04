package routes


import (
	http "net/http"
        controllers "meenu_be/controllers"
	sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func addItemRoutes(mux *http.ServeMux, db *sql.DB) {

	mux.HandleFunc("/items/new", func(w http.ResponseWriter, r *http.Request) {
		controllers.CreateNewItem(w, r, db);
	}) 
}
