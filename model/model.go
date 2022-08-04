package model

type Hero struct {
	ID     string    		`json:"id"`
	SuperHero string 		`json:"superhero"`
	Publisher string 		`json:"publisher"`
	Alter_Ego string 		`json:"alter_ego"`
	First_appearance string 	`json:"first_appearance"`
	Characters string 		`json:"characters"`
}

type Heroes struct {
	Heroes []Hero `json:"heroes"`
}

func (heroes *Heroes) GetAllBHeroes() Heroes {
	return *heroes
}

func (heroes *Heroes) GetHeroByID(id string) Hero {
	for _, hero := range *&heroes.Heroes {
		if hero.ID == id {
			return hero
		}
	}
	return Hero{}
}

func (heroes *Heroes) AddHero(ID, SuperHero, Publisher, Alter_Ego, First_appearance, Characters string) Hero {
	var hero *Hero
	if ID != "" {
		hero = NewHero(ID, SuperHero, Publisher, Alter_Ego, First_appearance, Characters)
		*&heroes.Heroes = append(*&heroes.Heroes, *hero)
	}
	return *hero
}

func (heroes *Heroes) UpdateHero(ID, SuperHero, Publisher, Alter_Ego, First_appearance, Characters string) Hero {
	hero, heroAux := Hero{}, Hero{}
	
	if ID != "" {
		hero = heroes.GetHeroByID(ID)
		//index := heroes.indexOf(ID)
		if hero == heroAux {
			return heroAux
		}
		if SuperHero != "" {
			hero.SuperHero = SuperHero
		}
		if Publisher != "" {
			hero.Publisher = Publisher
		}
		if Alter_Ego != "" {
			hero.Alter_Ego = Alter_Ego
		}
		if First_appearance != "" {
			hero.First_appearance = First_appearance
		}
		if Characters != "" {
			hero.Characters = Characters
		}
		//*heroes[index] = hero
	}
	return hero
}

func NewHero(ID, SuperHero, Publisher, Alter_Ego, First_appearance, Characters string) *Hero {
	return &Hero {
		ID : ID,
		SuperHero : SuperHero,
		Publisher : Publisher,
		Alter_Ego : Alter_Ego,
		First_appearance : First_appearance,
		Characters : Characters,
	}
}

func (heroes *Heroes)indexOf(ID string) (int) {
	for k, hero := range *&heroes.Heroes {
		if ID == hero.ID {
			return k
		}
	}
	return -1    //not found.
 }