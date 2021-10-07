package flag

import "path/filepath"

// Add flag vars to be referenced by main and other packages
var ConfigLocation string = filepath.Join("/", "etc", "ipwatcher")

var ConfigFile string = filepath.Join(ConfigLocation, "config.json")
