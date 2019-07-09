package api

import (
	"fmt"
	"io"
	"log"
	"os"
)

func Copy(from string, to string, offset int, limit int) {

	fileSrc, errSrcOpen := os.Open(from)
	if errSrcOpen != nil {
		log.Panic(errSrcOpen)
	}

	fileDest, errDistCreate := os.Create(to)

	if errDistCreate != nil {
		log.Panic(errDistCreate)
	}

	fi, _ := fileSrc.Stat() // todo handle error

	if limit == 0 {
		limit = int(fi.Size())
	}

	_, errSeek := fileSrc.Seek(int64(offset), 1)

	if errSeek != nil {
		log.Panic(errSeek)
	}

	readOffset := 0

	for readOffset < limit {

		bufSize := 1024

		if limit-readOffset < 1024 {
			bufSize = limit - readOffset
		}

		buf := make([]byte, bufSize)
		read, errSrc := fileSrc.Read(buf)

		_, errDist := fileDest.Write(buf)

		if errDist != nil {
			log.Panic(errDist)
		}

		readOffset += read

		fmt.Printf("copied %d%% \n", int(float32(readOffset)/float32(limit)*100))

		if errSrc == io.EOF {
			break
		}
		if errSrc != nil {
			fmt.Printf("failed to read: %v", errSrc)
		}
	}

}
