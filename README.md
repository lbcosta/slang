# S-Lang
| Instrução             | Sintaxe             | Descrição                                      |
|-----------------------|---------------------|------------------------------------------------|
| *Increment*           | `V <- V + 1`        | Incrementa o valor da variável V em 1          |
| *Decrement*           | `V <- V - 1`        | Decrementa o valor da variável V em 1          |
| *Conditional Branch*  | `IF V != 0 GOTO L`  | Se V for diferente de 0, desvia para o rótulo L|

- "S" é a linguagem usada por Martin D. Davis no seu livro *Computability, Complexity, and Languages: Fundamentals of Theoretical Computer Science*, 2ª ed.

# Regras da linguagem:
- Valores são sempre inteiros não negativos;
- Só existem as 3 instruções acima;
- A variável de saída **Y** sempre existe e sempre inicia com o valor 0;
- As variáveis locais **Z** iniciam com o valor 0 quando são usadas;

## Labels (rótulos)
- Instruções podem conter um *label* escrito entre []. Exemplo: `[B] Z <- Z - 1`
- Se uma instrução for do tipo `IF V != 0 GOTO L` e **L** não existir, o programa para.


# Como buildar e executar

## Build (compilar)

Requer Go instalado (versão >= 1.18).

Para compilar o executável na pasta `bin`:

```sh
make build
# ou manualmente:
go build -o bin/slang ./src/main.go
```

## Executando

Você pode passar argumentos iniciais para variáveis, por exemplo:

```sh
./bin/slang ./programs/example.slang x=5
```

## Exemplo de saída

```
2025/09/26 14:29:32 Counter  Instruction          X          Y         
2025/09/26 14:29:32 0        X <- X - 1           5          0         
2025/09/26 14:29:32 1        Y <- Y + 1           4          0         
2025/09/26 14:29:32 2        IF X != 0 GOTO A     4          1         
```

# Próximas implementações:
[] Macros
    [] Change Program to be recursive (Program P1 can call P1.Run(...P2.Run()...))

