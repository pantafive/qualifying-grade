package aggregate

type AggregateVersioner interface {
	Version() uint
	SetVersion(uint)
}

func NewVersionedAggregate(version uint) VersionedAggregate {
	return VersionedAggregate{version: version}
}

type VersionedAggregate struct {
	version uint
}

func (a VersionedAggregate) Version() uint {
	return a.version
}

func (a *VersionedAggregate) SetVersion(val uint) {
	a.version = val
}
