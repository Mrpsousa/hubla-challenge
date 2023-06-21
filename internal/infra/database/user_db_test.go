package database_test

import (
	"testing"

	"github.com/mrpsousa/api/internal/entity"
	"github.com/mrpsousa/api/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Roger", "email@email", "123456")
	userDB := database.NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound, "id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}

func TestFindByEmail(t *testing.T) {
	db, err := returnDBInstance()
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Roger", "email@email", "123456")
	userDB := database.NewUser(db)

	err = userDB.Create(user)
	assert.Nil(t, err)
	userFound, err := userDB.FindByEmail(user.Email)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)
}
