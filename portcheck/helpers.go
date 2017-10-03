package portcheck

import (
	"log"
	"net"
)

// Returns true if OK, false for NOK
func ping(ip string, port string) (status bool) {
	conn, err := net.Dial("tcp", ip+":"+port)
	if err != nil {
		log.Println("Connection error:", err)
	} else {
		status = true
		conn.Close()
	}
	log.Printf("TCP %s:%s; Status: %t", ip, port, status)
	return
}
