package main

import (
	"github.com/KirillBogatikov/a"
	"github.com/KirillBogatikov/b"
	"github.com/KirillBogatikov/m/c"
	"log"

	ma "github.com/KirillBogatikov/m/a"
	mb "github.com/KirillBogatikov/m/b"

	"github.com/KirillBogatikov/m"
)

func main() {
	log.Println("from separated: ")

	a.CallA()
	b.CallB()
	c.CallC()

	log.Println("from submodules: ")

	ma.CallA()
	mb.CallB()

	log.Println("from m: ")

	m.CallAll()
}
