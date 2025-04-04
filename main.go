package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"

	"go.rikki.moe/librespeedtest-cli/defs"
	"go.rikki.moe/librespeedtest-cli/speedtest"
)

// init sets up the essential bits on start up
func init() {
	// set logrus formatter and default log level
	formatter := &defs.NoFormatter{}

	// debug level is for --debug messages
	// info level is for non-suppress mode
	// warn level is for suppress modes
	// error level is for errors

	log.SetOutput(os.Stderr)
	log.SetFormatter(formatter)
	log.SetLevel(log.InfoLevel)
}

func main() {
	// define cli options
	app := &cli.App{
		Name:     "librespeed-cli",
		Usage:    "Test your Internet speed with LibreSpeed",
		Action:   speedtest.SpeedTest,
		HideHelp: true,
		Flags: []cli.Flag{
			cli.HelpFlag,
			&cli.BoolFlag{
				Name:  defs.OptionVersion,
				Usage: "Show the version number and exit",
			},
			&cli.BoolFlag{
				Name:    defs.OptionIPv4,
				Aliases: []string{defs.OptionIPv4Alt},
				Usage:   "Force IPv4 only",
			},
			&cli.BoolFlag{
				Name:    defs.OptionIPv6,
				Aliases: []string{defs.OptionIPv6Alt},
				Usage:   "Force IPv6 only",
			},
			&cli.BoolFlag{
				Name:  defs.OptionNoDownload,
				Usage: "Do not perform download test",
			},
			&cli.BoolFlag{
				Name:  defs.OptionNoUpload,
				Usage: "Do not perform upload test",
			},
			&cli.BoolFlag{
				Name: defs.OptionNoICMP,
				Usage: "Do not use ICMP ping. ICMP doesn't work well under Linux\n" +
					"\tat this moment, so you might want to disable it",
			},
			&cli.IntFlag{
				Name:  defs.OptionConcurrent,
				Usage: "Concurrent HTTP requests being made",
				Value: 3,
			},
			&cli.BoolFlag{
				Name: defs.OptionBytes,
				Usage: "Display values in bytes instead of bits. Does not affect\n" +
					"\tthe image generated by --share, nor output from\n" +
					"\t--json or --csv",
			},
			&cli.BoolFlag{
				Name:  defs.OptionMebiBytes,
				Usage: "Use 1024 bytes as 1 kilobyte instead of 1000",
			},
			&cli.StringFlag{
				Name: defs.OptionDistance,
				Usage: "Change distance unit shown in ISP info, use 'mi' for miles,\n" +
					"\t'km' for kilometres, 'NM' for nautical miles",
				Value: "km",
			},
			&cli.BoolFlag{
				Name: defs.OptionShare,
				Usage: "Generate and provide a URL to the LibreSpeed.org share results\n" +
					"\timage, not displayed with --csv",
			},
			&cli.BoolFlag{
				Name:  defs.OptionSimple,
				Usage: "Suppress verbose output, only show basic information\n\t",
			},
			&cli.BoolFlag{
				Name: defs.OptionCSV,
				Usage: "Suppress verbose output, only show basic information in CSV\n" +
					"\tformat. Speeds listed in bit/s and not affected by --bytes\n\t",
			},
			&cli.StringFlag{
				Name: defs.OptionCSVDelimiter,
				Usage: "Single character delimiter (`CSV_DELIMITER`) to use in\n" +
					"\tCSV output.",
				Value: ",",
			},
			&cli.BoolFlag{
				Name:  defs.OptionCSVHeader,
				Usage: "Print CSV headers",
			},
			&cli.BoolFlag{
				Name: defs.OptionJSON,
				Usage: "Suppress verbose output, only show basic information\n" +
					"\tin JSON format. Speeds listed in bit/s and not\n" +
					"\taffected by --bytes",
			},
			&cli.BoolFlag{
				Name:  defs.OptionList,
				Usage: "Display a list of LibreSpeed.org servers",
			},
			&cli.IntSliceFlag{
				Name: defs.OptionServer,
				Usage: "Specify a `SERVER` ID to test against. Can be supplied\n" +
					"\tmultiple times. Cannot be used with --exclude",
			},
			&cli.IntSliceFlag{
				Name: defs.OptionExclude,
				Usage: "`EXCLUDE` a server from selection. Can be supplied\n" +
					"\tmultiple times. Cannot be used with --server",
			},
			&cli.StringFlag{
				Name:  defs.OptionServerJSON,
				Usage: "Use an alternative server list from remote JSON file",
			},
			&cli.StringFlag{
				Name: defs.OptionLocalJSON,
				Usage: "Use an alternative server list from local JSON file,\n" +
					"\tor read from stdin with \"--" + defs.OptionLocalJSON + " -\".",
			},
			&cli.StringFlag{
				Name:  defs.OptionSource,
				Usage: "`SOURCE` IP address to bind to",
			},
			&cli.StringFlag{
				Name:  defs.OptionInterface,
				Usage: "network INTERFACE to bind to",
			},
			&cli.IntFlag{
				Name:  defs.OptionTimeout,
				Usage: "HTTP `TIMEOUT` in seconds.",
				Value: 15,
			},
			&cli.IntFlag{
				Name:  defs.OptionDuration,
				Usage: "Upload and download test duration in seconds",
				Value: 15,
			},
			&cli.IntFlag{
				Name:  defs.OptionChunks,
				Usage: "Chunks to download from server, chunk size depends on server configuration",
				Value: 100,
			},
			&cli.IntFlag{
				Name:  defs.OptionUploadSize,
				Usage: "Size of payload being uploaded in KiB",
				Value: 1024,
			},
			&cli.BoolFlag{
				Name: defs.OptionSecure,
				Usage: "Use HTTPS instead of HTTP when communicating with\n" +
					"\tLibreSpeed.org operated servers",
			},
			&cli.StringFlag{
				Name: defs.OptionCACert,
				Usage: "Use the specified CA certificate PEM bundle file instead\n" +
				    "\tof the system certificate trust store",
			},
			&cli.BoolFlag{
				Name:  defs.OptionSkipCertVerify,
				Usage: "Skip verifying SSL certificate for HTTPS connections (self-signed certs)",
			},
			&cli.BoolFlag{
				Name: defs.OptionNoPreAllocate,
				Usage: "Do not pre allocate upload data. Pre allocation is\n" +
					"\tenabled by default to improve upload performance. To\n" +
					"\tsupport systems with insufficient memory, use this\n" +
					"\toption to avoid out of memory errors",
			},
			&cli.BoolFlag{
				Name:    defs.OptionDebug,
				Aliases: []string{"verbose"},
				Usage:   "Debug mode (verbose logging)",
				Hidden:  true,
			},
			&cli.StringFlag{
				Name: defs.OptionTelemetryJSON,
				Usage: "Load telemetry server settings from a JSON file. This\n" +
					"\toptions overrides --" + defs.OptionTelemetryLevel + ", --" + defs.OptionTelemetryServer + ",\n" +
					"\t--" + defs.OptionTelemetryPath + ", and --" + defs.OptionTelemetryShare + ". Implies --" + defs.OptionShare,
			},
			&cli.StringFlag{
				Name: defs.OptionTelemetryLevel,
				Usage: "Set telemetry data verbosity, available values are:\n" +
					"\tdisabled, basic, full, debug. Implies --" + defs.OptionShare,
			},
			&cli.StringFlag{
				Name:  defs.OptionTelemetryServer,
				Usage: "Set the telemetry server base URL. Implies --" + defs.OptionShare,
			},
			&cli.StringFlag{
				Name:  defs.OptionTelemetryPath,
				Usage: "Set the telemetry upload path. Implies --" + defs.OptionShare,
			},
			&cli.StringFlag{
				Name:  defs.OptionTelemetryShare,
				Usage: "Set the telemetry share link path. Implies --" + defs.OptionShare,
			},
			&cli.StringFlag{
				Name: defs.OptionTelemetryExtra,
				Usage: "Send a custom message along with the telemetry results.\n" +
					"\tImplies --" + defs.OptionShare,
			},
			&cli.StringFlag{
				Name: defs.OptionUserAgent,
				Usage: "Set the user agent string to use when communicating\n" +
					"\twith the server.",
			},
		},
	}

	// run main function with cli options
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("Terminated due to error")
	}
}
