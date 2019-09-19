package main

import (
    "fmt"
    "os"
)
//import "os"

func main() {
    
    fmt.Println("selamat datang")

    os.Exit(100)

    defer fmt.Println("test")
    defer fmt.Println("halo")
    defer fmt.Println("test halo")
}