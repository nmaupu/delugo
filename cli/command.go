package cli

import (
	"github.com/gdm85/go-libdeluge"
	"github.com/jawher/mow.cli"
	"log"
	"os"
	"time"
)

var (
	verbose                  *bool
	host, username, password *string
	port                     *int
)

func Process() {
	app := cli.App("delugo", "Deals with a deluged torrents daemon")
	//app.Spec = "[--verbose] --address --port --username --password"

	verbose = app.BoolOpt("v verbose", false, "Verbose mode")
	host = app.StringOpt("a address", "127.0.0.1", "Host to connect to")
	port = app.IntOpt("p port", 58846, "Port to connect to")
	username = app.StringOpt("u username", "", "Username to use for connection")
	password = app.StringOpt("w password", "", "Password to use for connection")

	app.Command("m methods", "List all available methods", listMethods)
	app.Command("d daemons", "Get daemon version", daemonVersion)

	app.Run(os.Args)
}

func getDelugeClient() *delugeclient.Client {
	var logger *log.Logger
	if *verbose {
		logger = log.New(os.Stderr, "DELUGE: ", log.Lshortfile)
	}

	settings := delugeclient.Settings{
		*host, uint(*port),
		*username, *password,
		logger, time.Duration(0),
	}

	d := delugeclient.New(settings)

	err := d.Connect()
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}

	return d
}

func listMethods(cmd *cli.Cmd) {
	cmd.Action = func() {
		d := getDelugeClient()
		defer d.Close()

		methods, err := d.MethodsList()
		if err != nil {
			log.Fatalf("Unable to get methods list - %v\n", err)
		}
		log.Printf("Methods list : %s\n", methods)
	}
}

func daemonVersion(cmd *cli.Cmd) {
	cmd.Action = func() {
		d := getDelugeClient()
		defer d.Close()

		ver, err := d.DaemonVersion()
		if err != nil {
			log.Fatalf("Unable to get daemon version - %v\n", err)
		}
		log.Printf("Daemon version: %s\n", ver)
	}
}
