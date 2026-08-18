//go:build generate

package main

//go:generate go run github.com/bakito/crossplane-function-inventory-sdk/cmd/tag-const-gen -i fn.go
