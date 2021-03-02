package main

import (
	"bytes"
	"os"
)

/**

1: some reader type which will read a large chunk and return the lines in there
	read x bytes into buffer

2: some writer func which will write a large chunk

 */

type fileWriter struct {
	file *os.File
	bufSize int
	offset int
}

type fileReader struct {
	file *os.File
	bufSize int
	chunk []byte
}

func newFileReader(file *os.File)fileReader{
	return fileReader{
		file:    file,
		bufSize: 1024,
		chunk:   []byte{},
	}
}

func newFileWriter(file *os.File)fileWriter{
	return fileWriter{
		file:    file,
		bufSize: 1028,
	}
}

func (w *fileWriter) WriteLines(chunks [][]byte){
	buf := []byte{}
	for _, chunk := range(chunks){
		buf = append(buf, chunk...)
	}
	w.file.Write(buf)
}

func (r *fileReader) ReadLines() ([][]byte, error){
	buf := make([]byte, r.bufSize)
	c, err := r.file.Read(buf)
	if err != nil{
		return [][]byte{}, err
	}
	if c == r.bufSize{
		chunks := bytes.Split(buf, []byte("\n"))
		chunks[0] = append(r.chunk, chunks[0]...)
		r.chunk = chunks[len(chunks)-1]
		return chunks[:len(chunks)-1], err
	}else{
		chunks := bytes.Split(buf, []byte("\n"))
		chunks[0] = append(r.chunk, chunks[0]...)
		return chunks[:len(chunks)-1], err
	}
}

func (r fileReader) ResetPointer(){
	r.file.Seek(0, 0)
}
