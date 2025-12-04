package controllers

import (
	http "net/http"
	strings "strings"
	_ "time"
	fmt "fmt"
	io "io"
	_ "io/ioutil"
	_ "time"
	_ "encoding/json"
	models "meenu_be/models"
	sql "database/sql"
)


func CreateNewClient(w http.ResponseWriter, r *http.Request, db *sql.DB) {


	queries := r.URL.Query();
	clientName := queries.Get("clientName");
	clientDispName := queries.Get("clientDispName");

	if ( clientName == "" ) {
		io.WriteString(w, *models.MakeResp("Failed: client name not found in request", 403, nil));
		fmt.Println("Failed: client name not found in request at clientController.go:25");
		return;
	}


	clientsTable := new(models.ClientTable);
	clientsTable.SetDB(db);

	isExistingClient, _ := clientsTable.CheckClientExists(clientName)
	if isExistingClient {
		io.WriteString(w, *models.MakeResp("Failed: client name already exists", 403, nil));
		fmt.Println("Failed: client name already exists at clientController.go:36");
		return;
	}

	if ( clientDispName == "" ) {
		clientDispName = strings.Title(clientName);
	}


	clientsTable.CreateNewClient(clientName, clientDispName);

	io.WriteString(w, *models.MakeResp("Success: new client added successfully", 403, nil));

	
}
