package config

import (
    "encoding/json"
    "github.com/moogar0880/venom"
    "io/ioutil"
    "os"
    "path/filepath"
)

const (
	HostConfigDir = "/etc/audmon/"
)

// init for the config package will attempt to load config from the default file
// path or, if that does not exist, load it from the string literal of defaults
// above and then write that to said path.
func init() {
    // check if config file exists
    appConf := "/etc/audmon/audmon.conf.json"
    if confFile, _ := os.Stat(appConf); confFile != nil {
        // if the file exists on host, load that file
        LoadConfig()
        return
    }

    // if no config existed on host, load default and write config
    LoadDefaults()
}

// DefaultConfig represents the basic config values for the application. This JSON
// will be persisted to disk under /etc/audmon/audmon.conf.json and should that file
// exists on start of the application, it will prefer those values over the defaults.
var DefaultConfig = `
{
  "log": {
    "level": "trace",
    "file": "/var/log/audmon.log",
    "silent": false,
    "report_caller": false
  }
}
`

// AudmonConfig is a struct representing the overarching global configuration for
// the whole project including but not limited to the daemon, server, client, and
// any future services which may be added.
type AudmonConfig struct {
    Log LogConfig `json:"log"`
    Registry *venom.Venom
    Daemon DaemonArgs
}

// LogConfig is the structural representation of the config for the application's
// global Log instance. See pkg/log for detail.
type LogConfig struct {
    Level        string `json:"level"`
    File         string `json:"file"`
    Silent       bool   `json:"silent"`
    ReportCaller bool   `json:"report_caller"`
}

var AppConf = &AudmonConfig{Registry: &venom.Venom{}}

type DaemonArgs struct {
    Foreground bool
}

// LoadConfig makes the assumption there is a file in place at the default file
// path for config, /etc/audmon/audmon.conf.json
func LoadConfig() {
    data, err := ioutil.ReadFile("/etc/audmon/audmon.conf.json")
    if err != nil {
        panic(err)
    }
    if err = json.Unmarshal(data, AppConf); err != nil {
        panic(err)
    }
    AppConf.Registry = readConfig("/etc/audmon/audmon.conf.json")
}

// LoadDefaults will parse the literal string above into the DaemonConfig
// struct, write that to a file on the host and then set the
func LoadDefaults() {
    writeFromMemory()
    AppConf.Registry = readConfig("/etc/audmon/audmon.conf.json")
}

// writeFromMemory writes the literal string variable DefaultConfig to a
// file and populates an in memory struct, AudmonConfig.
func writeFromMemory() {
    if err := json.Unmarshal([]byte(DefaultConfig), AppConf); err != nil {
        panic(err)
    }
    hostFile := filepath.Join("/etc/audmon/", "audmon.conf.json")
    if _, err := os.Stat(hostFile); err != nil {
        if err = os.MkdirAll("/etc/audmon/", os.FileMode(os.O_RDWR)); err != nil {
            panic("config dir did not exist and could not be created: " + err.Error())
        }
        if _, err = os.Create(hostFile); err != nil {
            panic("config file did not exist and could not be created: " + err.Error())
        }
    }
    data, err := json.MarshalIndent(AppConf, "", "    ")
    if err != nil {
        panic("could not unmarshal default config: " + err.Error())
    }
    if err = ioutil.WriteFile(hostFile, data, os.FileMode(os.O_RDWR)); err != nil {
        panic("could not write audmon configuration: " + err.Error())
    }
}

// readConfig loads configuration values from a JSON file into an instance of
// venom.Venom and will panic on error.
func readConfig(config string) *venom.Venom {
    v := venom.Default()
    if err := v.LoadDirectory(config, false); err != nil {
        panic(err)
    }
    return v
}