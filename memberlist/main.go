package main

import (
	"flag"
	"log"
	"os"

	"github.com/hashicorp/memberlist"
)

var (
	name string
	port int
	join string
)

func main() {
	fs := flag.NewFlagSet("", flag.PanicOnError)
	fs.StringVar(&name, "name", "", "Name of this node")
	fs.IntVar(&port, "port", 0, "Port to listen on")
	fs.StringVar(&join, "join", "", "Address of node to join")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	conf := memberlist.DefaultLocalConfig()
	conf.Name = name
	conf.BindPort = port
	conf.AdvertisePort = port

	list, err := memberlist.Create(conf)
	if err != nil {
		log.Fatal(err)
	}

	node := list.LocalNode()
	log.Printf("Node at %s:%s\n", node.Address(), node.Port)

	if join != "" {
		_, err := list.Join([]string{join})
		if err != nil {
			log.Fatal(err)
		}
	}

	for _, member := range list.Members() {
		log.Printf("Member: %s(%s:%d)\n", member.Name, member.Addr.To4().String(), member.Port)
	}

	select {}
}
