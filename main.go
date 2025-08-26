package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample go -2 app here-2")
}

func currentHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting Hostname: %v\n", err)
		return ""
	}
	return hostname
}

func getIpAdress() ([]net.Addr, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("Error getting interface address: %v\n", err)
		return nil, err
	}

	return addrs, nil
}

func GetHome(w http.ResponseWriter, r *http.Request) {
	host := currentHost()

	ip, err := getIpAdress()
	if err != nil {
		http.Error(w, "Error finding the ip address", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "host: %s, ip: %v", host, ip)
}

func main() {
	http.HandleFunc("/", handler)

	http.HandleFunc("/ip", GetHome)

	http.ListenAndServe(":8500", nil)
}
