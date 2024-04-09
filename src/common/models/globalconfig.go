package models

type Header struct {
	Name string `yaml:"name"`
}

type GlobalConfig struct {
	Auth struct {
		Headers []Header `yaml:"headers"`
	} `yaml:"auth"`
	Logger struct {
		Headers []Header `yaml:"headers"`
	} `yaml:"logger"`
	PureEngage struct {
		Host string `yaml:"host"`
	} `yaml:"pureengage"`
	Services struct {
		Stt  Service `yaml:"stt"`
		Auth Service `yaml:"auth"`
		Conv Service `yaml:"conv"`
	} `yaml:"services"`
}

type Service struct {
	Host            string `yaml:"host"`
	Basepath        string `yaml:"basepath"`
	Active          bool   `yaml:"active"`
	DefaultProvider string `yaml:"defaultProvider"`
}
