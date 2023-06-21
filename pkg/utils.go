package pkg

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/pemistahl/lingua-go"
)

var (
	ErrInvalidID      = errors.New("id_is_required_or_invalid")
	ErrInvalidType    = errors.New("type_is_required_or_invalid")
	ErrInvalidDate    = errors.New("date_is_required_or_invalid")
	ErrInvalidProduct = errors.New("product_is_required_or_invalid")
	ErrInvalidValue   = errors.New("value_is_required_or_invalid")
	ErrInvalidSeller  = errors.New("seller_is_required_or_invalid")
)

func StringToInt8(str string) (int8, error) {
	num, err := strconv.ParseInt(str, 10, 8)
	if err != nil {
		return 0, errors.New("Failed type convert")
	}
	return int8(num), nil
}

func StringToFloat64(str string) (float64, error) {
	num, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, errors.New("Failed type convert")
	}
	return num, nil
}

func FileNameGenerate(file string) string {
	return uuid.New().String() + "-" + file
}

func ForeignProductValidate(text string) bool {
	languages := []lingua.Language{
		lingua.Portuguese,
		lingua.English,
	}

	detector := lingua.NewLanguageDetectorBuilder().
		FromLanguages(languages...).
		Build()

	if language, exists := detector.DetectLanguageOf(text); exists {
		if language.String() != "Portuguese" {
			return true
		}
	}
	return false
}

// TODO must be tested
func RemoveFolder(dirPath string) {
	err := os.RemoveAll(dirPath)
	if err != nil {
		log.Println("failed_to_remove_tmp_directory:", err)
		return
	}
}

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

func listFilesInDirectory(dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("Directory: %s\n", entry.Name())
		} else {
			fmt.Printf("File: %s\n", entry.Name())
		}
	}
}
