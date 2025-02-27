package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {
	// Получаем адрес Redis из переменной окружения
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}

	// Подключаемся к Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr, // Адрес Redis
		Password: "",        // Без пароля
		DB:       0,
	})

	// Проверяем соединение
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения к Redis: %v", err)
	}

	fmt.Println("Redis подключен!")

	// Добавляем тестовые данные
	//err = rdb.Set(ctx, "key", "Hello Eraga", 0).Err()
	//if err != nil {
	//	log.Fatalf("Ошибка при записи в Redis: %v", err)
	//}

	// Читаем тестовые данные
	//val, err := rdb.Get(ctx, "leader").Result()
	//if err != nil {
	//	log.Fatalf("Ошибка при чтении из Redis: %v", err)
	//}

	///////////////////
	err = rdb.ZAdd(ctx, "leader",
		redis.Z{Score: 150, Member: "Eraga"},
		redis.Z{Score: 100, Member: "SHeraga"},
		redis.Z{Score: 500, Member: "Araga"},
		redis.Z{Score: 50, Member: "ZHanaga"},
		redis.Z{Score: 1150, Member: "Inatay"},
		redis.Z{Score: 950, Member: "Samat"},
	).Err()
	if err != nil {
		log.Fatal(err)
	}
	//ranks, err := rdb.ZRevRank(ctx, "leader", "Eraga").Result()
	//if err != nil {
	//	fmt.Println("лощибка ")
	//	log.Fatal(err)
	//}
	topPlayers, err := rdb.ZRevRangeWithScores(ctx, "leader", 0, 9).Result()
	if err != nil {
		log.Fatal(err)
	}
	for _, player := range topPlayers {
		fmt.Printf("oynwy: %v, upai: %v\n", player.Member, player.Score)

	}
	err = rdb.ZAdd(ctx, "leader",
		redis.Z{Score: 15000, Member: "Eraga"},
		redis.Z{Score: 10000, Member: "SHeraga"},
		redis.Z{Score: 500, Member: "Araga"},
		redis.Z{Score: 5000, Member: "ZHanaga"},
		redis.Z{Score: 1150, Member: "Inatay"},
		redis.Z{Score: 950, Member: "Samat"},
	).Err()
	if err != nil {
		log.Fatal(err)
	}
	topPlayers, err = rdb.ZRevRangeWithScores(ctx, "leader", 0, 9).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	for _, player := range topPlayers {

		fmt.Printf("oynwy: %v, upai: %v\n", player.Member, player.Score)

	}
	//fmt.Println("Прочитанное значение из Redis:", ranks)
}
