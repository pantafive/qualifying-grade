package artifact

import (
	"time"

	"github.com/emacsway/qualifying-grade/grade/internal/domain/expertisearea"
	"github.com/emacsway/qualifying-grade/grade/internal/domain/member"
)

// Artifact is a good candidate for EventSourcing
type Artifact struct {
	id               TenantArtifactId
	status           Status
	name             Name
	description      Description
	url              Url
	expertiseAreaIds []expertisearea.ExpertiseAreaId
	authorIds        []member.TenantMemberId
	createdBy        member.TenantMemberId
	createdAt        time.Time
}
