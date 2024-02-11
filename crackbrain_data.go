package main

import (
	"fmt"
	"math/rand"
)

type CrackbrainTables struct {
	rnd rand.Source
}

func NewCrackbrainTables(i int64) *CrackbrainTables {
	return &CrackbrainTables{rnd: rand.NewSource(i)}
}

var NomiFemminili = []string{
	"Andrée",
	"Ange",
	"Angèle",
	"Angeline",
	"Angélique",
	"Anne",
	"Hannah",
	"Axelle",
	"Benjamine",
	"Christiane",
	"Christine",
	"Danièle",
	"Danielle",
	"Déborah",
	"Élisabeth",
	"Elizabeth",
	"Isabeau",
	"Isabelle",
	"Emmanuelle",
	"Esther",
	"Ève",
	"Gabrielle",
	"Jacqueline",
	"Jacquette",
	"Jeanne",
	"Janine",
	"Yvanne",
	"Joëlle",
	"Jordane",
	"Josée",
	"Josèphe",
	"Josèphine",
	"Judith",
	"Léa",
	"Madeleine",
	"Madeline",
	"Madelon",
	"Marlène",
	"Marie",
	"Mariamne",
	"Myriam",
	"Marianne",
	"Marthe",
	"Michèle",
	"Michèlle",
	"Natanaëlle",
	"Nicole",
	"Nicolle",
	"Noëlle",
	"Noémi",
	"Noémie",
	"Pascale",
	"Pascalle",
	"Paule",
	"Pauline",
	"Pierrette",
	"Perinne",
	"Perrine",
	"Rachel",
	"Rébecca",
	"Salomé",
	"Sara",
	"Sarah",
	"Séraphine",
	"Simone",
	"Stéphanie",
	"Étiennette",
	"Susanne",
	"Suzanne",
	"Suzette",
	"Emma", "Olivia", "Ava", "Isabella", "Sophia", "Mia", "Charlotte", "Amelia",
	"Harper", "Evelyn", "Abigail", "Emily", "Elizabeth", "Camila", "Sofia", "Avery", "Ella", "Scarlett",
	"Grace", "Hannah", "Harper", "Lily", "Ellie", "Nora", "Riley", "Penelope", "Layla", "Maya", "Zoey",
	"Eleanor", "Hazel", "Rutka", "Lillian", "Audrey", "Stella", "Aurora", "Caroline", "Brooklyn", "Violet",
	"Genesis", "Addison", "Autumn", "Willow", "Lucy", "Arianna", "Ariana", "Leah", "Alyssa", "Gabriella",
	"Naomi", "Alexis", "Alice", "Hailey Emilia", "Elizabeth", "Eva", "Sadie", "Clara", "Clara", "Eleanor",
	"Julia", "Luna", "Gianna", "Bella", "Isabelle", "Paisley", "Everly", "Adalyn", "Avery", "Serenity", "Piper",
	"Sydney", "Kaylee", "Ruby", "Peyton", "Annabelle", "Mackenzie", "Arielle", "Kennedy", "Madeline",
	"Aubree", "Sophie", "Liliana", "Jade", "Natalia", "Sophie", "Brielle", "Alexia", "Jasmine", "Valentina",
	"Savannah", "Naomi", "Madelyn", "Sofia", "Giorgia", "Francesca", "Martina", "Alessia", "Giulia",
	"Valentina", "Ludovica", "Veronica", "Elena", "Roberta", "Camilla", "Anna", "Emma", "Lena", "Mia",
	"Marie", "Sophie", "Mia", "Hannah", "Emilia", "Charlotte", "Johanna", "Franziska", "Laura", "Clara",
	"Amelie", "Lea", "Pauline", "Katharina", "Sophia", "Maria", "Leonie", "Yu", "Mei", "Ting", "Hua", "Fang",
	"Li", "Ying", "Jing", "Xin", "Xiu", "Yan", "Qing", "Xia", "Yue", "Lan", "Xue", "Rui", "Jia", "Zhen", "Hong",
	"Weiwei", "Yun", "Meiling", "Yaqi", "Xinyi", "Priya", "Sunita", "Rani", "Anjali", "Neha", "Nisha", "Kiran",
	"Amrita", "Divya", "Kavita", "Ritu", "Anita", "Yui", "Mei", "Hana", "Saki", "Riko", "Yuka", "Emi", "Aya",
	"Yumi", "Akane", "Bouchra",
	"Hasna",
	"Najet",
	"Najia",
	"Nezha",
	"Sagirah",
	"Zahra",
	"Zhora",
	"Dobryna", "Lyuba", "Militsa", "Svetlana", "Vera",
}

var NomiMaschili = []string{"William", "James", "John", "Robert", "Michael", "David", "Joseph",
	"Aaron",
	"Abel",
	"Abraham", "Abdelatif",
	"Abdeljilill",
	"Adil",
	"Ahjmed",
	"Alarbi",
	"Azedine",
	"Benyounes",
	"Brahim",
	"Driss",
	"Hachim",
	"Hamed",
	"Hamsa",
	"Hassan",
	"Hicham",
	"Jalloun",
	"Khaled",
	"Akrama",
	"Amara",
	"Bahari",
	"Behar",
	"Belaouf",
	"Belghazi",
	"Ben Bella",
	"Benarbia",
	"Bendjedid",
	"Benhamou",
	"Benida-Merah",
	"Benzai",
	"Benzine",
	"Boudiaf",
	"Boulmerka",
	"Boumedienne",
	"Bourrouag",
	"Diop",
	"Essaid",
	"Hacini",
	"Hecini",
	"Ilaes",
	"Kelkal",
	"Klouchi",
	"Krama",
	"Louahla",
	"Mammeri",
	"Matoub",
	"Merah",
	"Sahnine",
	"Saidi-Sief",
	"Selmi",
	"Yacine",
	"Zeroual",
	"Zidane",
	"Zouabri",
	"Kicham",
	"Larbi",
	"Mahjoub",
	"Mohamed",
	"Nourreddine",
	"Said",
	"Salheddine",
	"Smail",
	"Tahar",
	"Tashfin",
	"Adam",
	"André",
	"Ange",
	"Avit",
	"Axel",
	"Balthazar",
	"Baptiste",
	"Batiste",
	"Barnabé",
	"Bartholomé",
	"Bartolomé",
	"Barthélémy",
	"Benjamin",
	"Chrétien",
	"Christian",
	"Christophe",
	"Daniel",
	"David",
	"Davide",
	"Éloi",
	"Eloys",
	"Emannuel",
	"Emmanuel",
	"Esdras",
	"Etien",
	"Étienne",
	"Stéphane",
	"Stéphanne",
	"Ezéchiel",
	"Gabriel",
	"Gaspar",
	"Gáspárd",
	"Gédéon",
	"Isaak",
	"Isaie",
	"Jacques",
	"Jean",
	"Jan",
	"Yvan",
	"Jérémie",
	"Joachim",
	"Joël",
	"Jonathan",
	"José",
	"Josephe",
	"Jourdain",
	"Lazare",
	"Lévy",
	"Luc",
	"Lucas",
	"Marc",
	"Mathias",
	"Mathieu",
	"Matthieu",
	"Michel",
	"Nazaire",
	"Nicolas",
	"Noë",
	"Noël",
	"Pascal",
	"Pascale",
	"Paschal",
	"Pasquier",
	"Paul",
	"Paulin",
	"Pierre",
	"Raphaël",
	"Salomon",
	"Samuel",
	"Saul",
	"Siméon",
	"Simon",
	"Thomas",
	"Timothée",
	"Zacharie",
	"Christopher", "Thomas", "Daniel", "Steven", "Anthony", "Alistair", "Bartholomew", "Cedric",
	"Darian", "Emrys", "Fabian", "Giles", "Heathcliff", "Inigo", "Jago", "Kendric", "Leopold", "Uther",
	"Montgomery", "Peregrine", "Quillan", "Rafferty", "Sylvester", "Thaddeus", "Ulric", "Warrick",
	"Kevin", "Gislaine", "Charles", "Edward", "George", "Adam", "Benjamin", "Eric", "Paul", "Giuseppe",
	"Antonio", "Mario", "Luigi", "Giovanni", "Stefano", "Roberto", "Fabio", "Franco", "Carlo", "Angelo",
	"Davide", "Vittorio", "Davide", "Raffaele", "Juan", "José", "Antonio", "Francisco", "Manuel", "Miguel",
	"Luis", "Carlos", "David", "Fernando", "Javier", "Pedro", "Pablo", "Alejandro", "Daniel", "Rafael",
	"Gabriel", "Jorge", "Ignacio", "Sergio", "Victor", "Emilio", "Álvaro", "Diego", "Ricardo", "Yang", "Lin",
	"Guo", "Feng", "Zhou", "Wang", "Qian", "Hu", "Yu", "Shen", "Deng", "Gao", "Jiang", "Tang", "Ma", "Ivan",
	"Alexei", "Dmitry", "Nikolai", "Sergei", "Mikhail", "Andrei", "Vladimir", "Yuri", "Pavel", "Rohit", "Amit",
	"Rajesh", "Deepak", "Vikram", "Ravi", "Ankit", "Sameer", "Suresh", "Sanjay", "Pradeep", "Arvind",
	"Mohan", "Sunil", "Satish", "Sandeep", "Manish", "Rahul", "Ajay", "Vineet", "Yuto", "Takumi", "Ryota",
	"Ryo", "Yuki", "Kazuki", "Hiroki", "Daiki", "Shota", "Haruki", "Ahmed", "Mohammed", "Abdullah",
	"Omar", "Khalid", "Ali", "Hassan", "Youssef", "Saad", "Hamza",
	"Boris", "Radoslav", "Stanislav", "Vladimir", "Zoryn"}

var Cognomi = []string{"Abernathy", "Beckett",
	"Martin",
	"Bernard",
	"Thomas",
	"Petit",
	"Robert",
	"Richard",
	"Durand",
	"Dubois",
	"Moreau",
	"Laurent",
	"Simon",
	"Michel",
	"Lefebvre",
	"Leroy",
	"Roux",
	"David",
	"Bertrand",
	"Morel",
	"Fournier",
	"Girard",
	"Bonnet",
	"Dupont",
	"Lambert",
	"Fontaine",
	"Rousseau",
	"Vincent",
	"Muller",
	"Lefevre",
	"Faure",
	"Andre",
	"Mercier",
	"Blanc",
	"Guerin",
	"Boyer",
	"Garnier",
	"Chevalier",
	"Francois",
	"Legrand",
	"Gauthier",
	"Garcia",
	"Perrin",
	"Robin",
	"Clement",
	"Morin",
	"Nicolas",
	"Henry",
	"Roussel",
	"Mathieu",
	"Gautier",
	"Masson",
	"Gómez",
	"Moreno",
	"Rodriguez",
	"Cano",
	"Fernández",
	"García",
	"Suárez",
	"Marín",
	"Ruiz",
	"Alonso",
	"Iglesias",
	"Vidal",
	"Sánchez",
	"López",
	"Ramírez",
	"Álvarez",
	"Pérez",
	"Santos",
	"Gonzalez",
	"Ortiz",
	"Molina",
	"Rubio",
	"Castro",
	"Serrano",
	"Díaz",
	"Gil",
	"Torres",
	"Delgado",
	"Hernandez",
	"Romero",
	"Martínez",
	"Alegría",
	"Navarro",
	"Blanco",
	"Martin",
	"Ortega",
	"Torrente",
	"Torrez",
	"Torrico",
	"Valenzuela",
	"Valera",
	"Valeriano",
	"Valerio",
	"Valero",
	"Valiente",
	"Valladolid",
	"Valle",
	"Vallecillo",
	"Toribio",
	"Torno",
	"Toro",
	"Andric",
	"Bacid",
	"Balenovic",
	"Batelka",
	"Bilic",
	"Bogatec",
	"Broz",
	"Budisa",
	"Dezulovic",
	"Dragicevic",
	"Dubravka",
	"Frankovic",
	"Garaj",
	"Gotovac",
	"Gubec",
	"Horvat",
	"Ivanic",
	"Jarni",
	"Jovanovic",
	"Juric",
	"Kljujic",
	"Kozniku",
	"Krleza",
	"Kvaternik",
	"Leskovar",
	"Chippo",
	"El Brazi",
	"El Guerroui",
	"El Hadrioi",
	"El Kamche",
	"El Kalej",
	"El Khattavi",
	"Elfrougui",
	"El-Gharrouj",
	"El-Guerray",
	"Filali",
	"Hadda",
	"Hadji",
	"Haida",
	"Hissou",
	"Kaouch",
	"Khatabi",
	"Lahlafi",
	"Lahlou",
	"Lahsini",
	"Lahyani",
	"Lamaalem",
	"Leghzaoui",
	"Maazouzi",
	"Mellouk",
	"Morceli",
	"Mourit",
	"Moutawakel",
	"Naybet",
	"Neqrouz",
	"Nyambek",
	"Ouakili",
	"Ouaziz",
	"Rokki",
	"Sahere",
	"Sahiri",
	"Sediki",
	"Sellami",
	"Sghir",
	"Skah",
	"Sudani",
	"Lucic",
	"Mamic",
	"Maric",
	"Mavrovic",
	"Mestrovic",
	"Mladenovic",
	"Nancic",
	"Parmic",
	"Pinter",
	"Prosinecki",
	"Racan",
	"Sikavica",
	"Skansky",
	"Spegelj",
	"Stimac",
	"Talaja",
	"Tokic",
	"Vasilj",
	"Vigresic",
	"Vitezovic",
	"Vrdoljak",
	"Zamometic",
	"Carstairs", "Donnelly", "Everly", "Fairfax", "Granger", "Hawthorne", "Inglewood", "Jardine",
	"Kensington", "Larkspur", "Montague", "Seagull", "Northumberland", "Pendleton", "Radcliffe",
	"Somerset", "Townsend", "Underwood", "Winthrop", "Beauchamp", "Chateaubriand", "Dubois",
	"Fabre", "Lefebvre", "Moreau", "Navarre", "Poirier", "Remy", "Savard", "Tanguay", "Valois", "Yvon",
	"Zéphirin", "Auvray", "Barthelemy", "Cazaux", "Archibald", "Brixton", "Craven", "Davenport",
	"Eastwood", "Faraday", "Galloway", "Hollister", "Inverness", "Jardine", "Kellerman", "Lockwood",
	"MacAllister", "Northrop", "Pendragon", "Rutherford", "Sinclair", "Tremont", "Upton", "Banerjee",
	"Chakrabarti", "Desai", "Ghoshal", "Iyer", "Jaiswal", "Khanna", "Lobo", "Menon", "Nair", "Ojha",
	"Patel", "Qureshi", "Rastogi", "Shukla", "Trivedi", "Uppal", "Vaidya", "Wadhwani", "Yadav", "Alessi",
	"Basso", "Carbone", "Donati", "Esposito", "Farina", "Gallo", "Lombardo", "Marini", "Neri", "Orsini",
	"Pellegrino", "Rinaldi", "Santoro", "Terranova", "Vanni", "Zanetti", "Andreani", "Calabrese", "Lambert",
	"Giannini", "Al-Mahdi", "Barakat", "Dabous", "El-Hajj", "Fakhouri", "Ghazali", "Halabi", "Issa",
	"Jaber", "Khalaf", "Lutfi", "Mansour", "Nassar", "Odeh", "Qasem", "Rizk", "Saidi", "Taha", "Umran",
	"Zuhair", "Balle", "Christoffersen", "Dalgaard", "Enevoldsen", "Frandsen", "Gade", "Hansen",
	"Iversen", "Jensen", "Klausen", "Lindholm", "Madsen", "Nielsen", "Olsen", "Petersen", "Gutierrez",
	"Rasmussen", "Sørensen", "Thomsen", "Uldall", "Vangsgaard", "Cai", "Deng", "Fang", "Gao", "He",
	"Jiang", "Kuang", "Yao", "Zhang", "Zhuo", "Fujimura", "Goto", "Hashimoto", "Ishikawa", "Kikuchi",
	"Murakami", "Nakano", "Ozaki", "Sasaki", "Takayama", "Ueno", "Yamashita", "Yoshizawa",
	"Katzev", "Mikoyan", "Polyansky", "Simutenkov", "Dyachenko", "Baranov",
	"Landau", "Ben Zayr", "Stern", "Cohen", "Siegel", "Shapiro"}

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

func (t *CrackbrainTables) NomeCompleto(femmina bool) string {
	if femmina {
		return t.NomeFemminile() + " " + t.Cognome()
	} else {
		return t.NomeMaschile() + " " + t.Cognome()
	}
}

func (t *CrackbrainTables) NomeCasuale() string {
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

var RuoloPng = []string{
	"Parente/amico/civile/testimone",
	"Indagatore rivale, anti-PG",
	"Agente della minaccia, travestito",
	"Veggente, testimone paranormale",
	"Altra possibile vittima / redshirt"}

var BackgroundPng = []string{
	"guardia notturna",
	"infermiere generico",
	"biologo",
	"fisico teorico",
	"psichiatra",
	"bambino",
	"cacciatore di libri",
	"militare",
	"detective privato",
	"delinquente",
	"disoccupato",
	"professore",
	"sbirro",
	"sindacalista",
	"finta occupazione, agente del controspionaggio",
	"studente",
	"archeologo",
	"sacerdote/persona di chiesa",
	"giornalista",
	"avvocato",
	"paparazzo",
	"vip da due soldi",
	"conservatore",
	"corriere",
	"atleta",
	"cacciatore di criptidi",
	"allevatore",
	"sensitivo",
	"scrittore",
	"pittore",
	"musicista",
	"tassista",
	"sex worker",
	"vagabondo",
	"ingegnere (informatico se coerente con il periodo)",
	"impiegato"}

func (t *CrackbrainTables) GetBakgroundPng() string {
	return t.GetRandomItem(BackgroundPng)
}

var OggettoPersonale = []string{
	"colt M1911 “Rei Ayanami edition”",
	"moneta d’argento portafortuna",
	"cravatta nera con spilla",
	"due dadi truccati",
	"vecchia torcia a pile",
	"bicicletta rossa",
	"cocaina in una scatoletta di latta",
	"piastrine di riconoscimento",
	"utilitaria",
	"vecchia muscle car scassata",
	"puntatore laser",
	"chopstick tra i capelli",
	"pendaglio di giada cinese",
	"occhiali da sole di marca",
	"sigarette",
	"coltellino svizzero da boyscout",
	"tarocchi",
	"bibbia",
	"bibbia con fiaschetta nascosta",
	"nunchaku",
	"Canon e teleobbiettivo",
	"espressione seducente",
	"7 icone religiose diverse al collo",
	"cappello di stagnola",
	"scarpe da corsa fortunate",
	"antidolorifici",
	"rolex",
	"ciocca di capelli",
	"stilografica",
	"portachiavi in corda",
	"strumento musicale",
	"foto dell’amico perduto",
	"manette pelose",
	"accendino personalizzato",
	"Game Boy color",
	"erba (qualità buona)"}

func (t *CrackbrainTables) GetOggettoPersonale() string {
	return t.GetRandomItem(OggettoPersonale)
}

var Luoghi = []string{
	"libreria universitaria",
	"grand hotel",
	"laboratorio",
	"fast food, in centro",
	"quel locale dall’altra parte della città",
	"al parchetto, era così inquietante",
	"parcheggio dello stadio",
	"supermercato, tra i surgelati",
	"locale sinistro, tra brutti ceffi",
	"karaoke",
	"cimitero",
	"nel suo ufficio",
	"villa con piscina, che vista!",
	"in piazza, in mezzo alla gente",
	"in un ingorgo stradale",
	"su una limousine, con 3 sconosciuti",
	"piscina coperta",
	"discoteca popolare",
	"locale ricercato, hipster e storici",
	"nel suo appartamento",
	"in chiesa, di notte",
	"in periferia, sul cavalcavia",
	"in tv, sul quinto canale",
	"sala giochi, cabinato di Metal Slug",
	"grosso cantiere, sotto il grattacielo",
	"al cinema",
	"dentro un tombino, nelle fogne",
	"alla discarica",
	"blocco abitativo iper-popolato",
	"in metro",
	"magazzino merci, fuori città",
	"ospedale centrale",
	"clinica infame",
	"vicolo dietro a (ritira, aggiungi)",
	"al molo (c’è un molo?)",
	"squallido motel"}

func (t *CrackbrainTables) GetLuogo() string {
	return t.GetRandomItem(Luoghi)
}

var TipiDiOrrore = []string{
	"Demone, entità mistica, riportato nei testi sacri, risultato di un’invocazione, centro di culto",
	"Serial killer, psicopatico da film “slasher”",
	"Mutazione, mostro creato dalla scienza, vita sintetica",
	"Creatura mostruosa, predatore,mangia-uomini",
	"Spirito vendicativo, entità manipolatrice, incorporeo, possibile conseguenza di una maledizione",
	"Manifestazione di intelletto alieno, divinità extra-dimensionale,“lovecraftiana”"}

var TrattiDiOrrore = []string{
	"forza sovrumana",
	"invisibile",
	"controllo mentale",
	"pirocinesi",
	"immortale",
	"liquido, gelatinoso",
	"sciame",
	"telecinesi",
	"cannibale",
	"clonazione",
	"allucinatorio",
	"mimic, imitatore",
	"secrezioni acide - tossina",
	"magnetismo",
	"metacreatività",
	"manipolazione onirica",
	"contagio, riproduzione infettiva",
	"continua mutazione",
	"volante",
	"armamento militare d’avanguardia",
	"uncini, seghe, corpi macellati",
	"collegato a oggetto/luogo/persona",
	"burocrazia diabolica",
	"digitale",
	"risucchio vitale",
	"gigantesco, tentacolare",
	"segreto della comunità",
	"sensuale",
	"intelligenza occulta",
	"lenta, agonica corruzione",
	"inghiotte i corpi per interi",
	"sfortuna, malocchio",
	"perdita di umanità",
	"sanguisuga",
	"ombra",
	"natura quadridimensionale",
}

var Relazioni = []string{
	"potere",
	"sentimentale",
	"sottomissione",
	"sfida",
	"fiducia",
	"amicizia",
	"odio",
	"dominio",
	"fede assoluta",
	"disprezzo"}

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
