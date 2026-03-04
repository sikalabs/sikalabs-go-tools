package main

import (
	"os"
	"path/filepath"

	_ "github.com/sikalabs/gobble/cmd"
	gobble_root "github.com/sikalabs/gobble/cmd/root"
	_ "github.com/sikalabs/mon/cmd"
	mon_root "github.com/sikalabs/mon/cmd/root"
	_ "github.com/sikalabs/slr/cmd"
	slr_root "github.com/sikalabs/slr/cmd/root"
	_ "github.com/sikalabs/slu/cmd"
	slu_root "github.com/sikalabs/slu/cmd/root"
	_ "github.com/sikalabs/tergum/cmd"
	tergum_root "github.com/sikalabs/tergum/cmd/root"
	"github.com/spf13/cobra"
)

func main() {
	switch filepath.Base(os.Args[0]) {
	case "slu":
		slu_root.RootCmd.Execute()
	case "slr":
		slr_root.Cmd.Execute()
	case "tergum":
		tergum_root.Cmd.Execute()
	case "gobble":
		gobble_root.Cmd.Execute()
	case "mon":
		mon_root.Cmd.Execute()
	default:
		all()
	}
}

func all() {
	var Cmd = &cobra.Command{
		Use:   "sikalabs-go-tools",
		Short: "sikalabs-go-tools: embedded slu, slr, tergum, gobble and mon to single binary",
	}
	Cmd.AddCommand(slu_root.RootCmd)
	Cmd.AddCommand(tergum_root.Cmd)
	Cmd.AddCommand(gobble_root.Cmd)
	Cmd.AddCommand(mon_root.Cmd)
	Cmd.AddCommand(slr_root.Cmd)
	Cmd.Execute()
}
