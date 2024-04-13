-- Создание таблицы для комнат
CREATE TABLE rooms (
       id bigserial,
       hotel_id TEXT NOT NULL,
       room_id TEXT NOT NULL,
       category TEXT NOT NULL,
       description TEXT NOT NULL,
       status TEXT NOT NULL,
       created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
       updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
       PRIMARY KEY (hotel_id, room_id)
);

-- Создание индекса для room_id в таблице rooms
CREATE INDEX idx_room_id ON rooms (room_id);

-- Создание таблицы для брони комнат
CREATE TABLE orders (
        id BIGSERIAL PRIMARY KEY,
        hotel_id TEXT NOT NULL,
        room_id TEXT NOT NULL,
        user_email TEXT NOT NULL,
        from_date DATE NOT NULL,
        to_date DATE NOT NULL,
        created TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
        updated TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (hotel_id, room_id) REFERENCES rooms(hotel_id, room_id)
);
