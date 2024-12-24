# 1-BillionRowChallenge

### Aqui vou apresentar o processor de criar e otimizar um código em go para o 1BCR que consiste em processar um arquivo com 1 bihao de linhas da forma mais rápida possível as regras e o arquivo com as linhas pode ser gerado com o código do [GitHub do desafio](https://github.com/AlexanderYastrebov/1brc#submitting)  

### O objetivo é escrever um programa que recupere os valores de medições de temperatura de um arquivo de texto e calcule a temperatura mínima, média e máxima por estação meteorológica. Há apenas um detalhe: o arquivo contém 1.000.000.000 de linhas.
### Regras:
- Nenhuma dependência de bibliotecas externas pode ser usada. Isso significa nada de lodash, numpy, Boost, ou qualquer outra biblioteca. Você está limitado à biblioteca padrão da sua linguagem.

- As implementações devem ser fornecidas como um único arquivo de código-fonte. Tente mantê-lo relativamente curto; não copie e cole uma biblioteca inteira na sua solução como forma de burlar a regra.

- A computação deve ocorrer em tempo de execução da aplicação; você não pode processar o arquivo de medições durante o tempo de compilação.

- Os intervalos dos valores de entrada são os seguintes:
  - Nome da estação: string UTF-8 não nula com comprimento mínimo de 1 caractere e máximo de 100 bytes (ou seja, pode ser 100 caracteres de 1 byte, 50 caracteres de 2 bytes, etc.).
  - Valor de temperatura: número do tipo double não nulo, entre -99,9 (inclusivo) e 99,9 (inclusivo), sempre com uma casa decimal.

- Há um máximo de 10.000 nomes de estações únicos.

- As implementações não devem depender de especificidades de um conjunto de dados específico. Qualquer nome de estação válido, conforme as restrições acima, e qualquer distribuição de dados (número de medições por estação) devem ser suportados.

### Sobre esse repositório:
- Não uparei no git do challenge
- Gerei quatro arquivos com mil, 1 milhão, 100 milhoẽs e 1 bilhão respectivamente com o objetivo de facilitar os testes de pequenas mudanças.
- Os tempos mostrados são apenas do arquivo de 1 bilhão
- Os códigos serão versionados entre as mudças(só porque quero testar tags)
- Serão dois teste um no meu laptop e outra no meu computador

### versões
