package flagargs

// Parser handles passed non-flag arguments and updates flags with positional args
type Parser struct {
	knownArgs      []string
	extraArgsStart int
}

// NewParser returns a new Parser with the passed knownArgs
func NewParser(knownArgs ...string) *Parser {
	return &Parser{
		knownArgs:      knownArgs,
		extraArgsStart: len(knownArgs),
	}
}

// SplitArgs is a helper that creates a parser and splits the arguments in one call
func SplitArgs(args []string, knownArgs ...string) (known []string, extra []string) {
	parser := NewParser(knownArgs...)
	doubleDashPos := LocateDoubleDash(args)
	return parser.SplitArgs(args, doubleDashPos)
}

// SplitArgs simply splits the arguments into known and extra arguments
func (p *Parser) SplitArgs(args []string, doubleDashPos int) (known []string, extra []string) {
	ddash := doubleDashPos
	if ddash >= 0 && ddash < len(p.knownArgs) {
		// if `--` is supplied, treat all subsequent args as extra args
		p.extraArgsStart = ddash
	}
	if p.extraArgsStart >= len(args) {
		// if no extra args are passed, set the start to the last index to produce a 0-length slice
		p.extraArgsStart = len(args)
	}

	return args[:p.extraArgsStart], args[p.extraArgsStart:]
}

// ParseAndUpdateFlags splits the args between known- and extra args
// the known args are set on flags using the corresponding key passed in NewParser
func (p *Parser) ParseAndUpdateFlags(flags FlagSet, args []string) (extraArgs []string, err error) {
	var doubleDashPos int
	if flagsWithDash, ok := flags.(FlagSetWithDash); ok {
		doubleDashPos = flagsWithDash.ArgsLenAtDash()
	} else {
		doubleDashPos = LocateDoubleDash(args)
	}
	known, extra := p.SplitArgs(args, doubleDashPos)
	for i, arg := range known {
		err = flags.Set(p.knownArgs[i], arg)
		if err != nil {
			return
		}
	}
	return extra, nil
}

// LocateDoubleDash finds the first instance of "--" if present in the passed args
func LocateDoubleDash(args []string) int {
	for i, a := range args {
		if a == "--" {
			return i
		}
	}
	return -1
}
