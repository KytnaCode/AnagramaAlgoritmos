# Anagrama Algoritmo

Este repositorio agrupa tres algoritmos para detectar si un par de cadenas a y b son anagramas o no.

## Uso

### Requerimientos

* GNU Make
* Go >=1.22.2 

Para clonar el repostiorio ejecutar:

```bash
git clone https://github.com/KytnaCode/AnagramaAlgoritmos.git
make init
```

Utilizando el siguiente commando se ejecutaran las prueba unitarias y los benchmarks para cada algoritmo,
esto generara dos archivos, `test.log` y `bench.log` con la informaciòn de las pruebas y los benchmarks
respectivamente.

```bash
make all
```

el siguiente comando es equivalente

```bash
make
```

si solo deseas ejecutar las pruebas:

```bash
make test
```

para unicamente ejecutar los benchmarks:

```bash
make bench
```

## Algoritmos

### Simple Detection

#### Lenguaje Natural

`SimpleDetection` convierte las entradas *a* y *b* en listas de caracteres, luego las ordena y comprueba su igualdad,
si son iguales, *a* y *b* son anagramas, en caso contrario no lo son.

#### Descripciòn

```
Dado un par de cadenas a y b:

	Si difieren en longitud, a y b no son anagramas.
	En caso contrario, proseguir.
	Crear un slice A de caracteres con todos los caracteres de a.
	Eliminar las repeticiones en A.
	Crear dos variable resulta y resultB inicializados en 1.
	Por cada elemento v con indice i en A:
	    establecer resultA como si mismo multiplicado por el nùmero de veces que aparece v en a por el nùmero primo i.
		establecer resultB como si mismo multiplicado por el nùmero de veces que aparece v en b por el numero primo i.
	si resultA es igual a resultB, entonces a y b son anagramas.
	en caso contrario, a y b no son anagramas.
```

### Prime Factor Detection

#### Lenguaje Natural

`PrimeFactorDetection` convierte las entradas a y b en un producto de primos, asignando un primo distinto a cada
caracter distinto, si los productos obtenidos con a y b son iguales, a y b son anagramas, en caso contrario
no lo son.

#### Descripciòn

```
Dado un par de cadenas a y b:

	Si difieren en longitud, a y b no son anagramas.
	En caso contrario, proseguir.
	Crear un slice A de caracteres con todos los caracteres de a.
	Eliminar las repeticiones en A.
	Crear dos variable resulta y resultb inicializados en 1.
	Por cada elemento v con indice i en A:
	    Establecer resultA como si mismo multiplicado por el nùmero de veces que aparece v en a por el nùmero primo i.
        Establecer resultB como si mismo multiplicado por el nùmero de veces que aparece v en b por el numero primo i.
	Si resultA es igual a resultB, entonces a y b son anagramas.
	En caso contrario, a y b no son anagramas.
```

### Map Detection

#### Lenguaje Natural

`MapDetection` busca que existian la misma cantidad de cada caracter en las entradas a y b, si difiere an algùn caracter, entonces
no son anagramas, en caso contrario, lo son.

#### Descripciòn

```
Dado dos cadenas a y b:

	Si a y b difierien en longitud, entonces a y b NO son anagramas.
	En caso contrario, proseguir.
	Crear una lista A con los caracteres de a.
	Eliminar los caracteres repetidos en A.
	Por cada caracter v en A:

		Si la cantidad de veces que aparece v en a es distinta a la cantidad que aparece en b, entonces a y b NO son anagramas.
		Caso contrario proseguir.
	Si por cada elemento v no se concluyo que a y b no sean anagramas, entonces a y b son anagramas.
```

### Pruebas y Benchmarks

Resultados de las pruebas hechas con Go 1.22.2 el 30 de Abril de 2024:

test.log: 

```
=== RUN   TestSimpleDetection
--- PASS: TestSimpleDetection (0.00s)
=== RUN   TestPrimeFactorsDetection
--- PASS: TestPrimeFactorsDetection (0.00s)
=== RUN   TestMapDetection
--- PASS: TestMapDetection (0.00s)
PASS
ok  	Anagrama/algo	0.002s
```

bench.log:

```
goos: linux
goarch: amd64
pkg: Anagrama/algo
cpu: Intel(R) Core(TM) i5-6500 CPU @ 3.20GHz
BenchmarkControlCase-4            	1000000000	         0.3066 ns/op	       0 B/op	       0 allocs/op
BenchmarkSimpleDetection-4        	 5352294	       223.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkPrimeFactorDetection-4   	 1993088	       592.6 ns/op	       0 B/op	       0 allocs/op
BenchmarkMapDetection-4           	 2896156	       415.4 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	Anagrama/algo	5.179s
```

### Creditos

Estos algoritmos fueron hechos por *Alejandro Paz Gòmez*, o *KytnaCode* en github.
Este trabajo esta bajo THE UNLICENSE, eres libre de copiar, modificar, publicar, usar, compilar, vender, o distribuir este software. 
Màs informaciòn en UNLICENSE.txt
