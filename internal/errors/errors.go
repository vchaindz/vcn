/*
 * Copyright (c) 2018-2019 vChain, Inc. All Rights Reserved.
 * This software is released under GPL3.
 * The full license information can be found under:
 * https://www.gnu.org/licenses/gpl-3.0.en.html
 *
 */

package errors

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/vchain-us/vcn/pkg/meta"
)

func formatErrorURLCustom(domain string, status int) string {

	errorPage := meta.ErrorWikiURL()

	return fmt.Sprintf("%s%s-%d", errorPage, domain, status)

}

func formatErrorURLByEndpoint(resource string, verb string, status int) string {

	errorPage := meta.ErrorWikiURL()

	// get last part of endpoint
	x := strings.Split(resource, "/")
	resource = x[len(x)-1]

	return fmt.Sprintf("%s%s-%s-%d", errorPage, resource, strings.ToLower(verb), status)

}

// PrintErrorURLCustom takes custom domain and status code
func PrintErrorURLCustom(domain string, code int) {

	fmt.Print("Get help for this error at:\n")

	color.Set(meta.StyleError())
	fmt.Print(formatErrorURLCustom(domain, code))
	color.Unset()

	fmt.Println()
	return

}
