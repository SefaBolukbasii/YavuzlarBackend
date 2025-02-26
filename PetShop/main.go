package main

import (
	"errors"
	"fmt"
	"os"
	"petshop/domain"
	"petshop/service"
)

var devamMı bool
var userSession domain.User

func AdminGenelİşlemler() {
	var TusMenu int
	fmt.Println("1-Hayvan İşlemleri")
	fmt.Println("2-Müşteri İşlemleri")
	fmt.Println("3-Market İşlemleri")
	fmt.Println("4-Çıkış")
	fmt.Scan(&TusMenu)
	if TusMenu == 1 {
		AdminHayvanIslemleri()
	} else if TusMenu == 2 {
		AdminMusteriIslemleri()
	} else if TusMenu == 3 {
		AdminMarketIslemleri()
	} else if TusMenu == 4 {
		os.Exit(0)
	} else {
		fmt.Println("Hatalı tuşlama yaptınız")
	}
}
func AdminMusteriIslemleri() {
	var TusMenu int
	var MusteriId int
	fmt.Println("1-Müşteri Ekleme")
	fmt.Println("2-Müşteri Silme")
	fmt.Println("3-Müşteri Güncelleme")
	fmt.Println("4-Müşteri Listeleme")
	fmt.Println("5-Çıkış")
	fmt.Scan(&TusMenu)
	us := service.UserService{}
	ls := service.LogService{}
	if TusMenu == 1 {
		UyeOlPage()
		ls.LogAdd(userSession.ID, "Kullanıcı ekledi")
	} else if TusMenu == 2 {
		fmt.Print("Silmek istediğiniz müşterinin ID'sini giriniz ")
		fmt.Scan(&MusteriId)
		if MusteriId > 0 {
			if err := us.DeleteUser(MusteriId); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Müşteri silindi")
				ls.LogAdd(userSession.ID, "Kullanıcı sildi")
			}
		} else {
			fmt.Println("Geçerli bir ID giriniz")
		}

	} else if TusMenu == 3 {
		var newBalance int
		fmt.Print("Bakiyesini güncellemek istediğiniz müşterinin Id sini giriniz: ")
		fmt.Scan(&MusteriId)
		fmt.Print("Yeni Bakiye: ")
		fmt.Scan(&newBalance)
		usersList, err := us.ListUser()
		if err != nil {
			fmt.Println(err)
		}
		for _, a := range usersList {
			if a.ID == MusteriId {
				oldBalance := a.Balance
				if err := us.ChangeBalance(&a, oldBalance, newBalance); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Bakiye güncellendi")
					ls.LogAdd(userSession.ID, "Kullanıcı güncelledi")
					break
				}
			}
		}

	} else if TusMenu == 4 {
		usersList, err := us.ListUser()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("--------------------------")
		for _, a := range usersList {
			fmt.Println(a.ID, " - ", a.UserName, " - ", a.Balance)
		}
		fmt.Println("--------------------------")
	} else if TusMenu == 5 {
		os.Exit(0)
	} else {
		fmt.Println("Hatalı tuşlama yaptınız")
	}
	AdminGenelİşlemler()
}
func AdminHayvanIslemleri() {
	var TusHayvan int
	var AnimalSpecies string
	var AnimalName string
	var AnimalNameNew string
	var AnimalId int
	fmt.Println("1-Hayvan Ekleme")
	fmt.Println("2-Hayvan Silme")
	fmt.Println("3-Hayvan Güncelleme")
	fmt.Println("4-Hayvan Listeleme")
	fmt.Println("5-Geri")
	fmt.Println("6-Çıkış")
	as := service.AnimalService{}
	ls := service.LogService{}
	fmt.Scan(&TusHayvan)
	if TusHayvan == 1 {
		fmt.Println("Hayvanın Türü: ")
		fmt.Scan(&AnimalSpecies)
		fmt.Println("Hayvan Adı: ")
		fmt.Scan(&AnimalName)
		animal := domain.Animal{
			Name:    AnimalName,
			Species: AnimalSpecies,
		}
		if err := as.AddAnimal(animal); err != nil {
			fmt.Println(err)
		}
		ls.LogAdd(userSession.ID, "Hayvan eklendi")
	} else if TusHayvan == 2 {
		fmt.Println("Silinecek Hayvanın id'si: ")
		fmt.Scan(&AnimalId)
		if err := as.DeleteAnimal(AnimalId); err != nil {
			fmt.Println(err)
		}
		ls.LogAdd(userSession.ID, "Hayvan sildi")
	} else if TusHayvan == 3 {
		fmt.Println("Güncellenecek Hayvanın İsmi: ")
		fmt.Scan(&AnimalName)
		fmt.Println("Yeni İsmi: ")
		fmt.Scan(&AnimalNameNew)
		if err := as.UpdateAnimal(AnimalName, AnimalNameNew); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("İşlem Başarı İle Gerçekleşti")
			ls.LogAdd(userSession.ID, "Hayvan güncelledi")
		}
	} else if TusHayvan == 4 {
		animals, err := as.ListAnimals()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("--------------------------------------------------")
			for _, animal := range animals {
				fmt.Print(animal.ID, "-", animal.Name, "-", animal.Species, "\n")
			}
			fmt.Println("--------------------------------------------------")
		}
	} else if TusHayvan == 5 {

	} else if TusHayvan == 6 {
		os.Exit(0)
	} else {
		fmt.Println("Hatalı Tuşlama")
	}
	AdminGenelİşlemler()
}
func AdminMarketIslemleri() {
	var TusMenu int
	var ItemName string
	var ItemPrice int
	var ItemId int
	var NewItemPrice int
	fmt.Println("1-İtem Ekleme")
	fmt.Println("2-İtem Silme")
	fmt.Println("3-İtem Güncelleme")
	fmt.Println("4-İtem Listeleme")
	fmt.Println("5-Çıkış")
	fmt.Scan(&TusMenu)
	Is := service.ItemService{}
	ls := service.LogService{}
	if TusMenu == 1 {
		fmt.Print("İtem İsmi: ")
		fmt.Scan(&ItemName)
		fmt.Print("\nİtem fiyatı: ")
		fmt.Scan(&ItemPrice)
		if err := Is.AddItem(domain.Item{
			Name:  ItemName,
			Price: ItemPrice,
		}); err != nil {
			fmt.Println(err)
			AdminGenelİşlemler()
		}
		ls.LogAdd(userSession.ID, "İtem ekledi")
	} else if TusMenu == 2 {
		fmt.Print("İtem ID: ")
		fmt.Scan(&ItemId)
		if err := Is.DeleteItem(ItemId); err != nil {
			fmt.Println(err)
			AdminGenelİşlemler()
		}
		ls.LogAdd(userSession.ID, "İtem sildi")
	} else if TusMenu == 3 {
		fmt.Print("İtem ID: ")
		fmt.Scan(&ItemId)
		fmt.Print("\nİtem yeni fiyatı: ")
		fmt.Scan(&NewItemPrice)
		Items, err := Is.ListItems()
		if err != nil {
			fmt.Print(err)
		}
		for _, Item := range Items {
			if Item.ID == ItemId {
				if err := Is.UpdateItem(&Item, NewItemPrice); err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("fiyat güncellendi")
					ls.LogAdd(userSession.ID, "İtem güncelledi")
					break
				}
			}
		}
	} else if TusMenu == 4 {
		Items, err := Is.ListItems()
		if err != nil {
			fmt.Print(err)
		}
		for _, Item := range Items {
			println(Item.ID, " - ", Item.Name, " - ", Item.Price)
		}
	} else if TusMenu == 5 {
		os.Exit(0)
	} else {
		fmt.Println("Hatalı tuşlama yaptınız")
	}
	AdminGenelİşlemler()
}
func UyeOlPage() {
	var kulAD string
	var sifre string
	fmt.Print("Kullanıcı Adı: ")
	fmt.Scan(&kulAD)
	fmt.Println("")
	fmt.Print("Şifre: ")
	fmt.Scan(&sifre)
	if len(kulAD) != 0 && kulAD != " " && len(sifre) != 0 && sifre != " " {
		us := service.UserService{}
		if err := us.Register(kulAD, sifre, "musteri"); err != nil {
			fmt.Println(err)
			fmt.Println("Kayıt olunamadı")
		} else {
			fmt.Println("Kayıt oluşturuldu")
		}
	} else {
		fmt.Println("Kullanıcı Adı veya Şifre eksik")
	}
}
func GirisYap() (*domain.User, error) {
	var kulAD string
	var sifre string
	fmt.Print("Kullanıcı Adı: ")
	fmt.Scan(&kulAD)
	fmt.Println("")
	fmt.Print("Şifre: ")
	fmt.Scan(&sifre)
	ls := service.LogService{}
	if len(kulAD) != 0 && kulAD != " " && len(sifre) != 0 && sifre != " " {
		us := service.UserService{}
		kullanici, err := us.Login(kulAD, sifre)
		if err != nil {
			fmt.Println("Hatalı giriş")
			return nil, err
		}
		ls.LogAdd(kullanici.ID, "Giriş yaptı")
		return kullanici, nil
	} else {
		fmt.Println("Kullanıcı adı veya şifre eksik")
		return nil, errors.New("Kullanici adi veya sifre eksik")
	}
}
func MusteriGenelIslemleri() {
	var TusMenu int
	var NewBalance int
	var oldBalance int
	fmt.Println("1-Hayvan İşlemleri")
	fmt.Println("2-Market İşlemleri")
	fmt.Println("3-Bakiye Yükle")
	fmt.Println("4-Çıkış")
	fmt.Scan(&TusMenu)
	us := service.UserService{}
	ls := service.LogService{}
	if TusMenu == 1 {
		musteriHayvanIslemleri()
	} else if TusMenu == 2 {
		musteriMarketIslemleri()
	} else if TusMenu == 3 {
		fmt.Print("Yüklemek İstediğiniz tutarı giriniz: ")
		fmt.Scan(&NewBalance)
		oldBalance = userSession.Balance
		if err := us.ChangeBalance(&userSession, oldBalance, oldBalance+NewBalance); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Bakiye yüklendi")
			ls.LogAdd(userSession.ID, "Bakiye ekledi")
		}
		MusteriGenelIslemleri()
	} else if TusMenu == 4 {
		os.Exit(0)
	} else {
		fmt.Println("Hatalı tuşlama yaptınız")
	}
}
func musteriHayvanIslemleri() {
	var TusHayvan int
	var AnimalOldName string
	var AnimalNewName string
	var AnimalId int
	fmt.Println("1-Hayvan Sahiplenme")
	fmt.Println("2-Hayvan İsim Takma")
	fmt.Println("3-Hayvanlarımı Listele")
	fmt.Println("4-Hayvan Listeleme")
	fmt.Println("5-Geri")
	fmt.Println("6-Çıkış")
	as := service.AnimalService{}
	ls := service.LogService{}
	fmt.Scan(&TusHayvan)
	if TusHayvan == 1 {
		fmt.Print("Sahiplenmek istediğiniz hayvan Id: ")
		fmt.Scan(&AnimalId)
		if err := as.AnimalOwned(AnimalId, userSession.ID); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Hayvan sahiplenildi")
			ls.LogAdd(userSession.ID, "Hayvan sahiplendi")
		}

	} else if TusHayvan == 2 {
		fmt.Print("Hayvanın adı: ")
		fmt.Scan(&AnimalOldName)
		fmt.Print("\nYeni Adı: ")
		fmt.Scan(&AnimalNewName)
		if err := as.UpdateAnimal(AnimalOldName, AnimalNewName); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Başarı ile güncellendi")
			ls.LogAdd(userSession.ID, "Hayvan güncelledi")
		}

	} else if TusHayvan == 3 {
		Animals, err := as.MyAnimals(userSession.ID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("-----------------------------------")
		for _, Animal := range Animals {
			fmt.Println(Animal.ID, " - ", Animal.Name, " - ", Animal.Species)
		}
		fmt.Println("-----------------------------------")
	} else if TusHayvan == 4 {
		Animals, err := as.ListAnimals()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("-----------------------------------")
		for _, Animal := range Animals {
			if Animal.OwnerID == -1 {
				fmt.Println(Animal.ID, " - ", Animal.Name, " - ", Animal.Species)
			}

		}
		fmt.Println("-----------------------------------")
	} else if TusHayvan == 5 {

	} else if TusHayvan == 6 {
		os.Exit(0)
	} else {
		fmt.Println("Hatalı Tuşlama")
	}
	MusteriGenelIslemleri()
}
func musteriMarketIslemleri() {
	var TusMarket int
	var ItemId int
	var ItemDomain domain.Item
	fmt.Println("1-Ürün Satın Al")
	fmt.Println("2-Sipariş Geçmişi")
	fmt.Println("3-Geri")
	fmt.Println("4-Çıkış")
	fmt.Scan(&TusMarket)
	Is := service.ItemService{}
	ls := service.LogService{}
	if TusMarket == 1 {
		Items, err := Is.ListItems()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("------------Ürünler--------------")
		for _, Item := range Items {
			fmt.Println(Item.ID, " - ", Item.Name, " - ", Item.Price)
		}
		fmt.Println("---------------------------------")
		fmt.Println("")
		fmt.Print("Satın Almak İstediğiniz Ürün Id: ")
		fmt.Scan(&ItemId)
		for _, Item := range Items {
			if ItemId == Item.ID {
				ItemDomain = domain.Item{
					ID:    Item.ID,
					Name:  Item.Name,
					Price: Item.Price,
				}
				break
			}
		}
		if userSession.Balance > ItemDomain.Price {
			if err := Is.ToBuy(&ItemDomain, &userSession); err != nil {
				fmt.Println(err)
			}
			fmt.Println("Sipariş tamamlandı")
			ls.LogAdd(userSession.ID, "İtem aldı")
		} else {
			fmt.Println("Yetersiz bakiye")
		}

	} else if TusMarket == 2 {
		Items, err := Is.ListBuy(&userSession)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("--------------------------------")
		for _, Item := range Items {
			fmt.Println(Item.ID, " - ", Item.Name, " - ", Item.Price)
		}
		fmt.Println("---------------------------------")
	} else if TusMarket == 3 {

	} else if TusMarket == 4 {
		os.Exit(0)
	} else {
		fmt.Println("Hatalı seçim")
	}
	MusteriGenelIslemleri()
}
func main() {

	var Tus int
	devamMı = true
	for devamMı == true {
		fmt.Println("Hoşgeldiniz...")
		fmt.Println("1-Üye Ol")
		fmt.Println("2-Giriş Yap")
		fmt.Scan(&Tus)
		if Tus == 1 {
			UyeOlPage()
		} else if Tus == 2 {
			userSession, err := GirisYap()
			if err != nil {
				fmt.Println("Hatalı şifre veya kullanıcı adı")
			} else {
				if userSession.Role == "admin" {
					AdminGenelİşlemler()
				} else if userSession.Role == "müsteri" {
					MusteriGenelIslemleri()
				}
			}

		}
	}
}
