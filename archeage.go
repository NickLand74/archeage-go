package archeage

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type ArcheAge struct {
	client *http.Client
}

func New(c *http.Client) *ArcheAge {
	return &ArcheAge{c}
}

func (a *ArcheAge) post(url string, form io.Reader) (*goquery.Document, error) {
	return a.do(url, "POST", form)
}

func (a *ArcheAge) get(url string) (*goquery.Document, error) {
	return a.do(url, "GET", nil)
}

func (a *ArcheAge) do(url, method string, form io.Reader) (*goquery.Document, error) {
	req, err := http.NewRequest(method, url, form)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := a.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	return doc, nil
}

// ArcheAgeBot представляет бота в игре ArcheAge.
type ArcheAgeBot struct {
	Name       string      `json:"name"`
	Thumb      string      `json:"thumb"`
	UUID       string      `json:"uuid"`
	Server     string      `json:"server"`
	Level      string      `json:"level"`
	Race       string      `json:"race"`
	Expedition *Expedition `json:"expedition"`
	Stat       *Stat       `json:"stat"`
	Class      *Class      `json:"class"`
	Position   *Position   `json:"position"`
	X          float64     `json:"x"`
	Y          float64     `json:"y"`
}

// Node представляет шахтный узел.
type Node struct {
	X       float64
	Y       float64
	Z       float64
	IsMined bool
}

// getNearestMiningNode реализует логику поиска ближайшего шахтного узла.
func (a *ArcheAgeBot) getNearestMiningNode() *Node {
	return &Node{X: 100, Y: 200, Z: 300, IsMined: false} // Пример возвращаемого узла
}

// autoMine запускает цикл автоматической добычи ресурсов.
func (a *ArcheAgeBot) autoMine() {
	for {
		node := a.getNearestMiningNode() // Поиск ближайшего узла
		if node != nil && !node.IsMined {
			if err := a.moveToNode(node); err != nil {
				fmt.Println("Ошибка при перемещении:", err)
				return // Или можно выполнять повторную попытку
			}
			if err := a.mineAtNode(node); err != nil {
				fmt.Println("Ошибка добычи:", err)
			}
		} else {
			fmt.Println("Нет доступных узлов для добычи.")
		}

		time.Sleep(2 * time.Second) // Ожидание перед следующей попыткой
	}
}

// moveToNode перемещает персонажа к указанному узлу.
func (a *ArcheAgeBot) moveToNode(node *Node) error {
	// Логика перемещения. Можно использовать API или эмуляцию ввода
	fmt.Printf("Перемещение к узлу на координатах (%f, %f, %f)\n", node.X, node.Y, node.Z)
	// Эмуляция времени на перемещение
	time.Sleep(1 * time.Second)
	return nil
}

// mineAtNode выполняет действие добычи на узле.
func (a *ArcheAgeBot) mineAtNode(node *Node) error {
	if node.IsMined {
		return fmt.Errorf("узел уже был добыт")
	}

	// Симуляция добычи
	fmt.Printf("Добыча на узле: (%f, %f, %f)\n", node.X, node.Y, node.Z)
	time.Sleep(5 * time.Second) // Симуляция времени добычи

	// Пометить узел как добытый
	node.IsMined = true
	return nil
}

// Функция main для запуска программы.
func main() {
	bot := &ArcheAgeBot{
		Name:   "Мяснойпупс",
		Thumb:  "https://i.playground.ru/e/VuV5xie5XJiE9c3l_btlAw.jpeg", // Замените на URL вашего изображения
		UUID:   "uuid-1234",
		Server: "Мираж",
		Level:  "38",
		Race:   "Warborn",
		// Инициализируйте другие поля по мере необходимости
	}

	// Запуск авто-добычи
	fmt.Println("Запуск авто-добычи...")
	bot.autoMine()
}
