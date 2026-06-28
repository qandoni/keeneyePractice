package student_policy

type StudentScope struct {
	All       bool
	StudentID *int
	GroupIDs  []int
}
