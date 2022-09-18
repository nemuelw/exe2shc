# exe2shc 
exe to shellcode converter

## Installation
- Clone this repository
- Navigate to the project directory 
- Run the command :
    ```
    go build exe2shc.go
    ```

## Usage
- Navigate to the project directory
- Run the command :
    ```
    ./exe2shc -f <EXE-FILE>
    ```

## Note :
The shellcode is by default saved to a file named `shellcode.txt`

## Examples
Generate shellcode for the file test.exe and save to `shellcode.txt` :
```
./exe2shc -f text.exe
``` 
Generate shellcode for the file test.exe and save to `text.txt` instead of the default `shellcode.txt` :
```
./exe2shc -f text.exe -o test.txt
```