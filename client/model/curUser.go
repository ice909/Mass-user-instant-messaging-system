package model

import (
	"net"

	"github.com/ice909/go-common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}
