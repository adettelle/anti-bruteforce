package bucket

import (
	"fmt"
	"log"
	"time"
)

type LeakyBucket struct {
	capacity int           // Maximum number of requests the bucket can hold
	rate     time.Duration // Скорость "протекания" (например, 1 запрос в 100ms)
	bucket   chan struct{}
	ticker   *time.Ticker // Таймер для "протекания"
	Added    time.Time    // последняя попытка доступа
}

func NewLeakyBucket(capacity int, added time.Time) *LeakyBucket { // rate time.Duration,
	// TODO: attempt, attemptTime -> переести в config?
	attempt := 10
	attemptTime := 10 * time.Second // time.Minute

	rate := time.Duration(attempt) * attemptTime / time.Duration(capacity)

	lb := LeakyBucket{
		capacity: capacity,
		rate:     rate,
		bucket:   make(chan struct{}, capacity), // Буферизованный канал
		ticker:   time.NewTicker(rate),
		Added:    added,
	}

	go lb.leak()

	return &lb
}

// Функция "слива" запросов.
func (lb *LeakyBucket) leak() {
	for range lb.ticker.C { // Ждем тика от таймера
		select {
		case <-lb.bucket: // Если в ведре есть запрос, "выливаем" его
			fmt.Printf("Запрос обработан, bucket len = %v\n", len(lb.bucket))
		default:
			// Ведро пусто, ничего не делаем
		}
	}
}

func (lb *LeakyBucket) AddRequest() bool {
	select {
	case lb.bucket <- struct{}{}: // Пытаемся добавить запрос
		log.Println("request added in bucket")
		return true
	default:
		log.Println("bucket is full")
		return false // Ведро переполнено, запрос отбрасываем
	}
}

func (lb *LeakyBucket) GetBucketSize() int {
	return len(lb.bucket)
}
