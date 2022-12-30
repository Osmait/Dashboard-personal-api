DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS transactions;

CREATE TYPE TypeTransaction AS ENUM (
  'bill',
  'income'
);



CREATE TABLE users(
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  token VARCHAR(32) ,
  confirmed BOOLEAN  DEFAULT (false),
  created_at timestamptz NOT NULL DEFAULT (now())
);



CREATE TABLE account (
  id VARCHAR(32) PRIMARY KEY,
  name_account VARCHAR(255),
  bank VARCHAR(255),
  balance float,
  user_id VARCHAR(32) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE transactions (
  id varchar PRIMARY KEY,
  transaction_name varchar NOT NULL,
  transaction_description text,
  amount float NOT NULL,
  type_transation TypeTransaction NOT NULL,
  user_id varchar NOT NULL,
  account_id varchar NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);
ALTER TABLE account ADD FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE transactions ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE transactions ADD FOREIGN KEY (account_id) REFERENCES account (id);