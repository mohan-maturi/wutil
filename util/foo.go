package util

import "fmt"

func GetString(name string) string {
	w := "Hello World!!! This is " + name
	fmt.Printf("World Greeting is - %s\n", w)
	return w
}
