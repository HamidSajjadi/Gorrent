package main

import (
	bencode "github.com/jackpal/bencode-go"
	"github.com/pkg/errors"
	"io"
)

type info struct {
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
	Pieces      string `bencode:"pieces"`
}

type TorrentInfo struct {
	Announce string `bencode:"announce"`
	Info     info   `bencode:"info"`
}

func Open(r io.Reader) (*TorrentInfo, error) {
	var torrentData TorrentInfo
	err := bencode.Unmarshal(r, &torrentData)
	if err != nil {
		return nil, errors.Wrap(err, "error opening tracker while decoding bencode")
	}
	return &torrentData, err
}

func GetTracker(url string) {

}
