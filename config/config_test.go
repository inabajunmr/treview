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

	conf := config{Lang: []string{"go", "javascript"}}
	fmt.Println(path)
	write(path, conf)

	if !reflect.DeepEqual(read(path), conf) {
		t.Fatal("Not equal wrriten conf and read conf.")
	}

}
