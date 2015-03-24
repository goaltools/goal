// Package command is used for parsing input parameters.
package command

import ()

// IncorrectArgsErr is returned every time a user is trying
// to use input parameters we do not expect.
var IncorrectArgsErr error

// Type is a main type of command package.
// It is used for storage of parsed parameters.
type Type struct {
	action string
	params map[string]string
}

// NewType initializes and returns Type object. It expects even number of args.
// Otherwise, it an IncorIncorrectArgsErr error will be returned.
func NewType(args []string) (*Type, error) {
	// Make sure the number of arguments is even number
	// and it is more than zero.
	if len(args) == 0 || len(args)%2 != 0 {
		return nil, IncorrectArgsErr
	}

	// Save the arguments as dict.
	params := map[string]string{}
	for i := 0; i < len(args); i += 2 {
		params[args[i]] = args[i+1]
	}

	// The first argument is a requested action.
	return &Type{
		action: args[0],
		params: params,
	}, nil
}
