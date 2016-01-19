package mod

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"
	"github.com/ruflin/modbeat/config"
	"os"
	"path/filepath"
	"time"
)

func Scanner(config config.PathConfig, client publisher.Client) {

	duration, _ := time.ParseDuration(config.Period)

	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for {
		select {

		case <-ticker.C:
		}

		events := []common.MapStr{}

		logp.Debug("mod", "Scanning glob: %s", config.Glob)

		files, _ := filepath.Glob(config.Glob)
		for _, file := range files {
			logp.Debug("mod", "Get stats for file: %s", file)
			stat, _ := os.Lstat(file)

			event := common.MapStr{
				"@timestamp": common.Time(time.Now()),
				"type":       "modbeat",
				"name":       stat.Name(),
				"path":       file,
				"mode":       stat.Mode(),
				"modtime":    stat.ModTime(),
				"size":       stat.Size(),
			}
			events = append(events, event)
		}
		client.PublishEvents(events)
		logp.Debug("mod", "Events sent for %s files", len(events))
	}
}
