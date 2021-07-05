package cfg

import "github.com/bigmamallc/env"

func EnvCfg(cfg interface{}, prefix string) error {
	return env.SetWithEnvPrefix(cfg, prefix)
}

func MustEnvCfg(cfg interface{}, prefix string) {
	if err := EnvCfg(cfg, prefix); err != nil {
		panic(err)
	}
}
