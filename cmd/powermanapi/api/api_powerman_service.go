/* api_powerman_service.go: Powerman API service
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2020, Triad National Security, LLC
 * See LICENSE file for details.
 */

package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"

	pm "github.com/kraken-hpc/powerapi/pkg/powerman"
)

const urlBase = "/power/v1"

// PowermanApiService is a service that implents the logic for the DefaultApiServicer
type PowermanApiService struct {
	server string
	port   int
}

// NewPowermanApiService creates a default api service
func NewPowermanApiService(server string, port int) DefaultApiServicer {
	return &PowermanApiService{
		server: server,
		port:   port,
	}
}

func (s *PowermanApiService) connect() (*pm.Connection, error) {
	return pm.Connect(fmt.Sprintf("%s:%d", s.server, s.port), 0)
}

var pmToAPI = map[pm.Status]PowerState{
	pm.PM_ON:  POWERSTATE_ON,
	pm.PM_OFF: POWERSTATE_OFF,
}

var ERROR_NOTSUPPORTED = errors.New("operation not supported")

func (s *PowermanApiService) resetNode(conn *pm.Connection, name string, r ResetType) (ResetType, error) {
	switch r {
	case RESETTYPE_FORCE_OFF:
		if e := conn.NodeOff(name); e != nil {
			return ResetType(0), e
		}
		return RESETTYPE_FORCE_OFF, nil
	case RESETTYPE_FORCE_ON, RESETTYPE_ON:
		if e := conn.NodeOn(name); e != nil {
			return ResetType(0), e
		}
		return RESETTYPE_ON, nil
	case RESETTYPE_FORCE_RESTART, RESETTYPE_POWER_CYCLE:
		if e := conn.NodeCycle(name); e != nil {
			return ResetType(0), e
		}
		return RESETTYPE_POWER_CYCLE, nil
	case RESETTYPE_GRACEFUL_RESTART, RESETTYPE_PUSH_POWER_BUTTON, RESETTYPE_GRACEFUL_SHUTDOWN, RESETTYPE_NMI:
		return ResetType(0), ERROR_NOTSUPPORTED
	}
	return ResetType(0), ERROR_NOTSUPPORTED
}

func (s *PowermanApiService) nodeToURI(name string) string {
	return fmt.Sprintf("%s/ComputerSystems/%s", urlBase, name)
}

var reURI = regexp.MustCompile(fmt.Sprintf("^%s%s([a-zA-Z0-9-]+)/?$", regexp.QuoteMeta(urlBase), regexp.QuoteMeta("/ComputerSystems/")))

func (s *PowermanApiService) uriToNode(uri string) string {
	m := reURI.FindAllStringSubmatch(uri, 1)
	if len(m) != 1 {
		// not valid
		return ""
	}
	return m[0][1]
}

// AggregationServiceActionsAggregationServiceResetPost - Request aggregate system reset
func (s *PowermanApiService) AggregationServiceActionsAggregationServiceResetPost(ctx context.Context, aggregationResetBody AggregationResetBody) (ImplResponse, error) {
	conn, e := s.connect()
	defer conn.Disconnect()
	if e != nil {
		return Response(http.StatusInternalServerError, nil), errors.New("failed to connect to powerman server:" + e.Error())
	}
	ret := AggregationResetBody{
		ResetType:  aggregationResetBody.ResetType,
		TargetURIs: []string{},
	}
	for _, uri := range aggregationResetBody.TargetURIs {
		n := s.uriToNode(uri)
		if n == "" {
			// not a valid node, we just skip it
			continue
		}
		if t, e := s.resetNode(conn, n, aggregationResetBody.ResetType); e != nil {
			// failed, we won't report it
			continue
		} else {
			ret.ResetType = t
		}
		ret.TargetURIs = append(ret.TargetURIs, uri)
	}

	// we actually always return 200
	return Response(200, ret), nil
}

// ComputerSystemsGet - Get computer systems
func (s *PowermanApiService) ComputerSystemsGet(ctx context.Context) (ImplResponse, error) {
	conn, e := s.connect()
	defer conn.Disconnect()
	if e != nil {
		return Response(http.StatusInternalServerError, nil), errors.New("failed to connect to powerman server:" + e.Error())
	}
	csc := ComputerSystemCollection{
		Id:      urlBase + "/ComputerSystems",
		Name:    "powerman nodes",
		Systems: []ComputerSystem{},
	}
	conn.Reset()
	for {
		n, end := conn.Next()
		if end {
			break
		}
		stat, e := conn.NodeStatus(n)
		if e != nil {
			return Response(http.StatusInternalServerError, nil), errors.New("failed to get node state:" + e.Error())
		}
		csc.Systems = append(csc.Systems, ComputerSystem{
			Id:         s.nodeToURI(n),
			Name:       n,
			PowerState: pmToAPI[stat],
		})
	}
	return Response(200, csc), nil
}

// ComputerSystemsNameActionsComputerSystemResetPost - Request system reset
func (s *PowermanApiService) ComputerSystemsNameActionsComputerSystemResetPost(ctx context.Context, name string, resetRequestBody ResetRequestBody) (ImplResponse, error) {
	conn, e := s.connect()
	defer conn.Disconnect()
	if e != nil {
		return Response(http.StatusInternalServerError, nil), errors.New("failed to connect to powerman server:" + e.Error())
	}

	t, e := s.resetNode(conn, name, resetRequestBody.ResetType)
	if e != nil {
		// TODO don't just return 500 for every error
		return Response(http.StatusNotImplemented, nil), errors.New("node reset failed: " + e.Error())
	}
	return Response(200, ResetRequestBody{ResetType: t}), nil
}

// ComputerSystemsNameGet - Get a specific computer system state
func (s *PowermanApiService) ComputerSystemsNameGet(ctx context.Context, name string) (ImplResponse, error) {
	conn, e := s.connect()
	defer conn.Disconnect()
	if e != nil {
		return Response(http.StatusInternalServerError, nil), errors.New("failed to connect to powerman server:" + e.Error())
	}

	stat, e := conn.NodeStatus(name)
	if e != nil {
		return Response(http.StatusNotFound, nil), errors.New("failed to get node status: " + e.Error())
	}

	return Response(200, ComputerSystem{
		Id:         urlBase + "/ComputerSystems/" + name,
		Name:       name,
		PowerState: pmToAPI[stat],
	}), nil
}
