package work_queue

type Worker interface {
	Run() interface{}
}

type WorkQueue struct {
	Jobs         chan Worker
	Results      chan interface{}
	StopRequests chan int
	NumWorkers   uint
}

// Create a new work queue capable of doing nWorkers simultaneous tasks, expecting to queue maxJobs tasks.
func Create(nWorkers uint, maxJobs uint) *WorkQueue {
	q := new(WorkQueue)

	q.NumWorkers =  nWorkers
	q.Jobs = make(chan Worker , maxJobs)
	q.Results = make(chan interface{} , maxJobs)
	q.StopRequests = make(chan int , maxJobs)

	for i :=0;i < int(nWorkers);i++{

		go q.worker()

	}




	// TODO

	return q
}

//A worker goroutine that processes tasks from .Jobs unless .StopRequests has a message saying to halt now.
func (queue WorkQueue) worker() {


	running := true
	signal := len(queue.StopRequests)


	// Run tasks from the queue, unless we have been asked to stop.
	for running {



		if queue.Jobs != nil && signal == 0 {


			for c:= range queue.Jobs{

				if len(queue.StopRequests) == 0{

					queue.Results <- c.Run()
				}

			}
		}else{

			return

		}



		// TODO: listen on the .Jobs channel for incoming tasks
		// TODO: run tasks by calling .Run()
		// TODO: send the return value back on Results channel
		// TODO: exit (return) when a signal is sent on StopRequests
	}
}

func (queue WorkQueue) Enqueue(work Worker) {
	// TODO


	queue.Jobs <- work



}
////
func (queue WorkQueue) Shutdown() {
	// TODO

	for i:= int(queue.NumWorkers) ; i !=0;i-- {

		queue.StopRequests <- int(queue.NumWorkers)
	}






}
