package paasio

import (
	"io"
	"sync"
)

// MyReadCounter implementes the ReadCounter interface
type MyReadCounter struct {
	sync.Mutex
	r io.Reader
	readBytes int64
	readCalls int
}

// MyWriteCounter implementes the ReadCounter interface
type MyWriteCounter struct {
	sync.Mutex
	w io.Writer
	writeBytes int64
	writeCalls int
}

// MyRWCounter composes ReadCounter and WriteCounter
type MyRWCounter struct {
	sync.Mutex
	MyReadCounter
	MyWriteCounter
}

// NewReadCounter creates a ReadCounter object
func NewReadCounter(r io.Reader) ReadCounter {
	return &MyReadCounter{r:r}
}

// NewWriteCounter creates a WriteCounter object
func NewWriteCounter(w io.Writer) WriteCounter {
	return &MyWriteCounter{w:w}
}

// NewReadWriteCounter creates a ReadWriteCounter object
func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	var res = &MyRWCounter{}
	res.MyReadCounter.r = rw
	res.MyWriteCounter.w = rw
	return res
}

func (rw *MyReadCounter) Read(p []byte) (int, error) {
	rw.Lock()
	defer rw.Unlock()

	n, err := rw.r.Read(p)

	rw.readBytes += int64(n)
	rw.readCalls ++

	return n, err
}

// ReadCount returns the bytes read and no. of read method calls until now
func (rw *MyReadCounter) ReadCount() (int64, int) {
	rw.Lock()
	defer rw.Unlock()
	return rw.readBytes, rw.readCalls
}


func (rw *MyWriteCounter) Write(p []byte) (int, error) {
	rw.Lock()
	defer rw.Unlock()

	n, err := rw.w.Write(p)

	rw.writeBytes += int64(n)
	rw.writeCalls ++

	return n, err
}

// WriteCount returns the bytes write and no. of write method calls until now
func (rw *MyWriteCounter) WriteCount() (int64, int) {
	rw.Lock()
	defer rw.Unlock()
	return rw.writeBytes, rw.writeCalls
}
