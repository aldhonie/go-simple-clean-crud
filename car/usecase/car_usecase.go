package usecase

import (
	"context"
	"time"

	"github.com/aldhonie/go-simple-clean-crud/domain"
)

type carUsecase struct {
	carRepo        domain.CarRepository
	contextTimeout time.Duration
}

//NewCarUsecase will create new an  carUsecase object representation of domain.CarUsecase interface
func NewCarUsecase(a domain.CarRepository, timeout time.Duration) domain.CarUsecase {
	return &carUsecase{
		carRepo:        a,
		contextTimeout: timeout,
	}
}

func (a *carUsecase) Fetch(c context.Context, cursor string, num int64) (res []domain.Car, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, nextCursor, err = a.carRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return
}

func (a *carUsecase) GetByID(c context.Context, id int64) (res domain.Car, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.carRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (a *carUsecase) Update(c context.Context, ar *domain.Car) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	ar.UpdatedAt = time.Now()
	return a.carRepo.Update(ctx, ar)
}

func (a *carUsecase) GetByNamw(c context.Context, title string) (res domain.Car, err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()

	res, err = a.carRepo.GetByName(ctx, title)
	if err != nil {
		return
	}

	return
}

func (a *carUsecase) Store(c context.Context, m *domain.Car) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedCar, _ := a.GetByName(ctx, m.Name)
	if existedCar != (domain.Car{}) {
		return domain.ErrConflict
	}

	err = a.carRepo.Store(ctx, m)
	return
}

func (a *carUsecase) Delete(c context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(c, a.contextTimeout)
	defer cancel()
	existedCar, err := a.carRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedCar == (domain.Car{}) {
		return domain.ErrNotFound
	}
	return a.carRepo.Delete(ctx, id)
}
