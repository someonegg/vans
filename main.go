package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"vans/core"
)

var (
	srcFile  string
	destFile string
)

func main() {
	flag.Parse()
	if srcFile != "" && destFile != "" {
		if !filepath.IsAbs(srcFile) || !filepath.IsAbs(destFile) {
			fmt.Println("srcFile or destFile must be abs path")
			os.Exit(128)
		}

		err, code := core.Render(srcFile, destFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(code)
		}
		fmt.Println("success.")
	} else {
		fmt.Println("./vans -s <srcFile> -d <destFile>")
		os.Exit(1)
	}
}

func init() {
	flag.StringVar(&srcFile, "s", "", "s")
	flag.StringVar(&destFile, "d", "", "d")
}
