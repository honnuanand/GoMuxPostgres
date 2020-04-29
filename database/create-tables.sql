-- psql -v ON_ERROR_STOP=1 --username postgres  --host="127.0.0.1"  --port="5432" -d postgres
CREATE TABLE products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
);

