package main

import (
	"math/rand"
	"net"

	"github.com/satori/go.uuid"
)

func insanes(spawn int) {
	for i := 0; i < spawn; i++ {
		go insane()
	}
}

func insane() {
	buff := make([]byte, 1024)

	for {
		utilPause()

		var c net.Conn
		uuid := uuid.NewV4().String()

		switch rand.Intn(5) {
		case 0:
			c = utilWebSocketClient(buff)

		case 1:
			c = utilCreateSock()
			c.Write([]byte(httpRequest(uuid, "", true)))

		case 2:
			c = utilCreateRawClient()

		case 3:
			c = utilCreateSock()
			c.Write([]byte("<policy-file-request/>"))

		case 4:
			c = utilCreateSock()
		}

		switch rand.Intn(2) {
		case 0:
			c.Write([]byte(httpRequest(uuid, utilRandomEvent(), false)))

		case 1:
			c.Write([]byte(utilRandomEvent()))
		}

		c.Close()
	}
}
