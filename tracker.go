package main

import (
	bencode "github.com/jackpal/bencode-go"
	"github.com/pkg/errors"
	"io"
	"net/url"
	"strconv"
)

const Sha1Length = 20
const DownloaderPort = 6881

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
	peerID      string
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

func (tInfo *TorrentInfo) CreateTrackerURL(port int) (string, error) {
	base, err := url.Parse(tInfo.Announce)
	if err != nil {
		return "", err
	}
	values := url.Values{}
	values.Add("peer_id", tInfo.peerID)
	values.Add("port", strconv.Itoa(DownloaderPort))
	values.Add("uploaded", "0")
	values.Add("downloaded", "0")
	values.Add("left", "0")
	values.Add("event", "started")
	base.RawQuery = values.Encode()
	return base.String(), nil
}

func chunkPieces(pieces string) []string {
	sliceSize := len(pieces) / Sha1Length
	chunkedPieces := make([]string, sliceSize)
	for i := 0; i < sliceSize; i++ {
		chunkedPieces[i] = pieces[i*Sha1Length : (i+1)*Sha1Length]
	}

	return chunkedPieces
}
