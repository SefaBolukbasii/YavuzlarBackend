package service

import (
	"petshop/domain"

	jsondb "github.com/SefaBolukbasii/JsonDB"
)

type AnimalService struct {
	db *jsondb.Database
}

func CreateAnimalService(db *jsondb.Database) domain.IAnimal {
	return &AnimalService{db: db}
}
func (as *AnimalService) AddAnimal(animal domain.Animal) error {

	if err := as.db.Insert("Animals", map[string]any{
		"Name":    animal.Name,
		"Species": animal.Species,
		"OwnerId": -1,
	}); err != nil {
		return err
	}

	return nil
}
func (as *AnimalService) DeleteAnimal(id int) error {

	if err := as.db.Delete("Animals", "id", id); err != nil {
		return err
	}
	return nil
}
func (as *AnimalService) UpdateAnimal(oldName, newName string) error {

	if err := as.db.Update("Animals", "Name", oldName, newName); err != nil {
		return err
	}
	return nil
}
func (as *AnimalService) ListAnimals() ([]domain.Animal, error) {

	Animals, err := as.db.Select("Animals")
	if err != nil {
		return nil, err
	}
	var animalArray []domain.Animal
	for _, animal := range Animals {
		a := domain.Animal{
			ID:      animal["id"].(int),
			Name:    animal["Name"].(string),
			Species: animal["Species"].(string),
			OwnerID: animal["OwnerId"].(int),
		}
		animalArray = append(animalArray, a)
	}
	return animalArray, nil

}

func (as *AnimalService) AnimalOwned(AnimalId, UserId int) error {

	if err := as.db.Update("Animals", "OwnerId", -1, UserId); err != nil {
		return err
	}
	return nil
}

func (as *AnimalService) MyAnimals(UserId int) ([]domain.Animal, error) {
	Animals, err := as.ListAnimals()
	if err != nil {
		return nil, err
	}
	var animalArray []domain.Animal
	for _, animal := range Animals {
		if animal.OwnerID == UserId {
			animalArray = append(animalArray, animal)
		}
	}
	return animalArray, nil
}
