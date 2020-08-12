package main

import (
	"flag"
	"github.com/bukowa/goutils/pkg"
)

func main() {
	s := pkg.NewHttpServer()
	s.SetFlags()
	flag.Parse()
	s.ListenAndServe(nil)
}
