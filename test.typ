#set page(
  width:21cm,
  height:29.7cm,
  margin: (x: 1.5cm, y: 1.5cm),
  numbering: "1"
)

#set text(lang: "it", size: 11pt)
#set text(font: "Chivo")

#set par(
  justify: true,
  leading: 0.52em,

)

= Il Caso  Boyer

uno scenario per Crackbrain

== L'orrore

L'orrore è: *Creatura mostruosa, predatore,mangia-uomini*. I suoi tratti sono: *risucchio vitale - burocrazia diabolica* e il suo luogo preferito è *in metro*. 

L'orrore si manifesta attraverso la seguente creatura mortale: *Madelyn Torrez -- F*, 
che ha come background *sindacalista*. Come oggetto feticcio ha *una bibbia*, 
In forma umana ha il seguente aspetto: *completo grigio, spalle larghe*.
e quando è in forma "normale" ha un'indole *freddo - inafferrabile*, 

=== Note ulteriori sull'orrore:

Come si manifesta? Quali sono le relazioni con la vittima e i suoi avatar sulla terra? 

#figure(
block(
    width: 100%    ,
    height: 5cm,
    stroke: (0.5pt +black)
)
)

#grid(
columns: (1fr, 1fr),
gutter: 5mm,
[
=== La vittima

La vittima dell'orrore si chiama *Ange Richard -- F*, 
che ha come background *professore*. Come oggetto feticcio aveva *un nunchaku*, 
Chi la conosceva diceva che aveva questo aspetto: *sbronzo/a, giacca a quadri*
e la sua indole era  *affascinante - inafferrabile*. 
E' stata ritrovata in *un magazzino merci, fuori città*. 

Con l'orrore aveva una relazione di *sentimentale - fede assoluta*. 

#figure(
block(
    width: 100%    ,
    height: 6cm,
    stroke: (0.5pt +black)
)
)
],
[
=== l'agente dell'Orrore

L'agente principale al servizio dell'orrore si chiama *Adam Cohen -- M*, 
Questo è il suo aspetto: *assomiglia a qualcuno/a*
che ha come background *cacciatore di libri*. Come oggetto feticcio ha *una stilografica*, 
La sua indole è  *attento - dimenticabile*. 

Con l'orrore ha una relazione di *amicizia - potere*. 

#figure(
block(
    width: 100%    ,
    height: 6cm,
    stroke: (0.5pt +black)
)
)
])

#pagebreak()





== Altri Png

#let pngBlock(Nome:"", Background:"", Aspetto: "", Oggetto:"", Indole:"", Ruolo:"", Relazione:"") = {
figure(
block(
width: 100%,
[
#set align(left)
=== #Ruolo

- Nome *#Nome*
- Background *#Background*
- Aspetto *#Aspetto*
- Oggetto *#Oggetto*
- Indole *#Indole*
- Relazione *#Relazione*

#figure(
    block(
        width: 100%    ,
        height: 6cm,
        stroke: (0.5pt +black)
    )
)

]
)
)
}

Lista dei personaggi dello scenario

#grid(
columns: (1fr, 1fr),
gutter: 5mm,

pngBlock(Nome:       [ Madeleine Moreau -- F ], 
        Background: [ impiegato ] , 
        Aspetto   : [ parla poco, stringe un pupazzo ], 
        Oggetto   : [ dei tarocchi ], 
        Indole    : [ propositivo - avido ], 
        Ruolo     : [ Parente/amico/civile/testimone ], 
        Relazione : [ dominio]
),

pngBlock(Nome:       [ Hassan Bernard -- M ], 
        Background: [ delinquente ] , 
        Aspetto   : [ circondato/a da animali ], 
        Oggetto   : [ una colt M1911 “Rei Ayanami edition” ], 
        Indole    : [ perditempo - difficile ], 
        Ruolo     : [ Indagatore rivale anti-PG ], 
        Relazione : [ potere]
),

pngBlock(Nome:       [ Jacques Yadav -- M ], 
        Background: [ quadro aziendale ] , 
        Aspetto   : [ sbronzo/a, giacca a quadri ], 
        Oggetto   : [ due dadi truccati ], 
        Indole    : [ dimenticabile - freddo ], 
        Ruolo     : [ Agente della minaccia travestito ], 
        Relazione : [ potere]
),

pngBlock(Nome:       [ Eric Mansour -- M ], 
        Background: [ professore ] , 
        Aspetto   : [ noodles istantanei in mano ], 
        Oggetto   : [ dell’erba (qualità buona) ], 
        Indole    : [ affascinante - attento ], 
        Ruolo     : [ Veggente testimone paranormale ], 
        Relazione : [ sentimentale]
),

pngBlock(Nome:       [ Veronica Nakano -- F ], 
        Background: [ squatter ] , 
        Aspetto   : [ capelli a spazzola o raccolti in un codino, marziale, vene sul collo ], 
        Oggetto   : [ una bibbia ], 
        Indole    : [ intraprendente - mutevole ], 
        Ruolo     : [ Altra possibile vittima / redshirt ], 
        Relazione : [ amicizia]
),

pngBlock(Nome:       [ Leah Girard -- F ], 
        Background: [ ingegnere ] , 
        Aspetto   : [ imponente, raffreddato/a ], 
        Oggetto   : [ un cappello di stagnola ], 
        Indole    : [ dimenticabile - affascinante ], 
        Ruolo     : [ Png Extra 1 correlato a Adam Cohen ], 
        Relazione : [ amicizia]
),

pngBlock(Nome:       [ Uther Lambert -- M ], 
        Background: [ allevatore ] , 
        Aspetto   : [ peli ovunque (o barba se uomo), enormi braccia ], 
        Oggetto   : [ un pendaglio di giada cinese ], 
        Indole    : [ mutevole - aggressivo ], 
        Ruolo     : [ Png Extra 2 correlato a Madelyn Torrez ], 
        Relazione : [ fede assoluta]
),

pngBlock(Nome:       [ Martina Vidal -- F ], 
        Background: [ vip da due soldi ] , 
        Aspetto   : [ noodles istantanei in mano ], 
        Oggetto   : [ delle manette pelose ], 
        Indole    : [ deciso - affascinante ], 
        Ruolo     : [ Png Extra 3 correlato a Ange Richard ], 
        Relazione : [ fiducia]
),

pngBlock(Nome:       [ Feng Morin -- M ], 
        Background: [ militare ] , 
        Aspetto   : [ minuto/a, carino/a, sboccato/a ], 
        Oggetto   : [ una Canon e teleobbiettivo ], 
        Indole    : [ freddo - difficile ], 
        Ruolo     : [ Png Extra 4 correlato a Jacques Yadav ], 
        Relazione : [ odio]
),

pngBlock(Nome:       [ Nourreddine Gauthier -- M ], 
        Background: [ fisico teorico ] , 
        Aspetto   : [ cicatrice, minaccioso/a ], 
        Oggetto   : [ due dadi truccati ], 
        Indole    : [ manipolatore - saccente ], 
        Ruolo     : [ Png Extra 5 correlato a Madeleine Moreau ], 
        Relazione : [ disprezzo]
),

)

== Note

#block(
    width: 100%,
    height: 6cm,
    stroke: (0.5pt +black)
)

#pagebreak()

== Luoghi Extra


#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ una piazza, in mezzo alla gente]
)

#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ un grosso cantiere, sotto il grattacielo]
)

#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ un albergo per rottami umani]
)

#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ un appartamento, il suo]
)

#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ un laboratorio]
)

#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ un parcheggio dello stadio]
)

#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ un blocco abitativo iper-popolato]
)

#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ un cinema]
)

#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ un parchetto, era così inquietante]
)

#block(
    width: 100%    ,
    height: 2cm,
    inset: 0.2cm,
    stroke: (0.5pt +black),
    [ un locale dall’altra parte della città]
)


#pagebreak()

== Personaggi

#v(12pt)

#let personaggio(Nome:"", Background:"", Aspetto: "",Oggetto:"", Stile:"", Descrizione:"") = {
    figure(
        block(
        width: 100%    ,
        fill: rgb("#F5F5F5"),
        inset: 5pt,
        [
        #set align(left)
#text( size:14pt, weight: 700, [ #Nome (#Background) ])

- Aspetto: *#Aspetto*
- Stile: *#Stile* : #Descrizione
- Oggetto: *#Oggetto*

#block(
    width: 100%,
    height: 2cm,
    inset: 5pt,
    stroke: (0.5pt + gray),
    fill: white,
    [ _Legami_ ]
)

#block(
    width: 100%,
    height: 5cm,
    inset: 5pt,
    stroke: (0.5pt + gray),
    fill: white,
    [ _Note_ ]
)
        ]
    )
    )
}


#grid(
columns: (1fr, 1fr),
gutter: 5mm,

personaggio(Nome:       [ Sahnine Toribio -- M ], 
        Background:  [ impiegato ] , 
        Oggetto:    [ un Game Boy color ], 
        Aspetto:    [ al telefono, riflesso sugli occhiali ], 
        Stile:       [ Sapiente ], 
        Descrizione: [ Riflessivo, intelligente, logico, deduttivo. Maestro nel riconoscere pattern e collegamenti. ], 
),

personaggio(Nome:       [ Danielle Martínez -- F ], 
        Background:  [ detective privato ] , 
        Oggetto:    [ un coltellino svizzero da boyscout ], 
        Aspetto:    [ al telefono, moshi moshi ], 
        Stile:       [ Politico ], 
        Descrizione: [ Il politico costruisce e distrugge con la sua eloquenza e abilità attoriale. Diplomatico,carismatico. Bugiardo.  ], 
),

personaggio(Nome:       [ Ryota Ghoshal -- M ], 
        Background:  [ professore ] , 
        Oggetto:    [ un’espressione seducente ], 
        Aspetto:    [ filiforme, occhiaie, “burtonesco/a” ], 
        Stile:       [ Zero stile ], 
        Descrizione: [ non tutti nascono con talenti o inclinazioni. Stacci. ], 
),

personaggio(Nome:       [ Bartolomé Murakami -- M ], 
        Background:  [ impiegato ] , 
        Oggetto:    [ un pendaglio di giada cinese ], 
        Aspetto:    [ senza fiato, non parla la tua lingua ], 
        Stile:       [ Abile ], 
        Descrizione: [ L’abile ha sempre un asso nella manica, percorre la città inosservato ed è molto probabile che sappia come accendere un’auto senza usare le chiavi. ], 
),

personaggio(Nome:       [ Yun Lambert -- F ], 
        Background:  [ biologo ] , 
        Oggetto:    [ un’espressione seducente ], 
        Aspetto:    [ logorroico/a, razzista, psp in mano ], 
        Stile:       [ Medium ], 
        Descrizione: [ Il medium sfrutta il suo collegamento con il paranormale per individuare, comprendere e non farsi uccidere da quello che scienza e buon senso non possono spiegare. ], 
),

personaggio(Nome:       [ Javier Issa -- M ], 
        Background:  [ psichiatra ] , 
        Oggetto:    [ una ciocca di capelli ], 
        Aspetto:    [ capelli vistosi, fissato/a con gli alieni ], 
        Stile:       [ Politico ], 
        Descrizione: [ Il politico costruisce e distrugge con la sua eloquenza e abilità attoriale. Diplomatico,carismatico. Bugiardo.  ], 
),

personaggio(Nome:       [ Penelope Jardine -- F ], 
        Background:  [ archeologo ] , 
        Oggetto:    [ una cravatta nera con spilla ], 
        Aspetto:    [ filiforme, occhiaie, “burtonesco/a” ], 
        Stile:       [ Sapiente ], 
        Descrizione: [ Riflessivo, intelligente, logico, deduttivo. Maestro nel riconoscere pattern e collegamenti. ], 
),

personaggio(Nome:       [ Ting Baranov -- F(H) ], 
        Background:  [ studente ] , 
        Oggetto:    [ dei tarocchi ], 
        Aspetto:    [ un gatto/un cane/un pappagallo ], 
        Stile:       [ Sapiente ], 
        Descrizione: [ Riflessivo, intelligente, logico, deduttivo. Maestro nel riconoscere pattern e collegamenti. ], 
),

)
