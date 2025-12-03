package routes


import (
	http "net/http"
        controllers "meenu_be/controllers"
	sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func addBillRoutes(mux *http.ServeMux, db *sql.DB) {

	mux.HandleFunc("/bills/getBills", func(w http.ResponseWriter, r *http.Request) {
		controllers.FetchBills(w, r, db);
	}) 
	mux.HandleFunc("/bills/generateBill", func(w http.ResponseWriter, r *http.Request) {
		controllers.GenerateBill(w, r, db);
	}) 
}
