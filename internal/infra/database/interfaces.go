package database

import "github.com/mrpsousa/api/internal/entity"

type TransactionInterface interface {
	Create(transaction *entity.Transaction) error
	SaveFromFile(path string) error
	Save(line string) error
}

// 3. Exibir a lista de todas as transações de produtos importadas
// 4. Exibir o saldo final do produtor
// 5. Exibir o saldo final de um afiliado
