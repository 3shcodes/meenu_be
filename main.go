package main

import (
	http "net/http" 
	log "log"
	os "os"
	fmt "fmt" 
	godotenv "github.com/joho/godotenv" 
	routes "meenu_be/routes"
	database "meenu_be/database"
)

var (
	port string
)



func main() {

	godotenv.Load()

	port = os.Getenv("PORT")
	sqlType := os.Getenv("DBTYPE")
	dbConf := os.Getenv(sqlType);
	dbName := os.Getenv("DB_NAME");

	db := database.CreateInstance(sqlType, dbConf+"/"+dbName);
	defer db.Close();

	mux := http.NewServeMux();

	routes.InitiateRoutes(mux, db)

	log.Fatal(http.ListenAndServe(":"+port, mux))

	fmt.Println("ayo")
}
