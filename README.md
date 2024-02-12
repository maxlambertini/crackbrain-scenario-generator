# Crackbrain Scenario Generator

*Crackbrain Scenario Generator* è una semplice webapp scritta in Go che implementa le regole di generazione di uno scenario per il GDR Indie [Crackbrain](https://luca-negri.itch.io/crackbrain). 

L'app può essere usata così com'è. Se vi sentite smanettoni, essa espone tre endpoint:


| Endpoint | Effetto |
|----------|---------|
| `GET /`    | Mostra un nuovo scenario generato _pseudo_casualmente |
|`GET /api`  | Genera un nuovo scenario ed espone i dati in formato json. Utile se volete implementare il vostro **Crackbrain as a Service** |
| `GET /typst`  | Genera un nuovo scenario e in output restituisce un file **typst**, ovvero un simil-markdown pronto per essere processato dal sistema di page layout [typst](https://typst.app). Provatelo, è facile e potente! Un esempio del file generato lo trovate in [sample.typ](./sample.typ) e il PDF corrispondente in [sample.pdf](./sample.pdf)|

Tutti gli endpoint elencati possono prendere il parametro `?seed=<int64>` in query string. Essendo il generatore di numeri casuali **pseudocasuale**, assegnando un valore intero a 64 bit come numero di partenza questo genererà lo stesso scenario.

Se volete provare l'app andate qui: https://crackbrain.xoxarle.com 
