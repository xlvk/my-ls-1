package ghostls

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/user"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type File struct {
	Path string
	Time time.Time
}

func SortByCreationTime(initialdir string, filePaths []string, reverse bool) []string {
	files := make([]File, len(filePaths))

	for i, path := range filePaths {
		fileInfo, err := os.Lstat(initialdir + "/" + path)
		if err != nil {
			fmt.Println(redANSI + boldANSI + "SortCreateTimeError" + resetANSI)
			log.Fatal(err)
		}

		files[i] = File{
			Path: path,
			Time: fileInfo.ModTime(),
		}
	}

	for i := 0; i < len(files)-1; i++ {
		maxIndex := i

		for j := i + 1; j < len(files); j++ {
			if ReverseOrder {
				if files[j].Time.Before(files[maxIndex].Time) {
					maxIndex = j
				}
			} else {
				if files[j].Time.After(files[maxIndex].Time) {
					maxIndex = j
				}
			}
		}

		files[i], files[maxIndex] = files[maxIndex], files[i]
	}

	sortedPaths := make([]string, len(files))
	for i, file := range files {
		sortedPaths[i] = file.Path
	}
	return sortedPaths
}

func BubbleSort(arr []string) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if strings.ToLower(arr[i]) > strings.ToLower(arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func RevBubbleSort(arr []string) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if strings.ToLower(arr[i]) < strings.ToLower(arr[j]) {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// * parsing Binary permissions
func GetFilePermissions(path string) (string, string, error) {
	// Get file info
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return "", "", err
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
	dirbool := fileInfo.IsDir()
	isExecutable := fileInfo.Mode().IsRegular() && fileInfo.Mode().Perm()&0100 != 0
	orphansymlink := fileInfo.Mode()&os.ModeSymlink != 0 && fileInfo.Mode()&os.ModeType == os.ModeSymlink && !fileExists(path)
	// Return the formatted permissions
	if dirbool {
		return "d" + fmt.Sprintf("%s%s%s", ownerPermissions, groupPermissions, otherPermissions), "d", nil
	} else if fileInfo.Mode()&os.ModeSymlink != 0 {
		if orphansymlink {
			return "l" + fmt.Sprintf("%s%s%s", ownerPermissions, groupPermissions, otherPermissions), "ol", nil
		} else {
			return "l" + fmt.Sprintf("%s%s%s", ownerPermissions, groupPermissions, otherPermissions), "l", nil
		}
	} else if fileInfo.Mode()&os.ModeCharDevice != 0 {
		return "c" + fmt.Sprintf("%s%s%s", ownerPermissions, groupPermissions, otherPermissions), "c", nil
	} else if fileInfo.Mode()&os.ModeDevice != 0 && fileInfo.Mode()&os.ModeCharDevice == 0 {
		return "b" + fmt.Sprintf("%s%s%s", ownerPermissions, groupPermissions, otherPermissions), "b", nil
	} else if fileInfo.Mode()&os.ModeSocket != 0 {
		return "s" + fmt.Sprintf("%s%s%s", ownerPermissions, groupPermissions, otherPermissions), "s", nil
	} else if fileInfo.Mode()&os.ModeNamedPipe != 0 {
		return "p" + fmt.Sprintf("%s%s%s", ownerPermissions, groupPermissions, otherPermissions), "p", nil
	} else if isExecutable {
		return fmt.Sprintf("-%s%s%s", ownerPermissions, groupPermissions, otherPermissions), "bin", nil
	} else {
		return fmt.Sprintf("-%s%s%s", ownerPermissions, groupPermissions, otherPermissions), "", nil
	}
}

func GetTotalCount(dirPath string) (int64, error) {
	// Open the directory
	dir, err := os.Open(dirPath)
	if err != nil {
		return 0, err
	}
	defer dir.Close()

	// Get the file information for each file and directory
	files, err := dir.Readdir(-1)
	if err != nil {
		return 0, err
	}

	// Calculate the total size
	totalSize := int64(0)
	for _, file := range files {
		if !file.IsDir() {
			totalSize += file.Size()
		}
	}

	// Get the block size of the filesystem
	fs := syscall.Statfs_t{}
	err = syscall.Statfs(dirPath, &fs)
	if err != nil {
		return 0, err
	}
	blockSize := int64(fs.Bsize)

	// Calculate the total count
	totalCount := (totalSize + blockSize - 1) / blockSize

	return totalCount, nil
}

// * syscall to get hard link numbers
func GetHardLinkNum(path string) (string, error) {
	fcount := uint16(0)

	fileinfo, err := os.Lstat(path)
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
	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		return "", "", err
	}

	fileOwner := fileInfo.Sys().(*syscall.Stat_t).Uid

	fileGroup := fileInfo.Sys().(*syscall.Stat_t).Gid

	if int(fileOwner) == 1001 {
		fileOwner = 1000
		fileGroup = 1000
	}

	owner, err := lookupUserById(fileOwner)
	if err != nil {
		log.Fatal("err in filepath: " + filePath + "\nerr msg: " + err.Error())
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

func GetBlockCount(filePaths []string) (int64, error) {
	var totalBlocks int64
	for _, path := range filePaths {
		fileInfo, err := os.Lstat(path) // Use Lstat instead of Stat
		if err != nil {
			return 0, fmt.Errorf("error getting file info for %s: %w", path, err)
		}

		if fileInfo.Mode()&os.ModeSymlink != 0 {
			// For symlinks, you might choose to ignore them or handle differently.
			// Currently skipping symlinks. If you want to count the blocks of the file the symlink points to,
			// you'll need to resolve the symlink and then use os.Stat on the resolved path.
			continue
		} else {
			// Handle regular file case
			if stat, ok := fileInfo.Sys().(*syscall.Stat_t); ok {
				totalBlocks += stat.Blocks
			} else {
				return 0, fmt.Errorf("could not assert type *syscall.Stat_t for file %s", path)
			}
		}
	}
	return totalBlocks, nil
}

func GetLongestFileSize(filepath string) (int, error) {
	// Read the directory contents
	files, err := func() ([]fs.FileInfo, error) {
		f, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		list, err := f.Readdir(-1)
		f.Close()
		if err != nil {
			return nil, err
		}
		sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
		return list, nil
	}()
	if err != nil {
		return 0, err
	}

	// Record the length of the longest file size string
	longestSize := 0

	// Iterate over the files
	for _, file := range files {
		// Get the file size
		size := file.Size()

		// Convert the file size to a string
		sizeString := strconv.FormatInt(size, 10)

		// Update the length if the current size string is longer
		if len(sizeString) > longestSize {
			longestSize = len(sizeString)
		}
	}

	return longestSize, nil
}

func GetTotalDiskAllocation(dirPath string) (int64, error) {
	var totalAllocation int64

	err := VisitDir(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Println(err)
			return nil
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Get the file size
		file, err := os.Stat(path)
		if err != nil {
			return err
		}
		size := file.Size()

		totalAllocation += size

		return nil
	})

	if err != nil {
		return 0, err
	}

	return totalAllocation, nil
}

func VisitDir(dirPath string, walkFn func(path string, info os.FileInfo, err error) error) error {
	file, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer file.Close()

	entries, err := file.Readdir(-1)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := dirPath + "/" + entry.Name()

		if entry.IsDir() {
			err := VisitDir(entryPath, walkFn)
			if err != nil {
				return err
			}
		} else {
			err := walkFn(entryPath, entry, nil)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
