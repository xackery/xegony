// Code generaTed by fileb0x at "2017-12-30 10:25:32.497894 -0600 CST m=+0.004404401" from config file "box.yml" DO NOT EDIT.

package box

import (
	"log"
	"os"
)

// FileWwwREADMEMd is "/www/README.md"
var FileWwwREADMEMd = []byte("\x52\x65\x61\x64\x6d\x65\x2e")

func init() {

	f, err := FS.OpenFile(CTX, "/www/README.md", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write(FileWwwREADMEMd)
	if err != nil {
		log.Fatal(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatal(err)
	}
}