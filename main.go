package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	// Проверяем, что передан аргумент с путем к файлу
	if len(os.Args) < 2 {

		os.Exit(1)
	}

	filePath := os.Args[1]

	// Вычисляем MD5 хеш файла
	hash, err := calculateMD5(filePath)
	if err != nil {
		fmt.Printf("Ошибка при вычислении MD5: %v\n", err)
		os.Exit(1)
	}

	// Вставляем хеш в начало файла
	err = insertHashToFile(filePath, hash)
	if err != nil {
		fmt.Printf("Ошибка при записи хеша в файл: %v\n", err)
		os.Exit(1)
	}

}

// calculateMD5 вычисляет MD5 хеш файла
func calculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

// insertHashToFile вставляет хеш в начало файла
func insertHashToFile(filePath string, hash string) error {
	// Читаем содержимое исходного файла
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	// Создаем временный файл
	tempFile := filePath + ".tmp"
	file, err := os.Create(tempFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Записываем хеш и оригинальное содержимое
	hashLine := fmt.Sprintf("; MD5:%s\n", hash)
	if _, err := file.WriteString(hashLine); err != nil {
		return err
	}

	if _, err := file.Write(content); err != nil {
		return err
	}

	// Закрываем файл перед переименованием
	file.Close()

	// Заменяем оригинальный файл временным
	return os.Rename(tempFile, filePath)
}
