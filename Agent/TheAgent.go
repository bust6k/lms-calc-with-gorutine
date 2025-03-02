package Agent

import (
	"context"
	"fmt"
	API2 "project_yandex_lms/Agent/API"
	"project_yandex_lms/calc"
	"project_yandex_lms/structures"
	"strconv"
	"sync"
)

func CreateAgent(computing_power int) {
	chanOfTasks := make(chan structures.Task, computing_power)
	chanOfResults := make(chan float64, computing_power)
	errChan := make(chan error, computing_power)
	var wg sync.WaitGroup
	var stackMutex sync.Mutex
	var stack []float64

	wg.Add(computing_power)
	for i := 0; i < computing_power; i++ {
		go func() {
			defer wg.Done()
			task, err := API2.GetNewTask("http://localhost:8080/internal/task")
			if err != nil {
				errChan <- fmt.Errorf("ошибка при получении новой задачи: %v", err)
				return
			}
			ctx, cancel := context.WithTimeout(context.Background(), task.Operation_time)
			defer cancel()

			select {
			case <-ctx.Done():
				errChan <- fmt.Errorf("ошибка, горутина слишком долго выполняет свою работу")
				return
			default:
				chanOfTasks <- task

				expression := strconv.FormatFloat(task.Arg1, 'f', -1, 64) + task.Operation + strconv.FormatFloat(task.Arg2, 'f', -1, 64)

				result, err := calc.Calc(expression)
				if err != nil {
					errChan <- fmt.Errorf("ошибка при расчете выражения: %v", err)
					return
				}

				chanOfResults <- result
			}
		}()
	}

	go func() {
		wg.Wait()
		close(chanOfTasks)
		close(chanOfResults)
		close(errChan)
	}()

	for taksi := range chanOfTasks {
		result := <-chanOfResults

		stackMutex.Lock()

		switch taksi.Operation {
		case "*", "/":
			if len(stack) > 0 {
				last := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if taksi.Operation == "*" {
					stack = append(stack, last*result)
				} else {
					if result == 0 {
						errChan <- fmt.Errorf("деление на ноль")
						stackMutex.Unlock()
						continue
					}
					stack = append(stack, last/result)
				}
			} else {
				stack = append(stack, result)
			}
		case "+", "-":
			stack = append(stack, result)
		}

		stackMutex.Unlock()
	}

	for err := range errChan {
		fmt.Println("Ошибка:", err)
	}

	if len(stack) == 0 {
		fmt.Println("Ошибка: стек пуст, результат не вычислен")
		return
	}

	finalresult := stack[len(stack)-1]
	API2.PostTaskToServer(finalresult)

}
