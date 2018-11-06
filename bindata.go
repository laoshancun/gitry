// Code generated by go-bindata. DO NOT EDIT.
// sources:
// views\index.html
// views\layouts\layout.html

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataViewsindexhtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x92\xc1\x8e\xd3\x30\x10\x86\xcf\xe9\x53\x58\xde\x6a\xe3\x48\x95\xc3" +
		"\x22\xc4\xa1\x4d\x72\xd9\x1b\x12\x70\x58\x6e\x68\x0f\x6e\x32\x8d\xcd\xa6\x76\x64\x4f\x43\xab\x28\xcf\x00\x0f\xc1" +
		"\x53\xf0\x3c\x88\xe7\x40\xb6\xb3\x4b\x84\x2a\x2e\x13\xcd\xcc\x37\x1e\xff\xbf\x53\x34\x6a\xa8\x56\x49\x21\xef\xaa" +
		"\x71\xe4\xef\xc1\x39\xd1\xc2\x34\x15\xb9\xbc\xf3\x65\x75\x6c\x89\x6a\x4a\xaa\x8e\xed\xbd\xe8\xb1\x96\x82\x12\x87" +
		"\x97\x0e\x4a\xfa\x55\x35\x28\xb7\xe4\xf5\x9b\x57\xfd\x79\x27\x41\xb5\x12\xb7\xe4\xad\x4f\x28\xc9\xc3\xac\xee\x4f" +
		"\x48\xf0\xd2\x43\x49\x11\xce\x48\xc3\x49\x78\xc6\x7b\xd3\xc0\x15\x46\xaa\xa6\x01\xfd\x42\x7d\x32\x4f\x3e\x0b\xd8" +
		"\xfe\x84\x68\x74\xe8\xec\x51\x3f\x9c\xf6\x47\x85\xb4\xfa\xfd\xed\xfb\xaf\x9f\x3f\x8a\x3c\x76\x3d\xe7\x6a\xab\xfa" +
		"\xe5\xce\xfc\x8b\x18\x44\xac\xd2\x6a\x95\x24\x83\xb0\x64\xad\x8e\x6d\xb9\x66\xf4\x66\x21\x2a\xe3\x46\xb3\xb4\xee" +
		"\x54\xfd\x94\x6e\x0e\x27\x5d\xa3\x32\x9a\xb0\x8c\x8c\xab\x24\x49\xd6\xbc\x05\x7c\xf7\xf0\xf1\x03\xa3\xf9\x00\x56" +
		"\x1d\x2e\xb5\x97\xb0\x00\x2d\xb8\x99\x4d\xd4\xc1\x67\xe4\xf6\x96\x58\x70\x9f\xa9\x43\x81\x40\x1f\xb3\xd8\x4c\xfc" +
		"\x72\x2e\x10\x2d\x4b\x9d\xad\xd3\x4d\x60\x1a\x81\x82\x3e\x66\xbb\x19\x61\xf4\xe6\x45\x7f\xc6\x07\xd1\xb1\x40\x61" +
		"\x28\x3c\x63\x93\x8f\x53\x48\x62\xf4\x63\x7f\xcd\xf9\xbf\x20\x6f\x83\x97\x50\xce\xbb\xc2\x8b\xc4\x55\xf1\x78\x0f" +
		"\x84\x7d\xe5\x95\xdb\x44\x64\xcd\x7b\xe3\x90\xa5\x0b\x47\xd2\xcd\x48\x83\x33\x5b\x1f\x37\xf3\x8d\xb7\xe1\x33\x65" +
		"\x1c\x25\x68\x76\xd5\xb3\xda\x68\x67\x3a\xe0\x9d\x69\x43\x79\x17\xc5\xf1\x83\x50\xdd\x62\x02\xac\xfd\x77\x02\xac" +
		"\x35\x36\x34\xe6\x99\x67\x3f\x8a\x3c\x3e\x7b\x45\x56\x45\xee\x7f\xf2\x3f\x01\x00\x00\xff\xff\xdf\x31\xa0\xb2\xea" +
		"\x02\x00\x00")

func bindataViewsindexhtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataViewsindexhtml,
		"views/index.html",
	)
}

func bindataViewsindexhtml() (*asset, error) {
	bytes, err := bindataViewsindexhtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "views/index.html",
		size:        746,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1526654108, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataViewslayoutslayouthtml = []byte(
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x24\x90\x41\x6e\x03\x21\x0c\x45\xf7\x9c\xc2\x61\xd3\x15\xa0\x76\x11\x65" +
		"\xc1\x70\x8a\x5e\x80\x80\xa7\x90\x30\x90\x62\x4f\xa3\xd1\x68\xee\x5e\x25\xac\x6c\xff\xff\xf5\x24\x7f\x7b\x8a\x2d" +
		"\xf0\xf6\x40\x48\xbc\x14\x27\xec\x6b\x40\xf4\xec\xd5\xdc\xfd\x82\xcf\xd6\xef\x93\xfc\x5b\x51\x3a\x21\x6c\x42\x1f" +
		"\x9d\x00\xb0\x0b\xb2\x87\x90\x7c\x27\xe4\x49\xae\x3c\xab\x8b\x7c\x1b\x9c\xb9\xa0\xdb\x77\xfd\xfd\x5a\x8e\xc3\x9a" +
		"\xa1\xbc\xbc\x92\xeb\x1d\x3a\x96\x49\x12\x6f\x05\x29\x21\xb2\x84\xd4\x71\x9e\xa4\x31\x81\x48\x9f\xcf\xc5\x3f\x75" +
		"\xa8\xe6\xe9\x67\x33\x42\x3a\x10\x0d\xf4\x49\x29\xe0\x94\x09\x2a\x62\x24\xe0\x06\x57\x84\xd2\x7c\xc4\x08\x57\x9c" +
		"\x5b\x47\xf8\x59\x73\xc4\x0f\x82\x5c\x4b\xae\x08\x14\x7a\x7e\x30\x81\x52\x6f\xc0\x38\x81\x7a\x98\xa4\xb9\x91\xb9" +
		"\xfd\xae\xd8\x37\xf5\xa9\x2f\xfa\x4b\x2f\xb9\xea\x1b\x49\x67\xcd\x88\x39\x61\xcd\x78\xd7\x5e\x5b\xdc\x9c\xd8\x77" +
		"\xd8\x32\x96\x08\xc7\x01\x20\xac\x19\xaa\x35\xef\xde\xfe\x03\x00\x00\xff\xff\xaf\x68\x48\x2b\x47\x01\x00\x00")

func bindataViewslayoutslayouthtmlBytes() ([]byte, error) {
	return bindataRead(
		_bindataViewslayoutslayouthtml,
		"views/layouts/layout.html",
	)
}

func bindataViewslayoutslayouthtml() (*asset, error) {
	bytes, err := bindataViewslayoutslayouthtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name:        "views/layouts/layout.html",
		size:        327,
		md5checksum: "",
		mode:        os.FileMode(438),
		modTime:     time.Unix(1541522241, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"views/index.html":          bindataViewsindexhtml,
	"views/layouts/layout.html": bindataViewslayoutslayouthtml,
}

//
// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op:   "open",
					Path: name,
					Err:  os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op:   "open",
			Path: name,
			Err:  os.ErrNotExist,
		}
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"views": {Func: nil, Children: map[string]*bintree{
		"index.html": {Func: bindataViewsindexhtml, Children: map[string]*bintree{}},
		"layouts": {Func: nil, Children: map[string]*bintree{
			"layout.html": {Func: bindataViewslayoutslayouthtml, Children: map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
