//go:build wireinject
// +build wireinject

package main

import (
	// "github.com/go-redis/redis/v8"
	"github.com/google/wire"
)

//TODO: add your injects here
var _ = wire.NewSet()
