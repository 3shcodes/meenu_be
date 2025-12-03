package models

import (
   	time "time"
   	_ "fmt"
	sql "database/sql"
)

type Bill struct {
   Id        int
   BillID    int
   BillType  string // 'CUST', 'CORP'
   BillDate  time.Time
   ClientID  int
   TransactID int
}


type BillTable struct {
	db *sql.DB;	
}


func (bt *BillTable) SetDB(db *sql.DB) {

	bt.db = db
}


func (bt *BillTable) UpdateBill(transactionId int, billId int) {

	updQry := "UPDATE bills SET transact_id=? WHERE bill_id=?";
	stmt, err := bt.db.Prepare(updQry);
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(transactionId, billId);
	if err != nil {
		panic(err)
	}

}


func (bt *BillTable) CheckBillExists(billId int) bool {
	
	stmt, err := bt.db.Prepare("SELECT COUNT(*) as cnt FROM bills where bill_id=?");
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	cursor, err := stmt.Query(billId);
	if err != nil {
		panic(err)
	}
	defer cursor.Close();

	var cnt int
	for cursor.Next() {
		if err = cursor.Scan(&cnt); err != nil {
			panic(err)
		}
	}

	return cnt != 0;

}

func (bt *BillTable) CreateNewBill(billId int, date time.Time, clientId int) {


	insQuery := "INSERT INTO bills(bill_id, bill_type, bill_date, client_id) values(?,?,?,?);";
	stmt, err := bt.db.Prepare(insQuery);
	if err != nil {
		panic(err)
	}
	defer stmt.Close();

	if _, err = stmt.Exec(billId, "CORP", date, clientId); err != nil {
		panic(err)
	}
}


func FetchByTimeRange(from time.Time, to time.Time)  []*Bill{
	
	return nil;
}

func FetchByTransactID(id string)  []*Bill{

	return nil;
}
