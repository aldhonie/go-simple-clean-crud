package mysql

import (
	"context"
	"database/sql"
	"fmt"

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
			&t.Image,
			&t.UpdatedAt,
			&t.CreatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)

	}

	return result, nil
}

func (m *mysqlCarRepository) Fetch(ctx context.Context) (res []domain.Car, nextCursor string, err error) {
	query := `SELECT * FROM car`

	res, err = m.fetch(ctx, query)
	if err != nil {
		return nil, "", err
	}

	return
}

func (m *mysqlCarRepository) FetchByKeyword(ctx context.Context, keyword string) (res []domain.Car, err error) {
	query := `SELECT * FROM car WHERE name LIKE ?`

	res, err = m.fetch(ctx, query, `%`+keyword+`%`)
	if err != nil {
		return nil, err
	}

	return
}

func (m *mysqlCarRepository) GetByID(ctx context.Context, id int64) (res domain.Car, err error) {
	query := `SELECT * FROM car WHERE id=?`

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
	query := `SELECT * FROM car WHERE name = ?`

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

func (m *mysqlCarRepository) Store(ctx context.Context, c *domain.Car) (err error) {

	query := `INSERT INTO car SET name=? , brand=? , price=? , kondisi=? , quantity=? , description=? , specification=? , image=? , updated_at=? , created_at=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, c.Name, c.Brand, c.Price, c.Condition, c.Quantity, c.Description, c.Specification, c.Image, c.UpdatedAt, c.CreatedAt)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	c.ID = lastID

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
		err = fmt.Errorf("Anomali. Total Affected: %d", rowsAfected)
		return
	}

	return
}
func (m *mysqlCarRepository) Update(ctx context.Context, c *domain.Car) (err error) {

	query := `UPDATE car SET name=? , brand=?, price=?, kondisi=?, quantity=?, description=?, specification=?, image=?, updated_at=? WHERE id=?`

	stmt, err := m.Conn.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, c.Name, c.Brand, c.Price, c.Condition, c.Quantity, c.Description, c.Specification, c.Image, c.UpdatedAt, c.ID)
	if err != nil {
		return
	}
	logrus.Println(res)
	affect, err := res.RowsAffected()

	if err != nil {
		return
	}

	if affect != 1 {
		err = fmt.Errorf("Anomali. Total Affected: %d", affect)
		return
	}

	return
}
