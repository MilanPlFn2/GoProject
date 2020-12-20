package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

//Gorout prend en parametre le sommet de depart, le sommet d'arrive et le graph
//Cette fontion appelle la fonction dijkstra et la fonction affichage
func gorout(dep string, arriv string, graph map[string]map[string]int) {
	var precedent, distance = dijkstra(graph, dep, arriv)
	//dijkstra(graph, dep, arriv)
	affichage(dep, arriv, precedent, distance)
	wg.Done()
}

//testchemin prend en parametre le chemin du fichier matrice
//Cette fontion verifie si le fichier existe
//Retourne un booléen
func testchemin(chemin string) bool {
	if _, err := os.Stat(chemin); os.IsNotExist(err) {
		fmt.Println("Le fichier matrice n'existe pas !!!")
		return false
	}
	return true
}

//toutdijkstra prend en parametre le graph
//Cette fontion s'occupe de lancer toute les goroutine pour chaque ligne de dijkstra
func toutdijkstra(graph map[string]map[string]int) {
	fmt.Printf("|%7s | %7s | %7s | %7s |\n", "Départ", "Arrivée", "Prédécesseur", "Distance")
	fmt.Println("|--------|---------|--------------|----------|")
	for i := range graph {
		for v := range graph {
			wg.Add(1)
			go gorout(i, v, graph)
		}
	}
	wg.Wait()
	fmt.Println("|--------|---------|--------------|----------|")
}

//affichage prend en parametre le sommet de depart, le sommet d'arrive,le sommet précédent et le Cout du Chemin le Plus Court
//Cette fontion affiche | Sommet de Départ | Sommet d'Arrivée | Sommet précédent | Cout du Chemin le Plus Court |
func affichage(depart string, arrive string, precedent string, distance float64) {
	fmt.Printf("|   %-4s |    %-3s  |      %-6s  |    %-4.0f  |\n", depart, arrive, precedent, distance)
}

//dijkstra prend en parametre le graph, le sommet de depart et le sommet d'arrive
//Cette fontion trouve le chemin le plus court d'un sommet vers un autre en fonction de son cout
//Retourne un string (Le sommet précedent) et un float64 (Le cout du chemin le plus court)
func dijkstra(graph map[string]map[string]int, depart string, arrive string) (string, float64) {
	var nonVisited []string
	var distance_min = make(map[string]float64)
	var precedent = make(map[string]string)
// Initialisation des tableaux de distance, precedent et des noeuds non visités 
	for i := range graph {
		distance_min[i] = math.Inf(1)
		precedent[i] = depart
		nonVisited = append(nonVisited, i)
	}
	distance_min[depart] = 0 // Initialisation du poids du noeud de depart
	for true {
		var pointeur = min(distance_min, nonVisited)
		var index_pointeur = index(nonVisited, pointeur)
		nonVisited = append(nonVisited[:index_pointeur], nonVisited[(index_pointeur+1):]...) // retire le noeud de la liste "nonVisited"
		for v := range graph[pointeur] {
			if cherche(v, nonVisited) && (float64(int(distance_min[pointeur])+graph[pointeur][v]) < distance_min[v]) {
				precedent[v] = pointeur
				distance_min[v] = float64(int(distance_min[pointeur]) + graph[pointeur][v])
			}
		}
		if len(nonVisited) == 0 {
			break
		}
	}
	return precedent[arrive], distance_min[arrive]
}

//index prend en parametre un tableau de string et un element de ce tableau
//Cette fontion s'occupe de trouver l'index d'un element du tableau
//Retourne la valeur de l'index de l'element
func index(tableau []string, elem string) int {
	var k int = 0
	for i := range tableau {
		if tableau[i] == elem {
			k = i
		}
	}
	return k
}

//min prend en parametre un dictionnaire dont les clés sont des string et les valeurs sont des float et un tableau de string
//Cette fontion compare les poids de chaque noeud pour trouver le noeud de poids minimal parmi les noeuds non visités.
//Retourne le noeud de poids minimal
func min(dict map[string]float64, nvisited []string) string {
	var min float64 = math.Inf(1)
	var key = "Bonjour"
	for v := range nvisited {
		if dict[nvisited[v]] < min {
			min = dict[nvisited[v]]
			key = nvisited[v]
		}
	}
	return key
}

//cherche prend en parametre un noeud et un tableau de string
//Cette fontion s'occupe de verifier si le noeud et le tableau de string
//Retourne un booléen
func cherche(Noeud string, Tableau []string) bool {
	for i := range Tableau {
		if Tableau[i] == Noeud {
			return true
		}
	}
	return false
}

//recupfichier prend en parametre le chemin d'un fichier
//Cette fontion s'occupe d'ouvrir un fichier
//Retourne le fichier
func recupfichier(chemin string) *os.File {
	file, err := os.OpenFile(chemin, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

//parseur prend en parametre le fichier
//Cette fontion s'occupe de créer le graph sous la forme map[string]map[string]int
//Retourne le graph sous la forme map[string]map[string]int
func parseur(file *os.File) map[string]map[string]int {
	var graph = make(map[string]map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			var chemin = make(map[string]int)
			line := strings.Split(scanner.Text(), ": ")
			if len(line) < 2 {
				log.Fatal("Le fichier matrice n'est pas correcte")
			}
			if len(line[1]) < 7 {
				log.Fatal("Le fichier matrice n'est pas correcte")
			}
			line2 := strings.Split(strings.Split(strings.Split(line[1], "{ ")[1], " }")[0], ",")
			//fmt.Println(line[0])

			for i := 0; i < (len(line2)); i++ {
				//fmt.Println(line2[i],i)
				line3 := strings.Split(line2[i], ":")
				//fmt.Println(line3[0],line3[1])
				nbr, err := strconv.Atoi(line3[1])
				if err != nil {
					log.Fatal("Le cout du chemin n'est pas un nombre ")
				}
				chemin[line3[0]] = nbr
				graph[line[0]] = chemin
			}
			//fmt.Println("cc")
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	return graph
}

func main() {
	start := time.Now()
	// Code to measure
	if len(os.Args) == 1 {
		log.Fatal("Pas de fichier en parametre ")
	}
	chemin := os.Args[1]
	if testchemin(chemin) {
		file := recupfichier(chemin)
		graph := parseur(file)
		//fmt.Println(graph)
		//var r, c =dijkstra(graph, "A", "D")
		//fmt.Println(r,c)
		toutdijkstra(graph)
	}
	duration := time.Since(start)
	fmt.Println(duration)
}
