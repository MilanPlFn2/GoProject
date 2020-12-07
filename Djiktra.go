package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func dijkstra(graph map[string]map[string]int, depart string, arrive string) {
	var noeudVisited []string
	var distance_min = make(map[string]float64)
	for i := range graph {
		distance_min[i] = math.Inf(1)
	}
	distance_min[depart] = 0
	fmt.Println(distance_min)
	fmt.Println(noeudVisited)
	var condition = true
	for test := true; test; test = condition {
		noeudVisited = append(noeudVisited, "A")
		if len(noeudVisited) == len(graph) {
			condition = false
		}
	}
	fmt.Println(noeudVisited)
}

func main() {
	var graph = make(map[string]map[string]int)
	file, err := os.OpenFile("matrice", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var chemin = make(map[string]int)
		line := strings.Split(scanner.Text(), ": ")
		line2 := strings.Split(strings.Split(strings.Split(line[1], "{ ")[1], " }")[0], ",")
		//fmt.Println(line[0])
		for i := 0; i < (len(line2)); i++ {
			//fmt.Println(line2[i],i)
			line3 := strings.Split(line2[i], ":")
			//fmt.Println(line3[0],line3[1])
			nbr, err := strconv.Atoi(line3[1])
			if err != nil {
				log.Fatal(err)
			}
			chemin[line3[0]] = nbr
			graph[line[0]] = chemin
		}
		//fmt.Println("cc")
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(graph)
	dijkstra(graph, "A", "D")
}
