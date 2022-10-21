package main

type ServerRole int

const (
	candidate ServerRole = iota
	follower
	leader
)

type State struct {
	Role    ServerRole
	Me      Server
	Servers []Server
	Log     [][]byte
}
