package config

import (
	"log"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v3"
)

// Config represents the structure of the configuration file
type Config struct {
	URLs     []string `yaml:"urls"`
	Interval int      `yaml:"interval"`
	Timeout  int      `yaml:"timeout"`
	Logfile  string   `yaml:"logfile"`
}

var (
	config *Config
	mux    sync.RWMutex
)

// ReadConfig reads the configuration from the specified YAML file.
func ReadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// LoadConfig loads the configuration and stores it in a package-level variable.
func LoadConfig(path string) (*Config, error) {
	cfg, err := ReadConfig(path)
	if err != nil {
		return nil, err
	}
	mux.Lock()
	config = cfg
	mux.Unlock()
	return cfg, nil
}

// GetConfig returns the current configuration.
func GetConfig() *Config {
	mux.RLock()
	defer mux.RUnlock()
	return config
}

// WatchConfig sets up a file watcher on the configuration file to reload it on changes.
func WatchConfig(path string, onChange func(*Config)) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}

	if err = watcher.Add(path); err != nil {
		watcher.Close()
		return err
	}

	go func() {
		defer watcher.Close()
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Config file modified, reloading...")
					newCfg, err := ReadConfig(path)
					if err == nil {
						mux.Lock()
						config = newCfg
						mux.Unlock()
						if onChange != nil {
							onChange(newCfg)
						}
					} else {
						log.Println("Error reloading config:", err)
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	return nil
}
