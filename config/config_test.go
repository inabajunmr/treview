package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestWriteRead(t *testing.T) {
	dir := os.TempDir()
	path := dir + "/" + uuid.New().String()

	conf := config{Lang: []string{"go", "javascript"}}
	fmt.Println(path)
	write(path, conf)

	if !reflect.DeepEqual(read(path), conf) {
		t.Fatal("Not equal wrriten conf and read conf.")
	}

}

func TestLangOperate(t *testing.T) {
	dir := os.TempDir()
	path := dir + "/" + uuid.New().String()
	langs := []string{"go", "javascript"}
	SetLangs(path, langs)

	if !reflect.DeepEqual(GetLangs(path), langs) {
		t.Fatal("Not equal wrriten conf and read conf.")
	}
}
