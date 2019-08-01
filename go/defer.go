package main

import "fmt"
//import "os"

func main() {
    
    fmt.Println("selamat datang")

    //os.Exit(1)

    defer fmt.Println("test")
    defer fmt.Println("halo")
    defer fmt.Println("test halo")
}