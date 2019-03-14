package pdf

import (
	"bufio"
	"bytes"
	"fmt"
)

// PDF File Structure.
// https://en.wikipedia.org/wiki/PDF#Technical_overview

/*
Document represents a PDF document.
*/
type Document struct {
	Version string
}

/*
Parse takes content and returns a document that can be interacted with.
*/
func Parse(content []byte) *Document {
	bytesReader := bytes.NewReader(content)
	bufReader := bufio.NewReader(bytesReader)

	// Verify PDF version.
	versionRaw, _, _ := bufReader.ReadLine()
	if versionRaw[0] == 37 && versionRaw[1] == 80 && versionRaw[2] == 68 && versionRaw[3] == 70 && versionRaw[4] == 45 {
		fmt.Println("Valid pdf!")
		version := []byte{versionRaw[5], versionRaw[6], versionRaw[7]}
		fmt.Println(string(version))

		return &Document{
			Version: string(version),
		}
	}

	return nil
}
