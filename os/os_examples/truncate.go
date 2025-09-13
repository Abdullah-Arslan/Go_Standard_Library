package main
import "os"
func main() {
    os.Truncate("demo.txt", 0)
}