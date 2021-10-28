package pkg

import (
	"net/url"
	"strconv"
)

func TakeAllParams(url url.Values) map[string]int {
	params := make(map[string]int, 4)

	//all filter params
	filters := []string{"selled", "deleted", "feature", "category"}

	for param := range url {
		for _, filter := range filters {
			if param == filter {
				data, err := strconv.Atoi(url.Get(param))
				if err != nil {
					return map[string]int{"error": 0}
				}
				params[param] = data
				break
			}
		}
	}

	return params
}
