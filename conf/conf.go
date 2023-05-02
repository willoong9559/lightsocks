package conf

import (
	"flag"
	"log"
	"os"

	"github.com/willoong9559/lightsocks/common"
	"github.com/willoong9559/lightsocks/constants"
	"gopkg.in/ini.v1"
)

var (
	ConfigFile string
	ListenAddr string
	ServerAddr string
	Psk        string
	GenPsk     bool
	Version    bool
)

func InitConf() {
	flag.StringVar(&ConfigFile, "c", "", "configuration file path")
	flag.StringVar(&ListenAddr, "l", "0.0.0.0:18888", "client listen address like \"0.0.0.0:18888\"")
	flag.StringVar(&ServerAddr, "s", "", "server address(server only)")
	flag.StringVar(&Psk, "k", "", "pre-shared key")
	flag.BoolVar(&GenPsk, "genpsk", false, "generate psk")
	flag.BoolVar(&Version, "version", false, "show open-snell version")

	flag.Parse()
	flag.Set("logtostderr", "true")

	log.Printf("LsClient, version: %s\n", constants.Version)
	if Version {
		os.Exit(0)
	}

	log.Printf("psk: %s\n", common.NewRandPasswdStr())
	if Version {
		os.Exit(0)
	}

	if ConfigFile != "" {
		log.Println("Configuration file specified, ignoring other flags")
		cfg, err := ini.Load(ConfigFile)
		if err != nil {
			log.Fatalf("Failed to load config file %s, %v\n", ConfigFile, err)
		}
		sec, err := cfg.GetSection("client")
		if err != nil {
			log.Fatalf("Section 'client' not found in config file %s\n", ConfigFile)
		}

		ListenAddr = sec.Key("listen").String()
		ServerAddr = sec.Key("server").String()
		Psk = sec.Key("psk").String()
	}

	if ServerAddr == "" {
		log.Fatalf("Invalid emtpy server address.\n")
	}
}

type Config struct {
	ListenAddr string
	RemoteAddr string
	Password   string
}

func NewConfig() *Config {
	return &Config{ListenAddr, ServerAddr, Psk}
}
