package config

import (
	"Go-Git/gitlog"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Config map[string]string

// GetRootDirectory returns root directory path.
// It is guaranteed it returns validate root directory.
func (c *Config) GetRootDirectory() string {
	rootDir := (*c)["rootDir"]
	return rootDir
}

// SetConfig loads and verifies configuration
// from the provided filename, then sets up git branches.
func SetConfig(configFilename string) Config {
	config, err := LoadConfig(configFilename)
	if err != nil {
		log.Fatalf("Failed to load configuration file: %v. Make sure that the configuration file exists and is correct.", err)
	}
	printConfig(&config)

	rootDir, err := getRootDirectory(&config)
	if err != nil {
		log.Fatalf("Failed to get root directory: %v. Make sure the 'rootDir' key exists in the configuration and points to a valid directory.", err)
	}

	if !gitlog.IsGitRepository(rootDir) {
		log.Fatalf("Failed to get git branches from directory '%s': %v. Verify that it's a valid Git repository directory.", rootDir, err)
	}

	return config
}

// printConfig prints configures
func printConfig(c *Config) {
	for key, value := range *c {
		fmt.Printf("%s = %s\n", key, value)
	}
}

// getRootDirectory retrieves the root directory from the given Config.
// It returns the root directory as a string and an error if the rootDir key is not found in the Config.
func getRootDirectory(c *Config) (string, error) {
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
