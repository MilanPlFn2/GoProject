# GoProject

La fonction Dijkstra.go prend en parametre un fichier de matrice d'adjacence, de syntaxe
>String: { String:int } ou String: { String:int,String:int,... }
    
en respectant les espaces (Le fichier matrice fournit sert de fichier test donc est fonctionnel)

La fonction Dijkstra.go renvoie dans sa sortie : une liste de Dijkstra de tous les sommets de la matrice d'adjacence de type
>| Sommet de Départ | Sommet d'Arrivée | Sommet précédent | Cout du Chemin le Plus Court |
> Temps d'execution


Le fichier Dijkstra.go utilise des goroutines pour chaque sommet de depart et chaque sommet d'arrivee : c'est l'implementation optimale.

Le dossier test contient deux autres implementations de Dijkstra, moins rapides :
- Dijkstra-goroutine-sommet.go, qui utilise des goroutines juste pour le sommet de depart (une goroutine traite un sommet de depart, et tout ses sommets d'arrivee)
- Dijkstra-sans-goroutine.go, qui n'implemente pas les goroutines.

Temps requis pour executer le programme sur 253 sommets :
- Dijkstra.go => 5.61s
- Dijkstra-goroutine-sommet.go => 7.6s
- Dijkstra-sans-goroutine => 33.2s

Test realise sur un Ryzen 7 3700x (8cores/16threads) 4 Ghz

