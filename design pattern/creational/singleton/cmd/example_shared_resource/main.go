package main

func main() {
	var ch = make(chan struct{})
	new_class()
	go func() {
		new_class()
		ch <- struct{}{}
	}()
	<-ch
}
