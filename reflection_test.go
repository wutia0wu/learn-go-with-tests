package reflction

import "testing"

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one field",
			struct {
				Name string
			}{"Truman"},
			[]string{"Truman"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if len(got) != 1 {
				t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
			}

			if got[0] != test.ExpectedCalls[0] {
				t.Errorf("got %s want %s", got[0], test.ExpectedCalls[0])
			}
		})

	}

}
