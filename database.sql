CREATE TABLE products (
	id serial primary key,
	name TEXT,
	type INTEGER,
	description TEXT,
	curr_quantity INTEGER,
	min_quantity INTEGER 
);

CREATE TABLE order_products (
	order_id INTEGER,
	product_id INTEGER
);

CREATE TABLE orders (
	id serial primary key,
	approved BOOLEAN 
);

