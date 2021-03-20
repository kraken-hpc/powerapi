/* main.go: entry point for the powerman/PowerAPI gateway
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2020, Triad National Security, LLC
 * See LICENSE file for details.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	api "github.com/kraken-hpc/powerapi/cmd/powermanapi/api"
)

var flags struct {
	pmip   string
	pmport int
	ip     string
	port   int
	https  bool
}

func main() {
	flag.StringVar(&flags.pmip, "pmip", "127.0.0.1", "specify the IP where powermand is listening")
	flag.IntVar(&flags.pmport, "pmport", 10101, "specify the port that powermand is listening on")
	flag.StringVar(&flags.ip, "ip", "127.0.0.1", "specify the IP address to listen on")
	flag.IntVar(&flags.port, "port", 8269, "specify the TCP port to listen on")
	flag.Parse()

	log.Printf("starting powermanapi service on %s:%d talking to powermand at %s:%d", flags.ip, flags.port, flags.pmip, flags.pmport)

	PowermanApiService := api.NewPowermanApiService(flags.pmip, flags.pmport)
	DefaultApiController := api.NewDefaultApiController(PowermanApiService)

	router := api.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", flags.ip, flags.port), router))
}
