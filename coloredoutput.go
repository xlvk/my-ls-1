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

func BlueFormat(args ...any) string {
	return blueANSI + boldANSI + fmt.Sprintf("%s", args[0]) + resetANSI
}

func getColorizedFileType(fileType, FileName string) string {
	var colorCode string

	switch fileType {
	case ".txt", ".doc", ".docx", ".pdf":
		colorCode = blueANSI
	case ".jpg", ".png", ".gif", ".bmp":
		colorCode = magentaANSI
	case ".mp4", ".avi", ".mov", ".wmv":
		colorCode = redANSI
	case ".mp3", ".wav", ".flac", ".aac":
		colorCode = yellowANSI
	case ".zip", ".rar", ".tar.gz", ".7z":
		colorCode = cyanANSI
	case ".xls", ".xlsx", ".csv":
		colorCode = greenANSI
	case ".ppt", ".pptx":
		colorCode = magentaANSI
	case ".go", ".java", ".py", ".cpp":
		colorCode = yellowANSI
	case ".html", ".css", ".js", ".xml":
		colorCode = blueANSI
	case ".exe", ".app", ".sh", ".bat":
		colorCode = redANSI
	default:
		colorCode = resetANSI
	}
	return colorCode + boldANSI + fmt.Sprintf("%s", FileName) + resetANSI
}
