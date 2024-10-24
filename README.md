BMP Image Processing Tool
Overview
This tool allows users to read BMP image files, manipulate their headers, and apply various image transformations and filters. It supports several subcommands and options for ease of use.

Features
Header
The program can print the header information of a BMP file using the header subcommand. The output includes:

FileType: The file format (e.g., BM).
FileSizeInBytes: Total size of the BMP file.
HeaderSize: Size of the BMP header.
DIBHeaderSize: Size of the DIB header.
WidthInPixels: Image width in pixels.
HeightInPixels: Image height in pixels.
PixelSizeInBits: Number of bits per pixel.
ImageSizeInBytes: Size of the image data.
In case of an error (e.g., if the file is not a valid BMP), an error message is displayed with a non-zero exit status.

Example Usage:

bash

$ ./bitmap header sample.bmp
BMP Header:
- FileType BM
- FileSizeInBytes 518456
- HeaderSize 54
DIB Header:
- DibHeaderSize 40
- WidthInPixels 480
- HeightInPixels 360
- PixelSizeInBits 24
- ImageSizeInBytes 518402
Mirror
The program can mirror an image either horizontally or vertically using the --mirror option. Options can be combined, allowing for multiple mirroring operations.

Example Usage:

bash

$ ./bitmap apply --mirror=horizontal sample.bmp sample-mirrored-horizontal.bmp
Filter
Users can apply various filters to images using the --filter flag. Multiple filters can be specified, and they are applied sequentially:

--filter=blue: Retains only the blue channel.
--filter=red: Retains only the red channel.
--filter=green: Retains only the green channel.
--filter=grayscale: Converts the image to grayscale.
--filter=negative: Applies a negative filter.
--filter=pixelate: Applies pixelation (default block size: 20 pixels).
--filter=blur: Applies a blur effect.
Example Usage:

bash

$ ./bitmap apply --filter=negative sample.bmp sample-filtered-negative.bmp
Rotate
Images can be rotated using the --rotate option. The program supports various angles:

right or 90: Clockwise rotation.
left or -90: Counterclockwise rotation.
Multiple rotations can be combined.

Example Usage:

bash

$ ./bitmap apply --rotate=right --rotate=right sample.bmp sample-rotated-right-right.bmp
Crop
The program can crop images with the --crop option, specifying offsets and dimensions:

--crop=OffsetX-OffsetY-Width-Height (width and height are optional).
If the crop dimensions exceed the image size, an error is displayed.

Example Usage:

bash

$ ./bitmap apply --crop=20-20-100-100 sample.bmp sample-cropped-20-20-80-80.bmp
Combine Apply Options
All apply options can be combined and executed sequentially.

Example Usage:

bash

$ ./bitmap apply --mirror=horizontal --rotate=right --filter=negative sample.bmp sample-combined-output.bmp
Help
The program provides a help message when the -h or --help flag is used. This message includes information about all available commands and options.

Example Usage:

bash

$ ./bitmap --help