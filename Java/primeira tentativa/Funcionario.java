package com.example.Start;

public class Funcionario {
    protected String nome;
    protected int matricula;

    public String getNome() {
        return nome;
    }

    public void setNome(String nome) {
        this.nome = nome;
    }

    public int getMatricula() {
        return matricula;
    }

    public void setMatricula(int matricula) {
        this.matricula = matricula;
    }

    public Funcionario(){
        this.nome = "default";
        this.matricula = 0;
    }

    public double CalculaSalario(){
        return 0;
    }
}
