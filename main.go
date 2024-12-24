package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Dado struct{
	Nome string
	Counter int64
	Max float64
	Min float64
	Somatorio float64
}

func processarArquivo(_caminho string) (map[string]Dado, []string, error){
	file, err := os.Open(_caminho)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

	observatorios := make(map[string]Dado)
	// array com os nomes das cidades
	//depois ordena
	var cidades []string

	scanner := bufio.NewScanner(file);
	for scanner.Scan(){
		linha := scanner.Text()
		partes := strings.Split(linha,";")
		f,err := strconv.ParseFloat(partes[1],32)
		if err != nil{
			log.Panic(err)
		}

		//verifica se cidade já existe
		dado, jaTem := observatorios[partes[0]]
		if jaTem {
			dado.Counter = dado.Counter+1
			//faz verificações min max
			if dado.Max < f{
				dado.Max = f
			}
			if dado.Min > f{
				dado.Min = f
			}
			dado.Somatorio = dado.Somatorio + f
		} else {
			cidades = append(cidades, partes[0])
			dado = Dado{
				Nome: partes[0],
				Counter: 1,
				Max: f,
				Min: f,
				Somatorio: dado.Somatorio + f,
			}
		}
		observatorios[partes[0]] = dado
	}

	sort.Strings(cidades)
	return observatorios, cidades, scanner.Err()
}

func main(){
	inicio := time.Now()

	var caminho string = "./arqs/measurements1M.txt"
	// abre o arquivo 

	observatorios, cidades, err := processarArquivo(caminho)
	if err != nil{
		log.Panic(err)
	}

	

	println("|  CIDADE  |  MIN  |  MED  |  MAX  |")
	fmt.Println("|----------|-------|-------|-------|")


	for _, cidadeNome :=  range cidades{
		cidade := observatorios[cidadeNome]
		fmt.Printf("| %s | %.1fc | %.1fc | %.1fc |\n",
		cidade.Nome,
		cidade.Min,
		cidade.Somatorio/float64(cidade.Counter),
		cidade.Max)
	}


	duration := time.Since(inicio)
	//tempo de execução
	fmt.Printf("\nTempo de execução: %s\n", duration)
}