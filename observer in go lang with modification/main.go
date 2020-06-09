package main
func main() {
handler := customObserver.NewSubjectHandler()
      // add observer what you neeed here
	go handler.Notify() // notify all
}
