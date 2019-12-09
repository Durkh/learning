#include <stdio.h>

int recursive(int n){
    
    if(n){
        n = recursive(n-1);
    }  
    
    return n;
}




int main(){
    
    printf("%d", recursive(10));    
        
    
    return 0;
}
