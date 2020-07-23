package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gookit/color"
)

// Print help information
func help() {
	fmt.Printf(`
SemanticBumper - A semantic version bumper

USAGE:
	SemanticBumper <filename/command>
	
OPTIONS:
	filename - path to a valid .bumped file
	command - Must be an option from COMMANDS. See below
	
COMMANDS:
	init - Creates a version.bumped file with basic fields
	help - Displays help on SemanticBumper
		
For more help, visit: https://github.com/TheOtterlord/SemanticBumper
`)
}

// Read data from a file
func readfile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		color.Red.Println("File reading error", err)
		return ""
	}
	return string(data)
}

// Write data to a file
func writefile(filename string, contents string) bool {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return false
	}
	_, err = file.WriteString(contents)
	if err != nil {
		fmt.Println(err)
		file.Close()
		return false
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// Read a .bumped file and execute it
func bumper(fn string) {
	contents := readfile(fn)
	lines := strings.Split(contents, "\n")
	version := ""
	bumps := false
	// contains regex for semantc version
	re := regexp.MustCompile(`(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)(?:-((?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*)(?:\.(?:0|[1-9]\d*|\d*[a-zA-Z-][0-9a-zA-Z-]*))*))?(?:\+([0-9a-zA-Z-]+(?:\.[0-9a-zA-Z-]+)*))?`)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		command := strings.Trim(line, " ")
		if strings.HasPrefix(command, "version") {
			version = strings.Split(command, "version")[1]
			version = strings.Trim(version, "\r: ")
			// if version is not semver
			if !re.MatchString(version) {
				color.Red.Printf("Error: version %s on line %d does not match semver specification. For more information, see: https://semver.org\n", version, i)
				return
			}
			bumps = false
		} else if strings.HasPrefix(command, "bumps") {
			bumps = true
		} else if strings.HasPrefix(command, "//") {
			// do nothing
		} else if strings.HasPrefix(command, "-") {
			if bumps {
				file := strings.Split(command, "-")[1]
				file = strings.Trim(file, "\r ")
				raw := readfile(file)
				if raw != "" {
					new := re.ReplaceAllString(raw, version)
					if writefile(file, new) {
						fmt.Println("Successfully updated " + file + " to " + version)
					} else {
						color.Red.Println("Failed to update " + file + " to " + version)
					}
				}
			}
		} else if command == "" {
			// empty line
		} else {
			color.Red.Printf("Error: Failed to parse \"%s\" on line %d\n", command, i)
			return
		}
	}
}

func main() {
	fmt.Println("Running SemanticBumper version 1.0.0")
	// check that we only have our 1 command/filename
	argsLen := len(os.Args) - 1
	if argsLen != 1 {
		color.Red.Println("Error: expected 1 argument, received " + strconv.Itoa(argsLen))
		help()
		return
	}
	filename := os.Args[1]
	if filename == "help" {
		help()
		return
	}
	// if init command used: create a basic .bumped file
	if filename == "init" {
		if writefile("version.bumped", "version: your_version_name\nbumps:\n") {
			fmt.Println("Successfully created version.bumped")
		} else {
			color.Red.Println("Failed to create version.bumped")
		}
		return
	}
	// process .bumped file
	if !strings.HasSuffix(filename, ".bumped") {
		color.Warn.Println("Warning: \"" + filename + "\" should use .bumped extension")
	}
	bumper(filename)
}
