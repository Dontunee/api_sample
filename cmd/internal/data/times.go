package data

import (
	"time"
)

type Times struct {
	Tz []string
}

func (t *Times) GetTime() (map[string]string, error) {

	data := map[string]string{}

	if len(t.Tz) < 1 {
		data["current_time"] = time.Now().UTC().String()
		return data, nil
	}

	for _, value := range t.Tz {
		loc, err := time.LoadLocation(value)
		if err != nil {
			return nil, err
		}
		data[value] = time.Now().In(loc).String()
	}

	return data, nil
}
