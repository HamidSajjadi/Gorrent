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
	assert.IsType(t, TorrentInfo{}, *data.ToTorrentInfo())
	assert.NotEmpty(t, data.Announce, "torrent Info Announce should not be empty")
	data.ToTorrentInfo()
}

func TestTorrentInfo_CreateTrackerURL(t *testing.T) {
	tInfo := TorrentInfo{
		Announce: "http://test.com",
		peerID:   "random-peer-id",
	}
	expected := "http://test.com?downloaded=0&event=started&left=0&peer_id=random-peer-id&port=6881&uploaded=0"

	res, err := tInfo.CreateTrackerURL(6881)
	assert.Nil(t, err, "no error creating tracker file")
	assert.NotEmpty(t, res)
	assert.Equal(t, expected, res)
}

func TestTorrent_ToTorrentInfo(t *testing.T) {

}
