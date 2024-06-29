// After making changes, go mod init and go mod tidy
// (or "go mod init ___"? "caddy"?)

package main

import (
	caddycmd "github.com/caddyserver/caddy/v2/cmd"

	// plug in Caddy modules here
	_ "github.com/caddyserver/replace-response"

	_ "github.com/caddyserver/caddy/v2/modules/standard"
)

func main() {
	caddycmd.Main()
}