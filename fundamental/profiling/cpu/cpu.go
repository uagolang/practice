package cpu

import "strings"

// неефективна функція: конкатенація рядків за допомогою '+' у циклі
func concatInefficient(n int, s string) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s // кожна ітерація створює новий рядок і копіює дані
	}

	return result
}

// ефективна функція: використання strings.Builder
func concatEfficient(n int, s string) string {
	var builder strings.Builder

	// оцінюємо потрібний розмір, щоб уникнути ре-алокації
	builder.Grow(n * len(s))
	for i := 0; i < n; i++ {
		builder.WriteString(s) // додавання до буфера без зайвих алокацій
	}

	return builder.String()
}
