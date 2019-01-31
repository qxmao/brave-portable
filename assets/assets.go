// Code generated by go-bindata. DO NOT EDIT.
// sources:
// res/Brave.lnk (1.847kB)

package assets

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _braveLnk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x94\x4b\x68\x2b\x55\x18\xc7\x7f\x31\x15\xef\xdd\x55\x1a\xe4\x5e\x50\x1a\x2f\xa6\x08\x32\xb9\x79\x35\x24\xf1\x41\x9b\xd7\x4d\xc8\xc4\xa6\x99\x04\x22\x8e\x17\xa7\xc3\x5c\x6e\x24\xcd\xa4\x93\xb1\x69\xeb\xb6\xba\x28\x16\x94\x82\x0a\x2a\x2d\xc5\xa0\x08\x82\x82\x3b\x8b\xb8\x70\x27\x48\x17\x15\x04\x05\x6d\x5d\x28\xb5\xe2\xc2\x85\x0b\x15\x64\x4e\xa6\xe9\xc3\x82\x3b\x45\xf0\x7f\xc8\x79\x7c\xdf\x77\xfe\xdf\xff\x3b\x27\x73\x64\xc0\xe3\xbb\x03\x07\x9f\x88\x9e\xfc\x37\x53\xe0\x07\x3e\xb3\xff\x90\x5f\xee\xec\x79\x76\x8e\x7e\xad\x39\xe3\x4f\xaf\x0f\xd6\xb8\x18\x4e\x5c\x7c\x85\x8f\xf1\xca\xb7\x33\xbb\xfe\xc3\x54\x73\x74\xfb\xcb\x4b\x3c\x14\x0a\xbd\x79\x95\xeb\x99\x94\xca\x5f\x51\x21\x2c\xc6\x9b\xe5\x17\xfa\xa3\xa4\x2d\x6d\xd1\xe0\x11\x2e\x33\xc2\xcf\x1f\xdf\x2c\x2f\xf6\x1d\x7b\x10\xd8\xbe\xe6\x44\xbd\x72\x6e\x77\x75\x6e\xc2\x93\xc6\x42\x63\x11\x03\xf0\xa1\x12\x71\xf9\x56\xfb\x7e\xe6\x1c\xbe\xa0\xb1\x64\x90\x1d\x72\xae\xf6\x9d\x9f\xc3\xf9\xa3\xd7\x11\xbf\x71\x8e\x73\x6c\xea\x7b\xe6\x86\x9c\x41\x0c\x96\x04\xf7\x15\xa0\x08\xdc\xeb\x16\xed\x8c\x93\xee\x9e\x02\x70\x15\xf0\x02\xf1\x1b\xdf\x8d\x8c\x02\xd9\xe6\x03\x0b\xcf\x1a\xfe\x0c\x99\x94\x2a\xea\x52\x4f\xd4\x30\xc6\x34\x3a\x3a\x06\x5d\xba\xf8\xb1\xb9\x8d\x81\x9f\x22\x6d\x6c\x0c\x2c\xda\x18\xd8\x5c\x21\x48\x10\xf5\x4c\x7f\x52\xb0\x7a\x81\xd0\x4b\x64\x48\x9d\x89\x7a\xdc\xb5\xd4\xe9\x0a\xe6\x2e\x2a\xba\xf0\xae\xb0\x8c\xca\x34\x1d\x3a\x64\xd1\xb0\xd1\x50\x91\x31\xd1\xd1\x68\x9d\x61\x51\x30\xb9\x85\x4d\x0f\x0d\x4b\xe4\x3e\xf1\x49\x62\x6e\xd2\x73\x33\x1c\x73\xb6\x68\x0a\x26\x9b\x26\x26\xed\x0b\xf5\xfa\xbc\x70\x17\x6c\x05\xea\x4a\xae\x5a\xa9\xce\xe4\x8b\x72\x2e\xa0\x4e\x77\x3a\x59\xcd\xd6\x54\xd9\xd4\xb5\xd6\xe0\xf8\x14\xf3\x96\xdd\xd3\x2c\x63\xb0\x92\xd2\x96\xd9\xeb\x1a\x96\x13\xd9\x6a\xea\x9a\xdd\x34\xdb\xa7\x0f\xf8\x3f\x89\x00\x75\x14\x72\x54\xa9\x50\x65\x86\x3c\x45\x64\x72\x04\xfe\xf5\x5b\xfa\x1f\xff\x24\x9e\x1e\x3c\x25\x5b\x0d\x77\xad\x5b\xda\xca\xb2\xd4\xd1\x8f\xfd\x6f\x74\xaf\xfd\xf6\xe9\xc1\x53\xa5\x77\x5f\x8d\x4f\x3c\x7c\x58\xd9\x2c\x3c\xf6\x7e\xf2\x99\xaf\x7f\xb8\xfb\xf9\x97\xb2\x6b\x4f\xce\x7f\x78\xf4\x77\xfe\x8c\x07\x2e\xc3\xd6\x3a\x10\x56\x2a\xca\xc1\x5a\x23\xbf\x23\x27\x32\x1f\xfd\x3e\xb6\x31\xf1\xda\xfc\xe7\x0b\xc0\x88\x93\x68\x1c\xb8\x0e\x28\x48\x84\x91\x98\x44\x22\x22\x66\x51\xa2\x84\x88\x13\x25\x41\x98\x08\x31\x22\x48\xc4\x08\x89\x96\x10\x9e\x18\x71\x11\x1b\x26\x46\x82\x08\x09\x42\x44\x44\x4b\x0a\xab\x13\x19\x1e\x16\xbd\xe0\x8a\xa9\x3f\x28\x6f\x2e\x6f\x26\x4b\xef\xec\xee\xef\x49\xfb\x5f\xfc\x52\x07\xee\x3c\x16\x73\x3f\x9c\xfa\x47\x07\x99\x24\x4d\x43\x7c\x33\x11\xe2\xd4\x68\x50\xe7\x09\x4a\x64\x50\x88\x52\xa3\xcc\x0d\x62\xcc\x32\x4b\x9e\x1a\x31\xca\xcc\x0e\x13\x26\xdd\x84\x1f\xdc\x33\x9f\x7d\x6f\xbd\x53\x78\xbb\x30\x15\x7c\xeb\xd1\xa5\x17\xef\x03\x6e\xe3\xbe\xea\xde\xe7\x4a\x85\x41\x78\x65\xfc\xfc\x35\xfd\x19\x00\x00\xff\xff\xe1\x91\xe3\xa6\x37\x07\x00\x00")

func braveLnkBytes() ([]byte, error) {
	return bindataRead(
		_braveLnk,
		"Brave.lnk",
	)
}

func braveLnk() (*asset, error) {
	bytes, err := braveLnkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Brave.lnk", size: 1847, mode: os.FileMode(438), modTime: time.Unix(1547062414, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x31, 0xd4, 0x6d, 0xa9, 0x8b, 0xc9, 0xac, 0x3c, 0xe7, 0x1f, 0x25, 0xba, 0x79, 0x20, 0x8b, 0xf9, 0x53, 0x41, 0xb7, 0x36, 0xe2, 0x8a, 0x33, 0x2, 0xb7, 0x16, 0xfe, 0xa7, 0x74, 0xe8, 0xee, 0xe2}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"Brave.lnk": braveLnk,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"Brave.lnk": &bintree{braveLnk, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
