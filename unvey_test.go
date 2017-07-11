package unvey_test

import (
	"testing"

	"github.com/lozord/unvey"
)

func TestBlah(t *testing.T) {

	counter := struct {
		name  string
		value int
	}{}

	unvey.Run(t, unvey.Spec{
		Name: "when using a counter",
		BeforeEach: func(_ *testing.T) {
			counter.name = "gopher"
			counter.value = 1
		},
		It: []unvey.Case{{
			Name: "should increment correctly",
			Expectation: func(t *testing.T) {
				counter.value++

				if counter.value != 2 {
					t.Errorf("counter.value = %d, wanted 2", counter.value)
				}
			},
		}, {
			Name: "should not have changed name",
			Expectation: func(t *testing.T) {
				if counter.name != "gopher" {
					t.Errorf("counter.name = %q, wanted 'gopher'", counter.name)
				}
			},
		}},
		AndWhen: []unvey.Spec{{
			Name: "the counter name is updated",
			BeforeAll: func(_ *testing.T) {
				counter.name = "badger"
			},
			It: []unvey.Case{
				{
					Name: "should be what we expect",
					Expectation: func(t *testing.T) {
						if counter.name != "badger" {
							t.Errorf("counter.name = %q, wanted 'badger'", counter.name)
						}
					},
				}, {
					// Skip this test.
					Name: "should do quantum computations",
				},
			},
		}},
	})
}
