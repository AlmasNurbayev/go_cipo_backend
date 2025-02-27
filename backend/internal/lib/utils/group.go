package utils

import (
	"fmt"
	"strings"
)

type Record map[string]any

func GroupAndSum(arr []Record, groupKeys, sumKeys, addKeys []string) []Record {
	grouped := make(map[string]Record)

	for _, curr := range arr {
		// Create group key
		var groupKeyParts []string
		for _, k := range groupKeys {
			if val, ok := curr[k]; ok {
				groupKeyParts = append(groupKeyParts, fmt.Sprintf("%v", val))
			}
		}
		groupKey := strings.Join(groupKeyParts, "-")

		// Initialize group if not exists
		if _, exists := grouped[groupKey]; !exists {
			grouped[groupKey] = make(Record)
			for _, k := range groupKeys {
				grouped[groupKey][k] = curr[k]
			}
			for _, k := range sumKeys {
				grouped[groupKey][k] = 0.0
			}
			for _, k := range addKeys {
				grouped[groupKey][k] = []interface{}{}
			}
		}

		// Sum values
		for _, k := range sumKeys {
			if val, ok := curr[k].(float64); ok {
				grouped[groupKey][k] = grouped[groupKey][k].(float64) + val
			}
		}

		// Collect additional values
		for _, k := range addKeys {
			if val, ok := curr[k]; ok {
				grouped[groupKey][k] = append(grouped[groupKey][k].([]interface{}), val)
			}
		}
	}

	// Convert map to slice
	result := []Record{}
	for _, v := range grouped {
		result = append(result, v)
	}

	return result
}
