package cli

import (
	"flag"
	"fmt"
	"fuse/coder"
	"fuse/server"
	"fuse/util"
	"log/slog"
	"os"
	"runtime/debug"
	"strings"
)

var clientConfPath = "fuse.txt"

func Start() {
	var err error

	var standalone bool
	var baseURL string
	var listenAddress string
	var platform string
	var writeLog bool
	var passThrough bool
	flag.StringVar(&baseURL, "baseURL", "http://127.0.0.1:6667/", "base url")
	flag.StringVar(&listenAddress, "listenAddr", "127.0.0.1:6667", "listen addr")
	flag.StringVar(&platform, "platform", "tppstm", "server platform")
	flag.BoolVar(&writeLog, "writeLog", false, "save requests and responses to files")
	flag.BoolVar(&passThrough, "passThrough", false, "work as a proxy between client and original master server")
	flag.BoolVar(&standalone, "standalone", false, "standalone mode (do not attempt to patch and backup executable)")
	flag.Parse()

	opts := &slog.HandlerOptions{AddSource: true, Level: slog.LevelInfo}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, opts)))
	slog.Info("starting")

	bf, ok := debug.ReadBuildInfo()
	if !ok {
		slog.Error("cannot read build info")
		os.Exit(1)
	}

	commit := ""
	commitDate := ""
	dirty := false

	for _, k := range bf.Settings {
		switch k.Key {
		case "vcs.revision":
			commit = k.Value
		case "vcs.time":
			commitDate = k.Value
		case "vcs.modified":
			dirty = k.Value == "true"
		default:
		}
	}

	slog.Info("version", "commit", commit, "date", commitDate, "dirty", dirty)

	if !standalone {
		if err = util.Backup(); err != nil {
			slog.Error("backup", "error", err.Error())
			os.Exit(1)
		}

		customURL := false
		clientURL := "http://127.0.0.1:6667"
		_, err = os.Stat(clientConfPath)
		if err == nil {
			d, err := os.ReadFile(clientConfPath)
			if err != nil {
				slog.Error("read config", "error", err.Error(), "config path", clientConfPath)
				os.Exit(1)
			}

			clientURL = string(d)
			customURL = true
		}

		if err = util.Patch(clientURL); err != nil {
			slog.Error("patch", "error", err.Error())
			os.Exit(1)
		}

		slog.Info("patched executable", "server url", strings.TrimSpace(clientURL))

		if customURL {
			slog.Info("custom server url provided, not starting local server")
			slog.Info("You can start the game now")
			slog.Info("Press Enter to close this window")

			var a []byte
			_, _ = fmt.Scanln(&a)
			os.Exit(0)
		}
	}

	c := coder.Coder{}
	err = c.WithKey(nil)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("starting server", "listen address", listenAddress)

	server.Start(baseURL, listenAddress, platform, writeLog, passThrough)
}
