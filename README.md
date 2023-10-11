# GhostLs (My-Ls-1)

**Project GhostLS aims to replicate the Infamous command `ls`, and some of its flags**

Here is the file structure

```bash

.
├── coloredoutput.go
├── DirectorySearchers.go
├── globalvars.go
├── go.mod
├── LICENSE
├── longformatdisplay.go
├── main
│   └── main.go
├── Makefile
├── parseoptions.go
├── README.md
└── utils.go

```

***To Run this project***

1. `make ghostLS` after cloning and navigating
2. `./GhostLS [OPTIONS] [FLAGS]`

**Implemented Flags:**

* `-a` : Displays hidden files
* `-R` : Recusively searches a directory
* `-r` : Displays files in reverse order
* `-l` : Long format display
* `-t` : Sorts files by time
* `-o` : Long format display without group name

Project was Written in **Golang** using only allowed libraies

**Authors**

1. Fatima Abbas **fatabbas**
2. Mohamed Faris **mfaris**
3. Abdulrahman Khaled **akhaled**
