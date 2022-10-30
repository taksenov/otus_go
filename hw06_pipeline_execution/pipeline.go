// Package hw06pipelineexecution -- Otus Go HW06.
package hw06pipelineexecution

type (
	// In job input.
	In = <-chan interface{}
	// Out job output.
	Out = In
	// Bi channel.
	Bi = chan interface{}
)

// Stage -- job typing.
type Stage func(in In) (out Out)

// ExecutePipeline -- a few more code-lines and gitlab will go bankrupt.
func ExecutePipeline(in In, done In, stages ...Stage) Out {
	out := in

	for _, stage := range stages {
		ch := make(Bi)
		out = worker(done, ch, stage, out)
	}

	return out
}

func worker(done In, ch Bi, stage Stage, out Out) Out {
	go func(ch Bi, out Out) {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case val, ok := <-out:
				if !ok {
					return
				}
				ch <- val
			}
		}
	}(ch, out)

	return stage(ch)
}
