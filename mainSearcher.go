package ghostls

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func DirSearcher(orgPath string) {
	var Directories []string
	var fileArray []string
	files, err := os.ReadDir(orgPath)
	g, e := GetLongestFileSize(orgPath)
	if e != nil {
		RedPrintln("Error getting longest file size")
		log.Fatal(err)
	}
	LongS = g
	if err != nil {
		fmt.Println(redANSI + boldANSI + "Error Searching Directory" + resetANSI)
		log.Fatal(err)
	}
	if LongFormat|| DashO {
		bcount, err := GetBlockCount(orgPath)
		if err != nil {
			RedPrintln("ERROR GETTING BLOCKCOUNT IN MAINSEARCHER")
			log.Fatal(err)
		}
		fmt.Println("total "+  strconv.Itoa(int(bcount)))
	}
	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") && !DisplayHidden {
			continue
		}
		fileinfo, err := os.Lstat(orgPath + "/" + file.Name())
		if err != nil {
			fmt.Println(redANSI + boldANSI + "Error getting file Info for append" + resetANSI)
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
			fileArray = SortByCreationTime(orgPath, fileArray, true)
			Directories = SortByCreationTime(orgPath, Directories, true)
		} else {
			fileArray = SortByCreationTime(orgPath, fileArray, false)
			Directories = SortByCreationTime(orgPath, Directories, false)
		}
	}
	//* parse and print
	for _, v := range fileArray {
		todisplay := ""
		filestat, err := os.Lstat(orgPath + "/" + v)
		if err != nil {
			fmt.Println(redANSI + boldANSI + "Error getting file Info for print")
			log.Fatal(err)
		}

		_, dirbool, err := GetFilePermissions(orgPath + "/" + v)
		if err != nil {
			fmt.Println(redANSI + boldANSI + "DirSearch Error, cant get permissions")
			log.Fatal(err)
		}
		if !LongFormat && !DashO {
			if filestat.IsDir() || dirbool == "d" {
				fmt.Print(BlueFormat(v) + " ")
			} else {
				extension := getExtension(string(v))
				fmt.Println(extension)
				todisplay = getColorizedFileType(extension, string(v))
				fmt.Print(todisplay + " ")
			}
		} else if LongFormat || DashO {
			LongFormatDisplay(orgPath + "/" + v)
		}
	}
	//* Handle recursive search
	if RecursiveSearch {
		for _, dir := range Directories {
			fmt.Println()
			fmt.Println(orgPath + "/" + dir)
			mainPath := orgPath + "/" + dir
			DirSearcher(mainPath)
		}
	}
}
