package covargo

import "github.com/amattn/deeperror"

type Collection map[string]Item

func NewCollection() Collection {
	return Collection{}
}

// add a config item to our "library" of items
// will panic if ENV_VAR_NAME is empty
// CliFlag and JSONKey can be empty, and will default to ENV_VAR_NAME
// typically, CliFlag is shorter and easy to type and JSONKey is empty or a lowercase, snake case variant of ENV_VAR_NAME
func (col Collection) Add(evn_var_name, shortflag, longflag, json_key string) {

	if len(evn_var_name) == 0 {
		panic("evn_var_name len is zero, expected non-zero len string")
	}

	ci := Item{
		ENV_VAR_NAME: evn_var_name,
		Shortflag:    shortflag,
		Longflag:     longflag,
		Json_key:     json_key,
	}

	col[evn_var_name] = ci
}

func (col Collection) Get(key string) Item {
	return col[key]
}

func (col Collection) Load(key string) error {
	return deeperror.NewTODOError(3055333906)
}

func (col Collection) LoadAll() error {
	return deeperror.NewTODOError(3761192315)
}

func (col Collection) Contains(key string) bool {
	_, exists_in_map := col[key]
	return exists_in_map
}

func (col Collection) Remove(key string) {
	delete(col, key)
}

// convenience
func (col Collection) StringValue(key string) string {
	return col.Get(key).String()
}
