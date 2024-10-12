// Write file in Go
// The os.WriteFile function accepts the filename or fullpath with extension
// if any, data or content in []byte format and file permission code

// Syntax
// func WriteFile(name string, data []byte, perm os.FileMode) err

// Read file in Go
// The os.ReadFile function accepts the filename or full file path 
// and return []byte and eror
// It will throw an error if the file doesn't exist or the file
// doesn't have permission to read

// Syntax
// func ReadFile(name string) ([]byte, error)



package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
  // WriteFile
  data := []byte("dev.to")
  err := os.WriteFile("test.txt", data, 0777)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("file written successfully.")
  
  // ReadFile
  readData, err := os.ReadFile("test.txt")
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(string(readData))
}

















