package main

type Workable interface {
	Work()
}

type Eatable interface {
	Eat()
}

func startWork(w Workable) {
	w.Work()
}

func startEat(e Eatable) {
	e.Eat()
}

func goodInit() {
	human := Human{}
	robot := Robot{}

	startWork(&human)
	startWork(&robot)

	startEat(&human)
	// startEat(&robot) -> robot not implement this method because robot can't eat
}
