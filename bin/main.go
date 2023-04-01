package main

import (
	"net"
	"os"
	"strings"
)

const name = "sstp.vim"
const version = "0.1.0"

type Speach struct {
	Sender  string
	Event   string
	Script  string
	Option  string
	Charset string
}

func (s *Speach) String() string {
	return "NOTIFY SSTP/1.5\r\n" +
		"Sender: " + s.Sender + "\r\n" +
		"Event: " + s.Event + "\r\n" +
		"Script: " + s.Script + "\r\n" +
		"Option: " + s.Option + "\r\n" +
		"Charset: " + s.Charset + "\r\n" +
		"\r\n"
}

func main() {
	msg := strings.Join(os.Args[1:], " ")
	if err := Speak(msg); err != nil {
		panic(err)
	}
}

func Speak(msg string) error {
	if msg == "" {
		return nil
	}

	conn, err := net.Dial("tcp", "localhost:9801")
	if err != nil {
		return err
	}

	test := Speach{
		Sender:  name + "/v" + version,
		Script:  msg,
		Option:  "notranslate",
		Charset: "UTF-8",
	}

	_, err = conn.Write([]byte(test.String()))
	if err != nil {
		return err
	}

	_ = conn.Close()

	return nil
}
