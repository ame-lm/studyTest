package cal

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		desc    string
		intputA int
		inputB  int
		want    int
	}{
		{
			desc:    "",
			intputA: 12,
			inputB:  13,
			want:    25,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			c := new(IntCalculatorImpl)
			out := c.add(tC.intputA, tC.inputB)
			if out != tC.want {
				t.Errorf("c.add(%d,%d)=%d;want=%d", tC.intputA, tC.inputB, out, tC.want)
			}
		})
	}
}
