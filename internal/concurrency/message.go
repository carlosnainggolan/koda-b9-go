package concurrency

import (
	"fmt"
	"time"
)

type Message struct {
	Nama string
	Msg  string
}

func Run() {
	chn := make(chan Message)

	go Board(chn)
	message := []Message{
		{Nama: "Carlos", Msg: "Apa aja dah"},
		{Nama: "Carlos", Msg: "Apa aja dah"},
		{Nama: "Carlos", Msg: "Apa aja dah"},
		{Nama: "Carlos", Msg: "Apa aja dah"},
		{Nama: "Carlos", Msg: "Apa aja dah"},
	}

	for _, v := range message {
		Send(chn, v)
	}
}

func Send(chn chan Message, msg Message) {
	chn <- msg
}

func Board(chn chan Message) {
	for v := range chn {
		fmt.Printf("Pesan yang dari %s, isi pesan %s \n", v.Nama, v.Msg)
		time.Sleep(500 * time.Millisecond)
	}
}