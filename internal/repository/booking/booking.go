package booking

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/ViharevN/design_test_master/config"
	"github.com/ViharevN/design_test_master/internal/model"
)

type repository struct {
	connectionPool *pgxpool.Pool
}

func (r *repository) CreateOrder(ctx context.Context, tx pgx.Tx, order model.Order) error {
	sql := `INSERT INTO orders (hotel_id, room_id, user_email, from_date, to_date) VALUES ($1, $2, $3, $4, $5)`

	_, err := tx.Exec(ctx, sql, order.HotelID, order.RoomID, order.UserEmail, order.From, order.To)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	return nil
}

func (r repository) CreateRoom(ctx context.Context, room model.Room) error {
	sql := `INSERT INTO rooms (hotel_id, room_id, category, description, status) VALUES ($1,$2,$3,$4,$5)`

	_, err := r.connectionPool.Exec(ctx, sql, room.HotelID, room.RoomID, room.Category, room.Description, room.Status)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return fmt.Errorf("такая комната уже существует: %w", err)
		}
		return fmt.Errorf("failed ro create new room: %w", err)
	}

	return nil
}

func (r repository) GetAvailableRoomsByDate(ctx context.Context, date time.Time) (model.Availability, error) {
	sql := `
		SELECT r.hotel_id, r.room_id
		FROM rooms r
		LEFT JOIN orders o ON r.hotel_id = o.hotel_id AND r.room_id = o.room_id
			AND $1::date <@ daterange(o.from_date, o.to_date, '[]')
		WHERE o.id IS NULL
	`

	rows, err := r.connectionPool.Query(ctx, sql, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get available rooms: %w", err)
	}
	defer rows.Close()

	var rooms model.Availability
	for rows.Next() {
		var room model.RoomAvailability
		err = rows.Scan(&room.HotelID, &room.RoomID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan room: %w", err)
		}
		room.Date = date
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read rows: %w", err)
	}

	return rooms, nil
}

func (r repository) GetAvailableRoomsByRange(ctx context.Context, dateFrom, dateTo time.Time) (model.Availability, error) {
	sql := `
		SELECT r.hotel_id, r.room_id
		FROM rooms r
		LEFT JOIN orders o ON r.hotel_id = o.hotel_id AND r.room_id = o.room_id
			AND daterange($1::date, $2::date, '[]') && daterange(o.from_date, o.to_date, '[]')
		WHERE o.id IS NULL
	`

	rows, err := r.connectionPool.Query(ctx, sql, dateFrom, dateTo)
	if err != nil {
		return nil, fmt.Errorf("failed to get available rooms: %w", err)
	}
	defer rows.Close()

	var rooms model.Availability
	for rows.Next() {
		var room model.RoomAvailability
		err = rows.Scan(&room.HotelID, &room.RoomID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan room: %w", err)
		}
		room.Date = dateFrom
		rooms = append(rooms, room)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to read rows: %w", err)
	}

	return rooms, nil
}

func (r repository) CheckRoomAvailability(ctx context.Context, tx pgx.Tx, order model.Order) (bool, error) {
	sql := `SELECT COUNT(*) 
		FROM orders 
		WHERE hotel_id = $1 
		AND room_id = $2 
		AND daterange(from_date, to_date, '[]') && daterange($3::date, $4::date, '[]');
`

	var count int
	err := tx.QueryRow(ctx, sql, order.HotelID, order.RoomID, order.From, order.To).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("failed to check room availability: %w", err)
	}

	// Если count > 0, это означает, что есть заказы, которые пересекаются с указанными датами, поэтому комната недоступна
	return count == 0, nil
}

func NewRepository(configuration config.Config) (repository, error) {
	connectionPool, err := pgxpool.New(context.Background(), configuration.SQLConnectionUrl)
	if err != nil {
		return repository{}, fmt.Errorf("failed to create pgxpool: %w", err)
	}

	return repository{
		connectionPool: connectionPool,
	}, nil
}

func (self repository) Ping(ctx context.Context) error {
	return self.connectionPool.Ping(ctx)
}

func (self repository) Begin(ctx context.Context) (pgx.Tx, error) {
	return self.connectionPool.Begin(ctx)
}
