create table users(
    id serial primary key,
    name varchar,
    age int,
    document varchar,
    address varchar
);

INSERT INTO users(id,name,age,document,address) VALUES
(1,'Ana Pereira',25,'123.456.789-00','Rua das Flores, 123, São Paulo, SP');