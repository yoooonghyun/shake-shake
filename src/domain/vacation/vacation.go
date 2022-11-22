package domain

import (
	"shake-shake/src/base"
	"shake-shake/src/event"
	"time"

	"github.com/google/uuid"
)

type VacationType string

const (
	VacationTypeDayOff       VacationType = "day-off"
	VacationTypeAfternoonOff VacationType = "afternoon-off"
	VacationTypeMorningOff   VacationType = "morning-off"
)

type VacationState string

const (
	VacationStateCreated  VacationState = "created"
	VacationStateCanceled VacationState = "canceled"
)

type Vacation struct {
	event.EventRecorder
	Id              string        `bson:"id" json:"id"`
	MemberId        string        `bson:"memberId" json:"memberId"`
	CreatedAt       time.Time     `bson:"createdAt" json:"createdAt"`
	VacationStartAt time.Time     `bson:"vacationAt" json:"vacationStartAt"`
	VacationEndAt   time.Time     `bson:"vacationAt" json:"vacationEndAt"`
	Hours           int           `bson:"hours" json:"hours"`
	State           VacationState `bson:"state" json:"state"`
}

func CreateVacation(
	memberId string,
	vacationStartAt time.Time,
	vacationEndAt time.Time,
	hours int,
) (*Vacation, error) {
	vacation := Vacation{
		Id:              uuid.New().String(),
		MemberId:        memberId,
		CreatedAt:       time.Now(),
		VacationStartAt: vacationStartAt,
		VacationEndAt:   vacationEndAt,
		Hours:           hours,
		State:           VacationStateCreated,
	}

	if err := vacation.vaild(); err != nil {
		return nil, err
	}

	vacation.AddEvent(VacationCreated{Id: vacation.Id})
	return &vacation, nil
}

func (v *Vacation) Cancel() *Vacation {
	v.State = VacationStateCanceled
	v.AddEvent(VacationCanceled{Id: v.Id})
	return v
}

func (v *Vacation) vaild() error {
	if v.Id == "" {
		return base.ErrInvalidValue
	}
	return nil
}
