// Примеры шаблонов тестирования
package testingpattern

import (
	"fmt"
	"testing"
)

// BenchmarkFoo простая функция с benchmark
//
//	Имя функции начинается с префикса бенчмарка. Тестируемая функция (foo)
//
// вызывается в цикле for. b.N представляет собой переменное число итераций. При запуске
// бенчмарка Go пытается привести его в соответствие с запрошенным временем бенчмарка. Время
// тестирования по умолчанию установлено равным 1 секунде и может быть изменено с
// помощью флага -bench time. Время ожидания начинается с 1; если тест завершается менее чем
// за 1 секунду, время ожидания увеличивается, и тест выполняется снова, пока время
// ожидания примерно не совпадет со временем ожидания:
func BenchmarkFoo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			fmt.Println("any function")
		}()
	}
}

// BenchmarkResetTimer функция с benchmark и сбросом таймера
//
//	В некоторых случаях нам необходимо выполнить операции перед циклом тестирования. Эти
//
// операции могут занять довольно много времени (например, для создания большого фрагмента данных) и
// могут существенно повлиять на результаты тестирования.
//
//	В этом случае мы можем использовать метод сброса таймера перед входом в цикл
func BenchmarkResetTimer(b *testing.B) {
	func() {
		//ресурсоемкая функция
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		func() {
			//функция для которую выполняется тестирование
		}()
	}
}

// BenchmarkStopStartTimer функция с benchmark и паузой для исключения из таймеров ресурсоемкой функции
//
//	Можно сделать паузу для выполнения дорогостоящей операции при выполнения тестов
func BenchmarkStopStartTimer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		func() {
			//дорогостоящая операция
		}()
		b.StartTimer()

		func() {
			//выполнение функции для которой создан тест
		}()
	}
}
