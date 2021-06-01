package install

type Config struct {
	install_dir   string
	blocks_dir     string // defaults to install_dir/blocks if not set
	node_type string 	  // values = "full" or "lightweight"
	trusted_dns   []string
	trusted_nodes []string
}
