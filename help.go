package sopt

import (
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"
)

// SetDefaultHelp sets the default help text.
func (opt *Options) SetDefaultHelp() {
	opt.SetOption("", "h", "help", "Print this help message", false, false, VarTypeBool, nil)
	opt.hashelp = true
}

// PrintHelp builds and prints the help text based on available options.
func (opt *Options) PrintHelp() {
	w := &tabwriter.Writer{}
	w.Init(os.Stdout, 8, 8, 1, '\t', 0)
	w.Write([]byte("Usage:\n  "))
	name := filepath.Base(os.Args[0])
	w.Write([]byte(name))

	count := 0
	for _, g := range opt.groups {
		count += len(g.options)
	}

	if count > 0 {
		w.Write([]byte(" [OPTIONS]"))
	}

	if len(opt.commands) > 0 {
		w.Write([]byte(" [COMMAND]"))
	}

	w.Write([]byte("\n\n"))

	for _, g := range opt.groups {
		if len(g.options) > 0 {
			if g.Name == "default" {
				w.Write([]byte("Main options:\n"))
			} else {
				w.Write([]byte(fmt.Sprintf("%s options:\n", g.Name)))
			}

			for _, o := range g.options {
				if o.ShortName != "" && o.LongName != "" {
					w.Write([]byte(fmt.Sprintf("\t-%s,--%s\t%s\n", o.ShortName, o.LongName, o.Help)))
					continue
				}
			}
			w.Write([]byte("\n"))
		}

		if len(g.commands) > 0 {
			if g.Name == "default" {
				w.Write([]byte("Main commands:\n"))
			} else {
				w.Write([]byte(fmt.Sprintf("%s commands:\n", g.Name)))
			}

			for _, cmd := range g.commands {
				w.Write([]byte(fmt.Sprintf("\t%s\t%s", cmd, opt.commands[cmd].Help)))
				if len(opt.commands[cmd].Aliases) > 0 {
					w.Write([]byte(fmt.Sprintf(" (aliases: ")))
					for i, alias := range opt.commands[cmd].Aliases {
						if i == 0 {
							w.Write([]byte(fmt.Sprintf("%s", alias)))
						} else {
							w.Write([]byte(fmt.Sprintf(",%s", alias)))
						}
					}
					w.Write([]byte(fmt.Sprintf(")\n")))
				}
			}
			w.Write([]byte("\n"))
		}
	}

	w.Flush()
}
