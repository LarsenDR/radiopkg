package radiopkg

import (
	"log"
	"net"
	"os"
)

func main() {

	l, err := net.ListenUDP("udp", ":1024")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go copyToStderr(conn)
	}
}

func copyToStderr(conn net.Conn) {
	n, err := io.copy(os.Stderr, conn)
	log.Printf("Copied %d bytes; finnished with err %v", n, err)
}
