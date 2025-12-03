package models

import (
	time "time"
	sql "database/sql"
)

type Transaction struct {
   Id              int
   ClientID        int
   TransactionType string // 'BILL', 'CRED'
   Amount          float64
   HasProcessed    bool
   PostProcess     float64
   Status          string // 'PEND', 'PART', 'DONE'
   DoneOn          time.Time
   UsedFor         string
}

type TransactionTable struct {
	db *sql.DB
}

func (trt *TransactionTable) SetDB(db *sql.DB) {
	trt.db = db;
}

func (trt *TransactionTable) CreateNewTransaction(newBill *Bill, totalCost float64) int {

	insQuery := "INSERT INTO transactions(client_id, transaction_type, amount, has_processed, post_process, status, done_on, used_for ) values(?, 'BILL', ?, 0, ?, 'PEND', ?, NULL)"


	stmt, err := trt.db.Prepare(insQuery);
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec(newBill.ClientID, totalCost, totalCost, newBill.BillDate);
	if err != nil {
		panic(err)
	}

	trsnId, err := res.LastInsertId();
	if err != nil {
		panic(err)
	}

	newBill.TransactID = int(trsnId)

	return newBill.TransactID;

}

