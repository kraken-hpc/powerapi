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
	"log"
	"net/http"

	api "github.com/hpc/powerapi/cmd/powermanapi/api"
)

func main() {
	log.Printf("Server started")

	PowermanApiService := api.NewPowermanApiService("127.0.0.1", 10101)
	DefaultApiController := api.NewDefaultApiController(PowermanApiService)

	router := api.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8269", router))
}
