# powerctl

Powermanctl is a simple commandline tool to control power through the powerapi.  It mostly exists to test powerapi implementations.

To build:

```
go build .
```

To run:

```
./powerctl
```

`powerctl` takes a number of options and can operate as an interactive shell.  Optionally, a single command can be passed after options.

### Options
```
$ ./powerctl -h
Usage of ./powerctl
  -base string
        base url of the API on the endpoint (default "/power/v1")
  -https
        use HTTPS instead of HTTP
  -port int
        port that the PowerAPI server is listening on (default 8269)
  -server string
        ip or hostname of the PowerAPI server (default "localhost")
```

### Commands
```
$ ./powerctl help

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
```