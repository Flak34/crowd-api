package crowdapiv1

type TaskInputDataDTO struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type TaskOutputDataDTO struct {
	Name            string `json:"name"`
	Type            string `json:"type"`
	Value           string `json:"value"`
	WithAggregation bool   `json:"with_aggregation"`
}
