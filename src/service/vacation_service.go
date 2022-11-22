package service

import (
	"context"
	domain "shake-shake/src/domain/vacation"
	store "shake-shake/src/repository"
)

const kCollectionName = "vacation"

type VacationService struct {
	model domain.Vacation
	db    store.Store
}

func CreateVacationService() (*VacationService, error) {
	db, err := store.GetStore(&domain.Vacation{})

	if err != nil {
		return nil, err
	}

	return &VacationService{db: db, model: domain.Vacation{}}, nil
}

func (svc *VacationService) Create(ctx context.Context, v *domain.Vacation) error {
	vacation, err := domain.CreateVacation(v.MemberId, v.VacationStartAt, v.VacationEndAt, v.Hours)

	if err != nil {
		return err
	}

	return svc.db.Insert(ctx, svc.model, vacation)
}

func (svc *VacationService) ReadMany(ctx context.Context) ([]*domain.Vacation, error) {
	var result []*domain.Vacation
	var query = domain.Vacation{}
	err := svc.db.FindAll(ctx, svc.model, query, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc *VacationService) ReadOne(ctx context.Context, id string) (*domain.Vacation, error) {
	var (
		result = &domain.Vacation{}
		query  = domain.Vacation{Id: id}
	)

	err := svc.db.FindOne(ctx, svc.model, query, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc *VacationService) Delete(ctx context.Context, id string) error {

	err := svc.db.Delete(ctx, svc.model, &domain.Vacation{Id: id})

	return err
}
