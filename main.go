package main

import "log"

func main() {
	g := createGraph([]string{"A", "B", "C"})
	if err := addEdge("A", "B", &g); err != nil {
		log.Println(err.Error())
	}
	if err := addEdge("B", "A", &g); err != nil {
		log.Println(err.Error())
	}
	if err := addEdge("B", "C", &g); err != nil {
		log.Println(err.Error())
	}
	if err := addEdge("C", "A", &g); err != nil {
		log.Println(err.Error())
	}

	displayGraph(g)

	if err := deleteEdge("A", "B", &g); err != nil {
		log.Println(err.Error())
	}
	if err := deleteEdge("B", "C", &g); err != nil {
		log.Println(err.Error())
	}
	if err := deleteEdge("A", "C", &g); err != nil {
		log.Println(err.Error())
	}

	displayGraph(g)
}
