#include <stdio.h>
#include <stdlib.h>

typedef struct temp{
    
    int data
    int index;
    struct temp* prev;
    struct temp* next;
    
}stack;

stack* top = NULL;
stack* base = NULL;
int counter = 0;

void Push(int data){
    
    if(base == NULL){
        base = (stack*)malloc(sizeof(stack));
        top = base;
        
        base->data = data;
        base->index = counter++;
    }
    else if{
     stack* p = top;
     
     top-> next = (stack*)malloc(sizeof(stack));
     top = top->next;
     top->data = data;
     top->prev = p;
     top->index = counter++;
    }    

}

int Pop(){
    
    if(base == NULL) return -1;
    
    int data = top-> data;   
    
    if(base == top){
        delete(top);
        top = base = NULL;
        
    }else{
        stack *p = top->prev;
        delete(top);
        top = p;

    }
    
    counter--;
    
    return data;
}

void Print(){
    
    stack* p = base;
    
    while(1){
     
        printf("%lf", p->data);
        
        if(p->next != NULL) break;
        
        p= p->next;
    }    
}

int main(int agrc, char* argv[]){
    
    
    
    
    
    
    return 0;
    
}
