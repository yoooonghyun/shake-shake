package domain

import (
	"shake-shake/src/base"
	"shake-shake/src/event"

	"github.com/google/uuid"
)

type MemberState string

const (
	MemberStateUnknown      MemberState = "unknown"
	MemberStateNormal       MemberState = "normal"
	MemberStateDayOff       MemberState = "day-off"
	MemberStateAfterNoonOff MemberState = "afternoon-off"
	MemberStateMorningOff   MemberState = "morning-off"
)

type Member struct {
	event.EventRecorder
	Id           string      `bson:"id" json:"id" gorm:"primaryKey"`
	State        MemberState `bson:"state" json:"state"`
	Name         string      `bson:"name" json:"name"`
	DepartmentId string      `bson:"departmentId" json:"departmentId"`
	GroupId      string      `bson:"groupId" json:"groupId"`
	PrevGroupId  string      `bson:"prevGroupId" json:"prevGroupId"`
}

func CreateMember(
	name string,
	departmentId string,
) (*Member, error) {
	member := Member{
		Id:           uuid.New().String(),
		Name:         name,
		DepartmentId: departmentId,
		State:        MemberStateUnknown,
	}
	if err := member.valid(); err != nil {
		return nil, err
	}

	member.AddEvent(MemberCreated{Id: member.Id})
	return &member, nil
}

func (m *Member) UpdateState(
	state MemberState,
) (*Member, error) {
	m.State = state
	if err := m.valid(); err != nil {
		return nil, err
	}
	m.AddEvent(MemberStateChanged{Id: m.Id, State: state})
	return m, nil
}

func (m *Member) valid() error {
	if m.Id == "" {
		return base.ErrInvalidValue
	}
	return nil
}
