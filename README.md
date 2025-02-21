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
- A medição de tempo é feita pelo próprio programa comparando a hora que inicia e que termina

### versões 

### v0.1 
- Leitura do Arquivo:

    O arquivo é aberto e lido linha por linha usando um scanner.
    Cada linha é dividida em duas partes: o nome da cidade e o valor da temperatura (convertido para float64).

- Processamento dos Dados:

    Um mapa (observatorios) é usado para armazenar os dados de cada cidade.
    Para cada linha:
        Se a cidade já existe no mapa:
            O contador de medições é incrementado.
            Os valores de temperatura mínima e máxima são atualizados, se necessário.
            O somatório das temperaturas é incrementado.
        Se a cidade ainda não existe no mapa:
            Ela é adicionada ao mapa com os valores iniciais.
            O nome da cidade é adicionado a um slice (cidades) para manter a ordem.

- Ordenação:

    Após o processamento, o slice cidades é ordenado alfabeticamente para exibir os resultados em ordem.

- Exibição dos Resultados:

    Para cada cidade no slice ordenado:
        As estatísticas calculadas (mínima, média e máxima) são exibidas em formato tabular.
        O tempo total de execução é exibido ao final.

- Decisões:
    Usar um cicionário para agiizar o acesso ao dados e evitar ter que criar um novo objeto para cada linhas

- Tempo:
    Laptop: Após três execuções o tempo médio foi de: 3m36.1439s
    PC Após três execuções o tempo médio foi de: 2m31.8767

### v0.2 
- Mudanças:

  Uso de goroutines porém apenas uma para cada função
  
- Leitura do Arquivo:

  Uma go routine de leirura

- Processamento dos Dados:

    Uma go routine de processamento

- Tempo:
    Laptop: Após três execuções o tempo médio foi de: 12m20.2391s
    PC Após três execuções o tempo médio foi de: 8m31.4562s

### v0.2.1 
- Mudanças:

  Uso de goroutines com multíplos leitores
  
- Leitura do Arquivo:

  Uma go routine de leirura

- Processamento dos Dados:

    64 go routines de processamento

- Tempo:
    Laptop: Após três execuções o tempo médio foi de: a verificar
    PC Após três execuções o tempo médio foi de: 1m29.6823

  

