package main

func main() {

}

func countStudents(students []int, sandwiches []int) int {
	unableToEat := 0
	for i := 0; i < len(students); {
		if students[i] == sandwiches[i] {
			students = students[i+1:]
			sandwiches = sandwiches[i+1:]
			unableToEat = 0
		} else {
			students = append(students[i+1:], students[i])
			unableToEat++
		}

		if unableToEat == len(students) {
			return unableToEat
		}
	}
	return unableToEat
}
