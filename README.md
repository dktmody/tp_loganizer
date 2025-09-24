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
    "path": "/var/log/nginx/access.log",
    "type": "nginx-access"
  },
  {
    "id": "app-backend-2",
    "path": "/var/log/my_app/errors.log",
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
    "file_path": "/var/log/nginx/access.log",
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

### Fonctionnalités Bonus

#### Création automatique des répertoires

Pour tester cette fonctionnalité :

1. Spécifiez un chemin de sortie dans un répertoire inexistant, par exemple :
   ```bash
   go run main.go analyze --config config.json --output new_dir/report.json
   ```
2. Vérifiez que le répertoire `new_dir` est automatiquement créé et que le fichier `report.json` y est généré.

#### Ajout d'un horodatage dans les noms des fichiers de sortie

Pour tester cette fonctionnalité :

1. Exécutez la commande avec un chemin de sortie standard :
   ```bash
   go run main.go analyze --config config.json --output report.json
   ```
2. Vérifiez que le fichier généré contient un horodatage dans son nom, par exemple : `report_2025-09-24_15-30-00.json`.

---

