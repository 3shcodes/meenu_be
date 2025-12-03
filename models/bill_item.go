package models

import (
	_ "time"
	sql "database/sql"
)



type BillItem struct {
	
	Id int 
        BillId int
        ItemId int
        ItemCount int
        PriceId int

}

type BillItemTable struct {
	db *sql.DB
}

func (bit *BillItemTable) SetDB(db *sql.DB) {

	bit.db = db;
}

func (bit *BillItemTable) CreateBillItems(billId int, itemName string, priceId int, quantity float64) {

	insQuery := "INSERT INTO bill_items (bill_id, item_id, price_id, quantity) SELECT b.id as bill_id, it.id as item_id, ?, ? FROM bills b CROSS JOIN items it WHERE it.item_name = ? AND b.bill_id=?"


	stmt, err := bit.db.Prepare(insQuery);
	if err != nil {
		panic(err)
	}

	defer stmt.Close();

	_, err = stmt.Exec(priceId, quantity, itemName, billId);
	if err != nil {
		panic(err)
	}

}
