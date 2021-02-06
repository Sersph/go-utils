package doc

import (
	"fmt"
	_ "github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func readDir(path string) {
	dir, _ := ioutil.ReadDir(path)

	for _, info := range dir {
		fmt.Println(info.Name())
	}

}

func getAllFileDic(path string) (result map[string]os.File) {
	dir, _ := ioutil.ReadDir(path)
	dir2 := make([]os.FileInfo, 0)
	for _, info := range dir {
		if strings.Index(info.Name(), ".") != 0 {
			dir2 = append(dir2, info)
		}
	}
	for _, info := range dir2 {
		fmt.Println(info.Name())
	}
	return nil
}

func TestIdUtils(t *testing.T) {
	//getAllFileDic("/Users/winily/Projects/Open-Source/go-utils/")
}
