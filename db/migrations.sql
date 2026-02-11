CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price NUMERIC NOT NULL,
    stock INT NOT NULL,
    size TEXT,
    color TEXT,
    category_id INT REFERENCES categories(id) ON DELETE CASCADE
);
