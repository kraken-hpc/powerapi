/* powerman.go: go bindings for libpowerman
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2020, Triad National Security, LLC
 * See LICENSE file for details.
 */
// +build linux

package powerman

/*
#include <stdio.h>
#include <stddef.h>
#include <stdlib.h>
#include <errno.h>
#include <libpowerman.h>
#cgo LDFLAGS: -lpowerman
*/
import "C"
import (
	"fmt"
	"unsafe"
)

var errorMap = map[C.pm_err_t]error{
	C.PM_ENOADDR:      fmt.Errorf("failed to get address info for server"),
	C.PM_ECONNECT:     fmt.Errorf("connect failed"),
	C.PM_ENOMEM:       fmt.Errorf("out of memory"),
	C.PM_EBADHAND:     fmt.Errorf("bad server handle"),
	C.PM_ESERVEREOF:   fmt.Errorf("received unexpected EOF from server"),
	C.PM_ESERVERPARSE: fmt.Errorf("unexpected response from server"),
	C.PM_EUNKNOWN:     fmt.Errorf("server: unknown command (201)"),
	C.PM_EPARSE:       fmt.Errorf("server: parse error (202)"),
	C.PM_ETOOLONG:     fmt.Errorf("server: command too long (203)"),
	C.PM_EINTERNAL:    fmt.Errorf("server: internal error (204)"),
	C.PM_EHOSTLIST:    fmt.Errorf("server: hostlist error (205)"),
	C.PM_EINPROGRESS:  fmt.Errorf("server: command in progress (208)"),
	C.PM_ENOSUCHNODES: fmt.Errorf("server: no such nodes (209)"),
	C.PM_ECOMMAND:     fmt.Errorf("server: command completed with errors (210)"),
	C.PM_EQUERY:       fmt.Errorf("server: query completed with errors (211)"),
	C.PM_EUNIMPL:      fmt.Errorf("server: not implemented by device (213)"),
}

// Status is a value describing the state of a node
type Status C.pm_node_state_t

const (
	PM_UNKNOWN Status = C.PM_UNKNOWN
	PM_OFF     Status = C.PM_OFF
	PM_ON      Status = C.PM_ON
)

const PM_CONN_INET6 = C.PM_CONN_INET6

// StatusName maps node Status to a string
var StatusName = map[Status]string{
	PM_UNKNOWN: "UNKNOWN",
	PM_OFF:     "OFF",
	PM_ON:      "ON",
}

// StatusValue maps a string back to a Status Value
var StatusValue = map[string]Status{
	"UNKNOWN": PM_UNKNOWN,
	"OFF":     PM_OFF,
	"ON":      PM_ON,
}

// String gives a string representation of a Status
func (s Status) String() string {
	return StatusName[s]
}

func toError(r C.pm_err_t, errno error) error {
	switch r {
	case C.PM_ESUCCESS:
		return nil
	case C.PM_ERRNOVALID:
		return fmt.Errorf("system call failed: %v", errno)
	default:
		if _, ok := errorMap[r]; !ok {
			return fmt.Errorf("unknown error: %d", r)
		}
		return errorMap[r]
	}
}

// A Connection represents a connection to a powerman server
// To initialize a connection, call Connect()
// Connections should be Disconnect()'ed when done
type Connection struct {
	h      *C.pm_handle_t
	is     []*Iterator
	active bool
}

// Connect to the powerman service
func Connect(server string, flags int) (*Connection, error) {
	c := &Connection{}
	c.is = []*Iterator{}
	c.h = (*C.pm_handle_t)(C.malloc(C.sizeof_pm_handle_t))
	cs := C.CString(server)
	defer C.free(unsafe.Pointer(cs))
	r, errno := C.pm_connect(cs, nil, c.h, C.int(flags))
	if e := toError(r, errno); e != nil {
		C.free(unsafe.Pointer(c.h))
		return nil, e
	}
	c.active = true
	// we create a default iterator
	i, _ := c.NodeIteratorCreate()
	c.is = append(c.is, i)
	return c, nil
}

// Disconnect tears down the server connection and frees resources
func (c *Connection) Disconnect() {
	if !c.active {
		return
	}
	C.pm_disconnect(*c.h)
	C.free(unsafe.Pointer(c.h))
	for _, i := range c.is {
		i.destroy()
	}
	c.active = false
}

// NodeStatus returns the status of a specific node by name
func (c *Connection) NodeStatus(node string) (Status, error) {
	if !c.active {
		return PM_UNKNOWN, fmt.Errorf("NodeStatus called on an inactive connection")
	}
	ns := C.CString(node)
	defer C.free(unsafe.Pointer(ns))
	s := new(Status)
	r, errno := C.pm_node_status(*c.h, ns, (*C.pm_node_state_t)(s))
	if e := toError(r, errno); e != nil {
		return PM_UNKNOWN, e
	}
	return *s, nil
}

// NodeOn tells powerman to turn on a node
func (c *Connection) NodeOn(node string) error {
	if !c.active {
		return fmt.Errorf("NodeOn called on an inactive connection")
	}
	ns := C.CString(node)
	defer C.free(unsafe.Pointer(ns))
	r, errno := C.pm_node_on(*c.h, ns)
	return toError(r, errno)
}

// NodeOff tells powerman to turn a node off
func (c *Connection) NodeOff(node string) error {
	if !c.active {
		return fmt.Errorf("NodeOff called on an inactive connection")
	}
	ns := C.CString(node)
	defer C.free(unsafe.Pointer(ns))
	r, errno := C.pm_node_off(*c.h, ns)
	return toError(r, errno)
}

// NodeCycle tells powerman to cycle a node
func (c *Connection) NodeCycle(node string) error {
	if !c.active {
		return fmt.Errorf("NodeCycle called on an inactive connection")
	}
	ns := C.CString(node)
	defer C.free(unsafe.Pointer(ns))
	r, errno := C.pm_node_cycle(*c.h, ns)
	return toError(r, errno)
}

// Next gets the next node in the default iterator
func (c *Connection) Next() (node string, end bool) {
	return c.is[0].Next()
}

// All uses the default iterator to get a list of all known nodes
func (c *Connection) All() []string {
	return c.is[0].All()
}

// Reset resets the default iterator to the first node
func (c *Connection) Reset() {
	c.is[0].Reset()
}

type Iterator struct {
	c      *Connection
	i      *C.pm_node_iterator_t
	active bool
}

// NodeIteratorCreate initializes an iterator that can iterate through all node names
func (c *Connection) NodeIteratorCreate() (*Iterator, error) {
	if !c.active {
		return nil, fmt.Errorf("NodeIteratorCreate called on an inactive connection")
	}
	i := &Iterator{c: c}
	i.i = (*C.pm_node_iterator_t)(C.malloc(C.sizeof_pm_node_iterator_t))
	r, errno := C.pm_node_iterator_create(*c.h, i.i)
	if e := toError(r, errno); e != nil {
		return nil, e
	}
	i.active = true
	c.is = append(c.is, i)
	return i, nil
}

// Next returns the next node name, or returns end == true if there are no more nodes
func (i *Iterator) Next() (node string, end bool) {
	if !i.active || !i.c.active {
		// no errors, so just return end
		return "", true
	}
	r, _ := C.pm_node_next(*i.i)
	if r == nil {
		return "", true
	}
	return C.GoString(r), false
}

// Reset rewinds the iterator to the beginning of the list
func (i *Iterator) Reset() {
	if !i.active || !i.c.active {
		return //nop
	}
	C.pm_node_iterator_reset(*i.i)
}

// All gets a list of all known nodes
func (i *Iterator) All() (nodes []string) {
	i.Reset()
	for {
		n, end := i.Next()
		if end {
			break
		}
		nodes = append(nodes, n)
	}
	return
}

// destroy destroys an iterator and cleans up resources
// Note: this also destroys the connection
// because of the interaction with pm_disconnect, we don't export this function
// instead, we collect all iterators when Disconnect is called
func (i *Iterator) destroy() {
	if !i.active || !i.c.active {
		return // nop
	}
	C.pm_node_iterator_destroy(*i.i)
	C.free(unsafe.Pointer(i.i))
	i.active = false
	i.c.active = false
}

func main() {
	c, e := Connect("localhost", 0)
	if e != nil {
		fmt.Printf("connect error: %v\n", e)
		return
	}
	defer c.Disconnect()
	fmt.Printf("connect succeeded\n")
	node := "blah"
	s, e := c.NodeStatus(node)
	if e != nil {
		fmt.Printf("failed to get node status for %s: %v\n", node, e)
	} else {
		fmt.Printf("node status for %s: %s\n", node, s)
	}
	fmt.Println("Node list:")
	i, e := c.NodeIteratorCreate()
	if e != nil {
		fmt.Printf("failed to create node iterator: %v\n", e)
	}
	for {
		n, end := i.Next()
		if end {
			break
		}
		fmt.Println(n)
	}
}
