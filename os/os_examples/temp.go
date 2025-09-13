package main
import (
    "fmt"
    "os"
)
func main() {
    tmpFile, _ := os.CreateTemp("", "example-*.txt")
    defer os.Remove(tmpFile.Name())
    fmt.Println("Geçici dosya:", tmpFile.Name())
    tmpDir, _ := os.MkdirTemp("", "mytemp-*")
    defer os.RemoveAll(tmpDir)
    fmt.Println("Geçici dizin:", tmpDir)
}