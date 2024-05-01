package algo

import (
	"slices"
	"strings"
)

/*
Dado dos cadenas a y b:

	Si difieren en longitud, no es un anagrama.
	Si tienen la misma longitud, proseguir.
	Crear una lista con las letras de cada entrada.
	Ordenar alfabeticamente ambas listas.
	Si son las misma lista, es una anagrama.
	En caso contrario, no lo es.
*/
func SimpleDetection(a, b string) bool {
	// Si a y b difieren en longitud, no pueden ser anagramas.
	if len(a) != len(b) {
		return false
	}

	// Crear un par de listas de runas A y B cuyos elementos sean cada caracter de a y b respectivamente.
	A, B := []rune(a), []rune(b)

	// Ordenar cada lista
	slices.Sort(A)
	slices.Sort(B)

	// En caso de que A y B sean iguales, a y b son anagramas, por lo que se devuelve verdadero.
	// En caso de que A y B NO sean iguales, a y b NO son anagramas, por lo que se devuelve falso.
	//! Un par de `slices` no pueden ser comparadas utilizando el operador `==`, debe ser utilizado `slices.Equal([]T, []T)`.
	return slices.Equal(A, B)
}

// getPrimes crea y devuelve una lista con todos los numeros primos del 1 al 100.
// De esta manera no es necesario calcular los numeros primos necesarios para los algoritmos cada vez.
// Al ser una funciòn que crea la lista cada vez, esta lista no puede ser mutada globalmente, solo la copia devuelta.
func getPrimes() []int {
	return []int{
		2,
		3,
		5,
		7,
		11,
		13,
		17,
		19,
		23,
		29,
		31,
		37,
		41,
		43,
		47,
		53,
		59,
		61,
		67,
		71,
		73,
		79,
		83,
		89,
		97,
	}
}

/*
Dado un par de cadenas a y b:

	Si difieren en longitud, a y b NO son anagramas.
	En caso contrario, proseguir.
	Crear un slice A de caracteres con todos los caracteres de a.
	Eliminar las repeticiones en A.
	Crear dos variable resultA y resultB inicializados en 1.
	Por cada elemento v con indice i en A:


	Establecer resultA como si mismo multiplicado por el nùmero de veces que aparece v en a por el nùmero primo i.


		Establecer resultB como si mismo multiplicado por el nùmero de veces que aparece v en b por el numero primo i.
	Si resultA es igual a resultB, entonces a y b son anagramas.
	En caso contrario, a y b NO son anagramas.
*/
func PrimeFactorsDetection(a, b string) bool {
	// Si a y b difieren en longitud, no son anagramas.
	if len(a) != len(b) {
		return false
	}

	// Crear un slice A con todos los caracteres de a.
	A := []rune(a)

	// Eliminar repetidos
	slices.Sort(A)
	A = slices.Compact(A)

	// Inicializar resultA y resultB.
	resultA, resultB := 1, 1

	// Por cada caracter v con indice i en A:
	for i, v := range A {
		// resultA = resultA * Por las veces que v aparece en a * el nùmero primo nùmero i.
		resultA *= strings.Count(a, string(v)) * getPrimes()[i]
		// resultB = resultB * Por las veces que v aparece en b * el nùmero primo nùmero i.
		resultB *= strings.Count(b, string(v)) * getPrimes()[i]
	}

	// Si resultA es igual resultB entonces a y b son anagramas.
	// Si resultA NO es igual a resultB entonces a y b no son anagramas.
	return resultA == resultB
}

/*
Dado dos cadenas a y b:

	Si a y b difierien en longitud, entonces a y b NO son anagramas.
	En caso contrario, proseguir.
	Crear una lista A con los caracteres de a.
	Eliminar los caracteres repetidos en A.
	Por cada caracter v en A:


		Si la cantidad de veces que aparece v en a es distinta a la cantidad que aparece en b, entonces a y b NO son anagramas.
		Caso contrario proseguir.
	Si por cada elemento v no se concluyo que a y b no sean anagramas, entonces a y b son anagramas.
*/
func MapDetection(a, b string) bool {
	// Si difieren en longitud, entonces a y b NO son anagramas.
	if len(a) != len(b) {
		return false
	}

	// Crear una lista con los caracteres de a y b.
	A := []rune(a)

	// Eliminar los caracteres repetidos
	slices.Sort(A)
	A = slices.Compact(A)

	// Por cada elemento v en A:
	for _, v := range slices.Compact(A) {
		// Si la cantidad de apariciones de v en a y b difieren, entonces a y b NO son anagramas.
		if strings.Count(a, string(v)) != strings.Count(b, string(v)) {
			return false
		}
	}

	// Si se completo el bucle sin problemas, entonces a y b son anagramas.
	return true
}
