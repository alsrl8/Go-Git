package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Config map[string]string

// PrintConfig prints configures
func (c *Config) PrintConfig() {
	for key, value := range *c {
		fmt.Printf("%s = %s\n", key, value)
	}
}

// GetRootDirectory returns root directory path.
// If the rootDir is not configured, an error is returned.
func (c *Config) GetRootDirectory() (string, error) {
	rootDir, has := (*c)["rootDir"]
	if !has {
		return "", fmt.Errorf("no rootDir configure")
	}
	return rootDir, nil
}

// LoadConfig loads configure file with given filepath
func LoadConfig(configFilename string) (Config, error) {
	config := Config{}

	file, err := os.Open(configFilename)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") { // 주석 혹은 빈 줄은 건너뛴다.
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid config line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		config[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return config, nil
}
