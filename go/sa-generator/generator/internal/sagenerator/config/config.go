package config

type Config struct {
	TargetEnvs     []string
	TargetServices []string
}

func getConfig() *Config {
	c := &Config{
		TargetEnvs:     []string{"dev", "stg", "preprod", "prod"},
		TargetServices: []string{"a", "b", "c", "d", "e"},
	}

	return c
}
