package database

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/pkg"
	"gorm.io/gorm"
)

type Transaction struct {
	DB *gorm.DB
}

func NewTransaction(db *gorm.DB) *Transaction {
	return &Transaction{DB: db}
}

func (p *Transaction) Create(user *entity.Transaction) error {
	return p.DB.Create(user).Error
}

func (p *Transaction) Save(line string) error {
	tp, err := pkg.StringToInt8(line[:1])
	if err != nil {
		return err
	}
	value, err := pkg.StringToFloat64(line[56:66])
	if err != nil {
		return err
	}
	createdAt := line[1:26]
	product := line[26:56]
	seller := line[66:]

	transaction, err := entity.NewTransaction(tp, createdAt, product, seller, value)
	err = p.Create(transaction)
	if err != nil {
		return err
	}
	return nil
}

func (p *Transaction) SaveFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		line = bytes.TrimSuffix(line, []byte{'\n'})
		err = p.Save(string(line))
		if err != nil {
			return err
		}
	}

	return nil
}
