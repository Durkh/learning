package com.example.Start;

import java.util.Scanner;

public class Start {
    public static void main(String[] args){

        Assalariado assa = new Assalariado();
        SistemaGerenciaFolha sistema = new SistemaGerenciaFolha();

        assa.setSalario(1500);
        assa.setNome("Egídio");
        assa.setMatricula(41454);

        sistema.setFuncionario(assa);

        Horista hor = new Horista();

        hor.setHorasTrabalhadas(100);
        hor.setSalarioPorHora(3.50);
        hor.setMatricula(58574874);
        hor.setNome("Rondon");

        sistema.setFuncionario(hor);

        Comissionado comi = new Comissionado();

        comi.setPercentualComissao(30);
        comi.setVendasSemanais(3000);
        comi.setMatricula(6477834);
        comi.setNome("Eghone");

        sistema.setFuncionario(comi);

        System.out.println("Salario total cadastrado");
        System.out.println(sistema.CalculaFolhaTotal());

        System.out.println("digite o nome de um Funcionario: ");
        Scanner scanner = new Scanner(System.in);

        String nome = scanner.next();
        double salario = sistema.ConsultaSalarioFuncionario(nome);
        if(salario == 0){
            System.out.println("funcionario não existente");
        }else{
            System.out.println(salario);
        }


    }
}
