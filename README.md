# Libretto Universitario

## Panoramica

Questo progetto consiste in un'applicazione web che permette di visualizzare i voti universitari. 
Gli utenti possono visualizzare un grafico dell'andamento dei voti, l'andamento dei voti di laurea, e una tabella dettagliata degli esami sostenuti. 
L'applicazione è costruita con un backend Go e un frontend HTML, CSS con JavaScript per le visualizzazioni.

## Caratteristiche

- Visualizzazione grafica dei voti e dei voti di laurea.
- Tabella dei voti con dettagli degli esami.
- Selettore a tendina per passare tra grafici e tabella.
- Calcolo e visualizzazione della media ponderata e del voto d'ingresso alla laurea.

## Tecnologie

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS, JavaScript
- **Grafici**: Chart.js

## Installazione e Configurazione

### Prerequisiti

- [Go](https://golang.org/dl/) (versione 1.18 o superiore)
- Un browser web moderno

### Passaggi per l'installazione

1. **Clona il repository**

    ```bash
    git clone https://github.com/tuo-username/libretto-universitario.git
    cd libretto-universitario
    ```

2. **Compila ed esegui il server Go**

    Assicurati di avere Go installato e configurato correttamente. Poi, inserisci gli esami nel file go ed esegui il seguente comando per avviare il server:

    ```bash
    go run main.go
    ```

3. **Accedi all'applicazione**

    Apri il browser e vai all'indirizzo [http://localhost:8080](http://localhost:8080). Dovresti vedere l'applicazione in esecuzione con la possibilità di visualizzare i grafici e la tabella degli esami.

## Utilizzo

- **Selettore a tendina**: Utilizza il menu a tendina per passare tra la visualizzazione dei grafici e la tabella.
  - **Andamento dei Voti**: Mostra un grafico lineare con l'andamento dei voti degli esami.
  - **Andamento dei Voti di Laurea**: Mostra un grafico lineare con l'andamento del voto di laurea calcolato sulla base degli esami sostenuti.
  - **Esami Svolti**: Mostra una tabella dettagliata con i nomi degli esami, le date, i voti e i crediti.

## Struttura del Progetto

- **`main.go`**: File principale che contiene il server Go e la logica di backend.
- **`template.html`**: Template HTML per la visualizzazione dei dati e dei grafici.
- **`README.md`**: Questo file di documentazione.

## Contributi

Se desideri contribuire a questo progetto, apri una **issue** o invia una **pull request** con le tue modifiche e miglioramenti.

## Licenza

Questo progetto è concesso in licenza sotto la [Licenza MIT](https://opensource.org/licenses/MIT). Vedi il file [LICENSE](LICENSE) per i dettagli.

---

Grazie per aver visitato il progetto! 🎉
