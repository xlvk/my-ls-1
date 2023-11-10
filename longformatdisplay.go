package ghostls

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// * formatting output
func LongFormatDisplay(filepath string) {
	dest := ""
	filestats, err := os.Lstat(filepath)
	if err != nil {
		fmt.Println("LONG FORMAT ERROR")
		log.Fatal(err)
	}
	filerecord := ""
	//* check if symlink
	if filestats.Mode()&os.ModeSymlink != 0 {
		dest, err = os.Readlink(filepath)
		if err != nil {
			fmt.Println(redANSI + boldANSI + "Error getting link Info for long format print" + resetANSI)
			log.Fatal(err)
		}
	}
	//* parse binary permissions
	filepermissions, ftype, err := GetFilePermissions(filepath)
	if err != nil {
		RedPrintln("-L FILEPERM ERROR")
		log.Fatal(err)
	}
	filerecord += filepermissions + " "
	//* parse hard link number
	hardlinknum, err := GetHardLinkNum(filepath)
	if err != nil {
		RedPrintln("-L HARDLINKNUM ERROR")
		log.Fatal(err)
	}
	if len(hardlinknum) < 2 {
		hardlinknum = " " + hardlinknum
	}
	filerecord += hardlinknum + " "
	//* parse owner name and group name
	uname, gname, err := GetFileOwnerAndGroup(filepath)
	if err != nil {
		RedPrintln("-L UNAME ERROR")
		log.Fatal(err)
	}
	filerecord += " "+ uname + " "
	if !DashO {
		filerecord += gname + " "
	}
	//* parse file size
	filerecord += strings.Repeat(" ", LongS-len(strconv.Itoa(int(filestats.Size())))) + strconv.Itoa(int(filestats.Size())) + " "

	//* parse last mod date and time
	modtime := filestats.ModTime()
	_, month, day := modtime.Date()
	hour, min, _ := modtime.Clock()
	mainMin := strconv.Itoa(min)
	mainHr := strconv.Itoa(hour)
	if min < 10 {
		mainMin = "0" + mainMin
	}
	if hour < 10 {
		mainHr = "0" + mainHr
	}
	if len(strconv.Itoa(day)) < 2 {
		filerecord += string([]rune(month.String())[:3]) + "  " + strconv.Itoa(day) + " " + mainHr + ":" + mainMin + " "
	} else {
		filerecord += string([]rune(month.String())[:3]) + " " + strconv.Itoa(day) + " " + mainHr + ":" + mainMin + " "
	}
	var mainname string
	switch ftype {
	case "d":
		mainname = blueANSI + boldANSI + filestats.Name() + resetANSI
	case "l":
		mainname = cyanANSI + boldANSI + filestats.Name() + resetANSI + " -> " + dest
	case "ol":
		//* symlink that points to a file that doesnt exist
		mainname = blackBgANSI + redANSI + boldANSI + filestats.Name() + resetANSI + " -> " + dest
	case "c", "b":
		mainname = blackBgANSI + yellowANSI + boldANSI + filestats.Name() + resetANSI
	case "p":
		mainname = blackBgANSI + yellowANSI + filestats.Name() + resetANSI
	case "s":
		mainname = magentaANSI + filestats.Name() + resetANSI
	case "bin":
		mainname = greenANSI + boldANSI + filestats.Name() + resetANSI
	default:
		mainname = filestats.Name()
	}

	if ftype == "" {
		mainname = getColorizedFileType(getExtension(filestats.Name()), filestats.Name())
	}

	filerecord += mainname
	fmt.Println(filerecord)
}
