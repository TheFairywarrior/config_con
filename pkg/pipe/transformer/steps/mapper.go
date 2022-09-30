package steps

import (
	"fmt"
	"strings"

	"golang.org/x/exp/maps"
)

// MapperStep is a struct used to hold the config and logic for mapping one map structure to another.
// The config takes the shape of a map[string]string.
//
// Where the key is the path to the value in the original map.
// And the value is the path to the value within the new map.
//
// Example:
//
//	{
//		"foo.bar": "baz.qux",
//		"foo.baz": "baz.qux",
//	}
type MapperStep struct {
	Name      string
	MapConfig map[string]string
}

func NewMapperStep(name string, mapConfig map[string]string) MapperStep {
	return MapperStep{
		Name:      name,
		MapConfig: mapConfig,
	}
}

// mapConstructor takes a slice of "levels" for the map.
// Using recursion it creates a new map with the given levels.
func mapConstructor(keys []string) map[string]any {
	if len(keys) > 0 {
		key, keys := keys[0], keys[1:]
		return map[string]any{
			key: mapConstructor(keys),
		}
	}
	return nil
}

// build takes the MapConfig and usises the "mapConstructor" to build a new map.
// NOTE: this map has no values.
func (mapper MapperStep) build() map[string]any {
	newMapKeys := maps.Values(mapper.MapConfig)
	newMap := make(map[string]any)
	for _, keys := range newMapKeys {
		keyNames := strings.Split(keys, ".")
		key, keyNames := keyNames[0], keyNames[1:]
		newMap[key] = mapConstructor(keyNames)
	}
	return newMap
}

// getMapValue takes a map and slice of keys and then returns the value of the last key.
func getMapValue(data any, keys []string) (any, error) {
	if len(keys) > 1 {
		key, keys := keys[0], keys[1:]
		return getMapValue(data.(map[string]any)[key].(map[string]any), keys)
	}
	out, ok := data.(map[string]any)[keys[0]]
	if !ok {
		return nil, fmt.Errorf("key %s not found", keys[0])
	}
	return out, nil
}

// setMapValue takes a map and slice of keys and then sets the value of the last key.
func setMapValue(data map[string]any, keys []string, value any) {
	if len(keys) == 1 {
		data[keys[0]] = value
	} else {
		key, keys := keys[0], keys[1:]
		setMapValue(data[key].(map[string]any), keys, value)
	}
}

// AddData takes the newly made map and adds the data to it.
func (mapper MapperStep) AddData(data map[string]any, newData map[string]any) (map[string]any, error) {
	for dataKeys, newKeys := range mapper.MapConfig {
		keyNames := strings.Split(dataKeys, ".")
		valueNames := strings.Split(newKeys, ".")
		value, err := getMapValue(data, keyNames)
		if err != nil {
			return nil, err
		}
		setMapValue(newData, valueNames, value)
	}

	return newData, nil
}

// Process takes the map and maps the fields to a new map.
// First it checks if the map is not nil.
// Then it looks over MapConfig and and builds a new map with the new fields.
// Finally it takes the given data and puts it in the new app.
func (mapper MapperStep) Process(data any) (any, error) {
	newMap := mapper.build()
	return mapper.AddData(data.(map[string]any), newMap)
}
