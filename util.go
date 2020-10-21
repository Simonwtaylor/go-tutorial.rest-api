package main

// RemoveStringIndex -> Util -> will remove an item within a slice of strings and return it
func RemoveStringIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

// RemoveExampleDataIndex -> Util -> will remove an item within a slice of ExampleDatas and return it
func RemoveExampleDataIndex(s []ExampleData, index int) []ExampleData {
	return append(s[:index], s[index+1:]...)
}
