package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/leaanthony/clir"
	"github.com/leaanthony/winicon"
)

func generate(cli *clir.Cli) *clir.Command {
	cmd := cli.NewSubCommand("g", "Generate .ico file from image")
	sizes := "256"
	cmd.StringFlag("sizes", "Icon sizes to generate eg: -sizes 16,24,64", &sizes)
	cmd.Action(func() error {
		cli.PrintBanner()

		args := cmd.OtherArgs()
		if len(args) != 1 {
			return fmt.Errorf("usage: %s generate [flags] infile", os.Args[0])
		}

		// Check sizes
		var actualSizes []int

		for _, size := range strings.Split(sizes, ",") {
			parsedSize, err := strconv.Atoi(size)
			if err != nil {
				return fmt.Errorf("'%s' is not a valid size", size)
			}
			actualSizes = append(actualSizes, parsedSize)
		}

		infile, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer infile.Close()

		outfilename := filepath.Base(strings.TrimSuffix(args[0], filepath.Ext(args[0])) + ".ico")
		outfile, err := os.Create(outfilename)
		if err != nil {
			return err
		}

		err = winicon.GenerateIcon(infile, outfile, actualSizes)
		if err != nil {
			return err
		}

		fmt.Printf("Generated icon file: %s at sizes %v\n", outfilename, actualSizes)

		return nil

	})

	return cmd
}
