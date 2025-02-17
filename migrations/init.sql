CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

INSERT INTO users (username, password_hash) VALUES 
('user1', '08182008710cc0d57c774b85b4120804aea43aa32ed7304b020be43cfe39b8bc'),
('user2', '39e3ca2c8ac8450de86fe501c42ad6b35e8a272a685e4b9951148b47f42c59a7');

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

INSERT INTO items (item_type, price) VALUES 
('t-shirt', 80),
('cup', 20),
('book', 50),
('pen', 10),
('powerbank', 200),
('hoody', 300),
('umbrella', 200),
('socks', 10),
('wallet', 50),
('pink-hoody', 500);

CREATE TABLE inventory (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    item_id INT NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    quantity INT NOT NULL
);

CREATE TABLE receivers (
    id INT UNIQUE NOT NULL 
);

INSERT INTO receivers (id) VALUES 
('1'),
('2');

CREATE TABLE coin_history (
    id SERIAL PRIMARY KEY,
    sender_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id INT NOT NULL REFERENCES receivers(id) ON DELETE CASCADE,
    amount INT NOT NULL
);

INSERT INTO coin_history (sender_id, receiver_id, amount) VALUES 
('1', '2', 20);

CREATE TABLE coins_balance (
    id SERIAL PRIMARY KEY,
    balance INT NOT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO coins_balance (balance, user_id) VALUES 
(200, '1');
