Example about implement pointer lists with golang.\
The struct utilized for the axample is Persona:\

STRUCT\
type Persona struct {\
	Nome    string\
	Cognome string\
	Eta     int\
	Next    *Persona\
}\

DEFINED THE NEW LINKEDLIST TYPE OF THE PERSON STRUCTURE\
type LinkedList struct {\
	Head *Persona\
	Tail *Persona\
}\

I implemented the following functions:\
- func printList(list *LinkedList, emptyLine bool)
- func addHead(list *LinkedList, p Persona)
- func addTail(list *LinkedList, p Persona)
- func addInTheMiddle(list *LinkedList, p Persona, pos int)
