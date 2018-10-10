package config

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestWriteRead(t *testing.T) {
	// First day
	dir := os.TempDir()
	path := dir + "/" + uuid.New().String()

	conf := Config{Lang: []string{"go", "javascript"}}
	fmt.Println(path)
	Write(path, conf)

	if !reflect.DeepEqual(Read(path), conf) {
		t.Fatal("Not equal wrriten conf and read conf.")
	}

}
