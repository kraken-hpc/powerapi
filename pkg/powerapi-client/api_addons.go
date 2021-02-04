/* api_addons.go: api_addons contains useful add-on, non-generated functions to the client api
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2020, Triad National Security, LLC
 * See LICENSE file for details.
 */

package powerapi

import (
	"context"
	"fmt"
	"regexp"
)

func GetAPIBase(client *APIClient, ctx context.Context) string {
	vars, e := getServerVariables(ctx)
	if e != nil {
		return ""
	}
	if a, ok := vars["apiBase"]; ok {
		// ctx defines a custom apiBase
		return a
	}
	// we need the default
	index, e := getServerIndex(ctx)
	if e != nil {
		return ""
	}
	if len(client.cfg.Servers) > index {
		return client.cfg.Servers[index].Variables["apiBase"].DefaultValue
	}
	return ""
}

// NameToURI returns the URI for a Name based on its name
// returns an empty string on failure
func NodeToURI(client *APIClient, ctx context.Context, name string) string {
	base := GetAPIBase(client, ctx)

	return fmt.Sprintf("%s/ComputerSystems/%s", base, name)
}

var reURI *regexp.Regexp

// URIToNode returns the Name of a node based on its URI
// returns an empty string on failure
func URIToNode(client *APIClient, ctx context.Context, uri string) string {
	base := GetAPIBase(client, ctx)

	reURI = regexp.MustCompile(fmt.Sprintf("^%s/ComputerSystems/([a-zA-Z0-9-]+)/?$", regexp.QuoteMeta(base)))
	m := reURI.FindAllStringSubmatch(uri, 1)
	if len(m) != 1 {
		// not valid
		return ""
	}
	return m[0][1]
}
