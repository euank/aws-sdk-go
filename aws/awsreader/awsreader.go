package awsreader

import (
	"bytes"
	"crypto/sha256"
	"io"
	"sync"
)

// AWSReader is an interface that groups the basic Read and Close methods along
// with a caching checksum for the reader.
type AWSReader interface {
	io.Reader
	io.Closer
	Sha256Sum() []byte
}

type reader struct {
	io.Reader

	sync.Mutex
	checksum *[]byte
}

func (r *reader) Close() error {
	if closer, ok := r.Reader.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

func (r *reader) Sha256Sum() []byte {
	r.Lock()
	defer r.Unlock()

	if r.checksum != nil {
		return *r.checksum
	}

	hash := sha256.New()

	writeHash := func(sr io.Reader) {
		packet := make([]byte, 4096)
		for {
			n, err := sr.Read(packet)
			hash.Write(packet[0:n])
			if err == io.EOF {
				break
			}
			// TODO, if err != nil, handle
		}
	}

	if seeker, ok := r.Reader.(io.ReadSeeker); ok {
		seeker.Seek(0, 0)
		writeHash(seeker)
		seeker.Seek(0, 0)
	} else {
		var readerClone bytes.Buffer
		tee := io.TeeReader(r.Reader, &readerClone)
		writeHash(tee)
		r.Reader = &readerClone
	}

	checksum := hash.Sum(nil)
	r.checksum = &checksum

	return checksum
}

// New constructs a new AWSReader from an existing ReadSeeker. if the argument
// is also a Closer, that method will be exposed as well.
func New(s io.ReadSeeker) AWSReader {
	return &reader{Reader: s}
}

// NewInMemory constructs a new AWSReader that buffers its contents to memory
// from any existing Reader.  If the passed in reader has a Seek method then it
// will not buffer to memory, preferring to use Seeking.  If the passed in
// reader has a Close method then it will be exposed.
func NewInMemory(r io.Reader) AWSReader {
	return &reader{Reader: r}
}
