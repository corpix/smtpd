// The MIT License (MIT)

// Copyright © 2017 Dmitry Moskowski

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
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
