package main

type Worker interface {
	Work()
	Eat()
}

func startWoker(w Worker) {
	w.Work()
	w.Eat()
}

func badInit() {
	badHuman := BadHuman{}
	badRobot := BadRobot{}

	startWoker(&badHuman)
	startWoker(&badRobot)
}
