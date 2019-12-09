#ifndef _E_STACK__
#define _E_STACK__

 // bibilioteca de pilha duplamente encadeada, feita por Egídio Neto
 // implementação da pilha com dado genérico
 // possui índice


typedef struct temp{
    
    double data
    int index;
    struct temp* prev;
    struct temp* next;
    
}stack;

void Push(double data);
double Pop();
void Print();

#endif
