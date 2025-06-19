CREATE TYPE seat_status AS ENUM ('available', 'locked', 'booked');
CREATE TYPE transaction_status AS ENUM ('pending', 'success', 'failed', 'refunded');

CREATE TABLE city (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) NOT NULL,
    role VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE cinemas (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    id_city INT NOT NULL REFERENCES city(id),
    address TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE film (
    id SERIAL PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    duration INT NOT NULL, 
    description TEXT,
    rating VARCHAR(10),
    created_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE shcedule (
    id SERIAL PRIMARY KEY,
    cinema_id INT NOT NULL REFERENCES cinemas(id),
    movies_id INT NOT NULL REFERENCES film(id),
    start_time TIMESTAMP NOT NULL,
    screen_number INT NOT NULL,
    is_cancle BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE seat (
    id SERIAL PRIMARY KEY,
    schedules_id INT NOT NULL REFERENCES shcedule(id),
    codeseat VARCHAR(10) NOT NULL,
    status seat_status DEFAULT 'available',
    locked_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transaction (
    id UUID PRIMARY KEY,
    codetransaksi VARCHAR(50) UNIQUE NOT NULL,
    user_id INT NOT NULL REFERENCES users(id),
    jadwal_id INT NOT NULL REFERENCES shcedule(id),
    statustransaction transaction_status DEFAULT 'pending',
    totalamount INT NOT NULL,
    metodeamount VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE transactionitem (
    id SERIAL PRIMARY KEY,
    transaksi_id UUID NOT NULL REFERENCES transaction(id),
    seat_id INT NOT NULL REFERENCES seat(id),
    price INT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    update_at TIMESTAMP DEFAULT NOW()
);
