package main

import (
	"bufio"
	"fmt"

	"pault.ag/go/debian/control"
)

func main() {
	reader, err := NewArchiveContentsReader(
		"http://http.debian.net/debian",
		// "http://localhost/fnord/",
		"unstable-debug",
		"main",
		"amd64",
	)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	index, err := control.ParseBinaryIndex(bufio.NewReader(reader))
	if err != nil {
		panic(err)
	}

	for _, entry := range index {
		for _, buildId := range entry.DebugBuildIds {
			fmt.Printf("%s %s %s %s\n", buildId, entry.Source, entry.Version, entry.Filename)
		}
	}
}
