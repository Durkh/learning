package com.example.Start;

public class Assalariado extends Funcionario{
    private double salario;

    public Assalariado(){
        salario = 0;
    }

    public double getSalario() {
        return salario;
    }

    public void setSalario(double salario) {
        this.salario = salario;
    }

    @Override
    public double CalculaSalario(){
        return salario;
    }
}
