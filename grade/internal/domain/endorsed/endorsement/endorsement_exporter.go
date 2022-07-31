package endorsement

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/artifact"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/seedwork"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/shared"
)

type EndorsementExporter struct {
	RecognizerId      member.TenantMemberIdExporter
	RecognizerGrade   seedwork.Uint8Exporter
	RecognizerVersion uint
	EndorsedId        member.TenantMemberIdExporter
	EndorsedGrade     seedwork.Uint8Exporter
	EndorsedVersion   uint
	ArtifactId        seedwork.Uint64Exporter
	CreatedAt         time.Time
}

func (ex *EndorsementExporter) SetRecognizerId(val member.TenantMemberId) {
	val.Export(&ex.RecognizerId)
}

func (ex *EndorsementExporter) SetRecognizerGrade(val shared.Grade) {
	val.Export(&ex.RecognizerGrade)
}

func (ex *EndorsementExporter) SetRecognizerVersion(val uint) {
	ex.RecognizerVersion = val
}

func (ex *EndorsementExporter) SetEndorsedId(val member.TenantMemberId) {
	val.Export(&ex.EndorsedId)
}

func (ex *EndorsementExporter) SetEndorsedGrade(val shared.Grade) {
	val.Export(&ex.EndorsedGrade)
}

func (ex *EndorsementExporter) SetEndorsedVersion(val uint) {
	ex.EndorsedVersion = val
}

func (ex *EndorsementExporter) SetArtifactId(val artifact.ArtifactId) {
	val.Export(&ex.ArtifactId)
}

func (ex *EndorsementExporter) SetCreatedAt(val time.Time) {
	ex.CreatedAt = val
}
