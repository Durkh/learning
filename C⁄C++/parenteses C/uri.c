#include <stdio.h>


int main(){

    char str[1001];
    
    while(scanf("%s", str) != EOF){
        
        int controle = 0;
        int i;
        
        for(i = 0; str[i]!= '\0'; i++){
            
            if(str[i] == '(')
                controle++;
            else if(str[i] == ')')
                controle--;
            
            if(controle < 0){
                puts("incorrect");
                break;
            }
        }
        
        if(controle > 0)
            puts("incorrect");
        else if(controle == 0)
            puts("correct");
    }
    
    return 0;
}
