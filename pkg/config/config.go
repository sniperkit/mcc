package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	// external
	"github.com/cep21/xdgbasedir"
	"github.com/sniperkit/snk.golang.configor"
)

//
// Can be set at build via
// -ldflags "-X pkg/config.ProgramName=$(PROG_NAME) -X pkg/config.ExecutableName=$(BIN_BASENAME)"
//
var (
	ProgramName           string   = "mcc"
	ExecutableName        string   = ProgramName
	ConfigDirectoryPath   string   = "~/.mcc"
	Xdgbasedir            string   = "~"
	ConfigFileExtensions  []string = []string{"yaml", "yml", "json", "toml"}
	ConfigFilePrefixPaths []string = []string{Xdgbasedir, "."}
	ConfigFileBasenames   []string = []string{"." + ProgramName, ProgramName} // , "snk", "xcc", "sniperkit"} // let's see later how to handle it...
)

// ServiceConfig contains configuration information for a service
type ServiceConfig struct {
	Token string `json:"token" yaml:"token" toml:"token"`
	User  string `json:"username" yaml:"username" toml:"username"`
}

// OutputConfig sontains configuration information for an output
type OutputConfig struct {
	SpinnerIndex    int    `json:"spinner_index" yaml:"spinner_index" toml:"spinnerIndex"`
	SpinnerInterval int    `json:"spinner_interval" yaml:"spinner_interval" toml:"spinnerInterval"`
	SpinnerColor    string `json:"spinner_color" yaml:"spinner_color" toml:"spinnerColor"`
}

// Config contains configuration information
type Config struct {
	// Services     map[string]*ServiceConfig `yaml:"services"`
	GoRoot  string                   `json:"go_root" yaml:"go_root" toml:"goRoot"`
	GoSrc   string                   `json:"go_src" yaml:"go_src" toml:"go_src"`
	GoPath  string                   `json:"go_path" yaml:"go_path" toml:"goPath"`
	Outputs map[string]*OutputConfig `json:"outputs" yaml:"outputs" toml:"outputs"`
}

/*
// GetService returns the configuration information for a service
func (config *Config) GetService(name string) *ServiceConfig {
	if config.Services == nil {
		config.Services = make(map[string]*ServiceConfig)
	}

	service := config.Services[name]
	if service == nil {
		service = &ServiceConfig{}
		config.Services[name] = service
	}
	return service
}
*/

// GetOutput returns the configuration information for an output
func (c *Config) GetOutput(name string) *OutputConfig {
	if c.Outputs == nil {
		c.Outputs = make(map[string]*OutputConfig)
	}

	output := c.Outputs[name]
	if output == nil {
		output = &OutputConfig{}
		c.Outputs[name] = output
	}
	return output
}

// ReadConfig reads the configuration information
func ReadConfig() (*Config, error) {
	file := configFilePath()

	var c Config
	if _, err := os.Stat(file); err == nil {
		// Read and unmarshal file only if it exists
		f, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}

		err = yaml.Unmarshal(f, &c)
		if err != nil {
			return nil, err
		}
	}

	return &c, nil
}

// WriteConfig writes the configuration information
func (c *Config) WriteConfig() error {
	err := os.MkdirAll(configDirectoryPath, 0700)
	if err != nil {
		return err
	}
	fmt.Println("configFilePath=", configFilePath())
	configor.Load(c, configFiles...)
	/*
		data, err := yaml.Marshal(c)
		if err != nil {
			return err
		}
	*/

	return ioutil.WriteFile(configFilePath(), data, 0600)
}

func configFilePath() string {
	return path.Join(configDirectoryPath, fmt.Sprintf("%s.yaml", ProgramName))
}

func init() {
	baseDir, err := xdgbasedir.ConfigHomeDirectory()
	if err != nil {
		log.Fatal("Can't find XDG BaseDirectory")
	} else {
		configDirectoryPath = path.Join(baseDir, ProgramName)
	}
}
