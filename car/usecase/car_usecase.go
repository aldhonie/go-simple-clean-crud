package usecase

import (
	"context"
	"time"

	"github.com/aldhonie/go-simple-clean-crud/domain"
)

type carUsecase struct {
	carRepository  domain.CarRepository
	contextTimeout time.Duration
}

//NewCarUsecase will create new carUsecase object representation of domain.CarUsecase interface
func NewCarUsecase(a domain.CarRepository, timeout time.Duration) domain.CarUsecase {
	return &carUsecase{
		carRepository:  a,
		contextTimeout: timeout,
	}
}

func (a *carUsecase) Fetch(c context.Context) (res []domain.Car, nextCursor string, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)

	defer cancel()

	return a.carRepository.Fetch(ctx)
}

func (a *carUsecase) FetchByKeyword(c context.Context, keyword string) (res []domain.Car, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)

	defer cancel()

	return a.carRepository.FetchByKeyword(ctx, keyword)
}

func (a *carUsecase) GetByID(c context.Context, id int64) (res domain.Car, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)

	defer cancel()

	return a.carRepository.GetByID(ctx, id)
}

func (a *carUsecase) Update(c context.Context, dc *domain.Car) (err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)

	defer cancel()

	dc.UpdatedAt = time.Now()

	return a.carRepository.Update(ctx, dc)
}

func (a *carUsecase) GetByName(c context.Context, name string) (res domain.Car, err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)

	defer cancel()

	return a.carRepository.GetByName(ctx, name)
}

func (a *carUsecase) Store(c context.Context, dc *domain.Car) (err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)

	defer cancel()

	existedCar, _ := a.GetByName(ctx, dc.Name)

	if existedCar != (domain.Car{}) {
		return domain.ErrConflict
	}

	dc.UpdatedAt = time.Now()
	dc.CreatedAt = time.Now()

	return a.carRepository.Store(ctx, dc)
}

func (a *carUsecase) Delete(c context.Context, id int64) (err error) {

	ctx, cancel := context.WithTimeout(c, a.contextTimeout)

	defer cancel()

	existedCar, err := a.carRepository.GetByID(ctx, id)

	if err != nil {
		return
	}

	if existedCar == (domain.Car{}) {
		return domain.ErrNotFound
	}

	return a.carRepository.Delete(ctx, id)
}
