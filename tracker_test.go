package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	testFile := "./bunny.torrent"
	file, err := os.Open(testFile)
	assert.Nil(t, err, "no error for opening file")
	data, err := Open(bufio.NewReader(file))
	assert.Nil(t, err, "no error unmarshalling torrent file")
	assert.IsType(t, TorrentInfo{}, *data)
	assert.NotEmpty(t, data.Announce, "torrent info announce should not be empty")
}
