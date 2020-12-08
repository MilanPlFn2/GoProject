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

func dijkstra(graph map[string]map[string]int, depart string, arrive string) string {
	var nonVisited [] string
	var distance_min = make(map[string]float64)
	var precedent = make (map[string] string)
	for i := range graph {
		distance_min[i] = math.Inf(1)
		precedent[i]= depart
		nonVisited=append(nonVisited,i)
	}
	distance_min[depart] = 0
	for true{
		var pointeur = min(distance_min, nonVisited)
		var index_pointeur = index(nonVisited, pointeur)
		nonVisited = append(nonVisited[:index_pointeur], nonVisited[(index_pointeur+1):]...)
		for v := range graph[pointeur] {
			if (cherche(v, nonVisited) && (float64(int(distance_min[pointeur])+graph[pointeur][v]) < distance_min[v])) {
				precedent[v] = pointeur
				distance_min[v] = float64(int(distance_min[pointeur]) + graph[pointeur][v])
			}
		}
		if len(nonVisited) == 0 {
			break
		}
	}

	return precedent[arrive]

}
func index (tableau[]string, elem string) int{
	var k int =0
	for i := range tableau {
		if tableau[i] == elem {
			k=i
		}
	}
	return k
}
func min (dict map[string]float64, nvisited [] string) string {
	var min float64 = math.Inf(1)
	var key = "Bonjour"
	for v := range nvisited {
		if (dict[nvisited[v]]< min) {
			min = dict[nvisited[v]]
			key= nvisited[v]
		}
	}
	return key
}
func cherche(Noeud string,Tableau []string) bool  {
	for i := range Tableau{
		if Tableau[i] == Noeud{
			return true
		}
	}
	return false
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
	var r =dijkstra(graph, "C", "B")
	fmt.Println(r)
}
