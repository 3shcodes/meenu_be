package controllers

import (
	http "net/http"
	strings "strings"
	_ "time"
	fmt "fmt"
	io "io"
	ioutil "io/ioutil"
	_ "time"
	json "encoding/json"
	models "meenu_be/models"
	sql "database/sql"
)


func CreateNewItem(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	type NewItem struct {
		ItemName string
		ItemDispName string
		Price float64
	}

	defer r.Body.Close();

	body, err := ioutil.ReadAll(r.Body);
	if err != nil {
		panic(err)
	}

	var reqBd NewItem
	if err := json.Unmarshal(body, &reqBd); err != nil {
		panic(err)
	}

	if ( reqBd.ItemName == "" ) {
		io.WriteString(w, *models.MakeResp("Failed: itemname not in request", 403, nil));
		fmt.Println("Failed: itemname not in request in line no 39");
		return;
	}

	itemsTable := new(models.ItemTable)
	itemsTable.SetDB(db);

	itemExists, _ := itemsTable.CheckIfItemNameExists(reqBd.ItemName);
	if err != nil {
		panic(err)
	}

	if itemExists {
		io.WriteString(w, *models.MakeResp("Failed: itemname already exists", 403, nil));
		fmt.Println("Failed: itemname already exists in line no 39");
		return;

	}


	if ( reqBd.ItemDispName == "" ) {
		reqBd.ItemDispName = strings.Title(reqBd.ItemDispName);
	}


	itemId := itemsTable.CreateNewItem(reqBd.ItemName, reqBd.ItemDispName);

	if itemId == -1 {

		io.WriteString(w, *models.MakeResp("Failed: Error when inserting the item", 403, nil));
		fmt.Println("Failed: Error when inserting the item at itemController.go:69");
		return;
	}

	pricesTable := new(models.PriceTable)
	pricesTable.SetDB(db);

	pricesTable.InsertAndReturnPriceIfNotExists(reqBd.ItemName, reqBd.Price);


	io.WriteString(w, *models.MakeResp("Success: new client added successfully", 403, nil));

	
}
