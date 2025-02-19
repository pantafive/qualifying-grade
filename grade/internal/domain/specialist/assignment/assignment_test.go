package assignment

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/emacsway/grade/grade/internal/domain/member"
	"github.com/emacsway/grade/grade/internal/domain/seedwork/exporters"
)

func TestAssignmentExport(t *testing.T) {
	var actualExporter AssignmentExporter
	f := NewAssignmentFakeFactory()
	agg, _ := f.Create()
	agg.Export(&actualExporter)
	assert.Equal(t, AssignmentExporter{
		SpecialistId:      member.NewTenantMemberIdExporter(f.SpecialistId.TenantId, f.SpecialistId.MemberId),
		SpecialistVersion: f.SpecialistVersion,
		AssignedGrade:     exporters.Uint8Exporter(f.AssignedGrade),
		Reason:            exporters.StringExporter(f.Reason),
		CreatedAt:         f.CreatedAt,
	}, actualExporter)
}
