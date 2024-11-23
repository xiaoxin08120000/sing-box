//go:build !generate

package main

import "github.com/xiaoxin08120000/sing-box/log"

func main() {
	if err := mainCommand.Execute(); err != nil {
		log.Fatal(err)
	}
}
