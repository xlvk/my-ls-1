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
	widthh, err1 := strconv.Atoi(string(sArr[1]))
	if err1 != nil {
		log.Fatal(err)
	}
	return widthh
}
