package types

import (
	"encoding/json"
	"testing"
)

func TestCloudWatchEMFMetric(t *testing.T) {
	cloudWatchEmfMetric := CloudWatchEMFMetric{
		Name: "TestMetric",
		Unit: "Count",
	}
	expected := `{"Name":"TestMetric","Unit":"Count"}`
	actual, error := json.Marshal(cloudWatchEmfMetric)
	if error != nil {
		t.Errorf("failed to generate json")
	}
	if string(actual) != expected {
		t.Errorf("expected %s, got %s", expected, string(actual))
	}
}

func TestCloudWatchEMFMetricWithStorageResolution(t *testing.T) {
	cloudWatchEmfMetric := CloudWatchEMFMetric{
		Name:              "TestMetric",
		Unit:              "Count",
		StorageResolution: 60,
	}
	expected := `{"Name":"TestMetric","Unit":"Count","StorageResolution":60}`
	actual, error := json.Marshal(cloudWatchEmfMetric)
	if error != nil {
		t.Errorf("failed to generate json")
	}
	if string(actual) != expected {
		t.Errorf("expected %s, got %s", expected, string(actual))
	}
}
