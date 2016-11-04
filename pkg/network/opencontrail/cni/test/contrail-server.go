// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type vmResp struct {
	Vm   string `json:"vm"`
	Ip   string `json:"ip-address"`
	Plen int    `json:"plen"`
	Gw   string `json:"gateway"`
	Dns  string `json:"dns-server"`
	Mac  string `json:"mac-address"`
}

type statusResp struct {
	Status string `json:"status"`
}

var addr int = 3

func vmServer(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		fmt.Println("GET request")
		ip := "1.1.1." + strconv.Itoa(addr)
		addr += 1
		resp := vmResp{Vm: "VM", Ip: ip, Plen: 24, Gw: "1.1.1.1",
			Dns: "1.1.1.2", Mac: "00:00:00:00:00:01"}
		msg, _ := json.Marshal(resp)
		io.WriteString(w, string(msg))
		return

	case "POST":
		fmt.Println("POST request")
		resp := statusResp{Status: "OK"}
		msg, _ := json.Marshal(resp)
		io.WriteString(w, string(msg))
		return

	case "DELETE":
		fmt.Println("DELETE request")
		resp := statusResp{Status: "OK"}
		msg, _ := json.Marshal(resp)
		io.WriteString(w, string(msg))
		return
	default:
		fmt.Println("Unkown command")
	}
}

func main() {
	http.HandleFunc("/port/", vmServer)
	http.HandleFunc("/port", vmServer)
	http.ListenAndServe(":9060", nil)
}
