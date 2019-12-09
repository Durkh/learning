package com.example.Start;

public class Horista extends Funcionario{
    private double salarioPorHora;
    private double horasTrabalhadas;

    public double getSalarioPorHora() {
        return salarioPorHora;
    }

    public void setSalarioPorHora(double salarioPorHora) {
        this.salarioPorHora = salarioPorHora;
    }

    public double getHorasTrabalhadas() {
        return horasTrabalhadas;
    }

    public void setHorasTrabalhadas(double horasTrabalhadas) {
        this.horasTrabalhadas = horasTrabalhadas;
    }

    public Horista(){
        salarioPorHora = 0;
        horasTrabalhadas = 0;
    }

    @Override
    public double CalculaSalario(){
        return salarioPorHora * horasTrabalhadas;
    }
}
