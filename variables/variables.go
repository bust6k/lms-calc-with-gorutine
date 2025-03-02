package variables

import "project_yandex_lms/structures"

var (
	Count_Root_Id = 0
)

var Expressions []structures.Expression

var CurrentTask structures.Task
var TheTasks []structures.Task

var Operators = map[string]int{
	"+": 1,
	"-": 1,
	"*": 2,
	"/": 2,
}
