package main
import (
    "fmt"
    "os"
)
func main() {
    home, _ := os.UserHomeDir()
    fmt.Println("Home dizini:", home)
}