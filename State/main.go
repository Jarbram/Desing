package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Interfaz para el estado de la tarea
type TaskState interface {
	Display()
	SetComplete()
	SetIncomplete()
}

// Estado para una tarea pendiente
type PendingTask struct{}

func (p *PendingTask) Display() {
	fmt.Println("Esta tarea está pendiente.")
}

func (p *PendingTask) SetComplete() {
	fmt.Println("La tarea ahora está completada.")
}

func (p *PendingTask) SetIncomplete() {
	fmt.Println("La tarea todavía está pendiente.")
}

// Estado para una tarea completada
type CompletedTask struct{}

func (c *CompletedTask) Display() {
	fmt.Println("Esta tarea está completada.")
}

func (c *CompletedTask) SetComplete() {
	fmt.Println("La tarea ya está completada.")
}

func (c *CompletedTask) SetIncomplete() {
	fmt.Println("La tarea ahora está incompleta.")
}

// Estructura para la tarea que puede cambiar de estado
type Task struct {
	state TaskState
}

func (t *Task) SetState(state TaskState) {
	t.state = state
}

func (t *Task) Display() {
	t.state.Display()
}

func (t *Task) Complete() {
	t.state.SetComplete()
}

func (t *Task) Incomplete() {
	t.state.SetIncomplete()
}

func main() {
	// Crear una nueva tarea pendiente
	task := &Task{&PendingTask{}}

	// Ciclo de entrada de consola
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ingrese una acción (display/complete/incomplete/exit): ")
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		input = strings.TrimSpace(input)

		// Realizar la acción correspondiente según la entrada del usuario
		switch input {
		case "display":
			task.Display()
		case "complete":
			task.Complete()
		case "incomplete":
			task.Incomplete()
		case "exit":
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Entrada inválida.")
		}

		fmt.Println("Ingrese una acción (display/complete/incomplete/exit): ")
	}
}
