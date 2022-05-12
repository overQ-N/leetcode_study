package delete_columns_to_make_sorted

import "testing"

func TestMinDeletionSize(t *testing.T) {
	test1 := minDeletionSize([]string{"abc", "bce", "cae"})
	if test1 != 1 {
		t.Error("result must be equal to 1")
	}

	test2 := minDeletionSize([]string{"cba", "daf", "ghi"})
	if test2 != 1 {
		t.Error("result must be equal to 2")
	}

	test3 := minDeletionSize([]string{"a", "b"})
	if test3 != 0 {
		t.Error("result must be equal to 0")
	}
}
