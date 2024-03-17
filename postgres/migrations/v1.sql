--create database stori;

DROP TABLE IF EXISTS transactions;
CREATE TABLE transactions
(
    id               SERIAL    NOT NULL PRIMARY KEY,
    email            VARCHAR   NOT NULL,
    type_transaction VARCHAR   NOT NULL,
    transaction      decimal   not null,
    date             timestamp not null
);