DROP DATABASE IF EXISTS meenu_demo;
CREATE DATABASE meenu_demo;

USE meenu_demo;

CREATE TABLE users(
        id INT PRIMARY KEY AUTO_INCREMENT,
        username VARCHAR(20),
        display_name VARCHAR(100),
        email_id VARCHAR(100),
        phone VARCHAR(20),
        CONSTRAINT uq_username UNIQUE (username)
);

CREATE TABLE auths(
        id INT PRIMARY KEY AUTO_INCREMENT,
        user_id int,
        pwd varchar(64),
        CONSTRAINT fk_users_auth FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE clients(
        id INT PRIMARY KEY AUTO_INCREMENT,
        client_name VARCHAR(200),
        client_disp_name VARCHAR(200),
        first_transaction TIMESTAMP,
        created_on TIMESTAMP,
        CONSTRAINT uq_client_name UNIQUE (client_name)
);

CREATE TABLE transactions(
        id INT PRIMARY KEY AUTO_INCREMENT,
        client_id INT,
        transaction_type ENUM('BILL', 'CRED'), -- BILL, CRED, DEBT
        amount decimal(20,4),
        has_processed TINYINT(1),
        post_process decimal(20,4),
        status ENUM('PEND', 'PART', 'DONE'), -- PEND, DONE, PART
        done_on TIMESTAMP,
        used_for TEXT,
        CONSTRAINT fk_clients_transactions FOREIGN KEY (client_id) REFERENCES clients(id)
);

CREATE TABLE bills(
        id INT PRIMARY KEY AUTO_INCREMENT,
        bill_id INT,
        bill_type ENUM('CUST', 'CORP'), 
        bill_date TIMESTAMP,
        client_id INT,
        transact_id int,
        CONSTRAINT fk_clients_bills FOREIGN KEY (client_id) REFERENCES clients(id),
        CONSTRAINT fk_transactions_bills FOREIGN KEY (transact_id) REFERENCES transactions(id)
);

CREATE TABLE items(
        id INT PRIMARY KEY AUTO_INCREMENT,
        item_name varchar(20),
        item_disp_name varchar(100),
        CONSTRAINT uk_item_name UNIQUE (item_name)
);

CREATE TABLE prices(
        id INT PRIMARY KEY AUTO_INCREMENT,
        item_id INT,
        rate DECIMAL(20,4),
        is_default TINYINT(1),
        CONSTRAINT fk_items_prices FOREIGN KEY (item_id) REFERENCES items(id)
);

CREATE TABLE bill_items(
        id INT PRIMARY KEY AUTO_INCREMENT,
        bill_id INT,
        item_id INT,
        quantity INT,
        price_id INT,
        CONSTRAINT fk_bills_bill_items FOREIGN KEY (bill_id) REFERENCES bills(id),
        CONSTRAINT fk_items_bill_items FOREIGN KEY (item_id) REFERENCES items(id),
        CONSTRAINT fk_prices_bill_items FOREIGN KEY (price_id) REFERENCES prices(id)
);

CREATE VIEW bill_summary AS
SELECT 
        bt.bill_id,
        bt.bill_date as billed_on,
        it.item_disp_name, 
        pi.rate, 
        bi.quantity, 
        bi.quantity * pi.rate as item_cost 
FROM bill_items bi 
JOIN bills bt ON bt.id = bi.bill_id 
JOIN items it ON it.id = bi.item_id 
JOIN prices pi ON pi.id = bi.price_id
