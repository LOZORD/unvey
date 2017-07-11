// Package unvey exposes a BDD testing library.
// One of the benefits of unvey is that it doesn't use GLS!
package unvey

import (
	"sync"
	"testing"
)

// Case represents a testable unit: the test case name and how to test it.
type Case struct {
	// TODO(LOZORD): Rename to `Should`?
	Name        string
	Expectation func(*testing.T)
}

// Spec represents a recursive testing tree.
type Spec struct {
	// TODO(LOZORD): Rename to `When`?
	Name       string
	BeforeAll  func(*testing.T)
	BeforeEach func(*testing.T)
	It         []Case
	AfterEach  func(*testing.T)
	AfterAll   func(*testing.T)
	// Children test cases.
	AndWhen []Spec
}

// Run recursively executes a Spec using the given testing parameter.
func Run(t *testing.T, s Spec) {
	t.Run(s.Name, func(t *testing.T) {
		if s.BeforeAll != nil {
			s.BeforeAll(t)
		}

		var caseGroup sync.WaitGroup
		for _, testCase := range s.It {
			caseGroup.Add(1)
			go func(tc Case) {
				if s.BeforeEach != nil {
					s.BeforeEach(t)
				}

				t.Run(tc.Name, func(t *testing.T) {
					if tc.Expectation != nil {
						tc.Expectation(t)
					} else {
						t.SkipNow()
					}
				})

				if s.AfterEach != nil {
					s.AfterEach(t)
				}

				caseGroup.Done()
			}(testCase)
		}
		caseGroup.Wait()

		var childGroup sync.WaitGroup
		for _, nextDesc := range s.AndWhen {
			childGroup.Add(1)
			go func(ns Spec) {
				Run(t, ns)
				childGroup.Done()
			}(nextDesc)
		}
		childGroup.Wait()

		if s.AfterAll != nil {
			s.AfterAll(t)
		}

	})
}
