package handlers

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func currentHost() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Printf("Error getting Hostname: %v\n", err)
		return ""
	}
	return hostname
}

func getIpAdress() ([]string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Printf("Error getting interface address: %v\n", err)
		return nil, err
	}

	ip := []string{}
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil { // Check if it's an IPv4 address
				fmt.Printf("  IPv4: %s\n", ipnet.IP.String())
				ip = append(ip, ipnet.IP.String())
			} else { // It's an IPv6 address
				fmt.Printf("  IPv6: %s\n", ipnet.IP.String())
				ip = append(ip, ipnet.IP.String())
			}
		}
	}

	return ip, nil
}

func (h *Handler) Greet(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		http.Error(w, "missing name", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Hello %s!", name)
}

func (h *Handler) GetHome(w http.ResponseWriter, r *http.Request) {
	host := currentHost()

	ip, err := getIpAdress()
	if err != nil {
		http.Error(w, "Error in finding the ip address", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "host: %s, ip: %v", host, ip)
}

func (h *Handler) Dummy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Sample go -2 app here-2")
}

func (h *Handler) Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Healthy-1"))
}
