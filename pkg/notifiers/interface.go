package notifiers

// Notifier is a notifier.
type Notifier interface {
	Notify(string) error
	Type() string
	Vendor() string
}
