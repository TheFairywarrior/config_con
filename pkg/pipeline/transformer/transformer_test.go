package transformer

import (
	"testing"
)

func TestBaseTransformedData(t *testing.T) {
	data := map[string]any{
		"key1": "value1",
		"key2": "value2",
	}

	transformedData := BaseTransformedData{
		data: data,
	}

	// Test accessing data
	if transformedData.data["key1"] != "value1" {
		t.Errorf("Expected value1, got %v", transformedData.data["key1"])
	}

	if transformedData.data["key2"] != "value2" {
		t.Errorf("Expected value2, got %v", transformedData.data["key2"])
	}

	// Test modifying data
	transformedData.data["key1"] = "newvalue1"
	if transformedData.data["key1"] != "newvalue1" {
		t.Errorf("Expected newvalue1, got %v", transformedData.data["key1"])
	}

	// Test adding new data
	transformedData.data["key3"] = "value3"
	if transformedData.data["key3"] != "value3" {
		t.Errorf("Expected value3, got %v", transformedData.data["key3"])
	}

	// Test deleting data
	delete(transformedData.data, "key2")
	if _, ok := transformedData.data["key2"]; ok {
		t.Errorf("Expected key2 to be deleted")
	}
}
