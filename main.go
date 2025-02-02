package main

import (
	"io"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Dado struct {
	Nome      string
	Counter   int64
	Max       float64
	Min       float64
	Somatorio float64
}

func lerAqr(_CAMINHO_ string, dataChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(dataChan)

	file, err := os.Open(_CAMINHO_)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferSize := 64 * 1024 // 64kb
	buffer := make([]byte, bufferSize) 

	for {
		//faz uma leitura unica
		n, err := file.Read(buffer)
		if n > 0 {
			//coppia para o buffer por questao de integridade
			// o erro que estava dando era por causa disso
			dataCopy := make([]byte, n)
			copy(dataCopy, buffer[:n])
			dataChan <- dataCopy
		}

		if err != nil {
			if err == io.EOF {
				log.Println("Fim do arquivo:", err)
				break
			}
			log.Println("Erro ao ler arquivo:", err)
			break
		}
	}
}

func processarArq(dataChan chan []byte, wg *sync.WaitGroup, result map[string]Dado, cidadesOrdenadas *[]string) {
	defer wg.Done()
	var partialLine string // Armazena linhas incompletas entre chunks
	

	for chunk := range dataChan {
		lines := string(chunk)
		
		
		for _,chunk := range lines{
			
			switch chunk {
			case '\r':
				continue
			case '\n':
				partes := strings.Split(partialLine, ";")
				
				cidade := partes[0]
				valorStr := partes[1]
			
				// Tentar converter o valor para float
				valor, err := strconv.ParseFloat(valorStr, 64)
				if err != nil {
					fmt.Printf("Erro ao converter valor para float: %q, erro: %v\n", valorStr, err)
					log.Panic("erro na conversão de float")
					break
				}
			
				// Processar a cidade e o valor
				dado, jaTem := result[cidade]
				if jaTem {
					dado.Somatorio += valor
					dado.Counter++
					if dado.Max < valor {
						dado.Max = valor
					}
					if dado.Min > valor {
						dado.Min = valor
					}
				} else {
					(*cidadesOrdenadas) = append((*cidadesOrdenadas), cidade)
					dado = Dado{
						Nome:      cidade,
						Counter:   1,
						Max:       valor,
						Min:       valor,
						Somatorio: valor,
					}
				}
				result[cidade] = dado

				partialLine = ""


			default:
				partialLine += string(chunk)
			}
		}
	}
}

func main() {
	inicio := time.Now()
	var CAMINHO string = "./arqs/measurements.txt"
	observatoriosResult := make(map[string]Dado)
	var wg sync.WaitGroup
	canal := make(chan []byte,10)
	var cidades []string

	wg.Add(1)
	go lerAqr(CAMINHO, canal, &wg)

	wg.Add(1)
	go processarArq(canal, &wg, observatoriosResult, &cidades)

	// Aguarda as goroutines
	wg.Wait()

	sort.Strings(cidades)

	fmt.Println("|  CIDADE  |  MIN  |  MED  |  MAX  |")
	fmt.Println("|----------|-------|-------|-------|")
	for _, cidadeNome := range cidades {
		cidade := observatoriosResult[cidadeNome]
		fmt.Printf("| %s | %.1f | %.1f | %.1f |\n",
			cidade.Nome,
			cidade.Min,
			cidade.Somatorio/float64(cidade.Counter),
			cidade.Max)
	}

	fmt.Printf("\nTempo de execução: %s\n", time.Since(inicio))
}
