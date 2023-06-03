package templatemethod

import "fmt"

type WorkerInterface interface {
	GetUp()
	EatBreakfast()
	GoToWork()
	Work()
	ReturnHome()
	Relax()
	Sleep()
}

// Worker defines the worker.
type Worker struct {
	WorkerInterface
}

// NewWorker returns a new Worker.
func NewWorker(w WorkerInterface) *Worker {
	return &Worker{w}
}

// DailyRoutine is the template method for printing the workers daily routine.
func (w *Worker) DailyRoutine() {
	w.GetUp()
	w.EatBreakfast()
	w.GoToWork()
	w.Work()
	w.ReturnHome()
	w.Relax()
	w.Sleep()
}

// PostMan is a worker.
type PostMan struct {
}

// GetUp prints what the postman does to get up.
func (w *PostMan) GetUp() {
	fmt.Printf("Getting up\n")
}

// EatBreakfast prints what the postman does to eat breakfast.
func (w *PostMan) EatBreakfast() {
	fmt.Printf("Eating pop tarts\n")
}

// GoToWork prints what the postman does to get to work.
func (w *PostMan) GoToWork() {
	fmt.Printf("Cycle to work\n")
}

// Work prints what the postman does to work.
func (w *PostMan) Work() {
	fmt.Printf("Post letters\n")
}

// ReturnHome prints what the postman does to get home.
func (w *PostMan) ReturnHome() {
	fmt.Printf("Cycle home\n")
}

// Relax prints what the postman does to relax.
func (w *PostMan) Relax() {
	fmt.Printf("Collect stamps\n")
}

// Sleep prints what the postman does to sleep.
func (w *PostMan) Sleep() {
	fmt.Printf("Zzzzzzz\n")
}
