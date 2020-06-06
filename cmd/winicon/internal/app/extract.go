package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/leaanthony/clir"
	"github.com/leaanthony/winicon"
)

func extract(cli *clir.Cli) *clir.Command {
	cmd := cli.NewSubCommand("x", "Extract images from .ico")
	cmd.Action(func() error {
		cli.PrintBanner()

		args := cmd.OtherArgs()
		if len(args) != 1 {
			return fmt.Errorf("usage: %s extract [flags] infile", os.Args[0])
		}

		file, err := os.Open(args[0])
		if err != nil {
			return err
		}
		defer file.Close()

		icons, err := winicon.GetFileData(file)
		if err != nil {
			return err
		}

		badsizes := []string{}

		basefilename := filepath.Base(strings.TrimSuffix(args[0], filepath.Ext(args[0])))
		for _, icon := range icons {
			if icon.Format != "PNG" {
				badsizes = append(badsizes, fmt.Sprintf("%dx%d", icon.Width, icon.Height))
				continue
			}
			// Open file
			outfilename := fmt.Sprintf("%s.%dx%d.%s", basefilename, icon.Width, icon.Height, strings.ToLower(icon.Format))
			outfile, err := os.Create(outfilename)
			if err != nil {
				return err
			}
			// write out image data
			_, err = outfile.Write(icon.Data)
			if err != nil {
				return err
			}

			fmt.Printf("Extracted %dx%d image to file: %s\n", icon.Width, icon.Height, outfilename)

			// Close file
			outfile.Close()
		}

		if len(badsizes) > 0 {
			badsizesText := strings.Join(badsizes, ",")
			println("[Warning] Unable to extract these icons as they are not PNG: ", badsizesText)
		}
		imageText := "image"
		extractedImageCount := len(icons) - len(badsizes)
		if extractedImageCount > 1 {
			imageText = "images"
		}
		fmt.Printf("\nExtracted %d %s\n", extractedImageCount, imageText)

		return nil

	})

	return cmd
}
