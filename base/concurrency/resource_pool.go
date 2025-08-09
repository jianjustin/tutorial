package concurrency

type Resource struct {
	ID int
}
type Pool struct {
	resources chan chan *Resource // 资源池，元素是资源 channel
}

func NewPool(size int) *Pool {
	resources := make(chan chan *Resource, size)
	for i := 0; i < size; i++ {
		resCh := make(chan *Resource, 1)
		resCh <- &Resource{ID: i + 1}
		resources <- resCh
	}
	return &Pool{resources: resources}
}

func (p *Pool) Acquire() chan *Resource {
	return <-p.resources
}

func (p *Pool) Release(resCh chan *Resource) {
	p.resources <- resCh
}
