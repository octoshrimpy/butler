package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

// #cgo windows LDFLAGS: -Wl,--allow-multiple-definition -static
import "C"

var (
	version = "head" // set by command-line on CI release builds
	app     = kingpin.New("butler", "Your very own itch.io helper")

	dlCmd    = app.Command("dl", "Download a file (resumes if can, checks hashes)")
	pushCmd  = app.Command("push", "Upload a new version of something to itch.io")
	untarCmd = app.Command("untar", "Extract a .tar file")
	wipeCmd  = app.Command("wipe", "Completely remove a directory (rm -rf)")
	dittoCmd = app.Command("ditto", "Create a mirror (incl. symlinks) of a directory into another dir (rsync -az)")
	mkdirCmd = app.Command("mkdir", "Create an empty directory and all required parent directories (mkdir -p)")
	simonCmd = app.Command("simon", "Reproduce src into dst by symlinking as much as needed without touching existing files")
)

var appArgs = struct {
	json       *bool
	quiet      *bool
	verbose    *bool
	timestamps *bool
}{
	app.Flag("json", "Enable machine-readable JSON-lines output").Short('j').Bool(),
	app.Flag("quiet", "Hide progress indicators & other extra info").Short('q').Bool(),
	app.Flag("verbose", "Display as much extra info as possible").Short('v').Bool(),
	app.Flag("timestamps", "Prefix all output by timestamps (for logging purposes)").Bool(),
}

var dlArgs = struct {
	url  *string
	dest *string
}{
	dlCmd.Arg("url", "Address to download from").Required().String(),
	dlCmd.Arg("dest", "File to write downloaded data to").Required().String(),
}

var pushArgs = struct {
	src      *string
	repo     *string
	identity *string
	address  *string
}{
	pushCmd.Arg("src", "Directory or zip archive to upload, e.g.").Required().ExistingFileOrDir(),
	pushCmd.Arg("repo", "Repository to push to, e.g. leafo/xmoon:win64").Required().String(),
	pushCmd.Flag("identity", "Path to the private key used for public key authentication.").Default(fmt.Sprintf("%s/%s", os.Getenv("HOME"), ".ssh/id_rsa")).Short('i').ExistingFile(),
	pushCmd.Flag("address", "Specify wharf address (advanced)").Default("wharf.itch.zone").Short('a').Hidden().String(),
}

var untarArgs = struct {
	file *string
	dir  *string
}{
	untarCmd.Arg("file", "Path of the .tar archive to extract").Required().String(),
	untarCmd.Flag("dir", "An optional directory to which to extract files (defaults to CWD)").Default(".").Short('d').String(),
}

var wipeArgs = struct {
	path *string
}{
	wipeCmd.Arg("path", "Path to completely remove, including its contents").Required().String(),
}

var mkdirArgs = struct {
	path *string
}{
	mkdirCmd.Arg("path", "Directory to create").Required().String(),
}

var dittoArgs = struct {
	src  *string
	dst  *string
	link *bool
}{
	dittoCmd.Arg("src", "Directory to mirror").Required().String(),
	dittoCmd.Arg("dst", "Path where to create a mirror").Required().String(),
	dittoCmd.Flag("link", "Use symlinks instead of copying contents").Short('l').Bool(),
}

var simonArgs = struct {
	src *string
	dst *string
}{
	simonCmd.Arg("src", "Directory to mirror").Required().String(),
	simonCmd.Arg("dst", "Path where to create a mirror").Required().String(),
}

func must(err error) {
	if err != nil {
		Die(err.Error())
	}
}

func main() {
	app.HelpFlag.Short('h')
	app.Version(version)
	app.VersionFlag.Short('V')

	cmd, err := app.Parse(os.Args[1:])
	if *appArgs.timestamps {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	} else {
		log.SetFlags(0)
	}

	switch kingpin.MustParse(cmd, err) {
	case dlCmd.FullCommand():
		dl(*dlArgs.url, *dlArgs.dest)

	case pushCmd.FullCommand():
		push(*pushArgs.src, *pushArgs.repo)

	case untarCmd.FullCommand():
		untar(*untarArgs.file, *untarArgs.dir)

	case wipeCmd.FullCommand():
		wipe(*wipeArgs.path)

	case mkdirCmd.FullCommand():
		mkdir(*mkdirArgs.path)

	case dittoCmd.FullCommand():
		ditto(*dittoArgs.src, *dittoArgs.dst)

	case simonCmd.FullCommand():
		simon(*simonArgs.src, *simonArgs.dst)
	}
}
