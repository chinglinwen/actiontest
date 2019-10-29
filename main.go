package main

import "fmt"

func main() {
  fmt.Printf("envs: ", os.Environ())
	fmt.Println("Hello, action")
}
