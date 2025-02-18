package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

// Env chứa các thông tin cấu hình môi trường
type Env struct {
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPass     string
	PostgresName     string
	ElasticsearchURL string
	RedisAddr        string
	RedisPassword    string
	RedisDB          int
	WebSocketHost    string
}

// LoadConfig tải cấu hình từ biến môi trường
func LoadConfig() (*Env, error) {
	env := &Env{
		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresUser:     getEnv("POSTGRES_USER", "user"),
		PostgresPass:     getEnv("POSTGRES_PASS", "password"),
		PostgresName:     getEnv("POSTGRES_NAME", "database"),
		ElasticsearchURL: getEnv("ELASTICSEARCH_URL", "http://localhost:9200"),
		RedisAddr:        getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword:    getEnv("REDIS_PASSWORD", ""),
		RedisDB:          getEnvInt("REDIS_DB", 0),
		WebSocketHost:    getEnv("WEBSOCKET_HOST", "localhost:8080"),
	}

	return env, nil
}

// getEnv lấy giá trị từ biến môi trường hoặc trả về giá trị mặc định
func getEnv(key, defaultValue string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		val = defaultValue
	}
	return val
}

// getEnvInt lấy giá trị int từ biến môi trường
func getEnvInt(key string, defaultValue int) int {
	val, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	// Convert string to int
	parsedValue, err := strconv.Atoi(val)
	if err != nil {
		log.Printf("Error parsing %s as int: %v. Using default value %d.", key, err, defaultValue)
		return defaultValue
	}
	return parsedValue
}

// PrintConfig in ra cấu hình đã tải lên
func PrintConfig(env *Env) {
	fmt.Println("Postgres Host:", env.PostgresHost)
	fmt.Println("Postgres Port:", env.PostgresPort)
	fmt.Println("Postgres User:", env.PostgresUser)
	fmt.Println("Postgres Name:", env.PostgresName)
	fmt.Println("Elasticsearch URL:", env.ElasticsearchURL)
	fmt.Println("Redis Addr:", env.RedisAddr)
	fmt.Println("WebSocket Host:", env.WebSocketHost)
}
