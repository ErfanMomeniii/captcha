package captcha

import (
	"strconv"
	"testing"
)

func TestMath(t *testing.T) {
	c := New()
	for i := 0; i < 50; i++ {
		r, err := c.Math()
		if err != nil {
			t.Fatal(err)
		}
		if _, err := strconv.Atoi(r.Text); err != nil {
			t.Fatalf("answer %q not an integer", r.Text)
		}
	}
}

func TestEvalExpr(t *testing.T) {
	cases := []struct {
		a, b int
		op   byte
		want int
	}{
		{3, 5, '+', 8},
		{9, 4, '-', 5},
		{6, 7, '*', 42},
	}
	for _, tc := range cases {
		if got := evalExpr(tc.a, tc.b, tc.op); got != tc.want {
			t.Fatalf("%d%c%d = %d, want %d", tc.a, tc.op, tc.b, got, tc.want)
		}
	}
}
