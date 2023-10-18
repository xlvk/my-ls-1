package ghostls

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func TCmond() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(string(s), " ")
	// heigth, err := strconv.Atoi(string(sArr[0]))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	widthh, err1 := strconv.Atoi(string(sArr[1]))
	if err1 != nil {
		log.Fatal(err)
	}
	// if hight > heigth {
	// 	fmt.Println("Error: you need to make the hight of the terminal size beggier than this.")
	// 	os.Exit(0)
	// }
	// if width > widthh {
	// 	fmt.Println("Error: you need to make the width of the terminal size beggier than this.")
	// 	os.Exit(0)
	// }
	return widthh
}
