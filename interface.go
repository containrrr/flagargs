package flagargs

// FlagSet is the interface used for setting flag values
type FlagSet interface {
	Set(name, value string) error
}

// FlagSetWithDash is the interface used for locating the split between known- and extra args
type FlagSetWithDash interface {
	FlagSet
	ArgsLenAtDash() int
}
