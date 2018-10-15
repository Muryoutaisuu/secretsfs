package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hanwen/go-fuse/fuse/nodefs"
	//"github.com/hanwen/go-fuse/fuse/pathfs"

	"github.com/Muryoutaisuu/secretsfs/pkg/fio"
	"github.com/Muryoutaisuu/secretsfs/pkg/secretsfs"
)

func main() {
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() != 1 {
		usage()
		os.Exit(2)
	}
	mountpoint := flag.Arg(0)

	fms := fio.FIOMaps()
	sf := secretsfs.NewSecretsfs(fms)
	root := sf.Root()

	server, _, err := nodefs.MountRoot(mountpoint, root, nodefs.NewOptions())
	if err != nil {
		log.Fatalf("Mount fail: %v\n", err)
	}

	server.Serve()
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  %s MOUNTPOINT\n", os.Args[0])
	flag.PrintDefaults()
}
