package ghostls

import (
	"log"
	"os"
)

func IsBinaryExecutable(filename string) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the file mode has the executable permission
	// bitwise usage inspired by session with sayhusain and labdulla
	return fileInfo.Mode().Perm()&0111 == 0
}

func BubbleSort(arr []string) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// Swap elements
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
