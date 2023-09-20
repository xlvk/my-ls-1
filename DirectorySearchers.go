package ghostls

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func RecursiveSearchDir(filepath string) {
	files, err := os.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}
	OrangePrintln(filepath + ":")
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") && !DisplayHidden {
			continue
		}
		fileinfo, err := os.Stat(filepath + "/" + file.Name())
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			log.Fatal(err)
		}
		if fileinfo.IsDir() {
			BluePrintln(file.Name())
			RecursiveSearchDir(filepath + "/" + file.Name())
			continue
		}
		fmt.Print(file.Name() + " ")

	}
	fmt.Println()
}

func NormalSearchDir(filepath string) {
	files, err := os.ReadDir(filepath)
	var fileArray []string
	if err != nil {
		log.Fatal(err)
	}
	OrangePrintln(filepath + ":")
	for _, file := range files {
		fileinfo, err := os.Stat(filepath + "/" + file.Name())
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			log.Fatal(err)
		}
		if strings.HasPrefix(file.Name(), ".") && !DisplayHidden {
			continue
		}
		if fileinfo.IsDir() {
			BluePrintln(file.Name())
			continue
		}
		fileArray = append(fileArray, file.Name())
	}
	BubbleSort(fileArray)
	for _, v := range fileArray {
		fmt.Print(v + " ")
	}
	fmt.Println()
}
