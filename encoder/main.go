package main

import (
	"math/rand"

	log "github.com/sirupsen/logrus"
)

func main() {
	r := rand.Intn(100)

	log.Info(r)
}
