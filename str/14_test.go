package str

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	res := longestCommonPrefix([]string{"flower", "flow", "flight"})
	if res != "fl" {
		t.Fatalf("res should be equal `fl`,but value is %s get.", res)
	}

	res = longestCommonPrefix([]string{"dog", "racecar", "car"})
	if res != "" {
		t.Fatalf("res should be equal ``,but value is %s get.", res)
	}

	res = longestCommonPrefix([]string{"cir", "car"})
	if res != "c" {
		t.Fatalf("res should be equal `c`,but value is %s get.", res)
	}
}
