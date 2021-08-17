package sounds

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

func GetAudioBuffer(ps string) (buffer [][]byte, err error) {
	f, err := getDCAFile(ps)
	if err != nil {
		return
	}

	var framelen int16 // opus frame len
	buffer = make([][]byte, 0)
	for {
		// first part of the file
		err = binary.Read(f, binary.LittleEndian, &framelen)
		// EOF
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			err = f.Close()
			return
		}
		if err != nil {
			return
		}

		// read encoded PCM from DCA file
		InBuf := make([]byte, framelen)
		err = binary.Read(f, binary.LittleEndian, &InBuf)
		// check for EOF
		if err != nil {
			log.Println("3")
			return
		}

		buffer = append(buffer, InBuf)
	}
}

func getDCAFile(ps string) (file *os.File, err error) {
	fpath := "./media/" + ps + ".dca"
	file, err = os.Open(fpath)
	if err != nil {
		log.Println("could not open " + fpath)
		return
	}
	return
}
