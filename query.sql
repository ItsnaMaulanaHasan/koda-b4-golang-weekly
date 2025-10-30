CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    product_id INT REFERENCES products (id),
    quantity INT DEFAULT 1,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE histories (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP DEFAULT now(),
    no_invoice VARCHAR(20) UNIQUE NOT NULL,
    total NUMERIC(10, 2),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE product_history (
    id SERIAL PRIMARY KEY,
    history_id INT REFERENCES histories (id),
    product_id INT REFERENCES products (id),
    quantity int DEFAULT 1,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);