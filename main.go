package main

import (
	// "container/list"
	"fmt"
	"strconv"
	"strings"
)

type Element struct {
	value     int64
	next_byte *Element
	next_elt  *Element
}

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
		/* Dans le cas inverse, on insére l'élement au bon endroit. */
	} else {
		elt := *head
		prev := *head
		for elt != nil && elt.value < ip_octect {
			prev = elt
			elt = elt.next_elt
		}
		if elt == nil {
			elt = new(Element)
			elt.value = ip_octect
			elt.next_elt = nil
			prev.next_elt = elt
			insert(&(elt.next_byte), ip[1:])
		} else if elt.value == ip_octect {
			insert(&elt.next_byte, ip[1:])
		} else if elt.value > ip_octect {
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

func main() {
	var ip_table_head *Element

	insert(&ip_table_head, strings.Split("192.168.33.44", "."))
	insert(&ip_table_head, strings.Split("192.168.55.66", "."))
	// insert(&ip_list, strings.Split("55.168.55.66", "."))
	insert(&ip_table_head, strings.Split("88.168.55.66", "."))
	fmt.Println(ip_table_head)
}
