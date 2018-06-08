package covargo

import (
	"os"
	"testing"
)

func TestCollection(t *testing.T) {
	ev := "TestCollection_ENV_VAR"
	key := "TestCollection-key"

	expected := "hello-1017400967"
	os.Setenv(ev, expected)

	col := NewCollection()
	item := col.Add(key)
	item.SetEnvVar(ev)

	check_item(t, 1870534936, col.Get(key), expected)
}
