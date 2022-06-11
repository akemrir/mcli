package mcli

import "flag"

type parseOptions struct {
	cmdName       *string
	args          *[]string
	errorHandling flag.ErrorHandling

	disableGlobalFlags bool
}

// ParseOpt specifies options to customize the behavior of Parse.
type ParseOpt func(*parseOptions)

// WithArgs tells Parse to parse from the given args, instead of
// parsing from the command line arguments.
func WithArgs(args []string) ParseOpt {
	return func(options *parseOptions) {
		options.args = &args
	}
}

// WithErrorHandling tells Parse to use the given ErrorHandling.
// By default, Parse exits the program when an error happens.
func WithErrorHandling(h flag.ErrorHandling) ParseOpt {
	return func(options *parseOptions) {
		options.errorHandling = h
	}
}

// WithName specifies the command name to use when printing usage doc.
func WithName(name string) ParseOpt {
	return func(options *parseOptions) {
		options.cmdName = &name
	}
}

// DisableGlobalFlags tells Parse to don't parse and print global flags in help.
func DisableGlobalFlags() ParseOpt {
	return func(options *parseOptions) {
		options.disableGlobalFlags = true
	}
}
