package main

import (
	bencode "github.com/jackpal/bencode-go"
	"github.com/pkg/errors"
	"io"
)

const SHA1_LENGTH = 20

type files struct {
	Length int      `bencode:"length"`
	Path   []string `bencode:"path"`
}

type info struct {
	PieceLength int32   `bencode:"piece Length"`
	Length      int     `bencode:"length"`
	Name        string  `bencode:"name"`
	Pieces      string  `bencode:"pieces"`
	Files       []files `bencode:"files"`
}

type torrent struct {
	Announce string `bencode:"announce"`
	Info     info   `bencode:"info"`
}

type TorrentInfo struct {
	Announce    string
	Name        string
	PieceLength int32
	Pieces      []string
}

func Open(r io.Reader) (*torrent, error) {
	var torrentData torrent
	err := bencode.Unmarshal(r, &torrentData)
	if err != nil {
		return nil, errors.Wrap(err, "error opening tracker while decoding bencode")
	}
	return &torrentData, err
}

func (t *torrent) ToTorrentInfo() *TorrentInfo {
	torrentInfo := &TorrentInfo{
		Announce:    t.Announce,
		Name:        t.Info.Name,
		PieceLength: t.Info.PieceLength,
		Pieces:      chunkPieces(t.Info.Pieces),
	}

	return torrentInfo
}

func chunkPieces(pieces string) []string {
	sliceSize := len(pieces) / SHA1_LENGTH
	chunkedPieces := make([]string, sliceSize)
	for i := 0; i < sliceSize; i++ {
		chunkedPieces[i] = pieces[i*SHA1_LENGTH : (i+1)*SHA1_LENGTH]
	}

	return chunkedPieces
}
