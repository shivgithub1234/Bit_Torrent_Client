package main

import (
	"log"
	"os"

	"BITTORRENTCLIENT/torrentfiles"
)

func main() {
	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfiles.Open(inPath)
	if err != nil {
		log.Fatal(err)
		log.Println("For running the program, use the following command:")
		log.Println("go run main.go <torrent-file> <output-file>")
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
		log.Println("For running the program, use the following command:")
		log.Println("go run main.go <torrent-file> <output-file>")
	}
}

// for running the program
// go run main.go <torrent-file> <output-file>
