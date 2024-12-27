package main

type GoLimit struct {
	ch chan int
}

func NewGoLimit(limit int) *GoLimit {
	return &GoLimit{
		ch: make(chan int, limit),
	}
}

func (l *GoLimit) Add() {
	l.ch <- 1
}

func (l *GoLimit) Done() {
	<-l.ch
}

func (l *GoLimit) Close() {
	close(l.ch)
}

//func main() {
//	goLimit := NewGoLimit(2)
//	for i := 0; i < 10; i++ {
//		goLimit.Add()
//		go func(i int) {
//			defer goLimit.Done()
//			// do something
//			time.Sleep(time.Second * 1)
//			log.Println(i, "done")
//		}(i)
//	}
//	goLimit.Close()
//}
