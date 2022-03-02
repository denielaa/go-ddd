package aggregate_test

import (
	"testing"

	"github.com/denielaa/go-ddd/aggregate"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type testCase struct {
		test 						string
		name 						string
		expectedError 	error
	}

	testCases := []testCase{
		{
			test: "Empty name validation",
			name: "",
			expectedError: aggregate.ErrInvalidPerson,
		}, {
			test: "Valid name validation",
			name: "Deniel",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregate.NewCustomer(tc.name)

			if err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}