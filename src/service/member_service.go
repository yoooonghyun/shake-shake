package service

import (
	"context"
	domain "shake-shake/src/domain/member"
	store "shake-shake/src/repository"
)

type MemberService struct {
	model domain.Member
	db    store.Store
}

func CreateMemberService() (*MemberService, error) {
	db, err := store.GetStore(&domain.Member{})

	if err != nil {
		return nil, err
	}

	return &MemberService{db: db, model: domain.Member{}}, nil
}

func (svc *MemberService) Create(ctx context.Context, m *domain.Member) error {
	member, err := domain.CreateMember(
		m.Name,
		m.DepartmentId,
	)

	if err != nil {
		return err
	}

	return svc.db.Insert(ctx, svc.model, member)
}

func (svc *MemberService) ReadMany(ctx context.Context) ([]*domain.Member, error) {
	var result []*domain.Member
	var query = domain.Member{}
	err := svc.db.FindAll(ctx, svc.model, query, &result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc *MemberService) ReadOne(ctx context.Context, id string) (*domain.Member, error) {
	var (
		result = &domain.Member{}
		query  = domain.Member{Id: id}
	)

	err := svc.db.FindOne(ctx, svc.model, query, result)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (svc *MemberService) ShakeShake(ctx context.Context) ([]*domain.Member, error) {
	var result []*domain.Member
	var query = domain.Member{}
	err := svc.db.FindAll(ctx, svc.model, query, &result)

	return result, err
}
