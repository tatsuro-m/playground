package config

type Config struct {
	TargetEnvs     []string
	TargetServices []string
}

func GetConfig() *Config {
	c := &Config{
		TargetEnvs:     []string{"dev", "stg", "preprod", "prod"},
		TargetServices: []string{"webfront", "bff", "bff2", "rpcserver", "adminfront"},
	}

	return c
}
