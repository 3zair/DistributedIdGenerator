package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"time"
	"wz_id_generator/conf"
	"wz_id_generator/wz_id"
	"wz_rpc_go"
)

func init() {
	defaultT, err := time.Parse("2006-01-02 15:04:05", conf.C.App.StartTimeStr)
	if err != nil {
		fmt.Printf("default time parse err: %v\n", err)
		return
	}

	wz_id.DEFAULT_T = defaultT.Unix()
}

func main() {
	id := &wz_id.WzID{}
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
