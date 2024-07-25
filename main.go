package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"syscall"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"."}
	}
	for _, arg := range args {
		dir, err := filepath.Abs(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", arg, err)
			continue
		}

		var s syscall.Statfs_t
		if err := syscall.Statfs(dir, &s); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", dir, err)
			continue
		}
		fmt.Printf("%s:\n", dir)
		fmt.Printf("        block size: %12d bytes\n", s.Bsize)
		fmt.Printf("            blocks: %12d blocks (%s)\n", s.Blocks, blocksToHuman(s.Blocks, s.Bsize))
		fmt.Printf("       free blocks: %12d blocks (%s)\n", s.Bfree, blocksToHuman(s.Bfree, s.Bsize))
		fmt.Printf("  available blocks: %12d blocks (%s)\n", s.Bavail, blocksToHuman(s.Bavail, s.Bsize))
		fmt.Printf("             files: %12d\n", s.Files)
		fmt.Printf("        free files: %12d\n", s.Ffree)
	}
}

var units []string = []string{
	"bytes",
	"KiB",
	"MiB",
	"GiB",
	"TiB",
	"PiB",
	"EiB",
	"ZiB",
	"YiB",
}

func blocksToHuman(blocks uint64, blockSize int64) string {
	bytes := blocks * uint64(blockSize)
	exp := int(math.Floor(math.Log10(float64(bytes)) / (10 * math.Log10(2))))
	scaled := float64(bytes) / math.Pow(2, float64(10*exp))
	unit := units[exp]
	return fmt.Sprintf("%7.2f %s", scaled, unit)
}
