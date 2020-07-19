package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func help() {
	msg := `SemanticBumper - A semantic version bumper program

USAGE:
	SemanticBumper [filename]
	
OPTIONS:
	filename - The name of your .bumped file
	
For more help, visit: https://github.com/TheOtterlord/SemanticBumper
`
	fmt.Printf(msg + "\n")
}

func readfile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("File reading error", err)
		return ""
	}
	return string(data)
}

func writefile(filename string, contents string) bool {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	l, err := file.WriteString(contents)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return false
	}
	if l == 0 {
		fmt.Printf("No bytes written. Potential error")
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func bump(filename string) {
	contents := readfile(filename)
	lines := strings.Split(contents, "\n")
	version := ""
	bumps := false
	re := regexp.MustCompile(`(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		command := strings.Trim(line, " ")
		if strings.HasPrefix(command, "version") {
			version = strings.Split(command, "version")[1]
			version = strings.Trim(version, "\r: ")
			bumps = false
		} else if strings.HasPrefix(command, "bumps") {
			bumps = true
		} else if strings.HasPrefix(command, "//") {
			// it's a comment
			// do nothing
		} else if strings.HasPrefix(command, "-") {
			if bumps {
				file := strings.Split(command, "-")[1]
				file = strings.Trim(file, "\r ")
				raw := readfile(file)
				if raw != "" {
					new := re.ReplaceAllString(raw, version)
					if writefile(file, new) {
						fmt.Printf("Successfully updated " + file + " to " + version + "\n")
					} else {
						fmt.Printf("Failed to update " + file + " to " + version + "\n")
					}
				}
			} else if command == "" {
				// empty line
			} else {
				fmt.Printf("Error: Failed to parse \"" + command + "\"")
			}
		}
	}
}

func makeInitFile() {
	if writefile("version.bumped", "version: 1.0.0\nbumps:\n") {
		fmt.Printf("Successfully created version.bumped\n")
	}
}

func main() {
	fmt.Printf("Running SemanticBumper Version 0.1.0\n")
	if len(os.Args) < 2 {
		fmt.Printf("ERROR: No target file specified\n")
		help()
		return
	}
	filename := os.Args[1]
	if filename == "--help" || filename == "help" {
		help()
		return
	}
	if filename == "init" {
		makeInitFile()
		return
	}
	bump(filename)
}
