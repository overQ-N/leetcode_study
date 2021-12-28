package dp

import (
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	ans := longestCommonSubsequence("abcde", "ace")
	if ans != 3 {
		t.Fatalf("res:%d should be 3", ans)
	}
	// ans = longestCommonSubsequence("pmjghexybyrgzczy", "hafcdqbgncrcbihkd")

}
