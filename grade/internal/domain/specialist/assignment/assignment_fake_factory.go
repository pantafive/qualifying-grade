package assignment

import (
	"time"

	"github.com/emacsway/grade/grade/internal/domain/grade"
	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/uuid"
)

func NewAssignmentFakeFactory() AssignmentFakeFactory {
	idFactory := member.NewTenantMemberIdFakeFactory()
	idFactory.MemberId = uuid.ParseSilent("cf9462cf-51d3-4c0a-b5ef-53b3dfccc7f6")
	return AssignmentFakeFactory{
		SpecialistId:      idFactory,
		SpecialistVersion: 2,
		AssignedGrade:     1,
		Reason:            "Any",
		CreatedAt:         time.Now(),
	}
}

type AssignmentFakeFactory struct {
	SpecialistId      member.TenantMemberIdFakeFactory
	SpecialistVersion uint
	AssignedGrade     uint8
	Reason            string
	CreatedAt         time.Time
}

func (f AssignmentFakeFactory) Create() (Assignment, error) {
	specialistId, err := member.NewTenantMemberId(f.SpecialistId.TenantId, f.SpecialistId.MemberId)
	if err != nil {
		return Assignment{}, err
	}
	assignedGrade, err := grade.DefaultConstructor(f.AssignedGrade)
	if err != nil {
		return Assignment{}, err
	}
	reason, err := NewReason(f.Reason)
	if err != nil {
		return Assignment{}, err
	}
	return NewAssignment(specialistId, f.SpecialistVersion, assignedGrade, reason, f.CreatedAt)
}
