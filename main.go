package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"wz_id_generator/conf"
	"wz_id_generator/id_generator"
	"wz_rpc_go"
)

func main() {
	id := id_generator.NewIdGenerator()
	addr := fmt.Sprintf("%s:%d", "127.0.0.1", conf.C.App.Port)

	s := wz_rpc_go.NewServer(addr, func(conn net.Conn) wz_rpc_go.Codec {
		return &wz_rpc_go.JsonCodec{
			Enc: json.NewEncoder(conn.(io.Writer)),
			Dec: json.NewDecoder(conn.(io.Reader)),
		}
	})

	err := s.Register(id)
	if err != nil {
		fmt.Printf("register err: %v\n", err)
		return
	}

	fmt.Printf("serving...\n")
	err = s.Serve()
	if err != nil {
		fmt.Printf("serve err: %v\n", err)
		return
	}

}
