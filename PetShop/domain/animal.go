package domain

type Animal struct {
	ID      int
	Species string
	Name    string
	OwnerID int // Null olabilir
}

type IAnimal interface {
	AddAnimal(animal Animal) error
	DeleteAnimal(id int) error
	UpdateAnimal(oldName, newName string) error
	ListAnimals() ([]Animal, error)
	AnimalOwned(AnimalId, UserId int) error
	MyAnimals(UserId *int) ([]Animal, error)
}
