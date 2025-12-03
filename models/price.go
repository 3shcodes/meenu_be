package models

import (
	_ "time"
	sql "database/sql"
)

type Price struct {
	Id             	int
	ItemId     	int
	Rate 		float64
	IsDefault 	bool
}

type PriceTable struct {
	db *sql.DB
}

func (pt *PriceTable) SetDB(db *sql.DB) {
	pt.db = db;
}

func (pt *PriceTable) GetDefaultPriceByItemId(itemId int) {
}

func (pt *PriceTable) GetAllPriceByItemId(itemId int) {
}

func (pt *PriceTable) InsertAndReturnPriceIfNotExists(item_name string, rate float64) int {

	insQuery := "INSERT INTO prices (item_id, rate, is_default) SELECT it.id, ?, 1 FROM items it WHERE it.item_name=? AND NOT EXISTS(SELECT 1 FROM prices WHERE item_id=it.id and rate=?)"

	stmt, err := pt.db.Prepare(insQuery);
	if err != nil {
		panic(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(rate, item_name, rate);
	if err != nil {
		panic(err)
	}


	
	fQuery := "SELECT p.id FROM prices p JOIN items it ON it.id = p.item_id and item_name=?";
	
	stmt, err = pt.db.Prepare(fQuery);
	if err != nil {
		panic(err)
	}

	defer stmt.Close()
	cursor, err := stmt.Query(item_name);
	if err != nil {
		panic(err)
	}

	defer cursor.Close();


	resId := -1;
	for cursor.Next() {
		cursor.Scan(&resId);
	}


	return resId;

}








