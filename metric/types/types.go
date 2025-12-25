package types

type CloudWatchEMFMetric struct {
	Name              string `json:"Name"`
	Unit              string `json:"Unit"`
	StorageResolution int    `json:"StorageResolution,omitempty"`
}

type CloudWatchEMFMetrics struct {
	Namespace  string                `json:"Namespace"`
	Dimensions [][]string            `json:"Dimensions"`
	Metrics    []CloudWatchEMFMetric `json:"Metrics"`
}

type CloudWatchEMFRoot struct {
	Timestamp         int64                  `json:"Timestamp"`
	CloudWatchMetrics []CloudWatchEMFMetrics `json:"CloudWatchMetrics"`
}

type CloudWatchEMFOutput struct {
	Aws CloudWatchEMFRoot `json:"_aws"`
}
