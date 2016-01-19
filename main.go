package main

import (
	"github.com/elastic/beats/libbeat/beat"
	"github.com/ruflin/modbeat/beater"
)

func main() {
	beat.Run("modbeat", "", beater.New())
}
