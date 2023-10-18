package ghostls

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

func TCmond() int {
	var dimensions [4]uint16

	// Convert the dimensions array to a uintptr, which is safe
	dimensionsPtr := uintptr(unsafe.Pointer(&dimensions[0]))

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, os.Stdin.Fd(), uintptr(syscall.TIOCGWINSZ), dimensionsPtr)
	if err != 0 {
		fmt.Println(err)
		os.Exit(1)
	}
	width := int(dimensions[1])
	return width
}


func getExtension(filename string) string {
	lastDotIndex := -1
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			lastDotIndex = i
			break
		}
	}

	if lastDotIndex == -1 || lastDotIndex == len(filename)-1 {
		return ""
	}

	return filename[lastDotIndex:]
}