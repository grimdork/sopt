package sopt

import "errors"

var (
	// ErrMissingRequired is returned when a required option is missing.
	ErrMissingRequired = errors.New("missing required option")
	// ErrMissingArgument is returned when an option is missing an argument.
	ErrMissingArgument = errors.New("missing argument")
	// ErrLongShort is returned when a short option is longer than one character.
	ErrLongShort = errors.New("short option must be one character")
	// ErrUnknownOption is returned when an undefined option is encountered.
	ErrUnknownOption = errors.New("unknown option")
	// ErrEmptyLong is returned when a long option is empty.
	ErrEmptyLong = errors.New("long option without a string")
	// ErrUnknownType is returned when an unknown option variable type is encountered.
	ErrUnknownType = errors.New("unknown option type")
)
