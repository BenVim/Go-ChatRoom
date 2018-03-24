package main

import (
	"fmt"
	"net"
	"time"
)

var globalRoom *Room = NewRoom()

type Room struct {
	users map[string]net.Conn
}

func NewRoom() *Room {
	return &Room{
		users: make(map[string]net.Conn),
	}
}

func (r *Room) Join(user string, conn net.Conn) {
	_, ok := r.users[user]
	if ok {
		r.Leave(user)
	}

	r.users[user] = conn
	fmt.Printf("%s login success!", user)
	conn.Write([]byte(user + ":join this chat room!\n"))

}

func (r *Room) Leave(user string) {
	conn, ok := r.users[user]

	if !ok {
		fmt.Printf("%v user is not exist", user)
	}
	conn.Close()
	delete(r.users, user)
	fmt.Printf("%s leave", user)
}

func (r *Room) Broadcast(who string, msg string) {
	timeInfo := time.Now().Format("2018-03-25 23:00:00")
	toSend := fmt.Sprintf("%v %s:%s\n", timeInfo, who, msg)
	for user, conn := range r.users {
		if user == who {
			continue
		}
		conn.Write([]byte(toSend))
	}

}

func HandleConn(conn net.Conn)  {
	defer conn.Close()


}

func main() {
	fmt.Println("fuck")

}
