CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL
);

INSERT INTO users (id, username, password_hash) VALUES 
('537d0b8a-4037-40c1-8a66-32a497baa9f0', 'user1', '08182008710cc0d57c774b85b4120804aea43aa32ed7304b020be43cfe39b8bc'),
('b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f', 'user2', '39e3ca2c8ac8450de86fe501c42ad6b35e8a272a685e4b9951148b47f42c59a7');

CREATE TABLE shop (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    shopname TEXT UNIQUE NOT NULL
);

INSERT INTO shop (id, shopname) VALUES 
('45e6f2b0-d91c-4751-9a6c-5f6702080361', 'avito-shop');

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
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    item_id INT NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    quantity INT NOT NULL
);

CREATE TABLE receivers (
    id UUID UNIQUE NOT NULL 
);

INSERT INTO receivers (id) VALUES 
('537d0b8a-4037-40c1-8a66-32a497baa9f0'),
('b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f'),
('45e6f2b0-d91c-4751-9a6c-5f6702080361');

CREATE TABLE coin_history (
    id SERIAL PRIMARY KEY,
    sender_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    receiver_id UUID NOT NULL REFERENCES receivers(id) ON DELETE CASCADE,
    amount INT NOT NULL
);

INSERT INTO coin_history (sender_id, receiver_id, amount) VALUES 
('537d0b8a-4037-40c1-8a66-32a497baa9f0', 'b9f034ca-3b54-4f3e-8fd2-1c1a7d3a418f', 20);

CREATE TABLE coins_balance (
    id SERIAL PRIMARY KEY,
    balance INT NOT NULL,
    user_id UUID UNIQUE NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO coins_balance (balance, user_id) VALUES 
(200, '537d0b8a-4037-40c1-8a66-32a497baa9f0');