package main
import (
    "fmt"
    "os"
)
func main() {
    os.Setenv("USER", "Abdullah")
    result := os.ExpandEnv("Merhaba $USER, hoş geldin!")
    fmt.Println(result)
}