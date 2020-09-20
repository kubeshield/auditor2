// Package crds Code generated by go-bindata. (@generated) DO NOT EDIT.
// sources:
// auditor.kubeshield.to_dashboards.v1.yaml
// auditor.kubeshield.to_dashboards.yaml
package crds

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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// ModTime return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _auditorKubeshieldTo_dashboardsV1Yaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x4b\x8f\xdb\x36\x10\xbe\xfb\x57\x0c\xdc\xc3\x5e\x62\x1b\x41\x8b\xa2\xf0\x6d\xb1\x6d\x83\x45\xdb\x34\xd8\x4d\x73\xa7\xc5\xb1\xc4\x2e\x45\xb2\x9c\x91\x37\x6e\xd1\xff\x5e\x0c\x29\x59\xb2\x2d\xd9\x46\x80\x00\x9d\x1b\x5f\xf3\xf8\xe6\x49\x15\xcc\x27\x8c\x64\xbc\x5b\x83\x0a\x06\x3f\x33\x3a\x59\xd1\xf2\xe5\x07\x5a\x1a\xbf\xda\xbd\x9d\xbd\x18\xa7\xd7\xf0\xd0\x10\xfb\xfa\x09\xc9\x37\xb1\xc0\x1f\x71\x6b\x9c\x61\xe3\xdd\xac\x46\x56\x5a\xb1\x5a\xcf\x00\x8a\x88\x4a\x36\x3f\x9a\x1a\x89\x55\x1d\xd6\xe0\x1a\x6b\x67\x00\x56\x6d\xd0\x92\xdc\x01\x50\x21\x2c\x5f\x9a\x0d\x46\x87\x8c\x49\x8a\x53\x35\xae\x41\xf6\xa8\x32\x68\xf5\x0c\x20\x6f\x69\x45\xd5\xc6\xab\xa8\x69\xa9\x1a\x6d\xd8\xc7\x65\x7f\x6b\xc9\x7e\x46\x01\x0b\xe1\x5a\x46\xdf\x84\x35\x8c\x5f\xca\xdc\x5a\xe9\x85\x62\x2c\x7d\x34\xdd\x7a\xd1\x3d\x6a\x57\x47\x5a\xa4\xe3\x10\xa8\xf0\x1a\xd3\x32\x83\x71\x2f\x2f\x9e\xb0\x34\xc4\x31\x19\x9c\xce\xac\x21\xfe\x65\xfc\xfc\x57\x43\x9c\xee\x04\xdb\x44\x65\x87\x86\xa5\x6d\x32\xae\x6c\xac\x8a\x83\x83\x19\x00\x15\x3e\xe0\x1a\xde\x8b\xf2\x41\x15\x28\x7b\xbb\xec\xae\xa4\xfc\xa2\x45\x69\xf7\x56\xd9\x50\xa9\xb7\x99\x55\x51\x61\xad\xb2\x6d\x00\x3e\xa0\xbb\xff\xf0\xf8\xe9\xdb\xe7\xa3\x6d\x80\x10\x7d\xc0\xc8\x07\x18\x32\x0d\xe2\x61\xb0\x0b\xa0\x91\x8a\x68\x02\xa7\x40\xb9\x13\x86\xf9\x16\x68\x09\x04\x24\xe0\x0a\x3b\xd5\x50\xb7\x3a\x80\xdf\x02\x57\x86\x20\x62\x88\x48\xe8\xb8\xc7\xaa\x27\xbf\x05\xe5\xc0\x6f\xfe\xc4\x82\x97\xf0\x8c\x51\xd8\x00\x55\xbe\xb1\x1a\x0a\xef\x76\x18\x19\x22\x16\xbe\x74\xe6\xef\x03\x6f\x02\xf6\x49\xa8\x55\x8c\x2d\xb6\x3d\x19\xc7\x18\x9d\xb2\xb0\x53\xb6\xc1\x37\xa0\x9c\x86\x5a\xed\x21\xa2\x48\x81\xc6\x0d\xf8\xa5\x2b\xb4\x84\xdf\x7c\x44\x30\x6e\xeb\xd7\x50\x31\x07\x5a\xaf\x56\xa5\xe1\x2e\x0f\x0a\x5f\xd7\x8d\x33\xbc\x5f\x15\xde\x71\x34\x9b\x86\x7d\xa4\x95\xc6\x1d\xda\x15\x99\x72\xa1\x62\x51\x19\xc6\x82\x9b\x88\x2b\x15\xcc\x22\xa9\xee\x38\x25\x53\xad\xbf\x89\x6d\xe6\xd0\xdd\x91\xae\xbc\x17\x0f\x13\x47\xe3\xca\xc1\x41\x0a\xb3\x0b\x1e\x90\x30\x03\x43\xa0\xda\xa7\xd9\x8a\x1e\x68\xd9\x12\x74\x9e\x7e\x7a\xfe\x08\x9d\xe8\xe4\x8c\x53\xf4\x13\xee\xfd\x43\xea\x5d\x20\x80\x19\xb7\xc5\x98\x9d\xb8\x8d\xbe\x4e\x3c\xd1\xe9\xe0\x8d\xe3\xb4\x28\xac\x41\x77\x0a\x3f\x35\x9b\xda\xb0\xf8\xfd\xaf\x06\x89\xc5\x57\x4b\x78\x50\xce\x79\x86\x0d\x42\x13\xb4\x62\xd4\x4b\x78\x74\xf0\xa0\x6a\xb4\x0f\x8a\xf0\xab\x3b\x40\x90\xa6\x85\x00\x7b\x9b\x0b\x86\x75\xed\xf4\x72\x46\x6d\x70\xd0\x95\xa1\x9e\xc6\xf3\x4b\x28\x36\xf6\x7c\x13\xc0\x30\xd6\x23\xdb\x97\x38\x65\x72\x5d\x79\x98\x38\x3f\x89\x9d\x43\x35\x91\xa4\x55\xdc\x66\x68\x63\x11\x6a\xc5\x45\x25\xa9\xf0\x51\xbc\x5c\x07\xde\x9f\xa3\x72\x4c\xf3\x79\xf7\x08\x9c\x77\x8b\x83\x26\xfa\x10\x73\xb4\x84\x7b\xd7\x32\x93\xfa\x08\xa6\x0e\xd6\xe0\x69\x18\xf6\x84\x3b\x8c\xfb\xde\xa6\xe5\xc4\xc5\x49\xb4\x32\x4d\xf8\xf4\xfc\x8a\x8a\x51\xed\x47\x6f\x08\xe8\xa9\x5e\xdd\x06\xeb\xef\xdd\x75\x49\x4b\xc9\x8c\xc3\x7b\xd8\xa0\xa4\x63\xc0\xb8\xf5\xb1\x46\xfd\x65\x16\x5d\x16\x26\xc6\x48\x25\x3d\xa4\xfa\x89\xf4\x49\xb6\x00\x45\x85\xc5\x0b\x6a\xd8\xfa\x08\x4a\xd7\x86\x52\x5d\x4f\x49\xe6\xed\xd7\x04\xf8\x10\x22\x37\xe1\xdb\x4d\x1e\xd3\x51\x7b\x14\x67\x93\x8a\xb7\xf1\x07\xca\xda\x54\x65\x09\x8c\x4b\x8b\xfb\x0f\x8f\x79\x8a\xa0\x2f\x8c\xb9\x23\x75\xdf\x09\xa7\x5e\xe7\xbe\xbe\xf6\x2e\xea\xa5\x9f\xb6\xc4\x21\x1d\xd4\x9a\xd2\xea\x7a\x81\xc8\x94\x27\xa4\x0b\x17\xc6\x2c\xe8\xe2\x4b\xf2\x31\xb7\x73\xec\x35\xca\x9e\x90\x48\x51\xc6\x4d\x67\x74\x26\x79\x39\x28\x0a\xa7\x25\x66\x88\x50\xea\x2b\x3e\xe2\x15\x8e\x37\x20\x03\xb7\x45\x2a\x0c\xa2\xf1\x7d\x3f\x29\xde\x04\xd3\xd3\xf0\x5d\x6e\xc9\xa9\xce\x0d\x93\xd1\x38\x62\xe5\x8a\x6b\x06\xa5\xa2\xd7\x85\x37\x42\xf0\xd6\x14\xfb\x3e\xbe\xff\xa0\xdc\xd5\xa5\x17\xcb\x78\x9a\xda\xab\x89\x17\x6a\x69\xa6\x41\xe6\x78\x69\xbf\xd2\xab\xcc\xd6\x48\x03\x1e\x2d\xcd\x49\x83\x2b\x3c\x73\x95\xee\xec\xea\x22\xa3\x37\x98\x5a\xb5\xf5\x65\xdf\x5c\xc9\xa8\x4c\x37\x3a\xf0\x7a\xc1\xc9\x74\xb5\xec\x64\x9a\x28\x3e\xa3\x0e\x26\x78\x35\x5c\x19\x27\x20\xdc\x1c\xb4\xf0\xdc\x6c\x7a\x06\x2a\x76\x95\x4c\x43\x93\x1c\xad\x60\xbe\x9a\x8b\xcb\x8c\xd3\x46\x3e\x2d\x37\x24\x18\xf5\x2c\x97\xf0\xb3\x8f\x80\x9f\x55\x1d\x2c\xbe\x81\x79\xf0\x9a\x56\xd6\x97\x73\x78\x4d\x93\x75\x12\x76\x85\x65\x3b\xbf\x1d\x46\x6d\x5f\x0e\x25\x08\x02\xc2\x35\x27\x33\xfb\x00\x56\x86\xb1\xab\x3c\xdb\xd7\xda\xa7\xa9\x81\xb3\x26\x43\xc6\xd4\xaa\x3b\x4f\x77\xdc\xdd\xb5\x60\xcc\x0c\x7a\x0b\xff\x6f\x31\x37\x31\x33\xde\xce\xa6\x4d\x74\x3d\xa6\xf5\x62\x30\xa8\x8c\x1c\x5f\x94\x3d\x25\x75\x5c\xde\x22\x4f\xae\x37\x8d\xc3\xac\xb8\xa1\x5b\x07\x62\xbf\x21\xf9\x72\xe8\x77\xe8\x5a\x4b\xce\x0d\x3d\x1e\x7f\xce\x1e\x74\x7d\xaa\xf6\x94\x3e\x8b\xe8\x18\xca\xfe\xb4\x93\x30\x02\x90\x8c\x3d\xed\x27\xb5\x4b\x9c\x47\x69\x6a\x31\x22\x05\x2f\x1d\xba\x8d\xfe\xee\xfc\x8e\x06\x9c\xdf\x8c\x70\x7c\xad\x4c\x51\x89\x42\xed\x3f\x07\xbc\x83\xba\xe1\x76\x1a\xdb\x1f\x7a\x68\xfe\x67\x9d\x07\xab\x4c\x89\x8a\xd7\xf2\x89\xfd\xfe\xbb\x09\x9f\xc9\x07\xb7\xc4\x78\xdd\x17\x67\x9b\x19\x88\x35\x70\x6c\x72\x45\x21\xf6\x51\x95\x38\xdc\x19\xa4\x62\xe7\x89\xd6\xa3\xf0\xcf\xbf\xb3\xff\x02\x00\x00\xff\xff\x9d\xd8\x27\xd2\x34\x12\x00\x00")

func auditorKubeshieldTo_dashboardsV1YamlBytes() ([]byte, error) {
	return bindataRead(
		_auditorKubeshieldTo_dashboardsV1Yaml,
		"auditor.kubeshield.to_dashboards.v1.yaml",
	)
}

func auditorKubeshieldTo_dashboardsV1Yaml() (*asset, error) {
	bytes, err := auditorKubeshieldTo_dashboardsV1YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "auditor.kubeshield.to_dashboards.v1.yaml", size: 4660, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _auditorKubeshieldTo_dashboardsYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x58\xdd\x6f\x1c\x35\x10\x7f\xbf\xbf\x62\x14\x90\x02\xa8\x77\x47\x54\x54\xc1\xbe\xa0\x28\xd0\x2a\xa2\x94\x2a\x69\xfb\x52\x15\xc9\xb7\x9e\xbb\x35\xf1\xda\xc6\x33\x7b\x69\x8a\xf8\xdf\xd1\xd8\xfb\x95\xcb\xde\x5d\x40\xaa\x84\x9f\x6e\xc7\xf6\x7c\xfc\xe6\xd3\xa7\x82\x79\x87\x91\x8c\x77\x05\xa8\x60\xf0\x23\xa3\x93\x2f\x5a\xdc\x7c\x4f\x0b\xe3\x97\xdb\xb3\x15\xb2\x3a\x9b\xdd\x18\xa7\x0b\xb8\x68\x88\x7d\x7d\x85\xe4\x9b\x58\xe2\x4f\xb8\x36\xce\xb0\xf1\x6e\x56\x23\x2b\xad\x58\x15\x33\x80\x32\xa2\x12\xe2\x1b\x53\x23\xb1\xaa\x43\x01\xae\xb1\x76\x06\x60\xd5\x0a\x2d\xc9\x19\x00\x15\xc2\xe2\xa6\x59\x61\x74\xc8\x98\x44\x39\x55\x63\x01\x42\xa3\xca\xa0\xd5\x33\x80\x4c\xd2\x8a\xaa\x95\x57\x51\xd3\x42\x35\xda\xb0\x8f\x8b\xe1\xd4\x82\xfd\x8c\x02\x96\xc2\x75\x13\x7d\x13\x0a\x98\x3e\x94\xb9\xb5\xd2\x4b\xc5\xb8\xf1\xd1\x74\xdf\xf3\xee\x52\xfb\x75\x4f\x8b\xb4\x1d\x02\x95\x5e\x63\xfa\xcc\x60\x9c\xcb\x8d\x2b\xdc\x18\xe2\x98\x0c\x4e\x7b\xd6\x10\xff\x32\xbd\xff\xd2\x10\xa7\x33\xc1\x36\x51\xd9\xb1\x61\x89\x4c\xc6\x6d\x1a\xab\xe2\x68\x63\x06\x10\x22\x12\xc6\x2d\xbe\x75\x37\xce\xdf\xba\xe7\xa2\x14\x15\xb0\x56\x96\x44\x1b\x2a\x7d\xc0\x02\x5e\x89\x6d\x41\x95\x28\x57\xa8\x59\xc5\xd6\x47\xad\x7d\xc4\x8a\x1b\x2a\xe0\xaf\xbf\x67\x00\x5b\x65\x8d\x4e\x1a\xe5\x4d\x1f\xd0\x9d\xbf\xbe\x7c\xf7\xf4\xba\xac\xb0\x56\x99\x28\x82\x7d\xc0\xc8\x3d\x46\xd9\x6b\x7d\xbc\xf4\x34\x00\x8d\x54\x46\x13\x12\x47\x38\x15\x56\xf9\x0c\x68\x89\x10\x24\xe0\x0a\x61\x9b\x69\xa8\x81\x92\x18\xf0\x6b\xe0\xca\x10\x44\x4c\x26\x3a\x1e\x40\xec\x96\x5f\x83\x72\xe0\x57\x7f\x60\xc9\x0b\xb8\x16\x18\x22\x01\x55\xbe\xb1\x1a\x4a\xef\xb6\x18\x19\x22\x96\x7e\xe3\xcc\xa7\x9e\x33\x01\xfb\x24\xd2\x2a\xc6\x16\xf2\x6e\x19\xc7\x18\x9d\xb2\x02\x42\x83\x4f\x40\x39\x0d\xb5\xba\x83\x88\x22\x03\x1a\x37\xe2\x96\x8e\xd0\x02\x7e\xf5\x11\xc1\xb8\xb5\x2f\xa0\x62\x0e\x54\x2c\x97\x1b\xc3\x5d\x86\x94\xbe\xae\x1b\x67\xf8\x6e\x59\x7a\xc7\xd1\xac\x1a\xf6\x91\x96\x1a\xb7\x68\x97\x64\x36\x73\x15\xcb\xca\x30\x96\xdc\x44\x5c\xaa\x60\xe6\x49\x71\xc7\x29\xcd\x6a\xfd\x45\xef\xaa\xd3\x91\xa6\x7c\x27\x5e\x25\x8e\xc6\x6d\x7a\x72\x8a\xbb\xbd\xb8\x4b\xd4\x81\x21\x50\xed\xb5\xac\xff\x00\xaf\x90\x04\x95\xab\x9f\xaf\xdf\x40\x27\x34\xb9\xe0\x3e\xe6\x09\xed\xe1\x1a\x0d\xc0\x0b\x50\xc6\xad\x31\x66\xc7\xad\xa3\xaf\x13\x47\x74\x3a\x78\xe3\x38\x7d\x94\xd6\xa0\xbb\x0f\x3a\x35\xab\xda\xb0\x78\xfa\xcf\x06\x89\xc5\x3f\x0b\xb8\x50\xce\x79\x86\x15\x42\x13\xb4\x62\xd4\x0b\xb8\x74\x70\xa1\x6a\xb4\x17\x8a\xf0\xb3\xc3\x2e\x08\xd3\x5c\x20\x3d\x0e\xfc\xb8\xbc\x75\x6b\x2a\x3d\x64\xa5\x9a\x75\x8f\x02\x50\xab\x8f\x2f\xd1\x6d\xb8\x2a\xe0\xd9\xd3\x9d\xbd\xa0\x58\x42\xb2\x80\xdf\xdf\xab\xf9\xa7\x0f\x5f\xbd\x9f\xab\xf9\xa7\x6f\xe7\x3f\x7c\xf8\xe6\x7d\xfb\xe3\xeb\x1f\xbf\xdc\xb9\x33\xa9\x64\x47\xce\x0e\xec\xc9\x5d\x71\x3c\xa6\x77\x6c\xec\x2e\x09\xc0\x30\xd6\x0f\x88\xfb\x79\x0c\x08\xa4\x5a\x34\xb9\xbb\x13\xb6\x7d\xe1\x92\x2a\xa1\xb8\x2d\x09\x8d\x45\xa8\x15\x97\x95\xe4\xdf\x1b\x09\xb1\x3a\xf0\xdd\x43\x8b\xc7\xeb\xe4\xa4\xbb\x02\xce\xbb\x79\xaf\x85\xee\x83\x9d\x16\x70\xee\x5a\x56\x52\xa7\xc1\xd4\xc1\x1a\x24\xc0\x2d\xc6\xbb\x3d\x6c\x7b\x3e\x8b\xc9\x03\x7b\x10\xca\x6b\x8f\x9f\x76\x0f\xa8\x18\xd5\x94\x78\x01\x39\x15\xc5\xc7\x00\xf9\x5b\x77\x58\x6a\x80\x24\x62\x7f\x1b\x56\x28\xb9\x1f\x30\xae\x7d\xac\x51\xff\x7b\x3b\x0e\x0b\x12\x23\xa4\x54\xf7\x35\x65\x47\xf2\x1e\xa6\x00\x65\x85\xe5\x0d\x6a\x58\xfb\x08\x4a\xd7\x86\x52\xd3\x48\xf9\xec\xed\xe7\x81\x74\xa7\x2f\x1e\x34\xb4\x9b\x73\xf6\x47\xe6\x64\x34\xed\x51\x5c\x59\x9b\x0a\x38\x81\x71\xe9\xe3\xfc\xf5\x65\x9e\x58\xe8\x3f\x44\xd6\x3d\x45\x5f\x08\x97\x41\xdb\xa1\x74\x0f\x2e\x19\x24\xbb\xbd\xee\xe8\x15\x9a\xd6\xe7\x58\xda\xe7\x95\x67\xb0\xbd\xdb\x53\x9a\x77\x71\x24\x99\x96\xa7\x02\x1c\x74\xc9\xd8\x4b\x4c\x28\xe3\xf6\x61\x9b\x97\xdc\x1b\xa5\xfa\x6e\xd9\x18\xe3\x92\x1a\x95\xf4\x97\xf3\xd7\x97\x07\x79\x1e\xc4\x03\x1e\x13\x8f\x30\x8a\xba\x57\xc3\x04\xfa\x08\x70\xae\xc6\xb7\x72\x67\x4f\x71\x36\x4e\x35\xe3\x88\x95\x2b\xf1\xa0\x11\xa9\x84\x75\x41\x8c\x10\xbc\x35\xe5\xdd\x10\xc5\x6f\x29\x8f\x06\xd2\xd2\x65\xba\x4c\x7d\xda\xc4\xbd\x91\x9c\xd7\x28\x3b\xbc\x74\x71\xe9\x33\x66\x6d\xa4\x8f\x4f\x17\x59\x91\x7f\x90\x63\xaa\xc2\xbd\x45\x5d\x24\x0c\xa6\x52\xab\xb2\x3e\xe4\x8f\x83\x59\x93\xd7\xa3\x5c\x76\xac\x90\xe4\x75\xa4\x9c\xe4\xb5\xa7\xa8\x4c\x3a\x94\xe0\xd6\x70\x65\x9c\x98\x7e\xd0\x88\x21\x59\xe1\x7a\x34\xee\x83\x8a\x5d\x7d\xd2\xd0\x24\xc7\x2a\x38\x59\x9e\x88\x93\x8c\xd3\x46\x1e\x3e\x47\x93\x68\xf4\x7e\x58\xc0\x73\x1f\x01\x3f\xaa\x3a\x58\x7c\x02\x27\xc1\x6b\x5a\x5a\xbf\x39\x81\xdb\x34\x84\x27\x51\xdd\x60\x77\x98\x71\x3b\x93\xfb\xcd\x98\xbf\xd8\x2e\x3c\x73\xba\xb2\x0f\x60\x65\x82\xeb\xf1\x38\xc8\x52\xfb\xd4\xeb\xb9\xd5\x62\xfc\xec\x69\x55\x3d\x49\x67\xdc\x69\x77\x64\xd0\xff\xff\x12\x43\x93\x33\xdb\x63\x59\xb4\x89\xaa\x1f\xea\x3a\x1f\x0d\x0f\x0f\x36\x0f\xc8\x9c\x96\x36\x25\x67\x9e\x27\xc6\xa3\xe3\x67\x7e\x79\x3e\x62\x00\xf5\xab\xf4\xce\xd5\x2f\xd0\xb5\x9a\xef\x9a\x75\x7f\x0c\x79\x70\xbc\xeb\x23\xb5\xa7\xf4\x2a\x44\xc7\xb0\x19\x76\x3b\xfe\x0f\xe0\x90\xf1\xa3\x7d\x89\x76\x41\x7f\x29\x2d\x27\x46\xa4\xe0\xa5\x6f\xb6\xb1\xdb\xed\x9f\xd2\x88\xef\x13\xb8\xad\x4c\x59\x3d\xe0\x6a\xa8\x7b\xd8\x80\x77\x50\x37\xdc\x4e\x44\x77\x7d\x7f\xcb\xcf\xaa\xdd\x40\x94\x19\x4d\x71\x21\x2f\xd5\x67\xdf\x4d\x7a\x47\xde\xb0\x1b\x8c\x87\x91\xdf\x21\x6d\xbb\x7f\x79\xb6\x67\xca\x86\x4a\x9d\x0d\xb4\xe4\x86\x79\xfb\x5f\xcb\x68\x1b\x20\xe3\x55\x00\xc7\x06\xdb\xbf\x11\x7c\x54\x1b\x6c\x29\xff\x04\x00\x00\xff\xff\xff\x71\x1e\x0c\x3d\x12\x00\x00")

func auditorKubeshieldTo_dashboardsYamlBytes() ([]byte, error) {
	return bindataRead(
		_auditorKubeshieldTo_dashboardsYaml,
		"auditor.kubeshield.to_dashboards.yaml",
	)
}

func auditorKubeshieldTo_dashboardsYaml() (*asset, error) {
	bytes, err := auditorKubeshieldTo_dashboardsYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "auditor.kubeshield.to_dashboards.yaml", size: 4669, mode: os.FileMode(420), modTime: time.Unix(1573722179, 0)}
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
	"auditor.kubeshield.to_dashboards.v1.yaml": auditorKubeshieldTo_dashboardsV1Yaml,
	"auditor.kubeshield.to_dashboards.yaml":    auditorKubeshieldTo_dashboardsYaml,
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
	"auditor.kubeshield.to_dashboards.v1.yaml": {auditorKubeshieldTo_dashboardsV1Yaml, map[string]*bintree{}},
	"auditor.kubeshield.to_dashboards.yaml":    {auditorKubeshieldTo_dashboardsYaml, map[string]*bintree{}},
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
