package controllers


import (
	http "net/http"
	time "time"
	fmt "fmt"
	io "io"
	ioutil "io/ioutil"
	_ "time"
	json "encoding/json"
	models "meenu_be/models"
	sql "database/sql"
)

func GenerateBill(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	type SubBillItem struct {
		ItemName string
		Quantity float64
		Rate float64 
	}

	type NewBill struct {
		BillId int
		Date string
		ClientName string
		BillItems []SubBillItem
	}


	defer r.Body.Close();

	body, err := ioutil.ReadAll(r.Body);
	if err != nil {
		panic(err)
	}

	var someVar NewBill
	if err := json.Unmarshal(body, &someVar); err != nil {
		panic(err)
	}

	fmt.Println(someVar)
	fmt.Println(someVar.BillId)
	fmt.Println(someVar.Date)
	fmt.Println(someVar.ClientName)
	fmt.Println(someVar.BillItems)


	for _, v := range someVar.BillItems {
		fmt.Println(v.ItemName);
		fmt.Println(v.Quantity);
		fmt.Println(v.Rate);
		fmt.Println();
	}

	
	billsTable := new(models.BillTable);
	billsTable.SetDB(db);

	clientsTable := new(models.ClientTable);
	clientsTable.SetDB(db);

	pricesTable := new(models.PriceTable);
	pricesTable.SetDB(db);

	billItemsTable := new(models.BillItemTable);
	billItemsTable.SetDB(db);

	transactionsTable := new(models.TransactionTable);
	transactionsTable.SetDB(db);

	if billsTable.CheckBillExists(someVar.BillId) {
		io.WriteString(w, *models.MakeResp("Failed", 403, nil));
		return;
	}



	isExistingClient, clientId := clientsTable.CheckClientExists(someVar.ClientName) 
	if !isExistingClient {
		io.WriteString(w, *models.MakeResp("Failed", 403, nil));
		return;
	}

	t, err := time.Parse("02-01-2006", someVar.Date);
	if err != nil {
		panic(err)
	}




	billsTable.CreateNewBill(someVar.BillId, t, clientId);

	var totalAmnt float64;

	for _, v := range someVar.BillItems {
		price_id := pricesTable.InsertAndReturnPriceIfNotExists(v.ItemName, v.Rate);
		billItemsTable.CreateBillItems(someVar.BillId, v.ItemName, price_id, v.Quantity);

		totalAmnt += v.Quantity * v.Rate;
	}



	newBill := &models.Bill {

		BillID: someVar.BillId,
		BillType: "CORP",
		BillDate: t,
		ClientID: clientId,
		TransactID: -1,
	}

	transactId := transactionsTable.CreateNewTransaction(newBill, totalAmnt);

	billsTable.UpdateBill(transactId, someVar.BillId);

	io.WriteString(w, *models.MakeResp("Success", 200, nil));


}


func FetchBills(w http.ResponseWriter, r *http.Request, db *sql.DB) {

	// queries := r.URL.Query();
	// fetchType := queries.Get("type")
	// bill_id := queries.Get("bill_id");
	// clnt_id := queries.Get("client_id");
	//
	// fmt.Println(fetchType+" "+bill_id+" "+clnt_id);
	//
	//
	// billsTable := new(models.BillTable);
	// billsTable.SetDB(db);
	//
	// var allBills []*models.Bill
	// if fetchType == "bill_id" {
	// 	allBills = billsTable.FetchByBillID(bill_id);
	// } else if fetchType == "client_id" {
	// 	allBills = billsTable.FetchByClient(clnt_id);
	// }
	//
	// fmt.Println(clnt_id);
	//
	//
	//
	// io.WriteString(w, *models.MakeResp("Success", 200, allBills));
	io.WriteString(w, *models.MakeResp("Success", 200, nil));
}
