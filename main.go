package main

import (
	"fmt"
	"os"
	"os/signal"
	"net"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/rpc"
)

func listenSignal(exit chan int) {
	sig := make(chan os.Signal)
	signal.Notify(sig)

	for {
		<-sig
		exit <- 1
	}
}


type Service struct {}

// this will be called by kfchen_echo
func (this *Service) Echo(msg string) string {
	return fmt.Sprintf("%s from echo", msg)
}

// this will be called by kfchen_getPerson
func (this *Service) GetPerson(params interface{}) map[string]interface{} {
	p, ok := params.(map[string]interface{})
	if !ok {
		log.Fatal("fail")
	}

	name := p["name"].(string)

	if name == "kfchen" {
		return map[string]interface{}{
			"name": "kfchen",
			"desc": "good man",
		}
	} else {
		return map[string]interface{}{
			"invalid": false,
		}
	}
}


var endpoint string = "127.0.0.1:9010"
var server *rpc.Server

func startRPCServer() {
	server = rpc.NewServer()
	service := new(Service)
	if err := server.RegisterName("kfchen", service); err != nil {
		panic(err)
	}

	// All APIs registered, start the HTTP listener
	var (
		listener net.Listener
		err      error
	)
	if listener, err = net.Listen("tcp", endpoint); err != nil {
		panic(err)
	}

	cors := make([]string, 0)
	vhosts := make([]string, 0)
	go rpc.NewHTTPServer(cors, vhosts, server).Serve(listener)
	log.Println("RPC HTTP endpoint opened", "url", fmt.Sprintf("http://%s", endpoint), "cors", strings.Join(cors, ","), "vhosts", strings.Join(vhosts, ","))
	log.Println("Press Ctrl+C to exit")
}


func main() {
	startRPCServer()

	exit := make(chan int)
	<-exit

	server.Stop()
}