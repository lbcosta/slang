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

# Próximas implementações:
[] Macros

