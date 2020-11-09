package main

import (
	"fmt"
	"net"
)

func localAddresses() {
	iface_map := make(map[string][]string)
	ifaces, err := net.Interfaces()
	if err != nil {
			fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
			return
	}
	for _, i := range ifaces {
			addrs, err := i.Addrs()
			if err != nil {
					fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
					continue
			}
			for _, a := range addrs {
					switch v := a.(type) {
					case *net.IPAddr:
							fmt.Printf("%v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())

					case *net.IPNet:
							if _, ok := iface_map[i.Name]; !ok {
									iface_map[i.Name] = []string {fmt.Sprintf("%s", v)}
							} else {
								iface_map[i.Name] = append(iface_map[i.Name], fmt.Sprintf("%s", v))
							}
							fmt.Printf("%v : %s [%v/%v]\n", i.Name, v, v.IP, v.Mask)
					}
			}
	}

	for iface, addrs := range iface_map {
		fmt.Printf("%v, %v \n", iface, addrs)
	}
}

func main() {
	localAddresses()
}