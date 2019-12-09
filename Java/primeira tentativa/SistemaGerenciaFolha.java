package com.example.Start;

import java.util.ArrayList;

public class SistemaGerenciaFolha {
    private ArrayList<Funcionario> funcionario = new ArrayList();

    public void setFuncionario(Funcionario buffer){
        funcionario.add(buffer);
    }

    public double CalculaFolhaTotal(){
        double buffer=0;

        for(Funcionario i: funcionario) buffer += i.CalculaSalario();

        return buffer;
    }

    public double ConsultaSalarioFuncionario(String nome){
        for(Funcionario i: funcionario){
            if(nome.equals(i.getNome())) return i.CalculaSalario();
        }
        return 0;
    }

}

