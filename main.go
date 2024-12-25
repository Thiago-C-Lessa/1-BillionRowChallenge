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
	"sync"
)

type Dado struct{
	Nome string
	Counter int64
	Max float64
	Min float64
	Somatorio float64
}


func lerAqr(_CAMINHO_ string, dataChan chan string, wg *sync.WaitGroup){
	defer wg.Done()

	file, err := os.Open(_CAMINHO_)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close()

	scanner := bufio.NewScanner(file);
	for scanner.Scan(){
		linha := scanner.Text()
		dataChan <- linha
	}
	close(dataChan)
}

//maps são passados por referência não precisa de ponteiros
func processarArq(dataChan chan string, wg *sync.WaitGroup, result map[string]Dado, cidadesOrdenadas *[]string, /*mu *sync.Mutex, contador *int64*/){
	defer wg.Done()
	// array com os nomes das cidades
	//depois ordena
	//mu.Lock()
	for linha := range dataChan{
		partes := strings.Split(linha,";")
		f,err := strconv.ParseFloat(partes[1],32)
		if err != nil{
			log.Panic(err)
		}
		//verifica se cidade já existe
		dado, jaTem := result[partes[0]]
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
			(*cidadesOrdenadas) = append((*cidadesOrdenadas), partes[0])
			dado = Dado{
				Nome: partes[0],
				Counter: 1,
				Max: f,
				Min: f,
				Somatorio: dado.Somatorio + f,
			}
		}
		result[partes[0]] = dado
		//(*contador) = (*contador)+1
		//fmt.Println((*contador))
		//mu.Unlock()
	}


	
}

func main(){
	inicio := time.Now()
	var CAMINHO string = "./arqs/measurements1B.txt"
	observatoriosResult := make(map[string]Dado)
	var wg sync.WaitGroup
	//var mu sync.Mutex
	canal := make(chan string,100)
	var cidades []string
	//var quantos int64 = 0

	wg.Add(1)
	go lerAqr(CAMINHO,canal,&wg)

	wg.Add(1)
	go processarArq(canal,&wg,observatoriosResult,&cidades/*,mu,&quantos*/)

// faz com que espere as routines
	wg.Wait()
		
	sort.Strings(cidades)

	println("|  CIDADE  |  MIN  |  MED  |  MAX  |")
	fmt.Println("|----------|-------|-------|-------|")


	for _, cidadeNome :=  range cidades{
		cidade := observatoriosResult[cidadeNome]
		fmt.Printf("| %s | %.1fc | %.1fc | %.1fc |\n",
		cidade.Nome,
		cidade.Min,
		cidade.Somatorio/float64(cidade.Counter),
		cidade.Max)
	}


	//tempo de execução
	fmt.Printf("\nTempo de execução: %s\n", time.Since(inicio)/*,quantos*/)
}