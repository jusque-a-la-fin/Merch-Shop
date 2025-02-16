CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

CREATE TABLE shop (
    id SERIAL PRIMARY KEY,
    name TEXT UNIQUE NOT NULL
);

INSERT INTO shop (name) VALUES ('avito-shop');

CREATE TABLE items (
    id SERIAL PRIMARY KEY,
    -- Тип предмета или название предмета
    item_type TEXT UNIQUE NOT NULL,
    price INT NOT NULL
);

CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    item_id INT NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    quantity INT NOT NULL
);

CREATE TABLE receivers (
    id INT NOT NULL 
);

CREATE TABLE coin_history (
    id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id INT NOT NULL REFERENCES receivers(id) ON DELETE CASCADE,
    amount INT NOT NULL
);

CREATE TABLE coins_balance (
    id SERIAL PRIMARY KEY,
    balance INT NOT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);
