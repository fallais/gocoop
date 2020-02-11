package door

// Door operation contract.
type Door interface {
	Open() error
	Close() error
	Stop() error
}
