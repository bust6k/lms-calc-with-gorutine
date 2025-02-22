package Agent

import (
	"fmt"
	API2 "project_yandex_lms/Agent/API"
	"project_yandex_lms/calc"
	"strconv"
	"sync"
)

func CreateAgent(computing_power int) {
	chanOfTasks := make(chan float64, computing_power)
	var finalresult float64
	var mutex sync.Mutex
	for i := 0; i < computing_power; i++ {

		go func() {
			task, err := API2.GetNewTask("http:localhost:8080/internal/task")
			if err != nil {
				fmt.Errorf("ошибка при получении новой задачи : %v", err)
			}
			expression := strconv.FormatFloat(task.Arg1, 'f', 0, 64) + task.Operation + strconv.FormatFloat(task.Arg2, 'f', 0, 64)
			ResultOfGoroutine, err := calc.Calc(expression)
			if err != nil {
				fmt.Errorf("ошибка при расчете выражения: %v", err)
			}

			mutex.Lock()
			chanOfTasks <- ResultOfGoroutine
			mutex.Unlock()

		}()
	}
	for cv := range chanOfTasks {
		finalresult += cv
	}
	API2.PostTaskToServer(finalresult)
}
