Example about implement pointer lists with golang.\
The struct utilized for the axample is Persona:

STRUCT\
type Persona struct {\
	\tNome    string\
	\tCognome string\
	\tEta     int\
	\tNext    *Persona\
}

DEFINED THE NEW LINKEDLIST TYPE OF THE PERSON STRUCTURE\
type LinkedList struct {\
	\tHead *Persona\
	\tTail *Persona\
}

I implemented the following functions:
- func printList(list *LinkedList, emptyLine bool)
- func addHead(list *LinkedList, p Persona)
- func addTail(list *LinkedList, p Persona)
- func addInTheMiddle(list *LinkedList, p Persona, pos int)
