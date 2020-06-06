package app

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/leaanthony/winicon"

	"github.com/leaanthony/clir"
)

// Info contains information about a .ico file
type Info struct {
	Filename      string
	FileSize      int64
	NumberOfIcons int
	Icons         []*winicon.Icon
}

func info(cli *clir.Cli) *clir.Command {
	cmd := cli.NewSubCommand("i", "Show .ico file information")
	var outputJSON bool = false
	cmd.BoolFlag("json", "Get file information in json format", &outputJSON)
	cmd.Action(func() error {
		args := cmd.OtherArgs()
		if len(args) != 1 {
			return fmt.Errorf("usage: %s info [flags] filename", os.Args[0])
		}

		// Store the full filename
		filename := args[0]
		fullFilename, err := filepath.Abs(filename)
		if err != nil {
			return err
		}

		result := Info{
			Filename: fullFilename,
		}

		// Read file
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		// Get length
		stat, err := file.Stat()
		if err != nil {
			return err
		}
		result.FileSize = stat.Size()

		icons, err := winicon.GetFileData(file)
		if err != nil {
			return err
		}

		result.Icons = icons
		result.NumberOfIcons = len(icons)

		if outputJSON == true {
			jsonData, err := json.Marshal(result)
			if err != nil {
				return err
			}
			println(string(jsonData))
			return nil
		}

		cli.PrintBanner()
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 0, '\t', 0)

		fmt.Fprintln(w, "Filename:\t", result.Filename)
		fmt.Fprintln(w, "File Size:\t", result.FileSize, "bytes")
		fmt.Fprintln(w, "Icon Count:\t", result.NumberOfIcons)
		fmt.Fprintln(w)

		for index, icon := range result.Icons {
			fmt.Fprintf(w, "Icon %d:\tSize: %dx%d\tFormat: %s\tBits Per Pixel: %d\tOffset: %d\n", index+1, icon.Width, icon.Height, icon.Format, icon.BitsPerPixel, icon.Offset)
		}
		w.Flush()

		return nil

	})

	return cmd
}
