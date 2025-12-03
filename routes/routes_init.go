package routes

import (
        http "net/http"
        fmt "fmt"
	sql "database/sql"
)

func InitiateRoutes(mux *http.ServeMux, db *sql.DB) {

        // addTestRoutes(mux);
        addBillRoutes(mux, db);

	fmt.Println("Routes initated");
}
