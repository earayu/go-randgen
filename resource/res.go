package resource

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

var _resource_default_zz_lua = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x90\x41\x8b\xfa\x30\x10\xc5\xef\x7e\x8a\xb9\xc5\x3f\x44\x49\x55\xfe\x58\x96\x3d\x88\x08\xde\x16\x54\xd8\x73\x6a\xd3\x3a\x30\x4d\x25\x9d\xec\x22\x8b\xdf\x7d\xc9\xb4\x5a\x37\xa7\x97\x97\xcc\xcc\xef\x0d\xdb\x82\x5c\x07\xef\xf0\x33\x01\x00\x08\xed\xb7\x5c\x32\xa3\x61\x61\x34\x2c\x8d\x86\xdc\xdc\xb5\x3c\xce\x66\x70\xdc\x7f\x7c\xc2\x76\xbf\x39\x6c\xb6\xa7\xdd\x01\x8e\xbb\xd3\x9b\x3c\x9d\x2f\x36\x74\x8e\xa5\x56\x45\xae\xd6\x4a\x83\x22\xcb\xe8\xb3\xa4\x0a\xf4\x36\xdc\xd4\xd8\xe7\x6a\x03\x23\x63\xeb\xc1\xc7\xa6\x70\x41\xfc\xa7\x29\x6d\x56\x1a\xfe\x6b\x50\xd1\x97\xae\x4a\x95\xf7\xc9\xa4\x42\x47\xe5\x08\xcb\xb7\x6b\x8f\xae\x0a\xac\xd1\x73\x9a\x54\x51\x6b\x45\x94\x6d\x2c\xc8\x89\x72\x67\x6c\x2c\x4d\x57\x92\xe9\x9f\xea\x21\xd2\x51\x09\x7b\x2a\x1e\xa8\x2f\x1b\x5e\xaf\xce\xc7\xe6\x01\xdc\x61\xed\x65\x4e\x12\xae\x54\x82\x35\xe8\x7b\x02\x2b\x2d\xdb\x27\x56\x9f\xa8\x07\xf3\x91\x28\x7d\x67\xf4\xb7\x81\xb0\x6b\x2c\x91\xe8\x91\x23\x5b\xcc\xf3\x5c\x36\x95\xcd\x8d\xc9\x93\x98\xe5\xf3\x6c\xbd\x7c\xfd\x34\xe4\x18\xac\x07\x1a\x07\xf4\xf5\xdf\x61\xe4\x98\x5d\xe8\x43\xd4\x84\xdd\x45\xd6\xf7\x1b\x00\x00\xff\xff\xa2\x1b\x69\x7d\xea\x01\x00\x00")

func resource_default_zz_lua() ([]byte, error) {
	return bindata_read(
		_resource_default_zz_lua,
		"resource/default.zz.lua",
	)
}

var _resource_english_txt = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x1c\x90\xd1\x0d\xe4\x20\x0c\x44\xff\x5f\x23\xa9\x65\xcb\x70\x82\x37\xe6\x02\xf8\x04\xe4\x10\xdd\x9f\xbc\x3f\xcf\x03\x96\x47\xa3\x91\xd3\xdf\x89\x94\x82\xb4\x84\x74\x45\x06\x32\x11\x4e\xb9\x1e\x4e\xbd\xe4\x1d\xca\xa9\xda\x38\x95\xf3\x9d\x5c\xd2\x8e\x1f\xb9\xbc\x2a\x97\xbf\x25\x91\x72\x8a\xdf\x94\x13\xc9\x7f\xca\xf9\x7a\xe7\xdb\xbd\x72\xeb\xe4\xf6\xdc\x6e\x6e\xf7\xc4\xed\xf1\xc4\x24\x61\xf2\x4f\x31\x3d\x06\xa6\x3d\x54\xc7\x74\x63\x8a\xe5\x8a\xe5\x81\xf9\xe2\x73\x94\xc2\xe7\xa8\xe4\x2f\xb9\x91\x07\x79\x1e\x01\x32\x7f\xde\x31\x79\x9a\x2f\x4a\x7e\x94\xe2\xfe\x50\x55\x1a\x55\xa9\x9b\xe6\x93\x58\x36\xc7\xbf\xb8\xe1\x8f\x6c\xfc\xc1\x9b\xe2\x0d\xef\x44\x01\x5d\xa5\x94\x4d\xcf\xb7\x4d\x86\x6c\x86\x2a\xc3\x94\xe1\x55\xa7\x45\xf4\x50\x0c\x67\x6a\x29\x4c\x93\x48\x10\x83\x69\xda\x02\x5d\x83\x3b\x40\x9c\x3c\xc1\xc1\xcc\x55\x99\xce\xfb\x97\x25\x6d\xb2\x64\xb0\xc2\x63\xc5\xc5\x52\x56\x98\xac\x30\x59\xe6\x2c\xdb\xac\x1c\xeb\x3c\x8d\xf5\x2b\x77\xab\x18\x5b\x07\xdb\xdf\xa3\x6b\x8c\x1e\xf8\x1f\x00\x00\xff\xff\x65\xfb\xee\xbe\xbd\x01\x00\x00")

func resource_english_txt() ([]byte, error) {
	return bindata_read(
		_resource_english_txt,
		"resource/english.txt",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"resource/default.zz.lua": resource_default_zz_lua,
	"resource/english.txt": resource_english_txt,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"resource": &_bintree_t{nil, map[string]*_bintree_t{
		"default.zz.lua": &_bintree_t{resource_default_zz_lua, map[string]*_bintree_t{
		}},
		"english.txt": &_bintree_t{resource_english_txt, map[string]*_bintree_t{
		}},
	}},
}}
