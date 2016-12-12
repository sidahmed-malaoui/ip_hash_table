# Utilisation

On commenc par générer une liste d'@IP aléatoire grace au script random_ip_generator.py :

```python3 random_ip_generator.py 10000 > liste_ip_test.txt```

On copie le contenu de liste_ip_test.txt et on lance après le program principal qui est main.go :

```go run main.go```

Dés son lancement, il nous demande d'entrer une liste d'@IP à inserer dans notre blacklist (chaque @IP dans une ligne), et 
finir l'insertion par une ligne vide.

Après avoir insérer, il nous donne la main pour rechercher dans notre black list. On lui donne une @IP et il nous dit si elle
est dans notre blacklist ou pas (On finit la recherche par une ligne vide).

# Remarques
Il n'y a aucune limite sur le nombre d'@IP qu'on peut insérer.

L'accès et l'insertion dans notre black list est toujours O(1) (au pire cas c'est 1000 instruction, mais ça reste toujours fixe
même si on insére tous les @IPv4 qui existe dans le monde)
