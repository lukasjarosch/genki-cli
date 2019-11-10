// Code generated by vfsgen; DO NOT EDIT.

package internal

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	pathpkg "path"
	"time"
)

// Templates statically implements the virtual filesystem provided to vfsgen.
var Templates = func() http.FileSystem {
	fs := vfsgen۰FS{
		"/": &vfsgen۰DirInfo{
			name:    "/",
			modTime: time.Date(2019, 11, 10, 19, 3, 33, 868717316, time.UTC),
		},
		"/makefile.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "makefile.go.tmpl",
			modTime:          time.Date(2019, 11, 10, 14, 18, 30, 480751816, time.UTC),
			uncompressedSize: 1071,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x93\x51\x6b\xdb\x30\x10\xc7\x9f\xad\x4f\x71\x98\x50\xb6\x07\xc9\xef\x81\x8c\x84\x24\xa4\x61\xa4\x29\x5d\x5e\x06\x83\xa2\xc8\x17\x59\x8d\x2c\x05\x49\xf6\x28\xc1\xdf\x7d\x58\x76\xeb\x66\x64\xd9\x60\x2f\xb2\xa4\x3b\xf9\xff\xd3\xff\x4e\xab\xf5\xee\x79\xbe\xdd\x6c\xd6\xbb\xc9\xe8\x93\x2f\x50\x6b\x90\x2a\x80\xc3\x9a\x9e\xb8\xf3\x08\x94\xfa\xc2\xba\x00\xf7\xcb\xd9\xe2\x33\x59\x6c\xe7\x5f\x97\x4f\xcf\xeb\xcd\x6c\xb5\x9c\xa4\xe7\x33\xb0\xb9\x35\x07\x25\xd9\xa3\xb3\x2f\x28\x02\x5b\x58\x71\x44\xf7\x84\x52\xf9\xe0\x5e\xa1\x69\xb2\x2b\x49\x0f\xbc\x44\x7f\xe2\x02\xa1\x69\xe8\x95\xf8\x37\x74\xb5\x8a\xd1\xf4\x4d\x71\x37\x5b\x4d\xd2\x9a\x8e\xce\x03\x70\x43\x73\xac\x53\x42\xd8\xe3\xfd\xf6\xe1\xfb\x18\x5c\x65\x88\xab\xcc\x98\x24\x53\x14\x85\x85\x94\xd2\x2f\xe0\x03\x77\x41\x19\x09\xda\x0a\xae\xf5\x6b\x4a\x92\xa9\xb4\x6d\x2e\x88\x32\xbf\x06\x37\x88\x67\x25\x57\x86\x49\x0b\x94\x6a\x2b\xa9\xc6\x1a\xf5\x24\xc7\x7d\x25\x07\xd1\x3c\xde\x97\x74\x9f\x4b\xe9\x7d\xa5\x74\xde\x4a\x77\x41\x50\x25\x97\x08\xa3\xf3\x47\x0f\x9b\xf1\xfb\x7a\x37\x5b\x35\x29\x49\xfa\xe4\x78\x18\x18\xd0\x70\xfb\x04\xd0\x03\xb0\xac\x73\xfd\xa0\x34\xfe\x4e\x46\x5b\x57\x86\xe9\x1f\xcc\x11\xd6\x04\xae\x0c\xba\xbf\xe2\x4d\x7b\xbe\xd6\xc0\x1f\x24\x49\x12\x4a\x5d\xf9\x36\x33\xbc\x44\xf8\xaf\x7a\xbf\xff\x09\xc3\x4f\xeb\x8e\x50\x58\x1f\xba\xbd\x9b\x60\xc3\xad\xa3\x6f\x24\x8e\xd7\xab\xd1\x37\x40\xe7\x2f\xb5\xc0\xb2\xbd\x32\xb7\xfb\x00\x58\xf6\xcf\xbd\x32\x90\x08\x8d\xdc\x90\x38\x5e\x92\xc4\xad\xd6\xf5\xea\x04\x7b\x65\xb8\x53\xe8\x7b\xaa\x18\x02\x1a\xb8\xf4\x60\x30\xb4\xad\xa7\x80\x65\x8c\x31\x92\x4c\x5d\x09\xd4\x75\xbc\x83\x4a\x75\xca\x79\x40\x2a\xd1\x1c\x15\xf9\xb8\xb8\xd4\xac\xf2\x13\x8f\x95\x8e\xb1\x5e\x4d\x62\x68\x1f\x7b\x51\xed\x99\xb0\x65\xa6\xab\x23\xf7\x2f\xdc\x59\x2f\x8a\x2c\xe6\xc1\xdd\x1d\x48\x0b\xa5\xcd\xa1\x46\x93\x5b\x47\x7e\x05\x00\x00\xff\xff\x80\xee\x49\x9d\x2f\x04\x00\x00"),
		},
		"/readme.go.tmpl": &vfsgen۰CompressedFileInfo{
			name:             "readme.go.tmpl",
			modTime:          time.Date(2019, 11, 10, 19, 3, 33, 862050708, time.UTC),
			uncompressedSize: 231,

			compressedContent: []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\xcd\x31\x0b\xc2\x30\x10\x05\xe0\x3d\xbf\xe2\x41\xb7\x93\xf6\x07\x74\x70\x71\x70\x13\x27\xe7\x1e\xcd\x19\xa3\xf6\x02\x69\xd0\x21\xf4\xbf\xcb\x49\xc1\x41\x71\x4b\x8e\xf7\xbd\xd7\xa0\x56\x74\xbb\xa4\xe7\x18\xba\x63\x4e\x57\x19\x4b\x77\xe0\x49\xb0\x2c\x6e\x8b\xbd\xa8\x64\x2e\xe2\xf1\x8c\xe5\x82\x21\x88\xde\x62\x3b\xde\xe3\x80\xc7\x0f\x78\x92\x3c\xc7\xa4\x66\x5d\xd3\x60\x3d\x63\x92\xc2\x9e\x0b\xbb\x0d\x88\xac\xbc\x27\xfa\xb7\x6b\xb1\x75\x39\xe5\xb9\x27\x72\xb5\xb6\xc8\xac\x41\xbe\xcc\x27\x67\x12\x40\xfb\x6e\xb6\x8f\x21\x51\x6f\xcf\x57\x00\x00\x00\xff\xff\x66\x7b\x1c\x3b\xe7\x00\x00\x00"),
		},
	}
	fs["/"].(*vfsgen۰DirInfo).entries = []os.FileInfo{
		fs["/makefile.go.tmpl"].(os.FileInfo),
		fs["/readme.go.tmpl"].(os.FileInfo),
	}

	return fs
}()

type vfsgen۰FS map[string]interface{}

func (fs vfsgen۰FS) Open(path string) (http.File, error) {
	path = pathpkg.Clean("/" + path)
	f, ok := fs[path]
	if !ok {
		return nil, &os.PathError{Op: "open", Path: path, Err: os.ErrNotExist}
	}

	switch f := f.(type) {
	case *vfsgen۰CompressedFileInfo:
		gr, err := gzip.NewReader(bytes.NewReader(f.compressedContent))
		if err != nil {
			// This should never happen because we generate the gzip bytes such that they are always valid.
			panic("unexpected error reading own gzip compressed bytes: " + err.Error())
		}
		return &vfsgen۰CompressedFile{
			vfsgen۰CompressedFileInfo: f,
			gr:                        gr,
		}, nil
	case *vfsgen۰DirInfo:
		return &vfsgen۰Dir{
			vfsgen۰DirInfo: f,
		}, nil
	default:
		// This should never happen because we generate only the above types.
		panic(fmt.Sprintf("unexpected type %T", f))
	}
}

// vfsgen۰CompressedFileInfo is a static definition of a gzip compressed file.
type vfsgen۰CompressedFileInfo struct {
	name              string
	modTime           time.Time
	compressedContent []byte
	uncompressedSize  int64
}

func (f *vfsgen۰CompressedFileInfo) Readdir(count int) ([]os.FileInfo, error) {
	return nil, fmt.Errorf("cannot Readdir from file %s", f.name)
}
func (f *vfsgen۰CompressedFileInfo) Stat() (os.FileInfo, error) { return f, nil }

func (f *vfsgen۰CompressedFileInfo) GzipBytes() []byte {
	return f.compressedContent
}

func (f *vfsgen۰CompressedFileInfo) Name() string       { return f.name }
func (f *vfsgen۰CompressedFileInfo) Size() int64        { return f.uncompressedSize }
func (f *vfsgen۰CompressedFileInfo) Mode() os.FileMode  { return 0444 }
func (f *vfsgen۰CompressedFileInfo) ModTime() time.Time { return f.modTime }
func (f *vfsgen۰CompressedFileInfo) IsDir() bool        { return false }
func (f *vfsgen۰CompressedFileInfo) Sys() interface{}   { return nil }

// vfsgen۰CompressedFile is an opened compressedFile instance.
type vfsgen۰CompressedFile struct {
	*vfsgen۰CompressedFileInfo
	gr      *gzip.Reader
	grPos   int64 // Actual gr uncompressed position.
	seekPos int64 // Seek uncompressed position.
}

func (f *vfsgen۰CompressedFile) Read(p []byte) (n int, err error) {
	if f.grPos > f.seekPos {
		// Rewind to beginning.
		err = f.gr.Reset(bytes.NewReader(f.compressedContent))
		if err != nil {
			return 0, err
		}
		f.grPos = 0
	}
	if f.grPos < f.seekPos {
		// Fast-forward.
		_, err = io.CopyN(ioutil.Discard, f.gr, f.seekPos-f.grPos)
		if err != nil {
			return 0, err
		}
		f.grPos = f.seekPos
	}
	n, err = f.gr.Read(p)
	f.grPos += int64(n)
	f.seekPos = f.grPos
	return n, err
}
func (f *vfsgen۰CompressedFile) Seek(offset int64, whence int) (int64, error) {
	switch whence {
	case io.SeekStart:
		f.seekPos = 0 + offset
	case io.SeekCurrent:
		f.seekPos += offset
	case io.SeekEnd:
		f.seekPos = f.uncompressedSize + offset
	default:
		panic(fmt.Errorf("invalid whence value: %v", whence))
	}
	return f.seekPos, nil
}
func (f *vfsgen۰CompressedFile) Close() error {
	return f.gr.Close()
}

// vfsgen۰DirInfo is a static definition of a directory.
type vfsgen۰DirInfo struct {
	name    string
	modTime time.Time
	entries []os.FileInfo
}

func (d *vfsgen۰DirInfo) Read([]byte) (int, error) {
	return 0, fmt.Errorf("cannot Read from directory %s", d.name)
}
func (d *vfsgen۰DirInfo) Close() error               { return nil }
func (d *vfsgen۰DirInfo) Stat() (os.FileInfo, error) { return d, nil }

func (d *vfsgen۰DirInfo) Name() string       { return d.name }
func (d *vfsgen۰DirInfo) Size() int64        { return 0 }
func (d *vfsgen۰DirInfo) Mode() os.FileMode  { return 0755 | os.ModeDir }
func (d *vfsgen۰DirInfo) ModTime() time.Time { return d.modTime }
func (d *vfsgen۰DirInfo) IsDir() bool        { return true }
func (d *vfsgen۰DirInfo) Sys() interface{}   { return nil }

// vfsgen۰Dir is an opened dir instance.
type vfsgen۰Dir struct {
	*vfsgen۰DirInfo
	pos int // Position within entries for Seek and Readdir.
}

func (d *vfsgen۰Dir) Seek(offset int64, whence int) (int64, error) {
	if offset == 0 && whence == io.SeekStart {
		d.pos = 0
		return 0, nil
	}
	return 0, fmt.Errorf("unsupported Seek in directory %s", d.name)
}

func (d *vfsgen۰Dir) Readdir(count int) ([]os.FileInfo, error) {
	if d.pos >= len(d.entries) && count > 0 {
		return nil, io.EOF
	}
	if count <= 0 || count > len(d.entries)-d.pos {
		count = len(d.entries) - d.pos
	}
	e := d.entries[d.pos : d.pos+count]
	d.pos += count
	return e, nil
}
