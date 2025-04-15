package mem

import (
	"fmt"
	"math/rand"
)

// Data - структура, що займає трохи пам'яті
type Data struct {
	ID   int
	Name string
	Tags [10]string // додамо масив, щоб збільшити розмір
}

func (d Data) String() string {
	return fmt.Sprintf("id: %d, name: %s, tags: %v", d.ID, d.Name, d.Tags)
}

// неефективна функція: створює багато тимчасових об'єктів і слайсів
func processDataInefficient(count int) [][]byte {
	var result [][]byte // зберігатимемо серіалізовані дані

	// симуляція генерації даних
	allData := make([]*Data, count)
	for i := 0; i < count; i++ {
		tags := [10]string{}
		for j := 0; j < 10; j++ {
			tags[j] = fmt.Sprintf("tag-%d-%d", i, rand.Intn(1000))
		}
		allData[i] = &Data{
			ID:   i,
			Name: fmt.Sprintf("Name-%d", i),
			Tags: tags,
		}
	}

	// неефективна обробка: створюємо новий рядок для кожного поля при "серіалізації"
	for _, d := range allData {
		// симуляція простої "серіалізації" в []byte через форматований рядок
		result = append(result, []byte(d.String())) // кожен append може ре-алокувати result
	}

	// багато тимчасових рядків створено fmt.Sprintf
	// allData все ще існує, займаючи пам'ять до кінця функції
	return result
}

// ефективна функція: перевикористання буфера та уникнення зайвих алокацій
func processDataEfficient(count int) [][]byte {
	var result [][]byte
	if count > 0 {
		result = make([][]byte, 0, count) // пре-алокуємо ємність слайсу result
	}

	var buffer []byte

	// симуляція генерації та обробки в одному циклі
	for i := 0; i < count; i++ {
		tags := [10]string{}
		for j := 0; j < 10; j++ {
			// припустимо, теги генеруються на льоту, якщо це можливо
			tags[j] = fmt.Sprintf("tag-%d-%d", i, rand.Intn(1000))
		}

		// можна створити на стеку (не pointer), якщо не потрібен вказівник поза циклом
		data := Data{
			ID:   i,
			Name: fmt.Sprintf("Name-%d", i),
			Tags: tags,
		}

		// "серіалізація" з перевикористанням буфера (дуже спрощено)
		// в реальності тут була б ефективніша логіка (json.Marshal, protobuf, etc.)
		// або ручне формування байтового слайсу
		buffer = buffer[:0] // очистити буфер (зберігаючи capacity)
		buffer = append(buffer, []byte(data.String())...)

		// потрібно скопіювати дані з буфера, оскільки буфер перевикористовується
		dataCopy := make([]byte, len(buffer))
		copy(dataCopy, buffer)
		result = append(result, dataCopy)
	}

	// у цьому варіанті менше тимчасових рядків
	// allData не створюється окремо
	// буфер перевикористовується
	return result
}
