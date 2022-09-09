package steps

import (
	"strings"

	"golang.org/x/exp/maps"
)

// MapperStep is used to map a maps field to another one.
type MapperStep struct {
	Name      string
	MapConfig map[string]string
}

func MapConstructor(keys []string) map[string]any {
	if len(keys) > 0 {
		key, keys := keys[0], keys[1:]
		return map[string]any{
			key: MapConstructor(keys),
		}
	}
	return nil
}

// Build creates the "new" map with the new keys.
func (mapper MapperStep) Build() map[string]any {
	newMapKeys := maps.Values(mapper.MapConfig)
	newMap := make(map[string]any)
	for _, keys := range newMapKeys {
		keyNames := strings.Split(keys, ".")
		key, keyNames := keyNames[0], keyNames[1:]
		newMap[key] = MapConstructor(keyNames)
	}
	return newMap
}

// AddData takes the newly made map and adds the data to it.
func (mapper MapperStep) AddData(data map[string]any, newData map[string]any) map[string]any {
	mapKeys := maps.Keys(mapper.MapConfig)
	for _, key := range mapKeys {
		keyLevels := strings.Split(key, ".")

		var currentValue any
		for keyLevel := range keyLevels{
		}
	}


	return newData
}

// Process takes the map and maps the fields to a new map.
// First it checks if the map is not nil.
// Then it looks over MapConfig and and builds a new map with the new fields.
// Finally it takes the given data and puts it in the new app.
func (mapper MapperStep) Process(data any) (any, error) {
	return nil, nil
}
