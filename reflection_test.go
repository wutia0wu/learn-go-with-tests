package reflction

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	City string
	Age  int
}

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

		{
			"struct with two field",
			struct {
				Name string
				City string
			}{"Truman", "NewYork"},
			[]string{"Truman", "NewYork"},
		},

		{
			"struct with int field",
			struct {
				Name string
				Age  int
			}{"Truman", 22},
			[]string{"Truman"},
		},

		{
			"struct with struct field",
			Person{
				"Truman",
				Profile{"NewYork", 22},
			},
			[]string{"Truman", "NewYork"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})

	}

}
