package highaccessemployees

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindHighAccessEmployees_EmployeeWithHighAccessInOneHour(t *testing.T) {

	accessTimes := [][]string{
		{"a", "0549"},
		{"b", "0457"},
		{"a", "0532"},
		{"a", "0621"},
		{"b", "0540"},
	}
	want := []string{"a"}
	got := findHighAccessEmployees(accessTimes)

	assert.ElementsMatch(t, want, got)
}
func TestFindHighAccessEmployees_NoEmployeeWithHighAccessInOneHour(t *testing.T) {

	accessTimes := [][]string{
		{"a", "0549"},
		{"b", "1457"},
		{"a", "0532"},
		{"a", "1621"},
		{"b", "1540"},
	}
	want := []string{}
	got := findHighAccessEmployees(accessTimes)

	assert.ElementsMatch(t, want, got)
}
func TestFindHighAccessEmployees_MultipleEmployeesWithHighAccessInOneHour(t *testing.T) {

	accessTimes := [][]string{
		{"a", "0549"},
		{"b", "0457"},
		{"a", "0532"},
		{"c", "0549"},
		{"a", "0621"},
		{"b", "0540"},
		{"c", "0550"},
		{"c", "0601"},
	}
	want := []string{"a", "c"}
	got := findHighAccessEmployees(accessTimes)

	assert.ElementsMatch(t, want, got)
}
func TestFindHighAccessEmployees_EmployeeWithHighAccessWithDuplicateTimes(t *testing.T) {

	accessTimes := [][]string{
		{"a", "0549"},
		{"b", "0457"},
		{"a", "0549"},
		{"a", "0549"},
		{"b", "0540"},
	}
	want := []string{"a"}
	got := findHighAccessEmployees(accessTimes)

	assert.ElementsMatch(t, want, got)
}
func TestFindHighAccessEmployees_EmptyAccessTimes(t *testing.T) {

	accessTimes := [][]string{}
	want := []string{}
	got := findHighAccessEmployees(accessTimes)

	assert.ElementsMatch(t, want, got)
}

func TestFindHighAccessEmployees_AccessTimesWithSingleEntry(t *testing.T) {

	accessTimes := [][]string{
		{"a", "0549"},
	}
	want := []string{}
	got := findHighAccessEmployees(accessTimes)

	assert.ElementsMatch(t, want, got)
}

func TestFindHighAccessEmployees_AccessTimesWithTwoEntriesLessThanOneHourApart(t *testing.T) {

	accessTimes := [][]string{
		{"a", "0549"},
		{"a", "0555"},
	}
	want := []string{}
	got := findHighAccessEmployees(accessTimes)

	assert.ElementsMatch(t, want, got)
}
func TestFindHighAccessEmployees_AccessTimesWithTwoEntriesMoreThanOneHourApart(t *testing.T) {

	accessTimes := [][]string{
		{"a", "0549"},
		{"a", "0655"},
	}
	want := []string{}
	got := findHighAccessEmployees(accessTimes)

	assert.ElementsMatch(t, want, got)
}
