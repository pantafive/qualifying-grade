package gradelogentry

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
)

func TestRecognizerExportTo(t *testing.T) {
	var actualExporter GradeLogEntryExporter
	f := NewGradeLogEntryFakeFactory()
	agg, _ := f.Create()
	agg.ExportTo(&actualExporter)
	assert.Equal(t, GradeLogEntryExporter{
		EndorsedId:      member.NewTenantMemberIdExporter(f.EndorsedId.TenantId, f.EndorsedId.MemberId),
		EndorsedVersion: f.EndorsedVersion,
		AssignedGrade:   seedwork.NewUint8Exporter(f.AssignedGrade),
		Reason:          seedwork.NewStringExporter(f.Reason),
		CreatedAt:       f.CreatedAt,
	}, actualExporter)
}
