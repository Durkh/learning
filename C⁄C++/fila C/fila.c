
#include <stdio.h> 
#include <stdlib.h>


typedef struct temp{
    
    int Data;
    struct temp* pointer;
    
} Fila;

Fila* beginning = NULL;
Fila* ending = NULL;
Fila* pointer = NULL;

int head = 0 , tail = 0;

void Queue(int Data){
        
    if(beginning == NULL){
        ending = beginning = (Fila*)malloc(sizeof(Fila));
        tail = 1;
    }else{
        ending = ending->pointer = (Fila*)malloc(sizeof(Fila));
        tail++;
    }
    ending->Data = Data;
}

void ShowQueue(){
    int x = 1;
    pointer = beginning;
    
    if(head == tail) puts("Fila vazia");
    
    while(pointer->pointer != NULL){
        printf("%d. %d\n", x++, pointer->Data);
        pointer = pointer-> pointer;
    }    
}

void DeQueue(){
    
    if(beginning == NULL) 
        return;
    else if(beginning == ending){
        free(beginning);
        pointer = ending = beginning = NULL;
    }else{
        pointer = beginning->pointer;
        free(beginning);
        beginning = pointer;
    }
    head++;
}

int main(int argc, char* argv[]){
    
        int data;
        
        for(int i=0; i<10; i++){
            Queue(i*3);
        }
        
        ShowQueue();
        
        puts("======================================");
        DeQueue();
        
        ShowQueue();
        puts("======================================");
        
        DeQueue();
        
        ShowQueue();
        puts("======================================");
        
        
    return 0;
    
}
