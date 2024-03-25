package zbxctl

import (
	"strings"
)

func convertOperator(operator string) string {
	operator = strings.ToLower(operator)
	switch operator {
	case "like":
		return "0"
	case "equal":
		return "1"
	case "not like":
		return "2"
	case "not equal":
		return "3"
	case "exists":
		return "4"
	case "not exists":
		return "5"
	default:
		return "0"
	}
}

func convertZbxTagsFilter(tagsFilterConfig []ZbxTagsFilterProblemConfig) []zbxTagsFilterProblem {
	zTags := make([]zbxTagsFilterProblem, len(tagsFilterConfig))
	for i, zTag := range tagsFilterConfig {
		zTags[i].Operator = convertOperator(zTag.Operator)
		zTags[i].Value = zTag.Value
		zTags[i].Tag = zTag.Tag
	}
	return zTags
}
