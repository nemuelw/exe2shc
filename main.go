package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

// colors


func printUsage() {
	fmt.Println("Usage : ")
	os.Exit(1)
}

func fileExists(file string) bool {
	if _, err := os.Stat(file); err == nil {
		return true
	} else {
		return false
	}
}

func handleError(err error) {
	fmt.Printf("[!] Error : %v\n", err)
	os.Exit(1)
}

func main() {
	fmt.Println("[*] Starting exe2shc ... ")

	var exeFile string
	var shcFile string

	flag.StringVar(&exeFile, "f", "", "The executable file to be converted")
	flag.StringVar(&shcFile, "o", "shellcode.txt", "Output file where shellcode will be saved")
	flag.Parse()

	// check that exeFile is provided
	if len(exeFile) == 0 {
		printUsage()
	}

	// check that the file exists
	if fileExists(exeFile) != true {
		fmt.Println("[!] The supplied file does not exist !")
		os.Exit(1)
	}

	binaryData, _ := os.ReadFile(exeFile)
	hexData := hex.EncodeToString(binaryData)
	shellcode := ""
	for i, _ := range hexData {
		if i % 2 != 0 {
			continue
		}
		shellcode += "\\x"
		shellcode += hexData[i:i+2]
	}

	fmt.Println(shellcode)
}