package winicon

// IconFileHeader is the icon file's 48bit header
type IconFileHeader struct {
	_          uint16
	ImageType  uint16
	ImageCount uint16
}

// IconHeader is the header for the icon data
type IconHeader struct {
	Width        uint8
	Height       uint8
	Colours      uint8
	_            uint8
	Planes       uint16
	BitsPerPixel uint16
	Size         uint32
	Offset       uint32
}

// BitmapFileHeader stores information about a BMP Header
type BitmapFileHeader struct {
	HeaderField uint16
	Size        uint32
	_           uint16
	_           uint16
	Offset      uint32
}

// NewBitmapFileHeader creates a new bitmap file header based on the given data
func NewBitmapFileHeader(datasize uint32) *BitmapFileHeader {
	return &BitmapFileHeader{
		HeaderField: 0x4D42,
		Size:        uint32(datasize + 54),
		Offset:      54,
	}
}

// BitmapInfoHeader is the DIB image header format
// Credit: https://itnext.io/bits-to-bitmaps-a-simple-walkthrough-of-bmp-image-format-765dc6857393
type BitmapInfoHeader struct {
	HeaderSize      uint32 // An integer (unsigned) representing the size of the header in bytes. It should be '40' in decimal to represent BITMAPINFOHEADER header type.
	ImageWidth      uint32 // An integer (signed) representing the width of the final image in pixels.
	ImageHeight     uint32 // An integer (signed) representing the height of the final image in pixels.
	Planes          uint16 // An integer (unsigned) representing the number of color planes of the target device. Should be '1' in decimal.
	BitsPerPixel    uint16 // An integer (unsigned) representing the number of bits (memory) a pixel takes (in pixel data) to represent a color.
	Compression     uint32 // An integer (unsigned) representing the value of compression to use. Should be '0' in decimal to represent no-compression (identified by 'BI_RGB').
	ImageSize       uint32 // An integer (unsigned) representing the final size of the compressed image. Should be '0' in decimal when no compression algorithm is used.
	XpixelsPerMeter uint32 // An integer (signed) representing the horizontal resolution of the target device. This parameter will be adjusted by the image processing application but should be set to '0' in decimal to indicate no preference.
	YpixelsPerMeter uint32 // An integer (signed) representing the verical resolution of the target device (same as the above).
	TotalColors     uint32 // An integer (unsigned) representing the number of colors in the color pallet (size of the color pallet or color table). If this is set to '0' in decimal :- 2^BitsPerPixel colors are used.
	ImportantColors uint32 // An integer (unsigned) representing the number of important colors. Generally ignored by setting '0' decimal value.
}
