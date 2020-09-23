package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/aldhonie/go-simple-clean-crud/car/repository"
	"github.com/aldhonie/go-simple-clean-crud/domain"
	"github.com/sirupsen/logrus"
)

type mysqlCarRepository struct {
	Conn *sql.DB
}

//NewMysqlCarRepository will create an object that represent the car.Repository interface
func NewMysqlCarRepository(Conn *sql.DB) domain.CarRepository {
	return &mysqlCarRepository{Conn}
}

//Fetch data ...
func (m *mysqlCarRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []domain.Car, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]domain.Car, 0)

	for rows.Next() {
		t := domain.Car{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Brand,
			&t.Price,
			&t.Condition,
			&t.Quantity,
			&t.Description,
			&t.Specification,
			&t.ImageURL,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

	}

	return result, nil
}

func (m *mysqlCarRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Car, nextCursor string, err error) {
	query := `SELECT id, name, brand, price, condition, quantity, description, specification, image_url, updated_at, created_at 
		FROM car WHERE created_at > ? ORDER BY created_at LIMIT ?`

	decodedCursor, err := repository.DecodeCursor(cursor)
	if err != nil {
		return nil, "", domain.ErrBadParamInput
	}

	res, err = m.fetch(ctx, query, decodedCursor, num)
	if err != nil {
		return nil, "", err
	}

	if len(res) == int(num) {
		nextCursor = repository.EncodeCursor(res[len(res)-1].CreatedAt)
	}
	return
}

func (m *mysqlCarRepository) GetByID(ctx context.Context, id int64) (res domain.Car, err error) {
	query := `SELECT id, name, brand, price, condition, quantity, description, specification, image_url, updated_at, created_at FROM car WHERE ID = ?`

	list, err := m.fetch(ctx, query, id)
	if err != nil {
		return domain.Car{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}

	return
}

func (m *mysqlCarRepository) GetByName(ctx context.Context, title string) (res domain.Car, err error) {
	query := `SELECT id, name, brand, price, condition, quantity, description, specification, image_url, updated_at, created_at FROM car WHERE name = ?`

	list, err := m.fetch(ctx, query, title)
	if err != nil {
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, domain.ErrNotFound
	}
	return
}

func (m *mysqlCarRepository) Store(ctx context.Context, a *domain.Car) (err error) {
	query := `INSERT car SET name=? , brand=?, price=?, condition=?, quantity=?, description=?, specification=?, image_url=?, updated_at=?, created_at=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, a.Name, a.Brand, a.Price, a.Condition, a.Quantity, a.Description, a.Specification, a.ImageURL, a.UpdatedAt, a.CreatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	a.ID = lastID

	return
}

func (m *mysqlCarRepository) Delete(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM car WHERE id = ?"

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
func (m *mysqlCarRepository) Update(ctx context.Context, ar *domain.Car) (err error) {
	query := `UPDATE car SET name=? , brand=?, price=?, condition=?, quantity=?, description=?, specification=?, image_url=?, updated_at=?, created_at=? WHERE ID = ?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, ar.Name, ar.Brand, ar.Price, ar.Condition, ar.Quantity, ar.Description, ar.Specification, ar.ImageURL, ar.UpdatedAt, ar.ID)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("Weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}
