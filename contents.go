package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

func NewArchiveContentsReader(mirror, suite, section, arch string) (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf(
		"%s/dists/%s/%s/binary-%s/Packages.gz",
		mirror, suite, section, arch,
	))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Exit code %d from server", resp.StatusCode)
	}
	reader, err := gzip.NewReader(resp.Body)
	if err != nil {
		return nil, err
	}
	return reader, nil
}
