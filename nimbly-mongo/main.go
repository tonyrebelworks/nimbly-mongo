package main

import (
	"chi-rest/bootstrap"
	"chi-rest/lib/mysql"
	"chi-rest/lib/utils"
	"chi-rest/services/journeyplan"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/urfave/cli/v2"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
	config     utils.Config
	debug      = false
	host       string

	// app the base of skeleton
	app *bootstrap.App
)

// EnvConfigPath environtment variable that set the config path
const EnvConfigPath = "REBEL_CLI_CONFIG_PATH"

// setup initialize the used variable and dependencies
func setup() error {
	configFile := os.Getenv(EnvConfigPath)
	if configFile == "" {
		configFile = "./config.json"
	}

	log.Println(configFile)

	config = utils.NewViperConfig(basepath, configFile)
	host = config.GetString("app.host")

	debug = config.GetBool("app.debug")
	validator := bootstrap.SetupValidator(config)
	app = &bootstrap.App{
		Debug:     debug,
		Config:    config,
		DB:        mysql.Connect(config.GetString("db.mysql_dsn")),
		Validator: validator,
	}

	return nil
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	setup()
	cmd := &cli.App{
		Name:  "chi-rest",
		Usage: "journeyplan, cli",
		Commands: []*cli.Command{
			{
				Name:  "journeyplan",
				Usage: "Run the http 1/1 for API",
				// Flags:  journeyplan.Flags,
				Action: journeyplan.API{app}.Start,
				// After:  ListenSignal,
			},
		},
		Action: func(cli *cli.Context) error {
			fmt.Printf("%s version:%s\n", cli.App.Name, "1.0")
			return nil
		},
	}

	err := cmd.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
