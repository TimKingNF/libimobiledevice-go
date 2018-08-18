package main

import (
	"fmt"
	"log"
	"os"

	"github.com/alyyousuf7/libimobiledevice-go/cmd"
	"github.com/alyyousuf7/libimobiledevice-go/idevice"
	"github.com/alyyousuf7/libimobiledevice-go/lockdownd"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "idevice_id"
	app.Usage = "Prints device name or a list of attached devices."
	app.ArgsUsage = "[uuid]"
	app.HideVersion = true
	app.CustomAppHelpTemplate = cmd.AppHelpTemplate
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "enable communication debugging",
		},
		cli.BoolFlag{
			Name:  "list, l",
			Usage: "list UDID of all attached devices",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		if ctx.Bool("debug") {
			idevice.DebugMode(true)
		}

		if ctx.Bool("list") {
			udids, err := idevice.UDIDList()
			if err != nil {
				cli.HandleExitCoder(cli.NewExitError("Unable to retrieve device list", -1))
				return nil
			}

			for _, udid := range udids {
				fmt.Println(udid)
			}

			return nil
		}

		args := ctx.Args()
		if len(args) == 1 && len(args[0]) == 40 {
			udid := args[0]
			device, err := idevice.NewWithUDID(udid)
			if err != nil {
				cli.HandleExitCoder(cli.NewExitError(fmt.Sprintf("No device with UDID=%s attached", udid), -2))
			}
			defer device.Close()

			client, err := lockdownd.NewClient(device, "idevice_id")
			if err != nil {
				cli.HandleExitCoder(cli.NewExitError("Connecting to device failed!", -2))
			}
			defer client.Close()

			deviceName, err := client.DeviceName()
			if err != nil {
				cli.HandleExitCoder(cli.NewExitError("Could not get device name!", -2))
			}

			fmt.Println(deviceName)
			return nil
		}

		cli.ShowAppHelpAndExit(ctx, 0)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
