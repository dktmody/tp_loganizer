# GoLog Analyzer - Analyse de Logs Distribuée

## Description

GoLog Analyzer est un outil en ligne de commande (CLI) développé en Go. Il permet aux administrateurs système d'analyser des fichiers de logs provenant de diverses sources (serveurs, applications) de manière centralisée et concurrente. L'objectif est d'extraire des informations clés tout en gérant les erreurs de manière robuste.

## Fonctionnalités

- **Analyse concurrente des logs :**
  - Vérifie l'existence et la lisibilité des fichiers de logs.
  - Simule l'analyse avec un délai aléatoire (50 à 200 ms).
- **Gestion des erreurs personnalisées :**
  - Fichier introuvable ou inaccessible.
  - Erreur de parsing JSON.
- **Export des résultats :**
  - Génération d'un rapport JSON détaillé.
- **Interface CLI intuitive :**
  - Commande `analyze` avec des options pour spécifier les fichiers de configuration et de sortie.

## Prérequis

- **Go (version 1.25 ou supérieure)**
- Installer les dépendances :
  ```bash
  go mod tidy
  ```

## Installation

1. Clonez le dépôt :
   ```bash
   git clone https://github.com/dktmody/tp_loganizer.git
   cd loganizer
   ```
2. Compilez le projet :
   ```bash
   go build
   ```

## Utilisation

### Commande `analyze`

Analyse les fichiers de logs spécifiés dans un fichier de configuration JSON.

#### Exemple de fichier `config.json` :

```json
[
  {
    "id": "web-server-1",
    "path": "test_logs/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "test_logs/errors.log",
    "type": "custom-app"
  }
]
```

#### Exécution de la commande :

```bash
go run main.go analyze --config config.json --output report.json
```

- `--config` : Chemin vers le fichier de configuration JSON.
- `--output` : Chemin vers le fichier de sortie JSON.

### Résultats

Un fichier `report.json` sera généré avec le format suivant :

```json
[
  {
    "log_id": "web-server-1",
    "file_path": "test_logs/access.log",
    "status": "OK",
    "message": "Analyse terminée avec succès.",
    "error_details": ""
  },
  {
    "log_id": "invalid-path",
    "file_path": "/non/existent/log.log",
    "status": "FAILED",
    "message": "Fichier introuvable.",
    "error_details": "open /non/existent/log.log: no such file or directory"
  }
]
```

## Tests

1. **Vérifiez la compilation :**
   ```bash
   go build
   ```
2. **Exécutez les tests unitaires :**
   ```bash
   go test ./...
   ```

## Fonctionnalités Bonus

- Création automatique des répertoires pour les fichiers d'export.
- Ajout d'un horodatage dans les noms des fichiers de sortie.

---

Bon courage et bonne analyse de logs !
