package main

import "fmt"

func main() {
	c := make(chan bool)
	m := make(map[string]string)

	go func() {
		m["foo"] = "bar"
		c <- true
	}()

	m["bar"] = "foo"
	<-c

	for k, v := range m {
		fmt.Println(k, v)
	}
}
