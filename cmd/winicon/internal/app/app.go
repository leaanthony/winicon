package app

import (
	"fmt"
	"io"

	"github.com/leaanthony/clir"
)

// VERSION of the application
const VERSION = "v0.0.1"

// Run is the main entry point for the application
func Run(args []string, w io.Writer) error {

	app := clir.NewCli("WinIcon", "A utility for manipulating windows .ico files", VERSION)

	// Custom Banner
	app.SetBannerFunction(func(cli *clir.Cli) string {
		return fmt.Sprintf("%s %s  (c) 2020-Present Lea Anthony", cli.Name(), cli.Version())
	})

	// Add commands
	info(app)
	generate(app)
	extract(app)

	// Run!
	return app.Run()
}
