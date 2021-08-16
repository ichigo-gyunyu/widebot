package sounds

import (
	"encoding/binary"
	"io"
	"os"
)

func GetAudioBuffer(file string) (buffer [][]byte, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()

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
			return
		}

		buffer = append(buffer, InBuf)
	}
}
