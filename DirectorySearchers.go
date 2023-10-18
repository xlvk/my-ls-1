package ghostls

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func RecursiveSearchDir(filepath string) {
	var Directories []string
	var fileArray []string
	files, err := os.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") && !DisplayHidden {
			continue
		}
		fileinfo, err := os.Stat(filepath + "/" + file.Name())
		if err != nil {
			fmt.Println("STAT ERROR")
			log.Fatal(err)
		}
		if fileinfo.IsDir() {
			Directories = append(Directories, file.Name())
		}
		fileArray = append(fileArray, file.Name())
	}

	//* sort the arrays and proceed
	if !Timesort {
		if ReverseOrder {
			RevBubbleSort(fileArray)
			RevBubbleSort(Directories)
		} else {
			BubbleSort(fileArray)
			BubbleSort(Directories)
		}
	} else {
		if ReverseOrder {
			fileArray = SortByCreationTime(filepath, fileArray, true)
			Directories = SortByCreationTime(filepath, Directories, true)
		} else {
			fileArray = SortByCreationTime(filepath, fileArray, false)
			Directories = SortByCreationTime(filepath, Directories, false)
		}
	}
	maxLength := 0
	// AllLength := 0
	for _, file := range fileArray {
		length := len(file)
		if length > maxLength {
			maxLength = length
		}
		// AllLength += length
	}

	terminalWidth := TCmond()
	numColumns := terminalWidth / (maxLength + 4)
	for i, v := range fileArray {
		todisplay := ""
		filestat, err := os.Stat(filepath + "/" + v)
		if err != nil {
			fmt.Println("FILEARRAY ERR")
			log.Fatal(err)
		}
		permissions, err := GetFilePermissions(filepath + "/" + v)
		if permissions == "" {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		if !LongFormat && !DashO {
			padding := maxLength - len(string(v)) + 4
			if filestat.IsDir() || permissions == "rwx-rwx-r-x" {
				todisplay = BlueFormat(v)
				fmt.Print(todisplay + strings.Repeat(" ", padding))
			} else {
				extension := path.Ext(string(v))
				todisplay = getColorizedFileType(extension, string(v))
				fmt.Print(todisplay + strings.Repeat(" ", padding))
				// fmt.Print(filestat.Name() + " ")
			}
		} else if LongFormat || DashO {
			LongFormatDisplay(filepath + "/" + v)
		}
		if (i+1)%numColumns == 0 {
			fmt.Println()
		}
	}

	fmt.Println()
	for _, dir := range Directories {
		OrangePrintln(dir)
		RecursiveSearchDir(filepath + "/" + dir)
	}
}

func NormalSearchDir(filepath string) {
	var fileArray []string
	files, err := os.ReadDir(filepath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") && !DisplayHidden {
			continue
		}
		if err != nil {
			fmt.Println("STAT ERROR")
			log.Fatal(err)
		}
		fileArray = append(fileArray, file.Name())
	}

	//* sort the arrays and proceed
	if !Timesort {
		if ReverseOrder {
			RevBubbleSort(fileArray)
		} else {
			BubbleSort(fileArray)
		}
	} else {
		if ReverseOrder {
			fileArray = SortByCreationTime(filepath, fileArray, true)
		} else {
			fileArray = SortByCreationTime(filepath, fileArray, false)
		}
	}

	maxLength := 0
	// AllLength := 0
	for _, file := range fileArray {
		length := len(file)
		if length > maxLength {
			maxLength = length
		}
		// AllLength += length
	}

	terminalWidth := TCmond()
	numColumns := terminalWidth / (maxLength + 4)

	for i, v := range fileArray {
		todisplay := ""
		filestat, err := os.Stat(filepath + "/" + v)
		if err != nil {
			fmt.Println("FILEARRAY ERR")
			log.Fatal(err)
		}
		permissions, err := GetFilePermissions(filepath + "/" + v)
		if permissions == "" {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		if !LongFormat && !DashO {
			padding := maxLength - len(string(v)) + 4
			if filestat.IsDir() || permissions == "rwx-rwx-r-x" {
				todisplay = BlueFormat(v)
				fmt.Print(todisplay + strings.Repeat(" ", padding))
			} else {
				extension := path.Ext(string(v))
				todisplay = getColorizedFileType(extension, string(v))
				fmt.Print(todisplay + strings.Repeat(" ", padding))
				// fmt.Print(filestat.Name() + " ")
			}
		} else if LongFormat || DashO {
			LongFormatDisplay(filepath + "/" + v)
		}
		if (i+1)%numColumns == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
