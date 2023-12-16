// math.go
package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Add(a, b int) int {
	return a + b
}

type MyMockedObject struct {
	mock.Mock
}

// Add is a method on MockCalculator that satisfies the Calculator interface
func (m *MyMockedObject) SecondAdd(a, b int) int {
	args := m.Called(a, b)
	return args.Int(0)
}

func TestCalculatorUsingMock(t *testing.T) {

	// create an instance of our test object
	testObj := new(MyMockedObject)

	// set up expectations
	testObj.On("SecondAdd", 5, 3).Return(8)

	// Call the method we want to test
	result := testObj.SecondAdd(5, 3)

	testObj.AssertExpectations(t)

	// assert that the expectations were met
	assert.Equal(t, 8, result)

}

// TestAdd directly tests the Add function.
func TestAdd(t *testing.T) {
	got := Add(5, 3)
	want := 8

	assert.Equal(t, want, got, "Add(5, 3) should equal 8")
}
