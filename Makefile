MiddleElementOfLinkedList: MiddleElementOfLinkedList.o
	g++ -g -o MiddleElementOfLinkedList.exe MiddleElementOfLinkedList.o -pthread    

MiddleElementOfLinkedList.o: MiddleElementOfLinkedList/MiddleElementOfLinkedList.cpp
	g++ -g  -c -pthread MiddleElementOfLinkedList/MiddleElementOfLinkedList.cpp
