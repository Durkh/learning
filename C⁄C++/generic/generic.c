


#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct temp{
    void* dado;
    int indice;
    struct temp* next;
}pilhaGen;

int indice = 0;
pilhaGen* bottom = NULL;
pilhaGen* top = NULL;


typedef union{
    int x,y;
} Uteste;


void* Int2Void(const void* dado, const int dadoTam){

    int* buffer;
    buffer = (int*)malloc(dadoTam);
    *buffer = *((int*)dado);

    return (void*)buffer;
}

void* Float2Void(const void* dado, const int dadoTam){

    float* buffer;
    buffer = (float*)malloc(dadoTam);
    *buffer = *((float*)dado);

    return (void*)buffer;
}

void* Double2Void(const void* dado, const int dadoTam){

    double* buffer;
    buffer = (double*)malloc(dadoTam);
    *buffer = *((double*)dado);

    return (void*)buffer;
}

void* Char2Void(const void* dado, const int dadoTam){

    char* buffer;
    buffer = (char*)malloc(dadoTam);
    *buffer = *((char*)dado);

    return (void*)buffer;
}

void* Arr2Void(const void* dado, const int dadoTam){

    int* buffer;
    buffer = (int*)malloc(dadoTam);
    memcpy(buffer, dado, dadoTam);

    return (void*)buffer;
}

void* P2Void(const void* dado, const int dadoTam){

    int* buffer;
    buffer = (int*)malloc(dadoTam);
    buffer = *((int**)dado);

    return (void*)buffer;
}

void* Union2Void(const void* dado, const int dadoTam){

    Uteste* buffer;
    buffer = (Uteste*)malloc(dadoTam);
    *buffer = *((Uteste*)dado);

    return (void*)buffer;
}

void Push(void* dado, const int dadoTam, void* (*toVoid)(const void* , const int)) {

    if(bottom == NULL){
        bottom = (pilhaGen*) malloc(sizeof(pilhaGen));
        bottom->dado = toVoid(dado, dadoTam);
        bottom->indice = indice++;
        top = bottom;
    }else{
        top = top->next = (pilhaGen*) malloc(sizeof(pilhaGen));
        top->dado = toVoid(dado, dadoTam);
        top->indice = indice++;
    }

}

void pop(){

    free(top);
    top = bottom;
    indice--;

    while(top->next != NULL){
        top = top->next;
    }
}

int main(){

    printf("testes com:\n"
        "inteiro\n"
        "float\n"
        "double\n"
        "char\n"
        "ponteiro pra inteiro(aponta pro primeiro inteiro)\n"
        "array de inteiros(tam 5)\n"
        "union de inteiros\n"
        "TAMANHOS FIXOS PRA AJUDAR A PRINTAR PQ NINGUÉM É PERFEITO\n"
        "digite nessa ordem (n precisa digitar o ponteiro, ele é fixo)\n"
        "#PAZ\n\n");

    int inte;
    float flo;
    double doub;
    char carac;
    int* pInt = &inte;
    int arr[5];
    Uteste uni;

    scanf("%i %f %lf %c %d %d %d %d %d %d %d", &inte, &flo, &doub, &carac, &arr[0], &arr[1], &arr[2], &arr[3], &arr[4], &uni.x, &uni.y );
    puts("os 2 últimos são do Union, logo vai valer apenas o último número digitado");

    Push(&inte, sizeof(inte), Int2Void);
    Push(&flo, sizeof(flo), Float2Void);
    Push(&doub, sizeof(doub), Double2Void);
    Push(&carac, sizeof(carac), Char2Void);
    Push(&pInt, sizeof(pInt), P2Void);
    Push(arr, sizeof(arr), Arr2Void);
    Push(&uni, sizeof(uni), Union2Void);


    pilhaGen* ptr = bottom;

    int inte2;
    float flo2;
    double doub2;
    char carac2;
    int* pInt2;
    int arr2[5];
    Uteste* uni2;

    inte2 = *((int*)ptr->dado);
    ptr = ptr->next;
    flo2 = *((float*)ptr->dado);
    ptr = ptr->next;
    doub2 = *((double*)ptr->dado);
    ptr = ptr->next;
    carac2 = *((char*)ptr->dado);
    ptr = ptr->next;
    pInt2 = (int*)ptr->dado;
    ptr = ptr->next;
    memcpy(arr2, ptr->dado, sizeof(arr2));
    ptr = ptr->next;
    uni2 = (Uteste*)ptr->dado;

    ptr = bottom;

    printf("inteiro = %d\n", inte2);
    ptr = ptr->next;
    printf("float = %f\n", flo2);
    ptr = ptr->next;
    printf("double = %lf\n", doub2);
    ptr = ptr->next;
    printf("char = %c\n", carac2);
    ptr = ptr->next;
    printf("ponteiro que aponta pro primeiro int = %p || qnt= %i\n", (void*)pInt2, *pInt2);
    ptr = ptr->next;
    for(int i=0; i<5; i++){
            printf("arr[%i] = %i\n",i, arr2[i]);
    }
    ptr = ptr->next;
    printf("union x = %d\n", uni2->x);
    ptr = ptr->next;
    printf("union y = %d (o que vale)\n", uni2->y);

    return 0;
}
