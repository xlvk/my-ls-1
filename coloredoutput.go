package ghostls

import (
	"fmt"
)

const (
	yellowANSI  string = "\033[33m"
	redANSI     string = "\033[31m"
	greenANSI   string = "\033[32m"
	orangeANSI  string = "\033[38;5;208m"
	blueANSI    string = "\033[34m"
	boldANSI    string = "\033[1m"
	resetANSI   string = "\033[0m"
	magentaANSI string = "\u001b[35m"
	cyanANSI    string = "\u001b[36m"
	blackBgANSI string = "\033[40m"
)

func YellowPrintln(args ...any) {
	fmt.Print(yellowANSI)
	fmt.Print(boldANSI)
	fmt.Print(args...)
	fmt.Println(resetANSI)
}

func RedPrintln(args ...any) {
	fmt.Print(redANSI)
	fmt.Print(boldANSI)
	fmt.Print(args...)
	fmt.Println(resetANSI)
}

func GreenPrintln(args ...any) {
	fmt.Print(greenANSI)
	fmt.Print(boldANSI)
	fmt.Print(args...)
	fmt.Println(resetANSI)
}

func OrangePrintln(args ...any) {
	fmt.Print(orangeANSI)
	fmt.Print(boldANSI)
	fmt.Print(args...)
	fmt.Println(resetANSI)
}

func BlueFormat(argumet string) string {
	return blueANSI + boldANSI + argumet + resetANSI
}
func GreenFormat(argumet string) string {
	return greenANSI + boldANSI + argumet + resetANSI
}

func getColorizedFileType(fileType, FileName string) string {
	var colorCode string

	switch fileType {
	case ".jpg", ".png", ".gif", ".bmp":
		colorCode = magentaANSI + boldANSI
	case ".mp4", ".avi", ".mov", ".wmv":
		colorCode = magentaANSI + boldANSI
	case ".mp3", ".wav", ".flac", ".aac":
		colorCode = cyanANSI
	case ".zip", ".rar", ".tar.gz", ".7z":
		colorCode = redANSI + boldANSI
	default:
		colorCode = resetANSI
	}
	return colorCode + FileName + resetANSI
}
