package winicon

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/leaanthony/winicon/internal/winicon"
)

// Icon stores the data for a single icon
type Icon struct {
	Width        uint16
	Height       uint16
	Colours      uint8
	Planes       uint16
	BitsPerPixel uint16
	Data         []byte `json:"-"`
	Format       string
	Offset       uint32
	size         uint32
	ErrorMessage string `json:"-"`
}

// GetFileData reads in the given .ico filename and returns information
// about the icon/icons
func GetFileData(r io.Reader) ([]*Icon, error) {

	var result []*Icon

	// Parse the .ico file
	var header winicon.IconFileHeader
	err := binary.Read(r, binary.LittleEndian, &header)
	if err != nil {
		return nil, err
	}

	// Loop over icons
	for index := 0; index < (int)(header.ImageCount); index++ {
		// Read in icon headers
		var iconHeader winicon.IconHeader
		err = binary.Read(r, binary.LittleEndian, &iconHeader)
		if err != nil {
			return nil, err
		}
		icon := Icon{
			Width:        (uint16)(iconHeader.Width),
			Height:       (uint16)(iconHeader.Height),
			BitsPerPixel: iconHeader.BitsPerPixel,
			Planes:       iconHeader.Planes,
			Offset:       iconHeader.Offset,
			size:         iconHeader.Size,
		}

		// Width/Height of 256 is encoded as 0 in the icon header
		if icon.Width == 0 {
			icon.Width = 256
		}
		if icon.Height == 0 {
			icon.Height = 256
		}

		result = append(result, &icon)
	}

	// Loop over Icons to read in image data
	for _, icon := range result {
		icon.Data = make([]byte, icon.size)
		println("Allocating bytes to image data:", icon.size)
		_, err := r.Read(icon.Data)
		if err != nil {
			return nil, err
		}
		if string(icon.Data[1:4]) == "PNG" {
			icon.Format = "PNG"
		} else {
			icon.Format = "BMP"
			// Decode BMP
			decodeBMP(icon)
		}
	}

	return result, nil
}

func decodeBMP(icon *Icon) {
	if icon.BitsPerPixel != 32 {
		icon.ErrorMessage = fmt.Sprintf("BMP image at %d bits ber pixel unsupported.", icon.BitsPerPixel)
		return
	}
	// Prefix the icon data with the BMP header
	var dibheader winicon.BitmapInfoHeader
	iconDataReader := bytes.NewReader(icon.Data)
	binary.Read(iconDataReader, binary.LittleEndian, &dibheader)
	dibheader.ImageWidth = uint32(icon.Width)
	dibheader.ImageHeight = uint32(icon.Height)

	bmpHeader := winicon.NewBitmapFileHeader(icon.size)

	var icondata bytes.Buffer
	err := binary.Write(&icondata, binary.LittleEndian, bmpHeader)
	if err != nil {
		println("Error:", err.Error())
		icon.ErrorMessage = err.Error()
		return
	}

	err = binary.Write(&icondata, binary.LittleEndian, &dibheader)
	if err != nil {
		println("Error:", err.Error())
		icon.ErrorMessage = err.Error()
		return
	}

	bytesInPixels := uint32(icon.Height * icon.Width * (icon.BitsPerPixel / 8))
	err = binary.Write(&icondata, binary.LittleEndian, icon.Data[dibheader.HeaderSize:bytesInPixels+dibheader.HeaderSize])
	if err != nil {
		println("Error:", err.Error())
		icon.ErrorMessage = err.Error()
		return
	}
	icon.Data = icondata.Bytes()
}
