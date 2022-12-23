DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS account;
DROP TABLE IF EXISTS bill;
DROP TABLE IF EXISTS income;

CREATE TABLE users(
  id VARCHAR(32) PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  last_name VARCHAR(255) NOT NULL,
  email VARCHAR(255) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);



CREATE TABLE account (
  id VARCHAR(32) PRIMARY KEY,
  name_account VARCHAR(255),
  bank VARCHAR(255),
  balance integer,
  user_id VARCHAR(32) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE bill (
  id VARCHAR(32) PRIMARY KEY,
  bill_name VARCHAR(32),
  bill_description VARCHAR(255),
  amount integer,
  account_id VARCHAR(255) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE income (
  id VARCHAR(32) PRIMARY KEY,
  icome_name VARCHAR(255),
  icome_description VARCHAR(255),
  amount integer,
  account_id VARCHAR(32) NOT NULL,
  created_at timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE account ADD FOREIGN KEY (user_id) REFERENCES users (id);

ALTER TABLE bill ADD FOREIGN KEY (account_id) REFERENCES account (id);

ALTER TABLE income ADD FOREIGN KEY (account_id) REFERENCES account(id);
