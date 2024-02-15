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

= La Prospettiva  Horvat

uno scenario per Crackbrain

== L'orrore

L'orrore è: *Mutazione, mostro creato dalla scienza, vita sintetica*. I suoi tratti sono: *sensuale - continua mutazione* e il suo luogo preferito è *una serie di case occupate da diseredati*. 

L'orrore si manifesta attraverso la seguente creatura mortale: *Isabelle Torrico*, 
che ha come background *paparazzo*. Come oggetto feticcio ha *una foto dell’amico perduto*, 
e quando è in forma "normale" ha un'indole *freddo - saccente*, 

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

La vittima dell'orrore si chiama *Sydney Gómez*, 
che ha come background *sacerdote/persona di chiesa*. Come oggetto feticcio aveva *delle sigarette*, 
e chi la conosceva asserisce che la sua indole era  *difficile - dimenticabile*. 
E' stata ritrovata in *un blocco abitativo iper-popolato*. 

Con l'orrore aveva una relazione di *odio - fede assoluta*. 

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

L'agente principale al servizio dell'orrore si chiama *Siméon Balenovic*, 
che ha come background *fisico teorico*. Come oggetto feticcio ha *un coltellino svizzero da boyscout*, 
La sua indole è  *avido - naif*. 

Con l'orrore ha una relazione di *disprezzo - fede assoluta*. 

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

#let pngBlock(Nome:"", Background:"", Oggetto:"", Indole:"", Ruolo:"", Relazione:"") = {
figure(
block(
width: 100%,
[
#set align(left)
=== #Ruolo

- Nome *#Nome*
- Background *#Background*
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

pngBlock(Nome:       [ Ava Chippo ], 
        Background: [ sacerdote/persona di chiesa ] , 
        Oggetto   : [ una bicicletta rossa ], 
        Indole    : [ dimenticabile - furbo ], 
        Ruolo     : [ Parente/amico/civile/testimone ], 
        Relazione : [ sottomissione]
),

pngBlock(Nome:       [ Addison Valle ], 
        Background: [ musicista ] , 
        Oggetto   : [ un coltellino svizzero da boyscout ], 
        Indole    : [ deciso - furbo ], 
        Ruolo     : [ Indagatore rivale, anti-PG ], 
        Relazione : [ sottomissione]
),

pngBlock(Nome:       [ Bahari Yvon ], 
        Background: [ giornalista ] , 
        Oggetto   : [ un pendaglio di giada cinese ], 
        Indole    : [ affascinante - attento ], 
        Ruolo     : [ Agente della minaccia, travestito ], 
        Relazione : [ potere]
),

pngBlock(Nome:       [ Rébecca Landau ], 
        Background: [ psichiatra ] , 
        Oggetto   : [ un&#39;utilitaria ], 
        Indole    : [ aggressivo - attento ], 
        Ruolo     : [ Veggente, testimone paranormale ], 
        Relazione : [ fiducia]
),

pngBlock(Nome:       [ Guo Vitezovic ], 
        Background: [ musicista ] , 
        Oggetto   : [ uno strumento musicale ], 
        Indole    : [ saccente - freddo ], 
        Ruolo     : [ Altra possibile vittima / redshirt ], 
        Relazione : [ sentimentale]
),

pngBlock(Nome:       [ Angélique Vitezovic ], 
        Background: [ sindacalista ] , 
        Oggetto   : [ due dadi truccati ], 
        Indole    : [ deciso - difficile ], 
        Ruolo     : [ Png Extra 1 correlato a Addison Valle ], 
        Relazione : [ sentimentale]
),

pngBlock(Nome:       [ Brielle Yadav ], 
        Background: [ vip da due soldi ] , 
        Oggetto   : [ un rolex ], 
        Indole    : [ attento - aggressivo ], 
        Ruolo     : [ Png Extra 2 correlato a Guo Vitezovic ], 
        Relazione : [ odio]
),

pngBlock(Nome:       [ Thomas Elfrougui ], 
        Background: [ guardia notturna ] , 
        Oggetto   : [ un rolex ], 
        Indole    : [ difficile - perditempo ], 
        Ruolo     : [ Png Extra 3 correlato a Siméon Balenovic ], 
        Relazione : [ odio]
),

pngBlock(Nome:       [ Adam Pendragon ], 
        Background: [ fisico teorico ] , 
        Oggetto   : [ una foto dell’amico perduto ], 
        Indole    : [ inafferrabile - saccente ], 
        Ruolo     : [ Png Extra 4 correlato a Siméon Balenovic ], 
        Relazione : [ fede assoluta]
),

pngBlock(Nome:       [ Alessia Cano ], 
        Background: [ vip da due soldi ] , 
        Oggetto   : [ un coltellino svizzero da boyscout ], 
        Indole    : [ naif - inafferrabile ], 
        Ruolo     : [ Png Extra 5 correlato a Addison Valle ], 
        Relazione : [ sentimentale]
),

)

== Note

#block(
    width: 100%,
    height: 6cm,
    stroke: (0.5pt +black)
)

#pagebreak()

== Personaggi

#v(12pt)

#let personaggio(Nome:"", Background:"", Oggetto:"", Stile:"", Descrizione:"") = {
    figure(
        block(
        width: 100%    ,
        fill: rgb("#F5F5F5"),
        inset: 5pt,
        [
        #set align(left)
#text( size:14pt, weight: 700, [ #Nome (#Background) ])

- Stile: *#Stile* : #Descrizione
- Oggetto: #Oggetto

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

personaggio(Nome:       [ Perrine Moreno ], 
        Background:  [ atleta ] , 
        Oggetto:    [ una vecchia muscle car scassata ], 
        Stile:       [ Bruto ], 
        Descrizione: [ Il bruto affronta il caso di petto, facendo valere la propria prestanza fisica e capacità intimidatoria. ], 
),

personaggio(Nome:       [ Scarlett Maric ], 
        Background:  [ corriere ] , 
        Oggetto:    [ un accendino personalizzato ], 
        Stile:       [ Abile ], 
        Descrizione: [ L’abile ha sempre un asso nella manica, percorre la città inosservato ed è molto probabile che sappia come accendere un’auto senza usare le chiavi. ], 
),

personaggio(Nome:       [ Harper Carstairs ], 
        Background:  [ ingegnere (informatico se coerente con il periodo) ] , 
        Oggetto:    [ due dadi truccati ], 
        Stile:       [ Abile ], 
        Descrizione: [ L’abile ha sempre un asso nella manica, percorre la città inosservato ed è molto probabile che sappia come accendere un’auto senza usare le chiavi. ], 
),

personaggio(Nome:       [ Saidi-Sief Perrin ], 
        Background:  [ ingegnere (informatico se coerente con il periodo) ] , 
        Oggetto:    [ unìespressione seducente ], 
        Stile:       [ Zero stile ], 
        Descrizione: [ non tutti nascono con talenti o inclinazioni. Stacci. ], 
),

personaggio(Nome:       [ Elizabeth Nancic ], 
        Background:  [ avvocato ] , 
        Oggetto:    [ una bibbia ], 
        Stile:       [ Medium ], 
        Descrizione: [ Il medium sfrutta il suo collegamento con il paranormale per individuare, comprendere e non farsi uccidere da quello che scienza e buon senso non possono spiegare. ], 
),

personaggio(Nome:       [ Marie Garcia ], 
        Background:  [ tassista ] , 
        Oggetto:    [ un coltellino svizzero da boyscout ], 
        Stile:       [ Bruto ], 
        Descrizione: [ Il bruto affronta il caso di petto, facendo valere la propria prestanza fisica e capacità intimidatoria. ], 
),

personaggio(Nome:       [ Lily Jensen ], 
        Background:  [ finta occupazione, agente del controspionaggio ] , 
        Oggetto:    [ una bicicletta rossa ], 
        Stile:       [ Bruto ], 
        Descrizione: [ Il bruto affronta il caso di petto, facendo valere la propria prestanza fisica e capacità intimidatoria. ], 
),

personaggio(Nome:       [ Priya Morceli ], 
        Background:  [ professore ] , 
        Oggetto:    [ un Game Boy color ], 
        Stile:       [ Abile ], 
        Descrizione: [ L’abile ha sempre un asso nella manica, percorre la città inosservato ed è molto probabile che sappia come accendere un’auto senza usare le chiavi. ], 
),

)
