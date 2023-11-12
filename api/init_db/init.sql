CREATE DATABASE inventory;

USE inventory;

CREATE TABLE Products(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    quantity int,
    price float(10, 7),
    PRIMARY KEY(id)
);

INSERT INTO Products (id, name, quantity, price ) values 
    (1, "chair", 100, 200.00), (2, "desk", 50, 600.00);