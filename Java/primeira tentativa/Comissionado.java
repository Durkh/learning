package com.example.Start;

public class Comissionado extends Funcionario {
    private double vendasSemanais;
    private double percentualComissao;

    public double getVendasSemanais() {
        return vendasSemanais;
    }

    public void setVendasSemanais(double vendasSemanais) {
        this.vendasSemanais = vendasSemanais;
    }

    public double getPercentualComissao() {
        return percentualComissao;
    }

    public void setPercentualComissao(double percentualComissao) {
        this.percentualComissao = percentualComissao;
    }

    public Comissionado(){
        vendasSemanais = 0;
        percentualComissao = 0;
    }

    @Override
    public double CalculaSalario(){
        return vendasSemanais *(percentualComissao /100);
    }
}
