## Proposta
Implementar uma API com um endpoint `/encrypt` que receba uma string e a criptografe com o algoritmo de Cifra de César com um shift de +3 e um endpoint `/decrypt` que receba uma string criptografada e a retorne descriptografada.

### Cifra de César no nosso contexto
A Cifra de César é um dos algoritmos de criptogradia mais simples.  
O algoritmo troca cada letra de a-z do texto por outra letra do alfabeto, que se apresenta um número fixo de posições (deslocamento) mais a frente no mesmo alfabeto.  
Qualquer caractere que não seja uma letra minúscula não deve ser modificado.  
Por exemplo, a string "abc" com deslocamento de 3 posições se torna "def".  
A string "meli" com deslocamento de 3 posições se torna "phol".
