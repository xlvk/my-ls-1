package ghostls

import "fmt"

const (
	yellowANSI string = "\033[33m"
	redANSI    string = "\033[31m"
	greenANSI  string = "\033[32m"
	orangeANSI string = "\033[38;5;208m"
	blueANSI   string = "\033[34m"
	boldANSI   string = "\033[1m"
	resetANSI  string = "\033[0m"
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

func BluePrintln(args ...any) {
	fmt.Print(blueANSI)
	fmt.Print(boldANSI)
	fmt.Print(args...)
	fmt.Println(resetANSI)
}
