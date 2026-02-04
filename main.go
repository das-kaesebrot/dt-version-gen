package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	useZeroPadding = flag.Bool("use-zero-padding", false, "Whether to pad numbers with zeroes")
)

func setupFlags(f *flag.FlagSet) {
	f.Usage = func() {
		fmt.Printf("usage: %s [-h] [--use-zero-padding] [datetime_string]\n", os.Args[0])
		fmt.Println()
		fmt.Println("Script for generating SemVer-compatible formatted version numbers from an input date")
		fmt.Println("positional arguments:")
		fmt.Println("  datetime_string     The date string in ISO 8601 format to parse. If it's unset, the value from the environment variable CI_PIPELINE_CREATED_AT will be used.")
		fmt.Println()
		fmt.Println("options:")
		flag.PrintDefaults()
	}
}

func main() {
	logger := log.Default()
	setupFlags(flag.CommandLine)
	flag.Parse()

	var datetime string

	if flag.NArg() > 1 {
		flag.CommandLine.Usage()
		os.Exit(1)
	}

	datetime = os.Getenv("CI_PIPELINE_CREATED_AT")

	if flag.NArg() == 1 {
		datetime = flag.Arg(0)
	} else if datetime == "" {
		logger.Print("Arg can't be empty if CI_PIPELINE_CREATED_AT is not set!")
		os.Exit(2)
	}

	var parsedTime time.Time
	var err error

	// https://stackoverflow.com/questions/38596079/how-do-i-parse-an-iso-8601-timestamp-in-go
	if parsedTime, err = time.Parse("2006-01-02T15:04:05-0700", datetime); err != nil {
		logger.Fatalf("%v", err)
	}

	/*
	   if args.use_zero_padding:
	       minor = f"{dt.month:02}{dt.day:02}"
	       patch = f"{dt.hour:02}{dt.minute:02}{dt.second:02}"
	   else:
	       minor = f"{dt.month}{dt.day:02}"
	       patch = f"{dt.hour}{dt.minute:02}{dt.second:02}"

	   full = f"{major}.{minor}.{patch}"
	*/

	minorParseFormat := "%02d%02d"
	patchParseFormat := "%02d%02d%02d"
	if !*useZeroPadding {
		minorParseFormat = "%d%02d"
		patchParseFormat = "%d%02d%02d"
	}

	var major, minor, patch, full string
	major = strconv.Itoa(parsedTime.Year())
	minor = fmt.Sprintf(minorParseFormat, parsedTime.Month(), parsedTime.Day())
	patch = fmt.Sprintf(patchParseFormat, parsedTime.Hour(), parsedTime.Minute(), parsedTime.Second())
	full = fmt.Sprintf("%s.%s.%s", major, minor, patch)

	fmt.Printf("DATETIME_VERSION_MAJOR=%s\n", major)
	fmt.Printf("DATETIME_VERSION_MINOR=%s\n", minor)
	fmt.Printf("DATETIME_VERSION_PATCH=%s\n", patch)
	fmt.Printf("DATETIME_VERSION_FULL=%s\n", full)
}
