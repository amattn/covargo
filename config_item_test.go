package covargo

import (
	"flag"
	"io/ioutil"
	"os"
	"testing"
)

const (
	TESTING_ONLY_ENV_VAR = "COVARGO_TESTING_ONLY_ENV_VAR"
)

func TestLongFlag(t *testing.T) {
	key := "long test"
	expected := "TestLongFlag-1267897285"

	item := NewItem(key)

	item.SetCliValueFlags("", "somecoolflag", "-somecoolflag will do something")

	flag.CommandLine.Parse([]string{
		"-somecoolflag",
		expected,
	})

	check_item(t, 2299944751, item, expected)
}

func TestShortFlag(t *testing.T) {
	key := "short test"
	expected := "TestShortFlag-3162324001"

	item := NewItem(key)

	item.SetCliValueFlags("z", "", "-z will do something")

	flag.CommandLine.Parse([]string{
		"-z",
		expected,
	})

	check_item(t, 2600337503, item, expected)
}

func TestEnvVar(t *testing.T) {

	expected := "TestEnvVar-2052242827"

	os.Setenv(TESTING_ONLY_ENV_VAR, expected)

	item := NewItem(TESTING_ONLY_ENV_VAR)

	item.SetEnvVar(TESTING_ONLY_ENV_VAR)

	check_item(t, 2600337501, item, expected)
}

func TestFileContentsShortFlag(t *testing.T) {
	key := "TestFileContentsShortFlagKey"
	expected := "TestFileContentsShortFlag-3520918599"
	path := "/tmp/covargo4115216334.tmp.txt"
	ioutil.WriteFile(path, []byte(expected), 0666)

	item := NewItem(key)

	item.SetFileContentsFlags("fcsf", "", "", "file path to read contents to use as config variable (TestFileContents)")

	flag.CommandLine.Parse([]string{
		"-fcsf",
		path,
	})

	check_item(t, 4115216334, item, expected)
}

func TestFileContentsLongFlag(t *testing.T) {
	key := "TestFileContentsLongFlagKey"
	expected := "TestFileContentsLongFlag-1149492492"
	path := "/tmp/covargo2205809716.tmp.txt"
	ioutil.WriteFile(path, []byte(expected), 0666)

	item := NewItem(key)

	item.SetFileContentsFlags("", "file_contents_long_flag", "", "file path to read contents to use as config variable (TestFileContents)")

	flag.CommandLine.Parse([]string{
		"-file_contents_long_flag",
		path,
	})

	check_item(t, 7263672532, item, expected)
}

func TestFileContentsDefaultPath(t *testing.T) {
	key := "TestFileContentsDefaultPath-Key"
	expected := "TestFileContentsDefaultPath-3376226663"
	path := "/tmp/covargo2358454801.tmp.txt"
	ioutil.WriteFile(path, []byte(expected), 0666)

	item := NewItem(key)

	item.SetFileContentsFlags("", "", path, "file path to read contents to use as config variable (TestFileContents)")

	check_item(t, 4820722189, item, expected)
}

func TestUnparsedFlags(t *testing.T) {
	// key := "TestUnparsedFlags"
	// item := NewItem(key)
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

func check_item(t *testing.T, debug_num int64, item *Item, expected string) {
	item.LoadValue()

	candidate := item.StringValue()
	if candidate != expected {
		t.Error(debug_num, "expected", expected, "got", candidate)
	}
}
