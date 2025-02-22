/*
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

func readAqr(_CAMINHO_ string, dataChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(dataChan)
	var counter int64 = 0;

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
			counter++
			//coppia para o buffer por questao de integridade
			// o erro que estava dando era por causa disso
			dataCopy := make([]byte, n)
			copy(dataCopy, buffer[:n])
			dataChan <- dataCopy

			if len(dataChan) == cap(dataChan){
				fmt.Printf("Buffer Cheio, na leitura número: %v\n",counter)
			}
			switch{

			}
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

func processArq(dataChan chan []byte, wg *sync.WaitGroup, result map[string]Dado, cidadesOrdenadas *[]string) {
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
	var CAMINHO string = "E:/Dev/Go/arqs/measurements1B.txt"
	observatoriosResult := make(map[string]Dado)
	var wg sync.WaitGroup
	canal := make(chan []byte,100)
	var cidades []string

	wg.Add(1)
	go readAqr(CAMINHO, canal, &wg)

	wg.Add(1)
	go processArq(canal, &wg, observatoriosResult, &cidades)

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
*/

package main

import (
	"bytes"
	"fmt"
	"io"
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

const (
	NUM_WORKERS = 128        // Número de threads de processamento
	BUFFER_SIZE = 64 * 1024 * 1024 // 64 MB por leitura
)

func readAqr(_CAMINHO_ string, dataChan chan []byte, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(dataChan)
	//defer close(dataChan)

	file, err := os.Open(_CAMINHO_)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bufferSize := 64 * 1024 // 64kb
	buffer := make([]byte, bufferSize)
	var leftOver []byte

	for {
		//faz uma leitura unica
		n, err := file.Read(buffer)
		if n > 0 {
			leftOver = append(leftOver, buffer...)

			i := bytes.LastIndexByte(leftOver, '\n')

			if i == -1 {
				continue // Se não encontrou '\n', continue lendo mais dados
			}

			dataCopy := make([]byte, i) //make([]byte, n)
			copy(dataCopy, leftOver[:i])
			dataChan <- dataCopy				// o erro que estava dando era por causa disso

			leftOver = append([]byte{}, leftOver[i:]...)

			//if len(dataChan) == cap(dataChan){
			//	fmt.Printf("Buffer Cheio, na leitura número: %v\n",counter)
			//}

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

func processArq(dataChan chan []byte, wg *sync.WaitGroup, resultChan chan map[string]Dado) {
	defer wg.Done()
	result := make(map[string]Dado)

	for chunk := range dataChan {
		lines := string(chunk)
		lines = strings.ReplaceAll(lines, "\r", "")
		strs := strings.Split(lines, "\n")

		for _, lins := range strs {
			partes := strings.Split(lins, ";")
			if len(partes) < 2 {
				continue
			}

			cidade := partes[0]
			valorStr := partes[1]

			valor, err := strconv.ParseFloat(valorStr, 64)
			if err != nil {
				log.Printf("Erro ao converter valor: %q, erro: %v\n", valorStr, err)
				continue
			}

			dado, existe := result[cidade]
			if existe {
				dado.Somatorio += valor
				dado.Counter++
				if dado.Max < valor {
					dado.Max = valor
				}
				if dado.Min > valor {
					dado.Min = valor
				}
			} else {
				dado = Dado{
					Nome:      cidade,
					Counter:   1,
					Max:       valor,
					Min:       valor,
					Somatorio: valor,
				}
			}
			result[cidade] = dado
		}
	}
	resultChan <- result
}

func mergeResults(results []map[string]Dado) map[string]Dado {
	finalResult := make(map[string]Dado)
	for _, result := range results {
		for cidade, dado := range result {
			if existente, found := finalResult[cidade]; found {
				existente.Somatorio += dado.Somatorio
				existente.Counter += dado.Counter
				if existente.Max < dado.Max {
					existente.Max = dado.Max
				}
				if existente.Min > dado.Min {
					existente.Min = dado.Min
				}
				finalResult[cidade] = existente
			} else {
				finalResult[cidade] = dado
			}
		}
	}
	return finalResult
}

func main() {
	inicio := time.Now()
	CAMINHO := "E:/Dev/Go/arqs/measurements1B.txt"

	dataChan := make(chan []byte, 100)
	resultChan := make(chan map[string]Dado, NUM_WORKERS)

	var wg sync.WaitGroup

	wg.Add(1)
	go readAqr(CAMINHO, dataChan, &wg) // uma leitura já é bem rápidinho, em geral o canal fica cheio pq não foi consumido

	// Criando múltiplas goroutines para processamento
	for i := 0; i < NUM_WORKERS; i++ {
		wg.Add(1)
		go processArq(dataChan, &wg, resultChan)
	}

	// Fechando canal de leitura após todas as goroutines terminarem
	go func() {
		wg.Wait()
		//close(dataChan) para evitar erro de dar close num canal já fechado
		close(resultChan)
	}()

	// Coletando resultados de todas as goroutines
	var results []map[string]Dado
	for result := range resultChan {
		results = append(results, result)
	}

	// Mesclando resultados
	observatoriosResult := mergeResults(results)

	// Ordenando cidades
	var cidades []string
	for cidade := range observatoriosResult {
		cidades = append(cidades, cidade)
	}
	sort.Strings(cidades)

	// Exibindo resultados
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
