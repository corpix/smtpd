package smtpd


import (
	e "errors"
)

var (
	// ErrCmdContainsNonASCII indicates that received command line contains
	// non 7-bit ASCII characters.
	ErrCmdContainsNonASCII = e.New("command contains non 7-bit ASCII")

	// ErrCmdUnrecognized indicates that we failed to extract the actual
	// SMTP command name from the command line.
	// Probably reason is non-RFC compliant format of the command line.
	ErrCmdUnrecognized = e.New("unrecognized command")

	// ErrCmdHasNoArg indicates that we have found arguments while this
	// command should not receive any arguments.
	ErrCmdHasNoArg = e.New("SMTP command does not take an argument")

	// ErrCmdRequiresArg indicates that this command has one argument,
	// but we received this command without the argument.
	ErrCmdRequiresArg = e.New("SMTP command requires an argument")

	// ErrCmdRequiresArgs indicates that this command has one argument,
	// but we received this command without the argument.
	ErrCmdRequiresArgs = e.New("SMTP command requires an arguments")

	// ErrCmdRequiresAddress indicates that address is required for this command.
	ErrCmdRequiresAddress = e.New("SMTP command requires an address")

	// ErrCmdImpropperArgFmt indicates that command arguments has improper format.
	ErrCmdImpropperArgFmt = e.New("improper argument formatting")
)
