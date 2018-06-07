package covargo

var DefaultCollection Collection

func init() {
	DefaultCollection = NewCollection()
}

// add a config item to our "library" of items
// will panic if ENV_VAR_NAME is empty
// CliFlag and JSONKey can be empty, and will default to ENV_VAR_NAME
// typically, CliFlag is shorter and easy to type and JSONKey is empty or a lowercase, snake case variant of ENV_VAR_NAME
func Add(evn_var_name, shortflag, longflag, json_key string) {
	DefaultCollection.Add(evn_var_name, shortflag, longflag, json_key)
}

func Get(key string) Item {
	return DefaultCollection.Get(key)
}

func Load(key string) error {
	return DefaultCollection.Load(key)
}
func LoadAll() error {
	return DefaultCollection.LoadAll()
}

func Contains(key string) bool {
	return DefaultCollection.Contains(key)
}

func Remove(key string) {
	DefaultCollection.Remove(key)
}

func StringValue(key string) string {
	return DefaultCollection.StringValue(key)
}
