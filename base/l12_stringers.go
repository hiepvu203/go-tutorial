package main

import (
	"fmt"
	"strings"
)

type Info struct {
	Name string
	Age  int
}

func (i Info) String() string {
	return fmt.Sprintf("%s is %d years old", i.Name, i.Age)
}

type IPAddr [4]byte

func (i IPAddr) String() string {
	parts := make([]string, len(i))
	for i, v := range i {
		parts[i] = fmt.Sprint(v)
	}
	return strings.Join(parts, ".")
}

func main() {
	p := Info{Name: "Hiep", Age: 22}
	fmt.Println(p) // Hiep is 22 years old

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
