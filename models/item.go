package models

import (
	_ "time"
	sql "database/sql"
)

type Item struct {
   Id             int
   ItemName     string
   ItemDispName string
}


type ItemTable struct {
	db *sql.DB;	
}



func (it *ItemTable) SetDB(db *sql.DB) {
	it.db = db
}

func (it *ItemTable) CheckIfItemNameExists(itemName string) (bool, int) {
	
	stmt, err := it.db.Prepare("SELECT id FROM items where item_name=?");
	if err != nil {
		panic(err)
	}

	cursor, err := stmt.Query(itemName);
	if err != nil {
		panic(err)
	}

	cnt := 0;
	id := 0;
	for cursor.Next() {
		if err = cursor.Scan(&id); err != nil {
			panic(err)
		}
		cnt++;
	}

	return cnt != 0, id;

}

func (it *ItemTable) CreateNewItem(itemName, itemDispName string) int {

	stmt, err := it.db.Prepare("INSERT INTO items(item_name, item_disp_name ) values(?, ?)");
	if err != nil {
		panic(err)
	}
	defer stmt.Close();

	res, err := stmt.Exec(itemName, itemDispName);
	if err != nil {
		panic(err)
	}

	itemId, err := res.LastInsertId();
	if err != nil {
		panic(err)
	}

	return int(itemId);
}


