package main

import (
	"fmt"
	"net"
	"time"
	"bufio"
	"strings"
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

	r := bufio.NewReader(conn)
	line, err :=r.ReadString('\n')
	if err !=nil{
		fmt.Println(err)
		return
	}

	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields) !=2{
		conn.Write([]byte("user or password is error, is exit!"))
		return
	}

	user := fields[0]
	password:=fields[1]

	if password!="123"{
		fmt.Println("password error!")
		return
	}

	globalRoom.Join(user, conn)
	globalRoom.Broadcast("System", fmt.Sprintf("%s join room", user))

	for{
		conn.Write([]byte("send message:>>>"))
		line, err :=r.ReadString("\n")

		if err !=nil{
			break
		}

		line = strings.TrimSpace(line)
		fmt.Println(user, line)
		globalRoom.Broadcast(user,line)
	}

	globalRoom.Broadcast("System", fmt.Sprintf("%s Leave room", user))
	globalRoom.Leave(user)

}

func main() {
	fmt.Println("fuck")

}
