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
		out = func(in In, stage Stage) (out Out) {
			ch := make(chan interface{})

			go func(ch chan interface{}) {
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
			}(ch)

			return stage(ch)
		}(out, stage)
	}

	return out
}
