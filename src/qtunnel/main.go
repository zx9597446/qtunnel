package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"tunnel"
)

func waitSignal() {
	var sigChan = make(chan os.Signal, 1)
	signal.Notify(sigChan)
	for sig := range sigChan {
		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			log.Printf("terminated by signal %v\n", sig)
			return
		} else {
			log.Printf("received signal: %v, ignore\n", sig)
		}
	}
}

func main() {
	var faddr, baddr, cryptoMethod, secret string
	var clientMode bool
	flag.StringVar(&faddr, "listen", ":9001", "host:port qtunnel listen on")
	flag.StringVar(&baddr, "backend", "127.0.0.1:6400", "host:port of the backend")
	flag.StringVar(&cryptoMethod, "crypto", "rc4", "encryption method")
	flag.StringVar(&secret, "secret", "secret", "password used to encrypt the data")
	flag.BoolVar(&clientMode, "clientmode", false, "if running at client mode")
	flag.Parse()

	log.SetOutput(os.Stdout)
	t := tunnel.NewTunnel(faddr, baddr, clientMode, cryptoMethod, secret, 4096)
	log.Println("qtunnel started.")
	go t.Start()
	waitSignal()
}
