package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	api "github.com/jlowellwofford/powermanapi/pkg/powerapi-client"
)

var flags struct {
	server  string
	port    int
	https   bool
	apiBase string
}

// proto maps flag.https to the index of http/https
var proto = map[bool]int{
	false: 1,
	true:  0,
}

var client *api.APIClient
var ctx context.Context
var inShell = false

var prompt = "pm> "

func help() {
	fmt.Printf(`
commands:
  shell
        Start an interactve shell (default, ignored if already in a shell)
  on|1 <node>[ <node> ...]
        Turn a node(s) on
  off|0 <node>[ <node> ...]
        Turn a node(s) off
  [c]ycle <node>[ <node> ...]
        Power cycle node(s)
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
	fmt.Println("bye!")
	os.Exit(0)
}

func query(cmd []string) error {
	switch len(cmd) {
	case 1:
		// query all
		cs, r, e := client.DefaultApi.ComputerSystemsGet(ctx).Execute()
		if e != nil {
			return fmt.Errorf("api call failed: err(%v) response(%v)", e, r)
		}
		for _, n := range *cs.Systems {
			fmt.Printf("%s: %s\n", n.Name, *n.PowerState)
		}
	case 2:
		// query one
		cs, r, e := client.DefaultApi.ComputerSystemsNameGet(ctx, cmd[1]).Execute()
		if e != nil {
			return fmt.Errorf("api call failed: err(%v) response(%v)", e, r)
		}
		fmt.Println(cs.GetPowerState())
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
	cs, r, e := client.DefaultApi.ComputerSystemsGet(ctx).Execute()
	if e != nil {
		return fmt.Errorf("api call failed: err(%v) response(%v)", e, r)
	}
	for _, n := range *cs.Systems {
		fmt.Println(n.Name)
	}
	return nil
}

func nodeToURI(name string) string {
	return fmt.Sprintf("%s/ComputerSystems/%s", flags.apiBase, name)
}

var reURI *regexp.Regexp

func uriToNode(uri string) string {
	m := reURI.FindAllStringSubmatch(uri, 1)
	if len(m) != 1 {
		// not valid
		return ""
	}
	return m[0][1]
}

func reset(cmd []string, t api.ResetType) error {
	if len(cmd) < 2 {
		return fmt.Errorf("incorrect number of arguments. Usage: <on|off|cycle> <node> [<node> ...]")
	}
	if len(cmd) == 2 { // use a singleton call
		body := api.NewResetRequestBody(t)
		_, r, e := client.DefaultApi.ComputerSystemsNameActionsComputerSystemResetPost(ctx, cmd[1]).ResetRequestBody(*body).Execute()
		if e != nil {
			return fmt.Errorf("api call failed: err(%v) response(%v)", e, r)
		}
	} else { // use an aggregate call
		ns := []string{}
		for _, n := range cmd[1:] {
			ns = append(ns, nodeToURI(n))
		}
		body := api.NewAggregationResetBody()
		body.ResetType = &t
		body.TargetURIs = &ns
		ar, r, e := client.DefaultApi.AggregationServiceActionsAggregationServiceResetPost(ctx).AggregationResetBody(*body).Execute()
		if e != nil {
			return fmt.Errorf("api call failed: err(%v) response(%v)", e, r)
		}
		fmt.Printf("succeed:")
		for _, n := range *ar.TargetURIs {
			fmt.Printf(" %s", uriToNode(n))
		}
		fmt.Printf("\n")
	}
	return nil
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
		return reset(c, api.RESETTYPE_FORCE_ON)
	case "off", "0":
		return reset(c, api.RESETTYPE_FORCE_OFF)
	case "cycle", "c":
		return reset(c, api.RESETTYPE_POWER_CYCLE)
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
	flag.StringVar(&flags.server, "server", "localhost", "ip or hostname of the PowerAPI server")
	flag.IntVar(&flags.port, "port", 8269, "port that the PowerAPI server is listening on")
	flag.BoolVar(&flags.https, "https", false, "use HTTPS instead of HTTP")
	flag.StringVar(&flags.apiBase, "base", "/power/v1", "base url of the API on the endpoint")
	flag.Parse()

	ctx = context.Background()
	ctx = context.WithValue(ctx, api.ContextServerIndex, proto[flags.https]) // use http not https
	ctx = context.WithValue(ctx, api.ContextServerVariables, map[string]string{
		"server":  fmt.Sprintf("%s:%d", flags.server, flags.port),
		"apiBase": flags.apiBase,
	}) // set server host/port & apiBase

	configuration := api.NewConfiguration()
	client = api.NewAPIClient(configuration)
	reURI = regexp.MustCompile(fmt.Sprintf("^%s%s([a-zA-Z0-9-]+)/?$", regexp.QuoteMeta(flags.apiBase), regexp.QuoteMeta("/ComputerSystems/")))

	runShell(flag.Args())
}
