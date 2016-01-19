package beater

import (
	"fmt"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/cfgfile"
	"github.com/ruflin/modbeat/config"
	"github.com/ruflin/modbeat/mod"
)

type Modbeat struct {
	done chan struct{}
}

func New() *Modbeat {
	return &Modbeat{
		done: make(chan struct{}),
	}
}

func (mb *Modbeat) Config(b *beat.Beat) error {

	cfg := &config.ModbeatConfig{}
	err := cfgfile.Read(&cfg, "")
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	// Start all scanners
	for _, pathConfig := range cfg.Modbeat.Paths {
		// TODO: Handle proper shutdown -> waitgroups and done channel
		go func(pathConfig config.PathConfig) {
			mod.Scanner(pathConfig, b.Events)
		}(pathConfig)
	}

	<-mb.done

	return nil
}

func (mb *Modbeat) Setup(b *beat.Beat) error {
	return nil
}

func (mb *Modbeat) Run(b *beat.Beat) error {
	fmt.Println("Hello World")
	return nil
}

func (mb *Modbeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (mb *Modbeat) Stop() {
	close(mb.done)
}
