package database

import (
	fmt "fmt"
	sql "database/sql"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)



func CreateInstance(sqlType, dbConf string) *sql.DB {

	initQry := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT,
		display_name TEXT,
		email_id TEXT,
		phone TEXT,
		CONSTRAINT uq_username UNIQUE (username)
	);

	CREATE TABLE IF NOT EXISTS auths (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER,
		pwd TEXT,
		CONSTRAINT fk_users_auth FOREIGN KEY (user_id) REFERENCES users(id)
	);

	CREATE TABLE IF NOT EXISTS clients (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client_name TEXT,
		client_disp_name TEXT,
		first_transaction TIMESTAMP,
		created_on TIMESTAMP,
		CONSTRAINT uq_client_name UNIQUE (client_name)
	);

	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		client_id INTEGER,
		transaction_type TEXT CHECK (transaction_type IN ('BILL', 'CRED')),  -- BILL, CRED, DEBT
		amount DECIMAL(20, 4),
		has_processed BOOLEAN,
		post_process DECIMAL(20, 4),
		status TEXT CHECK (status IN ('PEND', 'PART', 'DONE')),  -- PEND, DONE, PART
		done_on TIMESTAMP,
		used_for TEXT,
		CONSTRAINT fk_clients_transactions FOREIGN KEY (client_id) REFERENCES clients(id)
	);

	CREATE TABLE IF NOT EXISTS bills (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bill_id INTEGER,
		bill_type TEXT CHECK (bill_type IN ('CUST', 'CORP')),
		bill_date TIMESTAMP,
		client_id INTEGER,
		transact_id INTEGER,
		CONSTRAINT fk_clients_bills FOREIGN KEY (client_id) REFERENCES clients(id),
		CONSTRAINT fk_transactions_bills FOREIGN KEY (transact_id) REFERENCES transactions(id)
	);

	CREATE TABLE IF NOT EXISTS items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		item_name TEXT,
		item_disp_name TEXT,
		CONSTRAINT uk_item_name UNIQUE (item_name)
	);

	CREATE TABLE IF NOT EXISTS prices (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		item_id INTEGER,
		rate DECIMAL(20, 4),
		is_default BOOLEAN,
		CONSTRAINT fk_items_prices FOREIGN KEY (item_id) REFERENCES items(id)
	);

	CREATE TABLE IF NOT EXISTS bill_items (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bill_id INTEGER,
		item_id INTEGER,
		quantity INTEGER,
		price_id INTEGER,
		CONSTRAINT fk_bills_bill_items FOREIGN KEY (bill_id) REFERENCES bills(id),
		CONSTRAINT fk_items_bill_items FOREIGN KEY (item_id) REFERENCES items(id),
		CONSTRAINT fk_prices_bill_items FOREIGN KEY (price_id) REFERENCES prices(id)
	);

	CREATE VIEW IF NOT EXISTS bill_summary AS
	SELECT 
	bt.bill_id,
	bt.bill_date AS billed_on,
	it.item_disp_name, 
	pi.rate, 
	bi.quantity, 
	bi.quantity * pi.rate AS item_cost 
	FROM 
	bill_items bi 
	JOIN 
	bills bt ON bt.id = bi.bill_id 
	JOIN 
	items it ON it.id = bi.item_id 
	JOIN 
	prices pi ON pi.id = bi.price_id;

	insert into items(item_name, item_disp_name) values('vanjaram', "Vanjaram");
	insert into items(item_name, item_disp_name) values('prawn', "Prawns");

	insert into clients(client_name, client_disp_name, created_on) values("marvel", "MARVELMMC", "CURRENT_TIMESTAMP");


	`

	if sqlType == "mysql" {
		db, err := sql.Open(sqlType, dbConf);
		if err != nil {
			panic(err)
		}
		return db;
	} else {
		fmt.Println(dbConf[1:]+".db")
		db, err := sql.Open(sqlType, dbConf[1:]+".db");
		if err != nil {
			panic(err)
		}
		_, err = db.Exec(initQry)
		if err != nil {
			panic(err)
		}

		return db;
	}




}
