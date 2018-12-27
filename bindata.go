// Code generated by go-bindata.
// sources:
// public/index.html
// DO NOT EDIT!

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

	"github.com/elazarl/go-bindata-assetfs"
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
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
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
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _publicIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xbc\x58\x6d\x73\xdb\xb8\x11\xfe\xee\x5f\xb1\xc3\xe6\x83\x34\xb6\x49\x59\xce\xd5\xe9\x85\x52\x27\x17\xf7\x6c\x5f\x9a\x97\xfa\x25\x99\x6b\xa7\xe3\x01\xc9\x25\x01\x0b\x04\x18\x60\x29\x59\xed\xdd\xcc\x7d\xef\xd7\xfe\xc2\xfb\x25\x1d\x80\xa4\x24\x32\x8e\x93\x9b\xce\xdc\x27\xe1\x65\xf9\xec\xee\xb3\x0f\x80\xb5\x63\x4e\xa5\x9c\xef\xc5\x1c\x59\x36\xdf\x03\x88\x4b\x24\x06\x29\x67\xc6\x22\xcd\x82\x9a\xf2\xc3\x67\x81\xdf\x20\x41\x12\xe7\xe7\xc8\x24\x71\x78\xc9\x31\x5d\xc0\x6b\x61\xad\xd0\x0a\x0e\xe1\x4c\x4b\xa6\x8a\x38\x6a\x8c\x9c\xb9\x14\x6a\x01\x06\xe5\x2c\xb0\xb4\x96\x68\x39\x22\x05\xc0\x0d\xe6\xb3\x80\x13\x55\xf6\xdb\x28\x2a\xd9\x7d\x9a\xa9\x30\xd1\x9a\x2c\x19\x56\xb9\x49\xaa\xcb\x68\xb3\x10\x1d\x87\xc7\xe1\x49\x94\x5a\xbb\x5d\x0b\x4b\xa1\xc2\xd4\xda\x00\x84\x22\x2c\x8c\xa0\xf5\x2c\xb0\x9c\x1d\x3f\x7b\x7a\xf8\xdd\xfb\x1f\x85\xb8\xba\xf8\x1e\x5f\x1d\x65\x67\xe5\x0f\x97\x2f\x16\xeb\xb4\x3e\x7f\x71\x7e\x59\x1c\x4f\xdf\x96\x37\xe9\x6a\x75\xa2\xd5\xf1\xe5\x8f\x59\xf1\xf4\x3d\xdb\x7f\x57\x5e\x5d\xdb\x7f\x45\xaf\xfe\xf8\x6c\x99\x64\x7f\xb9\xe3\x4f\xeb\x60\x0f\x00\x20\x35\xda\x5a\x6d\x44\x21\xd4\x2c\x60\x4a\xab\x75\xa9\x6b\x1b\xcc\xf7\xe2\xa8\xa1\x69\x2f\x4e\x74\xb6\xf6\x69\x66\x62\x09\xa9\x64\xd6\xce\x82\x54\x2b\x62\x42\xa1\xf1\x7c\x35\x7b\xcd\xa8\x6f\x67\x31\xa5\xa3\xa0\xdb\x01\x88\xf9\x14\x44\x36\x0b\x6e\xb9\xe7\xf6\x36\x75\xdc\xde\x26\x5a\x5b\x72\x58\x3d\xc6\xbf\x6b\x56\xe3\x88\x4f\x77\x00\x06\xe0\x42\x2b\x17\xdf\x8e\x8b\xbe\x4d\xc5\x0c\x2b\x0c\xab\x78\xcf\x02\x20\xae\xe6\xa7\x58\x6a\xe5\x68\x26\xb4\xc0\xf5\x0a\x88\x23\xbc\xaa\x13\x34\x0a\x09\x6d\xf4\xb6\x42\x65\xb9\xc8\x09\x9a\x60\xc1\x07\x6b\x61\xa5\xcd\x02\x48\x43\x86\x84\xa6\x14\x0a\x41\xe4\xc0\x60\x43\x09\x08\x0b\x96\x84\x94\xc0\xa4\x58\x22\x8c\x1c\x6e\x8c\xe5\xdc\xcd\x14\x5a\x1b\x47\x58\xf6\xa3\x01\xd0\xb9\x77\xbf\x01\x19\x03\x53\x19\x18\x64\xd9\xda\xf9\xb2\x68\x96\xe8\x2d\xc8\xb0\x3c\x17\x29\xe4\xda\xf8\xf9\xf9\xf5\xf5\x3b\x40\x95\x55\x5a\x28\xb2\x1d\x0e\xab\x2a\x29\x52\xe6\xe8\xd9\xfa\x77\x68\xe2\xab\x03\x08\xe3\xa8\xea\xb1\x1a\xed\xd4\xf8\x37\xd0\x7c\xed\x98\xda\x30\x0d\xc4\x85\x85\x04\x39\x5b\x0a\x6d\x0e\x3e\x09\x36\xd5\x2a\x17\x45\x6d\xd0\x02\x83\x38\xd5\x19\xce\xa3\x86\xfe\x38\xf2\xb3\x7e\xc2\x07\xb0\xe2\x22\xe5\x20\xec\x20\x9d\xda\x62\x36\x58\x4a\xd6\x0f\x57\x97\x34\x08\x6b\xeb\x96\x4a\x83\x1f\x6b\xb4\x64\x43\xb8\x18\x30\x32\xa8\xeb\xaf\xbf\xfc\xf7\xd7\x5f\xfe\xd3\xb8\x2f\x91\x29\xdb\x94\xa3\x91\x8a\x83\x1a\xb8\xef\x42\x1e\x2c\x0b\x0b\x2c\x91\xe8\xa2\x30\x58\xc9\x75\x03\xeb\xa0\x4a\xa6\x58\x81\x25\x2a\x82\x4a\x32\xca\xb5\x29\x61\xe5\xdc\x1b\x4c\xd1\x09\xcb\xc7\xeb\x48\x81\xe9\x64\x02\xcc\x31\x66\xd0\x56\x5a\x59\x3c\xf0\xf2\x51\x1a\xf2\xda\x10\x47\x03\xcc\x1f\x95\x4f\x7d\x13\x5b\xa0\x0a\xe1\x45\x4e\x68\x60\xad\x6b\x48\xa5\x48\x17\x2e\x95\x81\x6d\x53\x8b\x2b\xd2\x15\x5c\xa1\x59\x8a\x14\xbb\x82\x24\x35\x91\x56\x07\x9f\xaa\x11\x2c\x31\x43\xc3\xca\x18\xa4\xda\x28\xa1\x0a\x60\xf0\xcd\x64\x32\x08\xd9\x81\x6c\xb2\x25\x8e\xc3\x90\x0d\x3e\x08\xea\x3f\xd3\x19\xac\x04\x71\x3f\xc9\x99\x90\x98\x6d\x8b\x17\xc2\x07\x2e\x24\x6e\x0c\x85\x85\x4c\xaf\xda\xa8\xdb\xfb\x07\x6e\x2e\xa0\xd2\x52\x36\x95\xb4\x4d\x96\x03\x47\x15\x1a\xa1\x33\x91\x32\x29\xd7\x43\xc9\x29\x12\x72\xd7\x41\x1b\x2b\x66\x21\x5c\x28\xbf\xe1\x74\x42\xa2\xc4\x03\xc8\x34\xc4\x96\x8c\x56\xc5\x5c\x69\x8a\xa3\x76\x0c\x06\x73\x83\xb6\xc9\xa1\x62\x05\x42\x82\x29\xab\x2d\x82\x20\x5f\xfc\x81\x4f\xa5\x09\x92\x26\x56\xcc\x06\x11\xe4\x42\x09\xcb\x71\x13\x87\x50\xc5\x6f\x3a\xd3\xee\x7a\x9d\x0e\xcf\x33\x3f\x6e\xee\xef\xda\x0a\x55\xdc\x12\xc7\xdb\xc2\x20\x3a\xec\xdb\x96\xb0\x60\x7e\xe3\xf6\x7c\x18\xdd\x5e\x47\x66\x1c\xf1\xe3\x47\xfc\xf7\xa6\xed\xa4\x9b\x79\x41\xb4\x91\xb9\xf1\xa1\x50\x52\x28\x0c\x1e\x7e\x1a\xbc\x45\x61\x74\x5d\xf5\x5f\x06\xc9\x12\x94\xee\xee\x9c\x05\x8a\x95\x18\xcc\xdf\xb0\x12\xe3\xc8\x2f\xf7\x0c\x85\xaa\x6a\x02\x5a\x57\x38\x0b\x08\xef\x29\xe8\x21\x3b\x59\x19\x2d\x03\xcf\x85\x07\x72\xa2\x4d\x91\x6b\x99\xa1\x99\x05\x1f\xb4\x91\x59\xf0\x99\xc4\x00\xe2\xe6\xc8\xf8\xaf\x85\x5a\xea\x05\x06\xad\x2b\x5b\x27\xa5\xd8\x3a\x4b\x48\x41\x42\xea\xd0\xd6\x69\x8a\xd6\x06\xf3\x0b\x6f\x1d\x47\x0d\xc0\xc3\x88\x96\x74\xf5\x05\xbc\x8c\xa9\xc2\xbd\xb7\xfd\xe3\xdc\x07\x8d\x23\x97\xea\xb6\x00\xfc\x78\x7e\x89\xb6\x96\xf4\xed\x6e\x15\xe3\xca\xe0\xdc\xdf\x0d\xde\x77\x57\xf0\x43\xe3\x4d\xbb\x80\x77\x0f\x54\xf3\x9c\x35\x4b\x8d\x55\xd8\xde\x24\x71\xe4\xc0\xf6\x7a\x84\x6d\x55\x10\xdb\xd4\x88\x8a\xc0\x9a\x74\xdb\x55\xb1\x3b\x76\x1f\x16\x5a\x17\x12\x59\x25\xac\xef\xa8\xdc\x5a\x24\x45\x62\xa3\xbb\x8f\x35\x9a\x75\x74\x14\x1e\x4d\xc3\xa7\xed\xcc\xb7\x54\x77\x36\x98\xc7\x51\x03\x38\xff\x1c\xf6\xd7\x76\x6c\x77\xc3\x86\xed\xee\xc1\x7e\xed\x3a\xfd\xe6\xe2\x6f\x22\x99\x4c\x4f\x3e\x2e\xd7\x77\x57\xaf\xf3\xf3\xbb\xb7\xaf\xd9\x5f\x17\x79\xfd\xe1\xfd\xfd\xdf\xef\x6f\xde\xa9\x97\x3f\xbc\x38\x91\xd3\xf2\xe5\x87\x37\x17\xd5\xd9\x9f\xca\xb3\x97\xa7\xcf\x56\x67\x6f\x2e\xd2\x77\xa7\x27\xd7\xf7\xec\xf1\x7e\x6d\x9b\xcd\x36\x9d\x86\xca\x25\x33\xe0\x24\x71\x2d\x4a\x7c\xbe\x59\x61\x44\x58\x56\x64\x9f\x37\xf5\xcd\x6b\xe5\x1f\x08\x68\xe4\xd8\x4a\x62\x34\x86\x7f\xb7\x75\x76\xdf\x28\x98\xc1\x93\x51\xf0\x07\x2f\xf8\x71\xb8\x64\x72\x34\x86\x9f\x7e\x82\x56\xee\xcf\x5b\xd3\x27\x61\x81\x34\x0a\x22\x56\x89\xa8\xd3\xc3\x9f\xdd\x37\xb3\x00\xf6\x41\x1d\x6c\x9d\x8d\x0c\xda\xad\x0b\xf0\xe0\x43\x05\x8d\x43\xd7\xb5\x7b\xcb\xce\xc1\xcf\xed\xe8\xe7\xc7\x62\x67\x2a\xbb\x44\x32\xeb\x9b\x4a\xab\xef\x99\x90\xae\xa5\xf8\xbd\xf3\x69\xdc\x18\x4c\x84\x21\x0e\x33\x50\xb8\x82\x53\x46\x38\xda\xa4\xd2\xd9\x64\x22\xcf\x61\x06\xa3\xae\x50\xce\xa5\xfb\x1d\x8d\xe1\xb0\x03\xd8\xae\x8d\x21\x82\xa3\xc9\x64\x32\x44\x61\x89\x85\x19\xbc\x66\xc4\x43\x96\xd8\x91\x03\xed\x79\x7a\x9c\x5f\xd8\x87\xa0\x69\x16\x0d\xa6\x7a\x89\xc6\x75\x9e\x7a\x01\x2e\x4b\x87\xec\xb6\x2d\xa6\x5a\x65\x76\x1c\xec\x14\xa3\x1d\x84\xee\xc9\x1d\x6d\xb9\xf8\xca\xc2\x06\xff\x70\xf8\xfb\xfb\x9d\x20\x9d\x9b\x7f\x76\x77\x92\x7f\xdf\xd8\x92\x09\xe9\x7a\xa4\x30\x0c\xc1\xd7\xd4\xbd\x28\x42\xc1\xb4\x0b\x27\xd8\xc9\xd2\x36\x24\xe9\x9a\x46\x5f\xd4\xc3\x81\x6b\x9c\x26\x9f\xd3\xd5\x93\x51\xa6\xd3\xda\xf5\x5f\xe3\xd0\x37\xe2\x0f\x26\xe7\x52\x6b\xef\xf0\x71\xe8\x9b\xa7\x1d\x33\xdc\x25\x61\x70\xb4\xb6\x21\x63\x58\x19\x5c\xa2\xa2\x53\xcc\x59\x2d\x69\xd4\x8b\x68\xc7\x8f\xbf\xd9\xbf\xe0\x65\x57\xad\xde\x7e\x57\x9f\x19\x23\xd6\x17\xe8\x63\x95\xb9\x1e\xf4\xe6\x9c\xb9\xce\x1d\x95\xbf\x4e\x2a\xcc\xc2\x30\x0c\x7a\xfa\xea\xd4\xfb\x59\xa9\x6f\x8a\x3c\x83\x9e\x78\xff\x8f\xa2\x6d\xcb\xf6\x65\x2a\xb7\xbf\xdb\xeb\x32\x8e\x9a\xbf\x6f\xe3\xc8\xff\x77\xe0\x7f\x01\x00\x00\xff\xff\x2c\x05\x6a\xf9\x24\x10\x00\x00")

func publicIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_publicIndexHtml,
		"public/index.html",
	)
}

func publicIndexHtml() (*asset, error) {
	bytes, err := publicIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "public/index.html", size: 4132, mode: os.FileMode(420), modTime: time.Unix(1545906047, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"public/index.html": publicIndexHtml,
}

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
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
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

var _bintree = &bintree{nil, map[string]*bintree{
	"public": &bintree{nil, map[string]*bintree{
		"index.html": &bintree{publicIndexHtml, map[string]*bintree{}},
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
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
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

func assetFS() *assetfs.AssetFS {
	assetInfo := func(path string) (os.FileInfo, error) {
		return os.Stat(path)
	}
	for k := range _bintree.Children {
		return &assetfs.AssetFS{Asset: Asset, AssetDir: AssetDir, AssetInfo: assetInfo, Prefix: k}
	}
	panic("unreachable")
}