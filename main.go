package main

import (
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

type Element struct {
	value     int64
	next_byte *Element
	next_elt  *Element
}


/* Cette fonction affiche les adresses IP de notre black list. */
func print(elm *Element, stack [4]int64, byte_pos int) {
	for elm != nil {
		/* Si on est dans le dernier byte de notre adresse IP. */
		if elm.next_byte == nil {
			fmt.Print(stack[0], ".", stack[1], ".", stack[2], ".", elm.value)
			fmt.Println()
		/* Sinon on met la valeur du byte actuel de l'@IP, et on rappel à nouveau la fonction print
		   avec comme élément le prochain byte de l'@IP. */
		} else {
			stack[byte_pos] = elm.value
			print(elm.next_byte, stack, byte_pos+1)
		}
		elm = elm.next_elt
	}
}

/* Cette fonction insére une nouvelle adresse IP dans notre black list. */
func insert(head **Element, ip []string) {
	/* Condition d'arrêt de la fonction récursive. On s'arrête quand on a finit d'insérer tous les octets de
	   notre adresse IP. */
	if len(ip) == 0 {
		return
	}
	ip_octect, _ := strconv.ParseInt(ip[0], 10, 64)

	/* Dans le cas où la liste de cette octet d'adresse IP est déjà vide. */
	if *head == nil {
		*head = new(Element)
		(*head).value = ip_octect
		(*head).next_elt = nil
		insert(&((*head).next_byte), ip[1:])

	/* Dans le cas inverse, on insére l'élement au bon endroit. L'insertion est ordonnée.*/
	} else {
		elt := *head
		prev := *head
		for elt != nil && elt.value < ip_octect {
			prev = elt
			elt = elt.next_elt
		}
		/* Si on a atteint la fin de la liste, alors on crée un nouveau élément et on l'insére. */
		if elt == nil {
			elt = new(Element)
			elt.value = ip_octect
			elt.next_elt = nil
			prev.next_elt = elt
			insert(&(elt.next_byte), ip[1:])
		/* Sinon si l'ectet existe déjà alors on fait un appel récursive pour insérer le prochain octet. */
		} else if elt.value == ip_octect {
			insert(&elt.next_byte, ip[1:])
		/* Sinon si on a trouver le bon endroit pour insérer l'octet actuel, alors on l'insére et on
		   fait un appel récursive pour insérer l'octet suivant. */
		} else if elt.value > ip_octect {
			/* Cette condition gére le cas ou on veut insérer dans la tête de la liste. */
			if prev == elt {
				prev = new(Element)
				prev.value = ip_octect
				prev.next_elt = elt
				*head = prev
				insert(&(prev.next_byte), ip[1:])
			} else {
				tmp := new(Element)
				tmp.value = ip_octect
				tmp.next_elt = elt
				prev.next_elt = tmp
				insert(&(tmp.next_byte), ip[1:])
			}
		}
	}

}


/* Cette fonction cherche l'existance de L'@IP dans la blacklist. */
func search(ip_black_list *Element, ip []string) bool {
	var e *Element
	/* Condition d'arrêt de notre fonction récursive. Si on a trouver tous les octets de notre adresse IP
	   alors on s'arrête et on retourne true. */
    if len(ip) == 0 {
    	return true
    }

    for e = ip_black_list; e != nil; e = e.next_elt {
    	ip_octet, _ := strconv.ParseInt(ip[0], 10, 64)
    	if e.value == ip_octet {
    		break
    	}
    }

    /* Si on n'a pas trouvé l'actuel alors on retourne false. Sinon on continue de chercher les autres 
       octets de l'@IP .*/
    if e == nil {
    	return false
    } else {
    	return search(e.next_byte, ip[1:])
    }
}

func main() {
	var ip_black_list *Element
	var ip string
	n := 1

	fmt.Println("Insérez les @IP (une par ligne) et finissez l'insertion par une ligne vide :")
	for n != 0 {
		n, _ = fmt.Scanln(&ip)
		if matches, _ := regexp.MatchString("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}$", ip); !matches {
			fmt.Println("Ceci n'est pas une @IP valide")
		} else {
			insert(&ip_black_list, strings.Split(ip, "."))
		}
	}

	fmt.Println("Les adresses IP de notre black list")
	print(ip_black_list, [4]int64{0}, 0)

	fmt.Println("Entrez les @IP à chercher dans la blacklist (une par ligne) :")
	n = 1
	for n != 0 {
		n, _ = fmt.Scanln(&ip)
		if matches, _ := regexp.MatchString("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}$", ip); !matches {
			fmt.Println("Ceci n'est pas une @IP valide")
			continue
		}

		if exists := search(ip_black_list, strings.Split(ip, ".")); exists {
			fmt.Println("L'@IP est dans notre black list")
		} else {
			fmt.Println("L'@IP n'est pas dans notre black list")
		}
	}
}
