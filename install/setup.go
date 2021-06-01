package install

import "log"

func setup() {
	log.Println("Gathering configuration data")
	config := get_config()

	log.Println("Finding available nodes")

	log.Println("Connecting to available nodes")

	log.Println("Setup new block listener")

	if config.node_type == "full" {
		log.Println("Synchronize with neighbours")
	}else{

	}
}

func get_config() *Config{
	return nil
}
