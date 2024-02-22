package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
)

type CrackbrainTables struct {
	Seed int64
	Url  string
	rnd  rand.Source
}

func TextFileToArray(filename string) []string {
	fmt.Println("Loading data file: ", filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Error reading %s", filename))
	}
	lines := strings.Split(string(content), "\n")
	for i, s := range lines {
		lines[i] = strings.TrimSpace(s)
	}
	return lines
}

func NewCrackbrainTables(i int64) *CrackbrainTables {
	return &CrackbrainTables{Seed: i, rnd: rand.NewSource(i)}
}

func GetExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	return exPath
}

var NomiFemminili = TextFileToArray(fmt.Sprintf("%s/data/names_f.txt", GetExecutablePath()))
var NomiMaschili = TextFileToArray(fmt.Sprintf("%s/data/names_m.txt", GetExecutablePath()))
var Cognomi = TextFileToArray(fmt.Sprintf("%s/data/surnames.txt", GetExecutablePath()))
var Aspetto = TextFileToArray(fmt.Sprintf("%s/data/look.txt", GetExecutablePath()))

func (t *CrackbrainTables) GetRandomItem(data []string) string {
	l := len(data)
	idx := t.rnd.Int63() % int64(l)
	return data[idx]
}

func (t *CrackbrainTables) NomeFemminile() string {
	return t.GetRandomItem(NomiFemminili)
}

func (t *CrackbrainTables) NomeMaschile() string {
	return t.GetRandomItem(NomiMaschili)
}

func (t *CrackbrainTables) Cognome() string {
	return t.GetRandomItem(Cognomi)
}

func (t *CrackbrainTables) NomeCompleto(femmina bool) (string, string) {
	if femmina {
		return t.NomeFemminile() + " " + t.Cognome(), "F"
	} else {
		return t.NomeMaschile() + " " + t.Cognome(), "M"
	}
}

func (t *CrackbrainTables) NomeCasuale() (string, string) {
	idx := t.rnd.Int63() % 100
	return t.NomeCompleto(idx < 53)
}

var IndolePng = []string{
	"aggressivo", "mutevole", "difficile",
	"deciso", "intraprendente", "inafferrabile",
	"avido", "furbo", "dimenticabile",
	"manipolatore", "affascinante", "freddo",
	"naif", "propositivo", "sfortunato",
	"saccente", "perditempo", "attento"}

func (t *CrackbrainTables) Shuffle(array []string) {
	for i := len(array) - 1; i > 0; i-- {
		j := t.rnd.Int63() % int64(i+1)
		fmt.Println("shuffling ", j)
		array[i], array[j] = array[j], array[i]
	}
}

func (t *CrackbrainTables) NdShuffle(array []string, sliceSize int) []string {
	arrayTmp := make([]string, len(array))
	copy(arrayTmp, array)
	for i := len(arrayTmp) - 1; i > 0; i-- {
		j := t.rnd.Int63() % int64(i+1)
		arrayTmp[i], arrayTmp[j] = arrayTmp[j], arrayTmp[i]
	}
	return arrayTmp[0:sliceSize]
}

func (t *CrackbrainTables) GetIndolePng(n int) []string {
	return t.NdShuffle(IndolePng, n)
}

//"Agente della minaccia",
//"Agente della minaccia, travestito"
//	"Manifestazione dell’orrore",

var RuoloPng = TextFileToArray(fmt.Sprintf("%s/data/roles.txt", GetExecutablePath()))

var BackgroundPng = TextFileToArray(fmt.Sprintf("%s/data/backgrounds.txt", GetExecutablePath()))

func (t *CrackbrainTables) GetBakgroundPng() string {
	return t.GetRandomItem(BackgroundPng)
}

var OggettoPersonale = TextFileToArray(fmt.Sprintf("%s/data/objects.txt", GetExecutablePath()))

func (t *CrackbrainTables) GetOggettoPersonale() string {
	return t.GetRandomItem(OggettoPersonale)
}

func (t *CrackbrainTables) GetAspetto() string {
	return t.GetRandomItem(Aspetto)
}

var Luoghi = TextFileToArray(fmt.Sprintf("%s/data/places.txt", GetExecutablePath()))

func (t *CrackbrainTables) GetLuogo() string {
	s := t.GetRandomItem(Luoghi)
	s1 := ""
	s = strings.Replace(s, "YYY", t.Cognome(), 1)
	if strings.Contains(s, "XXX") {
		for {
			s1 = t.GetRandomItem(Luoghi)
			if s1 != s {
				break
			}
		}

		s = strings.Replace(s, "XXX", s1, 1)
	}
	return s
}

func (t *CrackbrainTables) GetLuoghi(n int) []string {
	lg := t.NdShuffle(Luoghi, n)
	for h := 0; h < n; h++ {
		s := lg[h]
		s1 := ""
		s = strings.Replace(s, "YYY", t.Cognome(), 1)
		if strings.Contains(s, "XXX") {
			for {
				s1 = t.GetRandomItem(Luoghi)
				if s1 != s {
					break
				}
			}

			s = strings.Replace(s, "XXX", s1, 1)
		}
		lg[h] = s
	}
	return lg
}

var TipiDiOrrore = TextFileToArray(fmt.Sprintf("%s/data/horror_types.txt", GetExecutablePath()))

var TrattiDiOrrore = TextFileToArray(fmt.Sprintf("%s/data/horror_traits.txt", GetExecutablePath()))

var Relazioni = TextFileToArray(fmt.Sprintf("%s/data/relations.txt", GetExecutablePath()))

func (t *CrackbrainTables) GetRelazioni(n int) []string {
	return t.NdShuffle(Relazioni, n)
}

func (t *CrackbrainTables) GetTrattiOrrore(n int) []string {
	return t.NdShuffle(TrattiDiOrrore, n)
}

func (t *CrackbrainTables) GetTipoOrrore() string {
	return t.GetRandomItem(TipiDiOrrore)
}

func (t *CrackbrainTables) GetPngBackground() string {
	return t.GetRandomItem(BackgroundPng)
}

func (t *CrackbrainTables) GetPngOggetti() string {
	return t.GetRandomItem(OggettoPersonale)
}

var EsempioTitolo = []string{
	"Il Caso ",
	"La Prospettiva ",
	"L'Enigma ",
	"Il Mistero ",
}

func (t *CrackbrainTables) GetTitolo() string {
	return t.GetRandomItem(EsempioTitolo) + " " + t.Cognome()
}

type StileInvestigativo struct {
	Nome        string `json:"nome"`
	Descrizione string `json:"descrizione"`
}

var StiliInvestigativi = []StileInvestigativo{
	{Nome: "Bruto", Descrizione: "Il bruto affronta il caso di petto, facendo valere la propria prestanza fisica e capacità intimidatoria."},
	{Nome: "Abile", Descrizione: "L’abile ha sempre un asso nella manica, percorre la città inosservato ed è molto probabile che sappia come accendere un’auto senza usare le chiavi."},
	{Nome: "Sapiente", Descrizione: "Riflessivo, intelligente, logico, deduttivo. Maestro nel riconoscere pattern e collegamenti."},
	{Nome: "Medium", Descrizione: "Il medium sfrutta il suo collegamento con il paranormale per individuare, comprendere e non farsi uccidere da quello che scienza e buon senso non possono spiegare."},
	{Nome: "Politico", Descrizione: "Il politico costruisce e distrugge con la sua eloquenza e abilità attoriale. Diplomatico,carismatico. Bugiardo. "},
	{Nome: "Zero stile", Descrizione: "non tutti nascono con talenti o inclinazioni. Stacci."},
}

func (t *CrackbrainTables) GetStileInvestigativo() *StileInvestigativo {
	return &StiliInvestigativi[t.rnd.Int63()%int64(len(StiliInvestigativi))]
}

var RiferimentoTitolo = []string{
	"bar preferito dalla vittima",
	"locale preferito dalla vittima",
	"bar preferito dalla minaccia",
	"locale preferito di uno dei giocatori",
	"bar preferito di uno dei png",
	"nome di un animale domestico di uno dei giocatori",
	"nome di un animale domestico di uno dei png",
	"nome della residenza della minaccia",
	"nome della residenza della vittima",
	"nome della residenza di uno dei png",
	"nome della residenza di uno dei giocatori",
}

/*
class Png(object):
    def __init__(self):
        self.nome = nomeCasuale()
        self.background = getPngBackground()
        self.oggetto = getPngOggetti()
        self.indole = " - ".join(getIndolePng())

    def toDict(self):
       return {
          "nome" : self.nome,
          "background" : self.background,
          "oggetto" : self.oggetto,
          "indole" : self.indole
       }


class Vittima(object):
    def __init__(self):
        self.personaggio = Png();
        self.luogoRitrovamento = getLuogo()

    def toDict(self):
       dict = self.personaggio.toDict()
       dict["luogoRitrovamento"] = self.luogoRitrovamento
       return dict


def getPngPerRuoli():
    res = {}
    for k in ruoloPng:
       res[k] = Png().toDict()

    for h in range(1,6):
       k = getRandomItem(ruoloPng)
       p = res[k]
       if p != None:
          l = "Png Extra {} legato a {}".format(h,p["nome"])
          pn = Png()
          res[l] = pn.toDict()
    return [ { "ruolo" : l, "png" : res[l] , "relazioni": getRelazioni(1)} for l in res.keys()]


class Orrore(object):
    def __init__(self):
        self.tipo = getTipoOrrore();
        self.tratti = " - ".join(getTrattiOrrore())
        self.luogo = getLuogo();
        self.manifestazione = Png()
        self.vittima = Vittima()
        self.relazione_vittima = getRelazioni()
        self.agente = Png()
        self.relazione_agente = getRelazioni(2)
        self.scagnozzo= Png()
        self.relazione_scagnozzo = getRelazioni(2)

    def toDict(self):
       return {
          "tipo" : self.tipo,
          "tratti" : self.tratti,
          "luogo" : self.luogo,
          "manifestazione" : self.manifestazione.toDict(),
          "vittima" : {
             "png": self.vittima.toDict(),
             "relazione" :self.relazione_vittima
          },
          "agente" : {
             "png": self.agente.toDict(),
             "relazione" :self.relazione_agente
          },
          "scagnozzo" : {
             "png": self.scagnozzo.toDict(),
             "relazione" :self.relazione_scagnozzo
          },
       }

class Scenario(object):
    def __init__(self):
        self.orrore = Orrore();
        self.pngs = getPngPerRuoli();

    def toDict(self):
       return {
          "orrore" : self.orrore.toDict(),
          "pngs" : self.pngs
       }

s = Scenario()
j = json.dumps(s.toDict(), indent=2,ensure_ascii=False)
with open ("test2.json","w", encoding="utf-8") as f:
   f.write(j)
*/
