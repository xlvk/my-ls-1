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
	var filePaths []string
	dir, err := os.Open(orgPath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()

	files, err := dir.Readdir(-1) // -1 to read all files
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() {
			filePaths = append(filePaths, orgPath+"/"+file.Name())
		}
	}
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
	if LongFormat || DashO {
		bcount, err := GetBlockCount(filePaths, orgPath)
		if err != nil {
			RedPrintln("ERROR GETTING BLOCKCOUNT IN MAINSEARCHER")
			log.Fatal(err)
		}
		fmt.Println("total " + strconv.FormatInt(int64(bcount), 10))
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
	if DisplayHidden {
		fileArray = append(fileArray, ".", "..")
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
	for i, v := range fileArray {
		todisplay := ""
		_, typePool, err := GetFilePermissions(orgPath + "/" + v)
		if err != nil {
			fmt.Println(redANSI + boldANSI + "DirSearch Error, cant get permissions")
			log.Fatal(err)
		}
		if !LongFormat && !DashO {
			switch typePool {
			case "d":
				todisplay = blueANSI + boldANSI + v + resetANSI
			case "l":
				todisplay = cyanANSI + boldANSI + v + resetANSI
			case "ol":
				//* symlink that points to a file that doesnt exist
				todisplay = blackBgANSI + redANSI + boldANSI + v + resetANSI
			case "c", "b":
				todisplay = blackBgANSI + yellowANSI + boldANSI + v + resetANSI
			case "p":
				todisplay = blackBgANSI + yellowANSI + v + resetANSI
			case "s":
				todisplay = magentaANSI + v + resetANSI
			case "bin":
				todisplay = greenANSI + boldANSI + v + resetANSI
			default:
				extension := getExtension(string(v))
				todisplay = getColorizedFileType(extension, string(v))
			}

			fmt.Print(todisplay + "  ")

			if i%8 == 0 && i != 0 {
				fmt.Println()
			}
		} else if LongFormat || DashO {
			LongFormatDisplay(orgPath + "/" + v)
		}
	}

	fmt.Println()
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
