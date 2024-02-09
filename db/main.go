package db

import (
	"log"

	"github.com/osaxon/funghi/types"
	"github.com/osaxon/funghi/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	dsn := utils.LoadEnv("DATABASE_URL")

	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&types.Funghi{})

	//Seed(db)

}

func Seed(db *gorm.DB) {
	var count int64
	db.Model(&types.Funghi{}).Count(&count)

	if count > 0 {
		err := db.Exec("DELETE FROM funghis")
		if err.Error != nil {
			log.Fatal(err.Error)
		}
	}

	db.Create(&types.Funghi{Name: "Fly Agaric", ScientificName: "Amanita muscaria", Edible: false})
	db.Create(&types.Funghi{Name: "Liberty Cap", ScientificName: "Psilocybe semilanceata", Edible: true})
	db.Create(&types.Funghi{Name: "Shaggy Ink Cap", ScientificName: "Coprinus comatus", Edible: true})
	db.Create(&types.Funghi{Name: "Chanterelle", ScientificName: "Cantharellus cibarius", Edible: true})
	db.Create(&types.Funghi{Name: "Porcini", ScientificName: "Boletus edulis", Edible: true})
	db.Create(&types.Funghi{Name: "Death Cap", ScientificName: "Amanita phalloides", Edible: false})
	db.Create(&types.Funghi{Name: "Chicken of the Woods", ScientificName: "Laetiporus sulphureus", Edible: true})
	db.Create(&types.Funghi{Name: "Hen of the Woods", ScientificName: "Grifola frondosa", Edible: true})
	db.Create(&types.Funghi{Name: "King Oyster", ScientificName: "Pleurotus eryngii", Edible: true})
	db.Create(&types.Funghi{Name: "Shiitake", ScientificName: "Lentinula edodes", Edible: true})
	db.Create(&types.Funghi{Name: "Maitake", ScientificName: "Grifola frondosa", Edible: true})
	db.Create(&types.Funghi{Name: "Reishi", ScientificName: "Ganoderma lucidum", Edible: true})
	db.Create(&types.Funghi{Name: "Turkey Tail", ScientificName: "Trametes versicolor", Edible: true})
	db.Create(&types.Funghi{Name: "Lacrymaria lacrymabunda", ScientificName: "Lacrymaria lacrymabunda", Edible: true})
	db.Create(&types.Funghi{Name: "Penny Bun", ScientificName: "Boletus edulis", Edible: true})
	db.Create(&types.Funghi{Name: "Cauliflower Fungus", ScientificName: "Sparassis crispa", Edible: true})
	db.Create(&types.Funghi{Name: "Hedgehog Fungus", ScientificName: "Hydnum repandum", Edible: true})
	db.Create(&types.Funghi{Name: "Honey Fungus", ScientificName: "Armillaria mellea", Edible: true})
	db.Create(&types.Funghi{Name: "Oyster Mushroom", ScientificName: "Pleurotus ostreatus", Edible: true})
	db.Create(&types.Funghi{Name: "Wood Blewit", ScientificName: "Clitocybe nuda", Edible: true})
	db.Create(&types.Funghi{Name: "Cep", ScientificName: "Boletus edulis", Edible: true})
	db.Create(&types.Funghi{Name: "Morel", ScientificName: "Morchella esculenta", Edible: true})
	db.Create(&types.Funghi{Name: "Girolle", ScientificName: "Cantharellus cibarius", Edible: true})
	db.Create(&types.Funghi{Name: "Black Trumpet", ScientificName: "Craterellus cornucopioides", Edible: true})
	db.Create(&types.Funghi{Name: "Horn of Plenty", ScientificName: "Craterellus cornucopioides", Edible: true})
	db.Create(&types.Funghi{Name: "St George's Mushroom", ScientificName: "Calocybe gambosa", Edible: true})
	db.Create(&types.Funghi{Name: "Pied de Mouton", ScientificName: "Hydnum repandum", Edible: true})
	db.Create(&types.Funghi{Name: "Puffball", ScientificName: "Lycoperdon perlatum", Edible: true})
	db.Create(&types.Funghi{Name: "Stinkhorn", ScientificName: "Phallus impudicus", Edible: true})
	db.Create(&types.Funghi{Name: "Truffle", ScientificName: "Tuber melanosporum", Edible: true})
	db.Create(&types.Funghi{Name: "Cordyceps", ScientificName: "Cordyceps militaris", Edible: true})
	db.Create(&types.Funghi{Name: "Lion's Mane", ScientificName: "Hericium erinaceus", Edible: true})
	db.Create(&types.Funghi{Name: "Matsutake", ScientificName: "Tricholoma matsutake", Edible: true})
	db.Create(&types.Funghi{Name: "Enokitake", ScientificName: "Flammulina velutipes", Edible: true})
	db.Create(&types.Funghi{Name: "Shimeji", ScientificName: "Hypsizygus tessellatus", Edible: true})
}
