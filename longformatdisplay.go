package ghostls

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

// TODO: INTENSIVE TESTING REQUIRED

//* parsing Binary permissions
func GetFilePermissions(path string) (string, error) {
	// Get file info
	fileInfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	// Get permission bits
	mode := fileInfo.Mode()
	permissions := mode.Perm()

	// Convert permission bits to string
	permissionString := strconv.FormatUint(uint64(permissions), 8)

	// Pad the permission string to ensure 3 digits
	for len(permissionString) < 3 {
		permissionString = "0" + permissionString
	}

	// Map permission bits to their respective symbols
	permissionSymbols := map[int]string{
		0: "---",
		1: "--x",
		2: "-w-",
		3: "-wx",
		4: "r--",
		5: "r-x",
		6: "rw-",
		7: "rwx",
	}

	// Extract permission symbols for owner, group, and others
	ownerPermissions := permissionSymbols[int(permissions>>6)&7]
	groupPermissions := permissionSymbols[int(permissions>>3)&7]
	otherPermissions := permissionSymbols[int(permissions)&7]

	// Return the formatted permissions
	return fmt.Sprintf("%s %s %s", ownerPermissions, groupPermissions, otherPermissions), nil
}

//* syscall to get hard link numbers
func GetHardLinkNum(path string) (string, error) {
	fcount := uint64(0)

	fileinfo, err := os.Stat(path)
	if err != nil {
		return "", err
	}

	if sys := fileinfo.Sys(); sys != nil {
		if stat, ok := sys.(*syscall.Stat_t); ok {
			fcount = stat.Nlink
		}
	}
	mainnum := strconv.Itoa(int(fcount))

	return mainnum, nil
}

func GetFileOwnerAndGroup(filePath string) (string, string, error) {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", "", err
	}

	fileOwner := fileInfo.Sys().(*syscall.Stat_t).Uid
	fileGroup := fileInfo.Sys().(*syscall.Stat_t).Gid

	owner, err := lookupUserById(fileOwner)
	if err != nil {
		return "", "", err
	}

	group, err := lookupGroupById(fileGroup)
	if err != nil {
		return "", "", err
	}

	return owner, group, nil
}

func lookupUserById(uid uint32) (string, error) {
	u, err := user.LookupId(strconv.Itoa(int(uid)))
	if err != nil {
		return "", err
	}
	return u.Username, nil
}

func lookupGroupById(gid uint32) (string, error) {
	g, err := user.LookupGroupId(strconv.Itoa(int(gid)))
	if err != nil {
		return "", err
	}
	return g.Name, nil
}

//* formatting output
func LongFormatDisplay(filepaths []string) {
	for _, filepath := range filepaths {
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
		uname, gname, err := GetFileOwnerAndGroup(filepath)
		if err != nil {
			log.Fatal(err)
		}
		filerecord += uname + " " + gname + " "
	}
}
