package service

import (
	"petshop/Database"
	"petshop/domain"
)

type AnimalService struct{}

func (as *AnimalService) AddAnimal(animal domain.Animal) error {
	Veritabani, err := Database.Connect()
	if err != nil {
		return err
	}
	if err := Veritabani.Db.Insert("Animals", map[string]any{
		"Name":    animal.Name,
		"Species": animal.Species,
		"OwnerId": -1,
	}); err != nil {
		return err
	}

	return nil
}
func (as *AnimalService) DeleteAnimal(id int) error {
	Veritabani, err := Database.Connect()
	if err != nil {
		return err
	}
	if err := Veritabani.Db.Delete("Animals", "id", id); err != nil {
		return err
	}
	return nil
}
func (as *AnimalService) UpdateAnimal(oldName, newName string) error {
	Veritabani, err := Database.Connect()
	if err != nil {
		return err
	}
	if err := Veritabani.Db.Update("Animals", "Name", oldName, newName); err != nil {
		return err
	}
	return nil
}
func (as *AnimalService) ListAnimals() ([]domain.Animal, error) {
	Veritabani, err := Database.Connect()
	if err != nil {
		return nil, err
	}
	Animals, err := Veritabani.Db.Select("Animals")
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
	Veritabani, err := Database.Connect()
	if err != nil {
		return err
	}
	if err := Veritabani.Db.Update("Animals", "OwnerId", -1, UserId); err != nil {
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
