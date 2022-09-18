// Author : Nemuel Wainaina

package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
)

func printUsage() {
	fmt.Println("Usage : ./exe2shc -f <EXE-FILE> [-o <SHC-FILE>]")
	fmt.Println("\t-f     exe file       Executable file to be converted")
	fmt.Println("\t-o     output file    Output file where shellcode will be saved")
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
	os.Exit(0)
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
		fmt.Println("\033[31m[!] The supplied file does not exist !\033[0m")
		os.Exit(1)
	}

	fmt.Printf("[*] Reading %v ... \n", exeFile)
	binaryData, _ := os.ReadFile(exeFile)
	fmt.Println("[*] Converting to shellcode ...")
	hexData := hex.EncodeToString(binaryData)
	shellcode := ""
	for i, _ := range hexData {
		if i % 2 != 0 {
			continue
		}
		shellcode += "\\x"
		shellcode += hexData[i:i+2]
	}
	fmt.Println("\033[32m[+]\033[0m Conversion complete")

	// check if shcFile already exists and delete it if so
	if fileExists(shcFile) == true {
		fmt.Printf("[*] File %v exists .\n", shcFile)
		fmt.Println("[*] Deleting file ... ")
		if err := os.Remove(shcFile); err != nil {
			handleError(err)
		}
		fmt.Printf("[*] %v deleted successfully\n", shcFile)
	}

	// write the shellcode to the file specified in shcFile
	if err := os.WriteFile(shcFile, []byte(shellcode), 0766); err != nil {
		handleError(err)
	}

	fmt.Printf("\033[32m[+]\033[0m Shellcode successfully written to %v\n", shcFile)

	fmt.Println("\033[32m[+] Done !\033[0m")
}