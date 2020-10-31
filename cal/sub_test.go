package cal

import "testing"

func TestSub(t *testing.T) {
	testCases := []struct {
		desc    string
		intputA int
		inputB  int
		want    int
	}{
		{
			desc:    "",
			intputA: 13,
			inputB:  12,
			want:    1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c := new(IntCalculatorImpl)
			out := c.sub(tC.intputA, tC.inputB)
			if out != tC.want {
				t.Errorf("c.sub(%d,%d)=%d;want=%d", tC.intputA, tC.inputB, out, tC.want)
			}
		})
	}
}
