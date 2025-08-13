package main

import "fmt"

type Persona struct {
	Nome    string
	Cognome string
	Eta     int
	Next    *Persona
}

type LinkedList struct {
	Head *Persona
	Tail *Persona
}

func main() {

	var listaPersone LinkedList

	var Persona1 = Persona{Nome: "Davide", Cognome: "Garau", Eta: 36}
	var Persona2 = Persona{Nome: "Mario", Cognome: "Rossi", Eta: 30}
	var Persona3 = Persona{Nome: "Luca", Cognome: "Bianchi", Eta: 25}
	var Persona4 = Persona{Nome: "Marco", Cognome: "Bianchi", Eta: 20}
	var Persona5 = Persona{Nome: "Daniele", Cognome: "Bianchi", Eta: 15}
	var Persona6 = Persona{Nome: "Lucia", Cognome: "Bianchi", Eta: 28}
	var Persona7 = Persona{Nome: "Uga", Cognome: "Suz", Eta: 31}
	var Persona8 = Persona{Nome: "Plerzo", Cognome: "Suz", Eta: 31}
	var Persona9 = Persona{Nome: "Cotroneo", Cognome: "Suz", Eta: 31}
	var Persona10 = Persona{Nome: "Vudro", Cognome: "Suz", Eta: 31}
	var Persona11 = Persona{Nome: "Zizza", Cognome: "Suz", Eta: 31}
	var Persona12 = Persona{Nome: "Fudo", Cognome: "Laz", Eta: 31}

	addInTheMiddle(&listaPersone, Persona12, 17)
	printList(&listaPersone, true)

	addInTheMiddle(&listaPersone, Persona9, 1)
	printList(&listaPersone, true)

	addInTheMiddle(&listaPersone, Persona10, 1)
	printList(&listaPersone, true)

	addInTheMiddle(&listaPersone, Persona11, 2)
	printList(&listaPersone, true)

	addTail(&listaPersone, Persona1)
	printList(&listaPersone, true)
	addHead(&listaPersone, Persona4)
	printList(&listaPersone, true)
	addTail(&listaPersone, Persona2)
	printList(&listaPersone, true)
	addHead(&listaPersone, Persona5)
	printList(&listaPersone, true)
	addTail(&listaPersone, Persona3)
	printList(&listaPersone, true)
	addHead(&listaPersone, Persona6)
	printList(&listaPersone, true)

	addInTheMiddle(&listaPersone, Persona7, 3)
	printList(&listaPersone, true)
	addInTheMiddle(&listaPersone, Persona7, 5)
	printList(&listaPersone, true)
	addInTheMiddle(&listaPersone, Persona7, 13)
	printList(&listaPersone, true)
	addInTheMiddle(&listaPersone, Persona8, 1)
	printList(&listaPersone, true)
}

func printList(list *LinkedList, emptyLine bool) {

	for p := list.Head; p != nil; p = p.Next {
		fmt.Printf("Nome: %s, Cognome: %s, Eta: %d\n", p.Nome, p.Cognome, p.Eta)
	}

	if emptyLine {
		fmt.Println()
	}
}

func addHead(list *LinkedList, p Persona) {

	if list.Head == nil {
		list.Head = &p
		list.Tail = &p
	} else {
		p.Next = list.Head
		list.Head = &p
	}
}

func addTail(list *LinkedList, p Persona) {

	if list.Head == nil {
		list.Head = &p
	} else {
		list.Tail.Next = &p
	}
	list.Tail = &p
}

func addInTheMiddle(list *LinkedList, p Persona, pos int) {

	if list.Head == nil {

		list.Head = &p
		list.Tail = &p
		return
	}

	var index int = 1
	var i *Persona = list.Head
	var x *Persona = nil

	for i = list.Head; i != nil && index < pos; i = i.Next {
		index++
		x = i
	}

	if i == list.Head {

		p.Next = list.Head
		list.Head = &p

	} else {

		x.Next = &p
		p.Next = i
	}

	if p.Next == nil {
		list.Tail = &p
	}
}
