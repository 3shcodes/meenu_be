package models

import (
   time "time"
   sql "database/sql"
)

type Client struct {
   Id             int
   ClientName     string
   ClientDispName string
   FirstTransaction time.Time
   CreatedOn       time.Time
}


type ClientTable struct {
	db *sql.DB;	
}



func (ct *ClientTable) SetDB(db *sql.DB) {
	ct.db = db
}

func (ct *ClientTable) CreateNewClient(clientName string, clientDispName string) {


	stmt, err := ct.db.Prepare("INSERT INTO clients(client_name, client_disp_name, created_on) values(?, ?, NOW())");
	if err != nil {
		panic(err)
	}
	defer stmt.Close();

	_, err = stmt.Exec(clientName, clientDispName);
	if err != nil {
		panic(err)
	}

}


func (ct *ClientTable) CheckClientExists(clientName string) (bool, int) {
	
	stmt, err := ct.db.Prepare("SELECT id FROM clients where client_name=?");
	if err != nil {
		panic(err)
	}

	cursor, err := stmt.Query(clientName);
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

