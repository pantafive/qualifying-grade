package endorsed

import (
	"errors"
	"time"

	"github.com/hashicorp/go-multierror"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/endorsement"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/endorsed/gradelogentry"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/grade"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/recognizer"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

var (
	ErrCrossTenantEndorsement = errors.New(
		"recognizer can't endorse cross-tenant members",
	)
	ErrEndorsementOneself = errors.New(
		"recognizer can't endorse himself",
	)
	ErrLowerGradeEndorses = errors.New(
		"it is allowed to receive endorsements only from members with equal or higher grade",
	)
	ErrAlreadyEndorsed = errors.New(
		"this artifact has already been endorsed by the recognizer",
	)
)

// FIXME: Move this constructor to tenant aggregate
func NewEndorsed(
	id member.TenantMemberId,
	createdAt time.Time,
) (*Endorsed, error) {
	versioned := seedwork.NewVersionedAggregate(0)
	eventive := seedwork.NewEventiveEntity()
	zeroGrade, _ := grade.DefaultConstructor(0)
	return &Endorsed{
		id:                 id,
		grade:              zeroGrade,
		VersionedAggregate: versioned,
		EventiveEntity:     eventive,
		createdAt:          createdAt,
	}, nil
}

type Endorsed struct {
	id                   member.TenantMemberId
	grade                grade.Grade
	receivedEndorsements []endorsement.Endorsement
	gradeLogEntries      []gradelogentry.GradeLogEntry
	createdAt            time.Time
	seedwork.VersionedAggregate
	seedwork.EventiveEntity
}

func (e *Endorsed) ReceiveEndorsement(r recognizer.Recognizer, aId artifact.ArtifactId, t time.Time) error {
	err := e.canReceiveEndorsement(r, aId)
	if err != nil {
		return err
	}
	ent, err := endorsement.NewEndorsement(
		r.GetId(), r.GetGrade(), r.GetVersion(),
		e.id, e.grade, e.GetVersion(),
		aId, t,
	)
	if err != nil {
		return err
	}
	e.receivedEndorsements = append(e.receivedEndorsements, ent)
	err = e.actualizeGrade(t)
	if err != nil {
		return err
	}
	return nil
}

func (e Endorsed) canReceiveEndorsement(r recognizer.Recognizer, aId artifact.ArtifactId) error {
	err := r.CanCompleteEndorsement()
	if err != nil {
		return err
	}
	return e.canBeEndorsed(r, aId)
}

func (e Endorsed) canBeEndorsed(r recognizer.Recognizer, aId artifact.ArtifactId) error {
	var errs error
	if !r.GetId().TenantId().Equal(e.id.TenantId()) {
		errs = multierror.Append(errs, ErrCrossTenantEndorsement)
	}
	if r.GetId().Equal(e.id) {
		errs = multierror.Append(errs, ErrEndorsementOneself)
	}
	if r.GetGrade().LessThan(e.grade) {
		errs = multierror.Append(errs, ErrLowerGradeEndorses)
	}
	for _, ent := range e.receivedEndorsements {
		if ent.IsEndorsedBy(r.GetId(), aId) {
			errs = multierror.Append(errs, ErrAlreadyEndorsed)
			break
		}
	}
	return errs
}

func (e Endorsed) CanBeginEndorsement(r recognizer.Recognizer, aId artifact.ArtifactId) error {
	err := r.CanReserveEndorsement()
	if err != nil {
		return err
	}
	return e.canBeEndorsed(r, aId)
}

func (e *Endorsed) actualizeGrade(t time.Time) error {
	if e.grade.NextGradeAchieved(e.getReceivedEndorsementCount()) {
		nextGrade, err := e.grade.Next()
		if err != nil {
			return err
		}
		reason, err := gradelogentry.NewReason("Achieved")
		if err != nil {
			return err
		}
		return e.setGrade(nextGrade, reason, t)
	}
	return nil
}
func (e Endorsed) getReceivedEndorsementCount() uint {
	var counter uint
	for _, v := range e.receivedEndorsements {
		if v.GetEndorsedGrade().Equal(e.grade) {
			counter += uint(v.GetWeight())
		}
	}
	return counter
}

func (e *Endorsed) setGrade(g grade.Grade, reason gradelogentry.Reason, t time.Time) error {
	gle, err := gradelogentry.NewGradeLogEntry(
		e.id, e.GetVersion(), g, reason, t,
	)
	if err != nil {
		return err
	}
	e.gradeLogEntries = append(e.gradeLogEntries, gle)
	e.grade = g
	return nil
}

func (e *Endorsed) DecreaseGrade(reason gradelogentry.Reason, t time.Time) error {
	previousGrade, err := e.grade.Next()
	if err != nil {
		return err
	}
	return e.setGrade(previousGrade, reason, t)
}

func (e Endorsed) Export(ex EndorsedExporterSetter) {
	ex.SetId(e.id)
	ex.SetGrade(e.grade)
	ex.SetVersion(e.GetVersion())
	ex.SetCreatedAt(e.createdAt)

	for i := range e.receivedEndorsements {
		ex.AddEndorsement(e.receivedEndorsements[i])
	}
	for i := range e.gradeLogEntries {
		ex.AddGradeLogEntry(e.gradeLogEntries[i])
	}
}

type EndorsedExporterSetter interface {
	SetId(member.TenantMemberId)
	SetGrade(grade.Grade)
	AddEndorsement(endorsement.Endorsement)
	AddGradeLogEntry(gradelogentry.GradeLogEntry)
	SetVersion(uint)
	SetCreatedAt(time.Time)
}
