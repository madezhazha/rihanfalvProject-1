package data

import "testing"

func Test_RsByCondition(t *testing.T) {
	conditions := []string{"刑法"}
	if _, err := RsByCondition(conditions); err != nil {
		t.Error(err)
	}
}
