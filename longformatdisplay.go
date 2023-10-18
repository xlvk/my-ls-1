package ghostls

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// TODO: INTENSIVE TESTING REQUIRED

// * formatting output
func LongFormatDisplay(filepath string) {
	filestats, err := os.Stat(filepath)
	if err != nil {
		fmt.Println("LONG FORMAT ERROR")
		log.Fatal(err)
	}
	filerecord := ""
	//* parse binary permissions
	filepermissions, err := GetFilePermissions(filepath)
	if err != nil {
		log.Fatal(err)
	}
	filerecord += filepermissions + " "
	//* parse hard link number
	hardlinknum, err := GetHardLinkNum(filepath)
	if err != nil {
		log.Fatal(err)
	}
	filerecord += hardlinknum + " "
	//* parse owner name and group name
	uname, gname, err := GetFileOwnerAndGroup(filepath)
	if err != nil {
		log.Fatal(err)
	}
	filerecord += uname + " "
	if !DashO {
		filerecord += gname + " "
	}
	//* parse file size
	filerecord += strconv.Itoa(int(filestats.Size())) + " "
	//* parse last mod date and time
	modtime := filestats.ModTime()
	_, month, day := modtime.Date()
	hour, min, _ := modtime.Clock()
	filerecord += string([]rune(month.String())[:3]) + " " + strconv.Itoa(day) + " " + strconv.Itoa(hour) + ":" + strconv.Itoa(min) + " "
	var mainname string
	if filestats.IsDir() || filepermissions == "rwx-rwx-r-x" {
		mainname = BlueFormat(filestats.Name())
	} else {
		mainname = filestats.Name()
	}
	filerecord += mainname
	fmt.Println(filerecord)
}
