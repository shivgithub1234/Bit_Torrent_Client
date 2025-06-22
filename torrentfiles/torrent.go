package torrentfiles

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"os"

	"BITTORRENTCLIENT/p2p"
	"github.com/jackpal/bencode-go"
)

const PORT uint16 = 6881 // as given in the article

type bencodeInfo struct{
	Pieces string `bencode:"pieces"`
	PieceLength int `bencode:"piece length"`
	Length int `bencode:"length"`
	Name string `bencode:"name"`
}

type bencodeTorrent struct{
	Announce string `bencode:"announce"`
	Info bencodeInfo  `bencode:"info"`
}

// Torrent file structure
type TorrentFile struct{
	Announce string
	InfoHash [20]byte
	PieceHashes [][20]byte
	PieceLength int
	Length int
	Name string
}
// Downloads and writesv into a file(outfile)
func (t *TorrentFile) DownloadToFile(path string) error{
	var peerID [20]byte // my peer id
	_,err:= rand.Read(peerID[:])
	if err!=nil{
		return err
	}
	peers,err := t.requestPeers(peerID,PORT)
	if err!=nil{
		return err
	}
	torrent:=p2p.Torrent{
		Peers:       peers,
		PeerID:      peerID,
		InfoHash:    t.InfoHash,
		PieceHashes: t.PieceHashes,
		PieceLength: t.PieceLength,
		Length:      t.Length,
		Name:        t.Name,
	}
	buf,err:=torrent.Download()
	if err!=nil{
		return err
	}
	outfile,err:=os.Create(path)
	if err!=nil{
		return err
	}
	defer outfile.Close()
	_,err = outfile.Write(buf)
	if err!=nil{
		return err
	}
	return nil
}

// open parsese the torrent file we jsut create and wrote the info into
func Open(path string)(TorrentFile,error){
	file,err:=os.Open(path)
	if err!=nil{
		return TorrentFile{}, err
	}
	defer file.Close()

	bto := bencodeTorrent{}
	err=bencode.Unmarshal(file,&bto)
	if err!=nil{
		return TorrentFile{}, err
	}
	return bto.toTorrentFile()
}


// Creates the SHA-1 Hashes
func (i *bencodeInfo) hash()([20]byte ,error){
	var buf bytes.Buffer
	err:=bencode.Marshal(&buf,*i)
	if err!=nil{
		return [20]byte{},err
	}
	h := sha1.Sum(buf.Bytes())
	return h,nil
}

// Calculate the individual hashes and also checks whether we have got malformed hahses
func (i *bencodeInfo) splitPieceHashes() ([][20]byte ,error){
	hashLen := 20 // lenght of each sha1 hash
	buf:=[]byte(i.Pieces) // size of total concat sha1 pieces
	if len(buf)%hashLen!=0{
		err:=fmt.Errorf("received malformed pieces of length %d", len(buf))
		return nil, err
	}
	numHashes := len(buf)/hashLen
	hashes := make([][20]byte,numHashes)

	for i:=0;i<numHashes;i++{
		copy(hashes[i][:], buf[i*hashLen:(i+1)*hashLen]) // splitting the concat hashes into a slice
	}
	return hashes,nil
}

func (bto *bencodeTorrent) toTorrentFile() (TorrentFile,error){
	infoHash,err:=bto.Info.hash()
	if err != nil {
		return TorrentFile{}, err
	}
	pieceHashes,err:=bto.Info.splitPieceHashes()
	if err != nil {
		return TorrentFile{}, err
	}
	t:= TorrentFile{
		Announce:    bto.Announce,
		InfoHash:    infoHash,
		PieceHashes: pieceHashes,
		PieceLength: bto.Info.PieceLength,
		Length:      bto.Info.Length,
		Name:        bto.Info.Name,
	}
	return t,nil
}
