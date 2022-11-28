package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config []struct {
	Chain       string `json:"chain"`
	Address     string `json:"address"`
	RPCEndpoint string `json:"rpc_endpoint"`
}

func LoadConfigFile(filename string) Config {
	fmt.Sprintf("Open Config file: %s", filename)

	rf, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	var cfg Config
	json.Unmarshal(rf, &cfg)

	return cfg
}
