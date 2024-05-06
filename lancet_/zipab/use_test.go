package zipab

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"log"
	"testing"
)

func TestName(t *testing.T) {
	err := fileutil.Zip("D:\\work\\code\\go\\tests\\v2", "./test.zip")
	if err != nil {
		log.Fatalln(err)
	}
}
