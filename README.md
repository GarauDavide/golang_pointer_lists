Example about implement pointer lists with golang. \
The struct utilized for the axample is Persona:

Struct
type Persona struct {
	Nome    string
	Cognome string
	Eta     int
	Next    *Persona
}

defined the new LinkedList type of the Person structure
type LinkedList struct {
	Head *Persona
	Tail *Persona
}

I have implement functions:
- func printList(list *LinkedList, emptyLine bool)
- func addHead(list *LinkedList, p Persona)
- func addTail(list *LinkedList, p Persona)
- func addInTheMiddle(list *LinkedList, p Persona, pos int)
