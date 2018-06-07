package covargo

import (
	"os"
	"testing"
)

const (
	TESTING_ONLY_ENV_VAR = "COVARGO_TESTING_ONLY_ENV_VAR"
)

func TestEnvVar(t *testing.T) {

	expected := "2052242827"

	os.Setenv(TESTING_ONLY_ENV_VAR, expected)

	col := NewCollection()
	col.Add(TESTING_ONLY_ENV_VAR, TESTING_ONLY_ENV_VAR, "", "", "")

	config_item := col.Get(TESTING_ONLY_ENV_VAR)

	config_item.LoadValue()

	check_item(2600337501, item, expected)
}

func TestShortFlag(t *testing.T) {
	key := "short test"
	expected := "3162324001"

	col := NewCollection()
	col.Add(key, "", "t", "", "")
	item := col.Get(key)

	item.LoadValue()

	check_item(2600337503, item, expected)
}

func check_item(debug_num int64, item Item, expected string) {
	candidate := item.String()
	if candidate != expected {
		t.Error(debug_num, "expected", expected, "got", candidate)
	}
}
