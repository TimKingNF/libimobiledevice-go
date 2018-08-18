package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/DHowett/go-plist"
	"github.com/alyyousuf7/libimobiledevice-go/cmd"
	"github.com/alyyousuf7/libimobiledevice-go/idevice"
	"github.com/alyyousuf7/libimobiledevice-go/lockdownd"
	"gopkg.in/urfave/cli.v1"
)

var knownDomains = []string{
	"com.apple.disk_usage",
	"com.apple.disk_usage.factory",
	"com.apple.mobile.battery",
	"com.apple.iqagent",
	"com.apple.purplebuddy",
	"com.apple.PurpleBuddy",
	"com.apple.mobile.chaperone",
	"com.apple.mobile.third_party_termination",
	"com.apple.mobile.lockdownd",
	"com.apple.mobile.lockdown_cache",
	"com.apple.xcode.developerdomain",
	"com.apple.international",
	"com.apple.mobile.data_sync",
	"com.apple.mobile.tethered_sync",
	"com.apple.mobile.mobile_application_usage",
	"com.apple.mobile.backup",
	"com.apple.mobile.nikita",
	"com.apple.mobile.restriction",
	"com.apple.mobile.user_preferences",
	"com.apple.mobile.sync_data_class",
	"com.apple.mobile.software_behavior",
	"com.apple.mobile.iTunes.SQLMusicLibraryPostProcessCommands",
	"com.apple.mobile.iTunes.accessories",
	"com.apple.mobile.internal",
	"com.apple.mobile.wireless_lockdown",
	"com.apple.fairplay",
	"com.apple.iTunes",
	"com.apple.mobile.iTunes.store",
	"com.apple.mobile.iTunes",
}

func main() {
	app := cli.NewApp()
	app.Name = "ideviceinfo"
	app.Usage = "Show information about a connected device."
	app.HideVersion = true
	app.CustomAppHelpTemplate = fmt.Sprintf(`%sKnown Domains:
   %s

`, cmd.AppHelpTemplate, strings.Join(knownDomains, "\n   "))
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, d",
			Usage: "enable communication debugging",
		},
		cli.BoolFlag{
			Name:  "simple, s",
			Usage: "use a simple connection to avoid auto-pairing with the device",
		},
		cli.StringFlag{
			Name:  "udid, u",
			Usage: "target specific device by its 40-digit device UDID",
		},
		cli.StringFlag{
			Name:  "domain, q",
			Usage: "query specified domain (default: none)",
		},
		cli.StringFlag{
			Name:  "key, k",
			Usage: "query specified key (default: all keys)",
		},
		cli.GenericFlag{
			Name: "format, f",
			Value: &cmd.EnumValue{
				Enum:    []string{"json", "plist", "xml"},
				Default: "json",
			},
			Usage: "output in json, plist or xml format",
		},
	}

	app.Action = func(ctx *cli.Context) error {
		if len(ctx.Args()) > 0 {
			cli.ShowAppHelpAndExit(ctx, 0)
			return nil
		}

		if ctx.Bool("debug") {
			idevice.DebugMode(true)
		}

		udid := ctx.String("udid")
		if udid != "" && len(udid) != 40 {
			cli.ShowAppHelpAndExit(ctx, 0)
			return nil
		}

		domain := ctx.String("domain")
		if domain != "" && len(domain) < 4 {
			cli.ShowAppHelpAndExit(ctx, 0)
			return nil
		}
		if domain != "" && !isDomainKnown(domain) {
			fmt.Printf("WARNING: Sending query with unknown domain \"%s\".\r\n", domain)
		}

		key := ctx.String("key")
		if key != "" && len(key) < 2 {
			cli.ShowAppHelpAndExit(ctx, 0)
			return nil
		}

		var (
			device *idevice.IDevice
			client *lockdownd.Client
			err    error
		)

		if udid == "" {
			device, err = idevice.New()
			if err != nil {
				cli.HandleExitCoder(cli.NewExitError("No device found, is it plugged in?", -1))
				return nil
			}
		} else {
			device, err = idevice.NewWithUDID(udid)
			if err != nil {
				cli.HandleExitCoder(cli.NewExitError(fmt.Sprintf("No device found with udid %s, is it plugged in?", udid), -1))
				return nil
			}
		}
		defer device.Close()

		if ctx.Bool("simple") {
			client, err = lockdownd.NewClient(device, "ideviceinfo")
		} else {
			client, err = lockdownd.NewClientWithHandshake(device, "ideviceinfo")
		}
		if err != nil {
			cli.HandleExitCoder(cli.NewExitError(fmt.Sprintf("Could not connect to lockdownd, error: %s", err), -1))
			return nil
		}
		defer client.Close()

		buf, err := client.Get(domain, key)
		if err != nil {
			cli.HandleExitCoder(cli.NewExitError(fmt.Sprintf("Could not query data, error: %s", err), -1))
			return nil
		}

		var data map[string]interface{}
		if _, err := plist.Unmarshal(buf, &data); err != nil {
			cli.HandleExitCoder(cli.NewExitError(fmt.Sprintf("Could not parse data, error: %s", err), -1))
			return nil
		}

		switch ctx.Generic("format").(*cmd.EnumValue).String() {
		case "plist":
			encoder := plist.NewEncoder(os.Stdout)
			encoder.Indent("  ")
			if err := encoder.Encode(data); err != nil {
				cli.HandleExitCoder(cli.NewExitError(fmt.Sprintf("Could not print data, error: %s", err), -1))
				return nil
			}

			// Leave an extra line
			fmt.Println()

		case "xml":
			encoder := plist.NewEncoderForFormat(os.Stdout, plist.XMLFormat)
			encoder.Indent("  ")
			if err := encoder.Encode(data); err != nil {
				cli.HandleExitCoder(cli.NewExitError(fmt.Sprintf("Could not print data, error: %s", err), -1))
				return nil
			}

			// Leave an extra line
			fmt.Println()

		default:
			encoder := json.NewEncoder(os.Stdout)
			encoder.SetIndent("", "  ")
			if err := encoder.Encode(data); err != nil {
				cli.HandleExitCoder(cli.NewExitError(fmt.Sprintf("Could not print data, error: %s", err), -1))
				return nil
			}
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func isDomainKnown(domain string) bool {
	for _, d := range knownDomains {
		if domain == d {
			return true
		}
	}
	return false
}
