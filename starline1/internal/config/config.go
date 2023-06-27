package config

import "flag"

type Config struct {
	PathToCSV string
	Port      string
}

func GetVars() *Config {
	path := flag.String("csv", "driver_positions.csv", "path to csv file")
	port := flag.String("port", "4444", "port to app")
	flag.Parse()
	return &Config{
		PathToCSV: *path,
		Port:      *port,
	}
}
