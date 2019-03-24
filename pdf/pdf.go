package pdf

import (
	"bufio"
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
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
Stream represents a stream inside the PDF-documet
*/
type Stream struct {
	Header  []byte
	Content []byte
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

/*
ParseFile takes content and returns a document that can be interacted with.
*/
func ParseFile(path string) *Document {
	p := fmt.Println

	// Discovery, find out where all the fun stuff is
	pdfBody, _ := ioutil.ReadFile(path)
	file, _ := os.Open(path)
	streams := findStreams(pdfBody, file)
	for i := 0; i < 1; i++ {
		// p(i, string(streams[i].Header))
		p(i, string(streams[i].Content))
	}

	// file.Seek(489, 0)
	// b1 := make([]byte, 57)
	// n1, err := file.Read(b1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%d bytes: %s\n", n1, string(b1))
	// b2 := make([]byte, 4)
	// n2, err := file.Read(b2)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%d bytes: %s\n", n2, string(b2))

	// b := bytes.NewBufferString("string in here")
	// r, err := zlib.NewReader(b)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// io.Copy(os.Stdout, r)
	// r.Close()
	// regexpPattern, parseErr := regexp.Compile("<<[^>]*>>stream")
	// if parseErr != nil {
	// 	log.Fatal(parseErr)
	// }

	// regexpResults := regexpPattern.FindAllSubmatch(content, -1)
	// regexpResults := regexpPattern.FindAllSubmatchIndex(content, -1)
	// p(regexpResults)

	// for i := 0; i < len(regexpResults); i++ {
	// 	// data := string(regexpResults[i][2])
	// 	// metaData := strings.Split(data, "/")
	// 	p(regexpResults[i])
	// }

	return &Document{
		Version: string("123"),
	}
}

func findStreams(content []byte, file *os.File) []Stream {
	// p := fmt.Println
	streamHeaderRegexp, _ := regexp.Compile("<<.*?>>stream")
	// streamEndRegexp, _ := regexp.Compile("endstream")

	file.Seek(126, 0)
	x := make([]byte, 16)
	r, _ := file.Read(x)
	fmt.Println(r)
	fmt.Println(string(x))

	streamHeaderIndices := streamHeaderRegexp.FindAllSubmatchIndex(content, -1)
	fmt.Println(streamHeaderIndices)
	// streamEndIndices := streamEndRegexp.FindAllSubmatchIndex(content, -1)

	headerIndicesLen := len(streamHeaderIndices)
	fmt.Println(headerIndicesLen)
	// endIndicesLen := len(streamEndIndices)
	// if headerIndicesLen != endIndicesLen {
	// 	log.Fatal("Something is up, go fetch a pizza")
	// }

	streams := make([]Stream, headerIndicesLen)

	for i := 0; i < headerIndicesLen; i++ {
		headerIndices := streamHeaderIndices[i]
		// endIndices := streamEndIndices[i]
		headerLen := headerIndices[1] - headerIndices[0]
		file.Seek(int64(headerIndices[0]), 0)
		headerBuffer := make([]byte, headerLen)
		headerBytesRead, _ := file.Read(headerBuffer)
		if headerBytesRead == 0 {
			log.Fatal("No bytes were read")
		}
		header := headerBuffer[2 : headerLen-8]

		file.Seek(int64(headerIndices[1]), 0)
		// contentLen := endIndices[0] - headerIndices[1]
		contentBuffer := make([]byte, 73)
		contentBytesRead, _ := file.Read(contentBuffer)
		if contentBytesRead == 0 {
			log.Fatal("No bytes were read")
		}
		content := contentBuffer
		// fmt.Printf("%d bytes: %s\n", res, string(buf[2:headerLen-8]))

		streams[i] = Stream{
			Header:  header,
			Content: content,
		}
	}

	return streams
}

func uncompress(content []byte) []byte {
	buf := bytes.NewBuffer(content)

	gRead, err := zlib.NewReader(buf)
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}

	if t, err := io.Copy(os.Stdout, gRead); err != nil {
		fmt.Println(t)
		return []byte{}
	}

	if err := gRead.Close(); err != nil {
		return []byte{}
	}
	return nil
}
