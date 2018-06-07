package covargo

import (
	"flag"
	"os"
	"testing"
)

const (
	TESTING_ONLY_ENV_VAR = "COVARGO_TESTING_ONLY_ENV_VAR"
)

func TestLongFlag(t *testing.T) {
	key := "long test"
	expected := "TestLongFlag-1267897285"

	col := NewCollection()
	item := col.Add(key)

	item.SetCliValueFlags("", "somecoolflag", "-somecoolflag will do something")

	flag.CommandLine.Parse([]string{
		"-somecoolflag",
		expected,
	})

	item.LoadValue()

	check_item(t, 2299944751, item, expected)
}

func TestShortFlag(t *testing.T) {
	key := "short test"
	expected := "TestShortFlag-3162324001"

	col := NewCollection()
	item := col.Add(key)

	item.SetCliValueFlags("z", "", "-z will do something")

	flag.CommandLine.Parse([]string{
		"-z",
		expected,
	})

	item.LoadValue()

	check_item(t, 2600337503, item, expected)
}

func TestEnvVar(t *testing.T) {

	expected := "TestEnvVar-2052242827"

	os.Setenv(TESTING_ONLY_ENV_VAR, expected)

	col := NewCollection()
	item := col.Add(TESTING_ONLY_ENV_VAR)

	item.SetEnvVar(TESTING_ONLY_ENV_VAR)

	item.LoadValue()

	check_item(t, 2600337501, item, expected)
}

func TestUnparsedFlags(t *testing.T) {
	// key := "TestUnparsedFlags"
	// item := MakeItem(key)
	// item.SetCliFlags("t", "", "-t will do something")

	// can't test this as the test package apparently called flag.Parse()
	// would need to manage a custom flagset to make this work... TODO

	// err := item.LoadValue()
	// log.Println("1832377877", err, flag.Parsed())
	// // we expect an error here:

	// if err == nil {
	// 	t.Error("err == nil, but we expect an error here.")
	// }
}

func check_item(t *testing.T, debug_num int64, item Item, expected string) {
	candidate := item.String()
	if candidate != expected {
		t.Error(debug_num, "expected", expected, "got", candidate)
	}
}
