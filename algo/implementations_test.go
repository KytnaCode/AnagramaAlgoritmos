package algo_test

import (
	"Anagrama/algo"
	"testing"
)

// TestData representa un caso de prueba.
type TestData struct {
	A        string // TestData.A representa la entrada a.
	B        string // TestData.B representa la entrada b.
	Expected bool   // TestData.Expected es la salida esperada.
}

// Casos de prueba, cada algoritmo deberà devolver el valor esperado en cada caso.
var testData = []TestData{
	// Casos solicitados:
	{"imperdonablemente", "imponderablemente", true},
	{"bulliciosamente", "escabullimiento", true},
	{"roca", "coro", false},
	{"enfriamiento", "refinamientos", false},

	// Casos extra:
	{"hola", "ola", false},
	{"aguacate", "gucteaaa", true},
}

// error da un test por fallido con informaciòn acerca del fallo
func error(t *testing.T, algorithmName string, result bool, data TestData) {
	t.Helper()

	t.Errorf(
		"algoritmo: `%v` no ha devuelto el resultado esperado: a: %v b: %v resultado: %v esperado: %v",
		algorithmName,
		data.A,
		data.B,
		result,
		data.Expected,
	)
}

func TestSimpleDetection(t *testing.T) {
	for _, data := range testData {
		if result := algo.SimpleDetection(data.A, data.B); result != data.Expected {
			error(t, "SimpleDetection", result, data)
		}
	}
}

func TestPrimeFactorsDetection(t *testing.T) {
	for _, data := range testData {
		if result := algo.PrimeFactorsDetection(data.A, data.B); result != data.Expected {
			error(t, "PrimeFactorsDetection", result, data)
		}
	}
}

func TestMapDetection(t *testing.T) {
	for _, data := range testData {
		if result := algo.MapDetection(data.A, data.B); result != data.Expected {
			error(t, "MapDetection", result, data)
		}
	}
}

// Caso de control, no prueba ningun algoritmo
func BenchmarkControlCase(b *testing.B) {
	for i := 0; i < b.N; i++ {
	}
}

func BenchmarkSimpleDetection(b *testing.B) {
	aString, bString := testData[0].A, testData[0].B

	b.ResetTimer() // Restablecer el temporizador para no incluir el tiempo que se toma en acceder a los valores de prueba.

	for i := 0; i < b.N; i++ {
		algo.SimpleDetection(aString, bString)
	}
}

func BenchmarkPrimeFactorDetection(b *testing.B) {
	aString, bString := testData[0].A, testData[0].B

	b.ResetTimer() // Restablecer el temporizador para no incluir el tiempo que se toma en acceder a los valores de prueba.

	for i := 0; i < b.N; i++ {
		algo.PrimeFactorsDetection(aString, bString)
	}
}

func BenchmarkMapDetection(b *testing.B) {
	aString, bString := testData[0].A, testData[0].B

	b.ResetTimer() // Restablecer el temporizador para no incluir el tiempo que se toma en acceder a los valores de prueba.

	for i := 0; i < b.N; i++ {
		algo.MapDetection(aString, bString)
	}
}
