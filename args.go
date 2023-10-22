/*
 * Copyright (c) Brandon Jordan
 */

package main

import "github.com/electrikmilk/args-parser"

func init() {
	args.CustomUsage = "[FILE]"
	args.Register(args.Argument{
		Name:        "version",
		Short:       "v",
		Description: "Print version information.",
	})
	args.Register(args.Argument{
		Name:        "help",
		Short:       "h",
		Description: "Print this usage information.",
	})
	args.Register(args.Argument{
		Name:         "action",
		Description:  "Print action definition. Leave empty to print all action definitions.",
		DefaultValue: "",
		ExpectsValue: true,
	})
	args.Register(args.Argument{
		Name:         "share",
		Short:        "s",
		Description:  "Shortcuts signing mode, passed to the `shortcuts` binary.",
		Values:       []string{"anyone", "contacts"},
		DefaultValue: "contacts",
		ExpectsValue: true,
	})
	args.Register(args.Argument{
		Name:        "debug",
		Short:       "d",
		Description: "Save generated plist. Print debug messages and stack traces.",
	})
	args.Register(args.Argument{
		Name:        "output",
		Short:       "o",
		Description: "Optional output file path. (e.g. /path/to/file.shortcut).",
	})
	args.Register(args.Argument{
		Name:        "import",
		Short:       "i",
		Description: "Opens compiled Shortcut after compilation. Ignored if unsigned.",
	})
	args.Register(args.Argument{
		Name:        "auto-inc",
		Short:       "a",
		Description: "Automatically include Cherri files in this directory.",
	})
	args.Register(args.Argument{
		Name:        "no-comments",
		Short:       "c",
		Description: "Don't include comments in the compiled Shortcut.",
	})
	args.Register(args.Argument{
		Name:        "no-ansi",
		Description: "Don't output ANSI escape sequences that format and color the output.",
	})
}