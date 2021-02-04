/* main.go: pmshell CLI go-based powerman utility
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2020, Triad National Security, LLC
 * See LICENSE file for details.
 */

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	pm "github.com/hpc/powerapi/pkg/powerman"
)

var flags struct {
	server string
	port   int
	ipv6   bool
}

var inShell bool
var conn *pm.Connection
var prompt = "pm> "

func help() {
	fmt.Printf(`
commands:
  shell
        Start an interactve shell (default, ignored if already in a shell)
  on|1 <node>
        Turn a node on
  off|0 <node>
        Turn a node off
  [c]ycle <node>
        Power cycle a node
  [q]uery [<node>]
		Get node state.  If <ndoe> is not specified, get the state of all nodes.
  [l]list|ls
		List all known nodes; do not query their statuses.
  quit|exit
        Quit the shell
  help|?
        Print this help
`)
}

func usage() {
	fmt.Printf(`
pmshell provides a shell-like interface to a powerman server.

Usage: %s <options> [cmd]

options:
`, os.Args[0])
	flag.PrintDefaults()
	help()
}

func exit() {
	if conn != nil {
		conn.Disconnect()
	}
	fmt.Println("bye!")
	os.Exit(0)
}

func query(cmd []string) error {
	switch len(cmd) {
	case 1:
		// query all
		nodes := conn.All()
		for _, n := range nodes {
			s, e := conn.NodeStatus(n)
			if e != nil {
				return e
			}
			fmt.Println(n + ": " + s.String())
		}
	case 2:
		// query one
		s, e := conn.NodeStatus(cmd[1])
		if e != nil {
			return e
		}
		fmt.Println(s)
	default:
		// incorrect
		return fmt.Errorf("query: incorrect number of arguments. Usage: query [<node>]")
	}
	return nil
}

func runShell(cmd []string) error {
	if inShell {
		return fmt.Errorf("you're already in a shell")
	}
	err := func(e error) {
		if e != nil {
			fmt.Printf("error: %v\n", e)
		}
	}
	inShell = true
	if len(cmd) != 0 && cmd[0] != "shell" {
		// we were given a command to run
		err(runCommand(cmd))
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s", prompt)
		scanner.Scan()
		s := scanner.Text()
		s = strings.Trim(s, "\n \t")
		cmd := strings.Split(s, " ")
		err(runCommand(cmd))
	}
}

func list(cmd []string) error {
	if len(cmd) != 1 {
		return fmt.Errorf("list: incorrect number of arguments. Usage: list")
	}
	nodes := conn.All()
	for _, n := range nodes {
		fmt.Println(n)
	}
	return nil
}

func on(cmd []string) error {
	if len(cmd) != 2 {
		return fmt.Errorf("on: incorrect number of arguments. Usage: on <node>")
	}
	return conn.NodeOn(cmd[1])
}

func off(cmd []string) error {
	if len(cmd) != 2 {
		return fmt.Errorf("off: incorrect number of arguments. Usage: off <node>")
	}
	return conn.NodeOff(cmd[1])
}

func cycle(cmd []string) error {
	if len(cmd) != 2 {
		return fmt.Errorf("cycle: incorrect number of arguments. Usage: cycle <node>")
	}
	return conn.NodeCycle(cmd[1])
}

func runCommand(c []string) error {
	if len(c) == 0 || c[0] == "" {
		// somenoe just hit enter?
		return nil
	}
	cmd := c[0]
	switch cmd {
	case "shell":
		return runShell(c)
	case "on", "1":
		return on(c)
	case "off", "0":
		return off(c)
	case "cycle", "c":
		return cycle(c)
	case "list", "l", "ls":
		return list(c)
	case "query", "q":
		return query(c)
	case "quit", "exit":
		exit()
		return nil
	case "help", "?":
		help()
		return nil
	default:
		return fmt.Errorf("error: unknown command, %s, help or ? for available commands", cmd)
	}
}

func usageExit(format string, a ...interface{}) {
	fmt.Printf("error: "+format+"\n", a...)
	usage()
	os.Exit(1)
}

func main() {
	var e error
	flag.StringVar(&flags.server, "server", "localhost", "specifies the powermand server address")
	flag.IntVar(&flags.port, "port", 10101, "specifies the powermand server port")
	flag.BoolVar(&flags.ipv6, "ipv6", false, "force using ipv6 to connect")
	flag.Parse()

	// try to connect
	cflags := 0
	if flags.ipv6 {
		cflags = pm.PM_CONN_INET6
	}
	if conn, e = pm.Connect(fmt.Sprintf("%s:%d", flags.server, flags.port), cflags); e != nil {
		usageExit("server connection failed: %v", e)
	}
	runShell(flag.Args())
}
