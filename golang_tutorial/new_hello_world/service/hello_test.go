package hello

import "testing"

func TestSay(t *testing.T) {
	subset := []struct {
		items  []string
		result string
	}{
		{items: []string{"saeed", "delaram"}, result: "Hello, saeed, delaram!"},
		{result: "Hello, world!"},
	}
	for _, subset_item := range subset {
		if s := Say(subset_item.items); s != subset_item.result {
			t.Errorf("wanted %s .%v, got %s", subset_item.result, subset_item.items, s)
		}
	}
}
