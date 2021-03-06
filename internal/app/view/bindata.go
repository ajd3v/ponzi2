// Code generated by "esc -o bindata.go -pkg view -include .*(ply|png) -modtime 1337 -private data"; DO NOT EDIT.

package view

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// _escFS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func _escFS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// _escDir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func _escDir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// _escFSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func _escFSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// _escFSMustByte is the same as _escFSByte, but panics if name is not present.
func _escFSMustByte(useLocal bool, name string) []byte {
	b, err := _escFSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// _escFSString is the string version of _escFSByte.
func _escFSString(useLocal bool, name string) (string, error) {
	b, err := _escFSByte(useLocal, name)
	return string(b), err
}

// _escFSMustString is the string version of _escFSMustByte.
func _escFSMustString(useLocal bool, name string) string {
	return string(_escFSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/data/addbutton.png": {
		local:   "data/addbutton.png",
		size:    170,
		modtime: 1337,
		compressed: `
H4sIAAAAAAAC/+oM8HPn5ZLiYmBg4PX0cAliYGBwAGEONgYGhlWZhfcYGBgKPV0cQypuvb15kfOAAo+D
4+N/1sd5WRwn9d+JXvBtxYKXUZNUPCepiV3lP3Nmj57GohLf6l/iDKjgDLPdnee3418Uzkvghog8b9j1
R3wDK6qyBhXPSSpbSxg6eHn8Iplv5oPEPF39XNY5JTQBAgAA//+yDEtFqgAAAA==
`,
	},

	"/data/erroricon.png": {
		local:   "data/erroricon.png",
		size:    605,
		modtime: 1337,
		compressed: `
H4sIAAAAAAAC/+oM8HPn5ZLiYmBg4PX0cAliYGBwAGEONgYGhlWZhfcYGJhUPF0cQypuvZ2+14/JUaT1
/b2IiEmmZlPYnFJ1MnJ8xOY7yjer3ziUW8EkFOfEPd9BXdEwIEBAW1pbVSvwI3v9AgluEyUjl4LvZzS+
rzY709ZYfW5vepuHAIsjGDH2rBNjjtVlM6s/3HSVK0bm7R77lT9D9jtFhXvcLmG0XvDv+vNKjYZ1YsxZ
5nah+eZrzjJtW7cvVXxD6Vfd3Q1zA+4v4r7B+VV3V8PcwPuTuG9IftXd2TA36H5jseOn5fuqXkjfX91u
96vkRMHX3/6Sj6v/Hte5xcrz/u8/52/urAxnmdq+170/96bhXi2bLbfEKjvBi6d7dzIszmYQNwipVVr0
XLuY2yY+7WLngYCvnR270p8xlZw3flHM6HxVXalq79fjrgvXtdvtvfcLLDRfpPUwJ0MJhDWZa8ICbol1
+zS53JUFnjC1rc7vOuT5cQZU7LirwgFWnvD7Iq2PlQQeMOmHfe0/3T7h9NWsjQ9yp0yN5nO49zD5TvfC
vfLihueY/hW/Or5nefrS78E71ZY2d4hN52MTrNjrqSsa+DyF15Zb7IHXkiLfqJtce74Gpb+4XBvLLTr3
3cqVXQY3oiOzG2dW2cu6KNfXJRz9GtR+ybFlHivP+8P8t/13T7ugk7DiF58ak/PV/xIVm9eoghx5P2ji
utaV+fcjuYOvBAneYjU+uuDeWfl/p/4VHn1gcLryaa45V6c4Q9Uz9juCEjfz61/2KXA0gdGh/894z80t
z/X5vV6BgYGBwdPVz2WdU0ITIAAA//+HYlJ7XQIAAA==
`,
	},

	"/data/icon.png": {
		local:   "data/icon.png",
		size:    477,
		modtime: 1337,
		compressed: `
H4sIAAAAAAAC/wDdASL+iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAYAAABXAvmHAAABpElEQVRo3u1Z
0bGCMBDcu6GEVwoFQBV24IcWQRW+DzqgCijAnvBDeA8haALZkMx4M4woN+Gyu7eECBDj3KM49yiY98jA
jWr47Fg3UCb6wPNgsqAB0J+fxz+BCfpj0FjQAOhTWdAA6FNZ0EDo01jQQOjTWFAi+lcAP8NxZbEgntFv
x+Jrwe/s+gXAbfha1uLn4aYk9BvD9YbBggbSPq0XlIA+AJwMOSeGI4ln7WPWxM2k+JshZ3cv+JhA6yif
aXS1oDxMQhu0770X1LP2jxrjG98I60L3Xc6z6kjI3R1JNxRfEIoHgGIYm+5CTNeowkjon43e+Hu+Mq5r
vkVkNO2vFfs536kXNALt7+oFpehzO/rO99LI0HdmQVnuEMqRZJckfLnNDnfKLIpvyXonvlKG177ntRBn
zeN1jaQJoP/WkTSRN6XKTUL32U7DUW7zml8iX+5gZNYzZrvN5/wKhv/aNCHnMfaCROo81o6kifn+ggWN
2HmsHEle0M9X9iljWwtNas0WXX6829jkd6MjaSprnrVekASc560jyeKpm1aUkij6fyw8AKZHneMZVCJr
AAAAAElFTkSuQmCCAQAA//8otrd/3QEAAA==
`,
	},

	"/data/refreshbutton.png": {
		local:   "data/refreshbutton.png",
		size:    687,
		modtime: 1337,
		compressed: `
H4sIAAAAAAAC/wCvAlD9iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAACdklEQVR42u2a
v28TMRTH7dIiBmAKkdqKqVI7MmRshuwM/R+y9I8oA1OljAwnZiT2kiUbqANdqDowhgllqMQUVWz80Ieh
rnQKd76z43Ody/tIXfLS5/e+efbZz6eUIAiCIAhCJIA3gF5nAQDeAZvrLADAB+DROgsA8Al4ui6J94FX
/M8UeNbWpPeAETDDzgx43qbEu0CGGz+A/RJ/L4HuqiQ/BOb4MQdeLPh7nbMNU08+Yzm+A7s5f4OC72Sp
Jn+2ZPJToFOR/B1nbUv+a/5xWJF8WiI4lv2Xgs8+Ao9z/noO/rIUFrw6TIBBwUZoXLQbNBUwqel7eF/J
d2uu9sclO8H3VecB4Ljmk6ObYunf3P3qBQK8BTZqjjMwvtKZCmaHV8Wg5H9PPcarsyjuxRRg5FL2gcas
mg6jmALY9vaTBse1LYyzmKc659IPNHbVVOi7+tzwiMOW4KXW+rwpAYzvS8/YggnQs9jGEYpw7BlbMAH2
LbaLCAJceMYWTIAdi20aQYCpZ2zF08pjIfqjlHpQYn6otf7d8CK8pZT6VWL+q7XebLoCWoWPAD8ttk6E
mDuesQUT4NpiO4ggwIFnbMEE+GaxHUYQ4NAztmACXFlsRxEEOPKMTbbCchiS43ACDZGGSj9uQ2SZlphn
8jfJdYd9m6KBy/7+mqImQOe2uMOvnnZb3GEqLF6MnJgEt4Et87dtPjspuTxJ82IkJ8KyV2M+tO5+cHWT
95wOvqR5Pb6wMM4bSDz9FyQWHpEhqyFbmVdkCnaMdV6SKntxatT0Dk9HFKOvbvv2PXXbvd1RSj3JdXKu
zXn+Sil1rrX+rARBEARBEARBEBriH/cd6xTO/96gAAAAAElFTkSuQmCCAQAA///Ye4OhrwIAAA==
`,
	},

	"/data/removebutton.png": {
		local:   "data/removebutton.png",
		size:    530,
		modtime: 1337,
		compressed: `
H4sIAAAAAAAC/wASAu39iVBORw0KGgoAAAANSUhEUgAAAEAAAABACAYAAACqaXHeAAAB2UlEQVR42u2a
TW7CMBCFn4Et56DqshIIJNbsuECl3qALhDhlq96g63IGiOB10SBFCJDtOGNPOt86wvM+Ow7+AQzDMAzD
MAzDiIXkKEObwxS/M0hQyBjAO8mFYPgVgNe67aw9Pya55R9fJNcCba7rtn5I7rJJqMPv6kIoIaER/kIe
CXfCdyrhRvg8Eq6G/T2SSngQvilhKzLbk9zQjyQSPMI3Jbyl+jo8KmjhWRBJHknOWrQ1JXkIEL6Seg3W
ARKiRoJEG8VKKD58ZKFer0PEsM8TvoveUtPzXRSuNnyKAOrDtwnSm/ANCbN6wvOhCnj2QHKqZX8gpFf7
0fMdStAXPqEEveETSNAf/krCd0D4I8mJRG2DQp258/ns+tL7If/tky2ltYfXLyFBeL0SOvgjpEdCxB5B
1RsJgeEPJCen0+kpYD1QroSI8NPIBVR5EgInvJurOrUSUoRXKyFleHUS2rzz6iX864MRkvOAb3erbazA
kRB1NBazGpwD8LkScwSwdM59Ri8JnfsAsARQeTz+AuBZYgRc7gbsCzoe34scj3tKkL4gsc99S0Rs6/rO
nLBFThq3RSqJfftaQlVE+EZRG5JzwfYWJDclrQWGGdocwTAMwzAMwzCMWH4BvSTd+ai+aPgAAAAASUVO
RK5CYIIBAAD//+DMCYgSAgAA
`,
	},

	"/data/roundedcorner_edges.ply": {
		local:   "data/roundedcorner_edges.ply",
		size:    653,
		modtime: 1337,
		compressed: `
H4sIAAAAAAAC/5yQXW7DIBCE3znFXsCIPzftaSJqJoklG6wNruKevkrUUgu3Dw4S0uhjZ9jdaVjEKfHo
M/lr1/ekpRJdGkfETJzmGBC6xBF8RDjjKu8GDHi8f4AzbtSKidMEzgudhuQz3Wqw1OCzBnHjiRtTXLnm
7uKZGKFGZwZiDd+HGaXr+xzkfkv6Mon+CxqBGI4X+AAWWqrHoSLU/6LUmLb9uUJJY92LtaTkqzsc2rfd
CY2STjujHa3E7ojvz5tVPzszCm2eXwZpockIQ1ZYcl8BAAD//zmz7guNAgAA
`,
	},

	"/data/roundedcorner_faces.ply": {
		local:   "data/roundedcorner_faces.ply",
		size:    1070,
		modtime: 1337,
		compressed: `
H4sIAAAAAAAC/6SRQY+bMBCF7/4Vc9tWKpbHNrbpsf0hEcGTXSRiR8Y0S399RXZFsg5dpSpchsebx/fM
aZjZIaZjm6Edu74H5IJ18XikkOFnojaTh/0MPwYKnhJIbh18Gac9iK9Qwfl85vu3Vzym528wxil1BId+
oO/wlOIUPPkupkBpd2g7Gt/sT4wGunzjF6VMr+DYKcUTpTzDYYhthtdSmEvhdymEu51wtxTutsZSyFdh
6l7aBIl8KT0nolCK+2GitddSFvTVMvRjfvdN/dp71wffdzQyCn73Qq2nxCrBNWqJGm4Gcbk2BrwqaOqm
RhBcSo1OgVhuthoq/HvKVpwVxjmzDKZGdO9xgkuljVIguNPW1s2DaUIIlAuctlIZU8J9wlZtw1mUdjHV
TtXqH+A245qmcba5wXz86D7P+3B21YpV3ZA+FKhRo9Jl3+r/f+6HwksugmQKFGioL88GkCmwgGDYnwAA
AP//PqGt/C4EAAA=
`,
	},

	"/data/squareplane.ply": {
		local:   "data/squareplane.ply",
		size:    625,
		modtime: 1337,
		compressed: `
H4sIAAAAAAAC/6SPQW7rIBiE95xidnlPaiycVGrVZXuQCMM4QSLg/kAT9/RVk8qp7F0DC9DHfAwMYVR9
kqMpMNl6j7bRyqbjkbHgTWgKHboRr4HRUbBpnp7xL9cO+j/WOJ1OTXc9apLsH5BTFUv0PvAFq/xejXAI
JvIaWykGXu7+oBSe8agGSQOljOhDMgXnORjn4HMO4sKJCyn+sqo9GIHQzdFeyDiHXaicXt0bS2xukeBz
+clVP/1q56Pzllkxut2BxlFU2+jLwHra6QnpBbqFvqe6SffZfyq/q3sLjRaby7pFq74CAAD//6eCpu5x
AgAA
`,
	},

	"/data/texturedsquareplane.ply": {
		local:   "data/texturedsquareplane.ply",
		size:    651,
		modtime: 1337,
		compressed: `
H4sIAAAAAAAC/5yPQU+EMBCF7/0V77aaCAHWRPGoP2RT2sFtUlqctgL+ehMwuytostle+vJm3jczvZ1E
67mTETIoY1DmhVC+68hFvDHJSBrNhFdLThOjyp+ecRdSg+IeGYZhyJullHt+f0DwiRWhNZZesIs0xsSk
w0eSTL2Vjpb2nSBL84xP4kgjHkXPvieOE1rrZcS4Nqa18bU23CbjNiG3SYW1EU+7tVIRqnPdmhCR1FEy
kjntfjBOG0VBkNOHI0lNLLIyL+aHsyr+F797ygshbgbVdV3/CbqVs4iL07ag7LrTfkh7FChRzX+F/XcA
AAD//zAgyZiLAgAA
`,
	},

	"/": {
		isDir: true,
		local: "",
	},

	"/data": {
		isDir: true,
		local: "data",
	},
}
