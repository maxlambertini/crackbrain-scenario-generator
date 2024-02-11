package main

import (
	"fmt"
	"strings"
)

type Person interface {
	GetName() string
}

type Png struct {
	Nome       string `json:"nome"`
	Background string `json:"background"`
	Oggetto    string `json:"oggetto"`
	Indole     string `json:"indole"`
}

func (p *Png) GetName() string {
	return p.Nome
}

func CreatePng(t *CrackbrainTables) *Png {
	var p = Png{
		Nome:       t.NomeCasuale(),
		Background: t.GetBakgroundPng(),
		Oggetto:    t.GetOggettoPersonale(),
		Indole:     strings.Join(t.GetIndolePng(2), " - "),
	}
	return &p
}

type Personaggio struct {
	Nome       string              `json:"nome"`
	Background string              `json:"background"`
	Oggetto    string              `json:"oggetto"`
	Stile      *StileInvestigativo `json:"stile"`
}

func CreatePersonaggio(t *CrackbrainTables) *Personaggio {
	var p = Personaggio{
		Nome:       t.NomeCasuale(),
		Background: t.GetBakgroundPng(),
		Oggetto:    t.GetOggettoPersonale(),
		Stile:      t.GetStileInvestigativo(),
	}
	return &p
}

func (p *Personaggio) GetName() string {
	return p.Nome
}

type Vittima struct {
	Png               *Png   `json:"png"`
	LuogoRitrovamento string `json:"luogoRitrovamento"`
	Relazione         string `json:"relazione"`
}

func (p *Vittima) GetName() string {
	return p.Png.GetName()
}

func CreateRandomVittima(t *CrackbrainTables) *Vittima {
	return &Vittima{
		Png:               CreatePng(t),
		LuogoRitrovamento: t.GetLuogo(),
		Relazione:         strings.Join(t.GetRelazioni(2), " - "),
	}

}

type PngPrincipale struct {
	Png       *Png   `json:"png"`
	Relazione string `json:"relazione"`
}

func (p *PngPrincipale) GetName() string {
	return p.Png.GetName()
}

func CreatePngPrincipale(t *CrackbrainTables) *PngPrincipale {
	return &PngPrincipale{
		Png:       CreatePng(t),
		Relazione: strings.Join(t.GetRelazioni(2), " - "),
	}

}

type PngSecondario struct {
	Ruolo     string `json:"ruolo"`
	Png       *Png   `json:"png"`
	Relazione string `json:"relazione"`
}

func (p *PngSecondario) GetName() string {
	return p.Png.GetName()
}

func CreatePngSecondario(t *CrackbrainTables, ruolo string) *PngSecondario {
	return &PngSecondario{
		Png:       CreatePng(t),
		Relazione: strings.Join(t.GetRelazioni(1), " - "),
		Ruolo:     ruolo,
	}
}

type Orrore struct {
	Tipo           string         `json:"tipo"`
	Tratti         string         `json:"tratti"`
	Luogo          string         `json:"luogo"`
	Manifestazione *Png           `json:"manifestazione"`
	Vittima        *Vittima       `json:"vittima"`
	Agente         *PngPrincipale `json:"agente"`
	Scagnozzo      *PngPrincipale `json:"scagnozzo"`
}

func CreateOrrore(t *CrackbrainTables) *Orrore {
	return &Orrore{
		Tipo:           t.GetTipoOrrore(),
		Tratti:         strings.Join(t.GetTrattiOrrore(2), " - "),
		Luogo:          t.GetLuogo(),
		Manifestazione: CreatePng(t),
		Vittima:        CreateRandomVittima(t),
		Agente:         CreatePngPrincipale(t),
		Scagnozzo:      CreatePngPrincipale(t),
	}
}

type Scenario struct {
	Titolo     string           `json:"titolo"`
	Orrore     *Orrore          `json:"orrore"`
	Pngs       []*PngSecondario `json:"pngs"`
	Personaggi []*Personaggio   `json:"personaggi"`
}

func CreateScenario(t *CrackbrainTables) *Scenario {
	orrore := CreateOrrore(t)
	var pngs []*PngSecondario = make([]*PngSecondario, 0)

	var mainPng []*PngSecondario = make([]*PngSecondario, 0)
	var otherPng []*PngSecondario = make([]*PngSecondario, 0)

	var personaggi []*Personaggio = make([]*Personaggio, 0)

	for _, ruolo := range RuoloPng {
		mainPng = append(mainPng, CreatePngSecondario(t, ruolo))
	}

	var persons []Person = make([]Person, 0)
	for _, p := range mainPng {
		persons = append(persons, p)
	}
	persons = append(persons, orrore.Agente)
	persons = append(persons, orrore.Manifestazione)
	persons = append(persons, orrore.Vittima)

	for h := 1; h <= 5; h++ {
		var idx = t.rnd.Int63() % int64(len(persons))
		var png = persons[idx]
		var ruolo = fmt.Sprintf("Png Extra %d correlato a %s", h, png.GetName())
		var newPng = CreatePngSecondario(t, ruolo)
		otherPng = append(otherPng, newPng)
	}

	for h := 1; h <= 8; h++ {
		var p = CreatePersonaggio(t)
		personaggi = append(personaggi, p)
	}

	pngs = append(pngs, mainPng...)
	pngs = append(pngs, otherPng...)

	return &Scenario{
		Personaggi: personaggi,
		Titolo:     t.GetTitolo(),
		Orrore:     orrore,
		Pngs:       pngs,
	}
}
