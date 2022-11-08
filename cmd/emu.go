package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
)

func main() {
	nes := SetupNES()
	SetupRenderer(&nes)

	wd, _ := os.Getwd()
	CreateRomAndLoad(&nes, filepath.Join(wd, "data", "pacman.nes"))

	flagDebug := flag.Bool("d", false, "runs in debug mode")
	flag.Parse()
	if *flagDebug {
		os.MkdirAll(filepath.Join(wd, "temp"), 0755)
		f, err := os.Create(filepath.Join(wd, "temp", "cpu.prof"))
		if err != nil {
			log.Fatal(err)
			return
		}
		pprof.StartCPUProfile(f)
		LogLn("Running in debug mode...")
	}

	CycleGlobal(&nes)
}
