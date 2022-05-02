package sopt

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ShowOptions shows the values of all options. Used for debugging.
func (opt *Options) ShowOptions() {
	for _, o := range opt.short {
		fmt.Printf("-%s: %v\n", o.ShortName, o.Value)
	}

	for _, o := range opt.long {
		fmt.Printf("--%s: %v\n", o.LongName, o.Value)
	}
}

// Parse the command line arguments from os.Args. Internally it calls ParseArgs.
// This will force an os.Exit(0) if help is requested, or if no args are given and emptyhelp is true.
func (opt *Options) Parse(emptyhelp bool) error {
	if len(os.Args) == 1 && emptyhelp {
		opt.PrintHelp()
		os.Exit(0)
	}

	return opt.ParseArgs(os.Args[1:])
}

// ParseArgs parses the supplied string slice as CLI arguments.
// It starts by parsing short and long options.
// - Short options start with a single dash ("-").
// - Short boolean options don't need to take a value.
// - Short boolean options require an equal sign ("=") after the option with a truthy or falsy value.
// - Truthy values are "true", "yes", "on", "1", and "t".
// - Falsy values are everything else.
// - Short options can be combined ("-a -b" can be written as "-ab").
// - Combined short options allow only the last one to take a value. The ones before must be booleans.
//
// - Long options start with a double dash ("--").
// - Long options are followed by either whitespace or an equal sign ("--foo bar" or "--foo=bar").
func (opt *Options) ParseArgs(args []string) error {
	unknown := []string{}

	//
	// Long options
	//

	for i, arg := range os.Args[1:] {
		if len(arg) < 2 {
			unknown = append(unknown, arg)
			continue
		}

		if arg[0] == '-' && arg[1] == '-' {
			arg = arg[2:]
			if arg == "" {
				return ErrEmptyLong
			}

			a := splitOption(arg)
			o, ok := opt.long[a[0]]
			if ok {
				switch o.Type {
				case VarTypeBool:
					t, v := isTruthy(a[1])
					// We have the form "--option=value"
					if t {
						o.Value = v
						continue
					}

					if len(args) > i+1 {
						t, v = isTruthy(args[i+1])
						// We have the form "--option value"
						if t {
							o.Value = v
							continue
						}
					}

					// It's a standalone boolean option, so just set it to teue. Phew!
					o.Value = true

				case VarTypeString:
					if a[1] != "" {
						o.Value = a[1]
						continue
					}

					if len(args) > i+1 && args[i+1][0] != '-' {
						o.Value = args[i+1]
						continue
					}

					return fmt.Errorf("--%s: %w", o.LongName, ErrMissingArgument)

				case VarTypeInt:
					if a[1] != "" {
						v, err := strconv.Atoi(a[1])
						if err != nil {
							return err
						}

						o.Value = v
						continue
					}

					// Since we allow negative numbers, we don't skip args starting with a minus.
					if len(args) > i+1 {
						v, err := strconv.Atoi(args[i+1])
						if err != nil {
							return err
						}

						o.Value = v
						continue
					}

					return fmt.Errorf("--%s: %w", o.LongName, ErrMissingArgument)

				case VarTypeFloat:
					if a[1] != "" {
						v, err := strconv.ParseFloat(a[1], 64)
						if err != nil {
							return err
						}

						o.Value = v
						continue
					}

					if len(args) > i+1 {
						v, err := strconv.ParseFloat(args[i+1], 64)
						if err != nil {
							return err
						}

						o.Value = v
						continue
					}

					return fmt.Errorf("--%s: %w", o.LongName, ErrMissingArgument)

				default:
					return fmt.Errorf("--%s: %w", o.LongName, ErrUnknownType)
				} // switch o.Type
			} else {
				return ErrUnknownOption
			} // if long option is defined
			continue
		} // if long option

		//
		// Short options
		//

		if arg[0] == '-' {
			s := arg[1:]
			a := splitOption(s)
			if a[1] != "" {
				s = a[0]
			}

			for _, c := range s {
				o, ok := opt.short[string(c)]
				if ok {
					switch o.Type {
					case VarTypeBool:
						if a[0] == string(c) && a[1] != "" {
							_, v := isTruthy(a[1])
							o.Value = v
							continue
						}

						if len(args) > i+1 {
							t, v := isTruthy(args[i+1])
							if t {
								o.Value = v
								continue
							}
						}

						o.Value = true

					case VarTypeString:
						if a[0] == string(c) && a[1] != "" {
							o.Value = a[1]
							continue
						}

						if len(args) > i+1 {
							o.Value = args[i+1]
							continue
						}

						return fmt.Errorf("-%c: %w", c, ErrMissingArgument)

					case VarTypeInt:
						if a[0] == string(c) && a[1] != "" {
							v, err := strconv.Atoi(a[1])
							if err != nil {
								return err
							}

							o.Value = v
							continue
						}

						if len(args) > i+1 {
							v, err := strconv.Atoi(args[i+1])
							if err != nil {
								return err
							}

							o.Value = v
							continue
						}
					} // switch o.Type
				} else {
					return ErrUnknownOption
				} // if short option is defined
			} // range s
			continue
		} // if short option

		unknown = append(unknown, arg)
	}

	// Parse the remaining unknown arguments as positional arguments, if applicable.
	for _, arg := range unknown {
		fmt.Printf("positional arg: %s\n", arg)
		unknown = unknown[1:]
	}
	opt.Remainder = unknown
	return nil
}

func splitOption(arg string) []string {
	a := strings.SplitN(arg, "=", 2)
	if len(a) == 1 {
		return []string{arg, ""}
	}

	return a
}

// isTruthy returns whether the supplied string is a truthy value.
// The second value is the decoded value, if applicable, false otherwise.
func isTruthy(s string) (bool, bool) {
	switch strings.ToLower(s) {
	case "true", "yes", "on", "1", "t":
		return true, true
	case "false", "no", "off", "0", "f":
		return true, false
	}

	return false, false
}
