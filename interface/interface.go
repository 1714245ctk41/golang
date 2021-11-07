package main

type MakeGun interface {
	design(name string)
	buildGun(name string) string
}
