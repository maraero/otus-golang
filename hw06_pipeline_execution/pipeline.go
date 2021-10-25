package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		ch := make(chan interface{})

		go func(in In, ch chan interface{}) {
			defer close(ch)

			for {
				select {
				case <-done:
					return
				case v, ok := <-in:
					if !ok {
						return
					}
					ch <- v
				}
			}
		}(out, ch)

		out = stage(ch)
	}

	return out
}
