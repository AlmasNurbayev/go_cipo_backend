package clients

import (
	"encoding/json"
	"log/slog"
	"sort"
	"sync"

	"github.com/AlmasNurbayev/go_cipo_backend/internal/config"
)

type result struct {
	ValueName string
	Status    int
	Items     []string
	Err       []error
}

// Передаем в функцию название категории обуви
// Получаем массив с необходимыми атрибутами и их значениями,
// такими как размер, вид модели, материал, цвет, сезон
func KaspiGetAllValues(cfg *config.Config, log *slog.Logger, categoryName string, token string) []result {

	queries := []string{"Shoes*Colour", "Shoes*Size", "Shoes*Season", "Shoes*Material", "Shoes*Gender", "Shoes*Model"}
	results := make([]result, len(queries)) // Создаём массив нужного размера
	var wg sync.WaitGroup

	for i, query := range queries {
		wg.Add(1)

		// Запускаем запрос на нужный список values
		go func(i int, query string) {
			defer wg.Done()
			var errors []error
			status, body, err := KaspiGetValues(cfg, log, query, categoryName, token)
			errors = append(errors, err)

			var items []string
			if status == 200 {
				var rawData []map[string]interface{}
				err = json.Unmarshal([]byte(body), &rawData)
				if err != nil {
					errors = append(errors, err)
				}
				for _, obj := range rawData {
					if item, ok := obj["name"].(string); ok {
						items = append(items, item)
					}
				}
				sort.Strings(items)
			}

			results[i] = result{ValueName: query, Status: status, Items: items, Err: errors} // Записываем результат по индексу
		}(i, query) //для запуска горутины передаем скобки и параметры
	}
	wg.Wait() // Ждём завершения всех горутин

	return results

}
