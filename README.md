# Devoir sur les entrées/sorties

## Description du sujet

Le but de ce devoir est de réaliser un programme utilisant une base de données
pour enregistrer des données en mode maître/détail, avec application d'une
intégrité référentielle.

Les données représentées seront des employés et leurs entreprises respectives,
dans une relation `1:n`, puis `m:n` via une table de relation.

Ce programme va présenter quelques points communs avec une application réelle :

- lancement depuis la ligne de commande sous forme de sous-commandes (Cobra)
- chargement de la configuration depuis l'environnement ou un fichier de configuration (Viper)
- accepter au moins un drapeau `-v` booléen permettant de le rendre bavard (PFlag):
  il affichera les opérations qu'il exécute et leur résultat
- exister en deux versions de schéma, entre lesquelles il sera en mesure d'effectuer
  une migration de schéma SQL montante et descendante.
- traitement des erreurs. Les sorties par `panic` ou `log.Fatal*` ne sont
  autorisées que dans le premier niveau de fonction (les callbacks Cobra du paquet `cmd`),
  pas dans les fonctions d'accès aux données (paquets `model` et `migrations`).
- documentation minimale (README.md, licence, commentaires `godoc`)
- respect à 100% des normes de codage, à valider avec `make lint`
- présence d'au moins un test de type `Example` associé à la documentation.
- un volume de code à réaliser non symbolique (cf. Étapes recommandées, _infra_).


## Organisation des fichiers

- `(projet)` : la racine, comportant `main.go`, `go.mod`, `go.sum`, `README.md`
- `(projet)/cmd` : les commandes et sous-commandes
- `(projet)/data` : les fichiers de données
- `(projet)/migrations` : les migrations (goose)
- `(projet)/model` : le modèle de données

## Liste des commandes à créer

- (commande racine): liste les commandes disponibles
- `v1`: implémente la première version du schéma, dans laquelle chaque personne 
  n'est employée que par une seule entreprise. Sous-commandes :
  - `v1 load`: charge les employés depuis `data/staff.yml`
  - `v1 lsCompany`: liste les entreprises et leurs effectifs à l'écran
  - `v1 chCompany <siret1> <siret2>`: change l'identifiant d'une entreprise
  - `v1 rmCompany`: supprime une entreprise et tous ses salariés
- `v2`:  implémente la seconde version du schéma, dans laquelle une personne peut avoir plusieurs employeurs.
  Sous-commandes:
  - `v2 addStaff <siret> <insee>`: ajoute une personne à une entreprise en créant une ligne
  - `v2 rmPerson <insee>`: supprime un salarié et le retire de toutes ses entreprises
  - `v2 chPerson <insee> "<given> <last>"`: modifie le nom d'une personne
- `migrate`: implémente les commandes de migration
  - `migrate up`: monte le schéma de version:
    - si la base n'est pas installée, crée la base et installe le schéma
    - si elle est en version 1, la convertit en version 2
    - si elle est en version 2, renvoie une erreur indiquant qu'il n'y a rien à faire
    - avec `goose` c'est la commande `up-by-one` (`goose.UpByOne()`)
  - `migrate down`: descend le schéma de version
    - si la base est en version 2, la convertit en version 1
    - si elle est en version 1, supprime la base
    - avec `goose`, c'est la commande `down` (`goose.Down()`)

Le but est de pouvoir enchaîner par exemple les commandes comme dans la commande
`make demo`:

```
./devoir10 migrate up              # Installe le schéma v1
./devoir10 migrate version         # Indique la version 1
./devoir10 v1 load                 # Charge les effectifs en v1
./devoir10 v1 lscompany            # Liste la base v1
./devoir10 v1 rmcompany 1          # Echoue proprement à effacer l'entreprise 1 (lignes présentes)
./devoir10 v1 rmcompany 3          # Indique que l'effacement n'a eu aucun effet (inexistante)
./devoir10 v1 chcompany 1 2        # Echoue à renuméroter l'entreprise 1 en 2 (2 existante)
./devoir10 v1 chcompany 3 1        # Echoue à renuméroter l'entreprise 3 en 1 (3 inexistant)
./devoir10 v1 chcompany 1 3        # Renumérote l'entreprise 1 et ses lignes en 3
./devoir10 migrate up              # Migre la base en v2 en chargeant les identités des personnes
./devoir10 version                 # Indique la version 2
./devoir10 v2 lscompany            # Liste la base v2
./devoir10 v2 rmperson 31          # Echoue à effacer le salarié 31 inexistant
./devoir10 v2 rmperson 21          # Efface le salarié 21 des effectifs et de la liste des identités
./devoir10 v2 chperson 11 John Doe # Change le nom de la personne 11 dans `people` uniquement
./devoir10 v2 chperson 21 Max Doe  # Échoue à renommer la personne 21 supprimée précédemment
./devoir10 v2 lscompany            # Liste la base avec la personne 11 non renommée dans `staff`
./devoir10 migrate down            # Inverse la migration de la base de v2 en v1
./devoir10 v1 lscompany            # Liste la base de nouveau
```

_Remarque_: Dans ce devoir aucune vérification des numéros SIRET et INSEE n'est 
à prévoir, ce ne sont que des chaînes numériques `[1-9][0-9]*`.


## Livrables

Le rendu est sous la forme d'une archive `devoir10.tgz`, sans historique de version.
Si vous développez le code par étapes (recommandé) en commettant chaque étape (recommandé),
exportez le résultat final après avoir commis la dernière étape, avec la commande:

`git archive -o ../devoir10.tgz HEAD`

Attention à bien avoir commis tout votre code avant de lancer cette commande, et
à vérifier que vous ne rendez pas une archive vide (de longueur 0).

Ne rendez **PAS** l'exécutable compilé, ni la base de données.

Le projet est à réaliser sur SQLite 3.x.

Vous êtes lead dév de ce projet, donc vous avez le choix de l'API :

- le paquet standard `database/sql`,
- la surcouche d'expérience développeur améliorée `github.com/jmoiron/sqlx`,
- ou l'ORM [Ent](https://entgo.io/) de Facebook

Le prix de la liberté:

- votre rendu doit justifier vos choix, comme vous aurez prochainement à le faire
  dans un cadre professionnel.
  Insérez votre explication dans le `README.md`: elle fait partie de la note.
- vous êtes libres de vous tirer une balle dans le pied avec ce choix, 
  figurativement parlant: choisissez bien.


## Détail des versions de schéma
### Première version

- La première version manipule des effectifs d'entreprises, composées de
  l'entreprise elle-même (`companies`), et de son personnel (`staff`),
  composées d'un identifiant personnel (`insee`), d'un nom en texte libre,
  et d'un salaire annuel brut en Euro.
- Dans cette version, on considère que chaque personne n'est employée que par
  une seule entreprise et au titre d'un seul contrat, donc: 
  - La relation entreprise/personnel est de cardinalité _1:n_.
  - chaque code `insee` est unique dans une même entreprise, donc pas besoin d'`id` 
    unique dans la table des lignes
- Les montants monétaires sont en `EUR`, locale `fr_FR`. Vous pouvez faire l'hypothèse 
  que les montants sont inférieurs à l'agrégat M3 de la masse monétaire EUR, donc 1E13:
  cela conditionne le choix de votre taille de colonne pour l'enregistrement des
  valeurs monétaires (salaires). Pour stocker les montants en base, utilisez 
  `Amount.MarshalBinary` pour obtenir une représentation inversible enregistrable
  dans un champ non-blob.

- table `companies`:
  - `siret`: identifiant unique, auto-incrémenté, clef primaire
  - `changed`: horodatage précis au moins à la seconde
  - `total_salary`: monétaire, égal au total des salaires annuels
- table `staff`, clef primaire composite `siret`+`insee`
  - `insee`: identifiant de la personne employée, longueur fixe, 13 chiffres de 0 à 9
  - `name` : nom et prénom complet de la personne, 40 caractères Unicode
  - `salary`: salaire annuel brut de la personne dans l'entreprise, monétaire

Le code doit être en mesure de:

- charger les entreprises et leur personnel depuis un fichier de données externe
- lister les entreprises et leur personnel depuis la base
- supprimer une personne d'une entreprise
- permettre la modification du SIRET d'une entreprise, que la base de données 
  doit reporter sur ses lignes par le jeu de l'intégrité référentielle
- empêcher par l'intégrité référentielle la suppression d'une entreprise employant 
  encore du personnel
- dans la réalité, il faudrait maintenir cohérente la masse salariale lorsqu'on
  ajoute ou supprime des lignes, mais le devoir n'inclut pas ces commandes, donc
  ignorez cette contrainte.
  

### Seconde version

Dans la seconde version, les numéros INSEE des personnes employées sont validés 
par référence à une table des personnes, et le champ `name` des lignes est tiré 
des noms des personnes dans cette table.

- `people`
  - `insee`: chaine de longueur fixe de 13 chiffres de 0 à 9, clef primaire
  - `given`: chaîne de caractères Unicode quelconque, longueur maximum 40.
  - `last`: chaîne de caractères Unicode quelconque, longueur maximum 40.
  
Pour cela, le code doit inclure une migration de schéma qui va:

- ajouter la table additionnelle `people`
- modifier le schéma de la table `staff` pour qu'elle inclue la contrainte de
  clef étrangère vers la table `people`, en relation _1:n_.
- copier les noms des personnes de `people` vers chaque ligne de `staff`  qui
  les référence, en remplaçant les données libres qui s'y trouvaient précédemment,
  par concaténation de la colonne `given`, d'un espace, et de `last`
- le tout sans perte de données, hormis le remplacement des noms complets.  
  
Comme dans la première version, le programme doit pouvoir charger
les entreprises, etc, (cf première version).

Il doit aussi maintenant :

- empêcher au niveau base de données l'insertion d'une ligne avec un SKU ne
  correspondant pas à un produit (`FOREIGN KEY`)
- empêcher la suppression d'une personne de `people` si elle est employée par
  une entreprise dans la table `staff` (`ON DELETE ...`)
- reporter par cascade de la contrainte d'intégrité référentielle les mises à 
  jour du numéro INSEE de la table `people` dans `staff` (`ON UPDATE ...`) 
- ne pas reporter par cascade les modifications des nom/prénom ni
  du salaire dans les lignes qui l'utilisent: la cascade ne doit concerner que le
  numéro INSEE.


## Notation

- si vous utilisez l'ORM plutôt que `sql` ou `sqlx` avec succès, vous bénéficiez
  d'un bonus pouvant aller jusqu'à 2 points supplémentaires... si votre code fonctionne:
  un code sur ORM ne fonctionnant pas ne rapport pas de points
- la présence de tests unitaires en complément d'au moins un test `Example*`
  peut donner un bonus jusqu'à 2 points.
  La présence du texte `Example*` est en revanche requise (cf. _infra_).
- la présence de code non documenté, mal formaté ou renvoyant des anomalies aux 
  linters indiqués retire jusqu'à 2 points.
- la note est en tout état de cause plafonnée à 20/20.


## Composants à utiliser

- Créez votre devoir sous forme d'un module Go (`go mod init devoir10_<votrenom>`)
- Utilisez `gopkg.in/yaml/v2` pour lire les fichiers YAML
- Utiliser `github.com/bojanz/currency` pour modéliser les montants monétaires.
  Attention, cela signifie que pour décoder les prix des fichiers, vous devrez
  faire implémenter l'interface `yaml.Unmarshaler` à vos types `Company` et `Staff`,
  pour que les salaires soient correctement chargés depuis YAML.
- Utilisez `spf13/cobra` pour créer les commandes, `spf13/viper` (inclus dans
  Cobra) pour charger la configuration depuis un fichier ou l'environnement,
  et `spf13/pflag` (inclus dans Cobra) pour les options.
- Utilisez SQLite3 pour ne pas avoir à déployer un serveur.
- Avec l'ORM `ent`, utilisez son outil de migration; sinon l'outil de migration
  `pressly/goose`.
  Vous n'avez pas besoin d'implémenter toutes les commandes de Goose :
  `up`, `down`, et `status` sont suffisantes, et`version` est recommandée pour
  vous simplifier la vie. Vous pouvez utiliser des migrations SQL ou Go (recommandées).
- Outils de vérification de code indiqués, à installer une fois votre projet initialisé:
  - `go fmt`: inclus avec Go
  - `golint`: https://github.com/golang/lint
  - `staticcheck`: https://staticcheck.io/docs
  
**Attention** Avec SQLite 3, votre code _doit_ exécuter la commande 
`PRAGMA foreign_keys=on;` à _chaque_ ouverture de connexion : par défaut SQLite 3
ignore l'intégrité référentielle, et ce choix est par connexion et non stocké
dans la base elle-même. Ce devoir ne peut _pas_ être réalisé tel que demandé
sans contraintes d'intégrité référentielle.

Pensez à l'inclure dans votre fonction partagée d'ouverture de connexion,
à créer dans votre paquet `cmd` en tant que:

```go
package cmd

import "database/sql"

func mustOpenDB() *sql.DB { /* ... votre code ... */ }
```


## Étapes recommandées

Le processus détaillé dans les étapes ci-après est probablement le plus simple
pour vous permettre d'arriver au bout du sujet sans blocage, mais vous pouvez
vous y prendre autrement.

Il est possible de réaliser l'ensemble du devoir en remplissant simplement les
parties marquées _// À vous..._ sans rajouter d'autres types ou fonctions,
mais ce n'est pas une obligation. 

À titre indicatif pour vous donner une idée du volume attendu, voici les métriques
du corrigé-type:

```
===============================================================================
Language            Files        Lines         Code     Comments       Blanks
===============================================================================
Go                     26         1741         1327          191          223
Makefile                1            6            5            0            1
YAML                    2           33           32            1            0
-------------------------------------------------------------------------------
Markdown                1          306            0          243           63
|- Go                   1            5            3            0            2
(Total)                            311            3          243           65
===============================================================================
Total                  30         2091         1367          435          289
===============================================================================
```

Le squelette fourni comporte déjà environ 25% des lignes de Go du corrigé-type.


#### Étape 1

- Modélisez la base de données des 2 étapes. Pensez aux clefs primaires,
  et aux contraintes d'intégrité référentielle.
- Créez un squelette d'application comportant les commandes et sous-commandes qui
  se contentent d'afficher ce qu'elles devraient respectivement faire.


#### Étape 2

- étendez le squelette pour qu'il soit capable de lire la configuration
   (codes d'accès à la base de données, chemin) depuis l'environnement ou un 
   fichier de configuration. Utilisez Cobra et Viper pour cela. Ajouter
   une option `-v` avec les flags standard ou avec PFlag.
- incorporez la bibliothèque de migration `goose` sous forme de sous-commandes 
  d'une commande `migrate` de votre programme (et non comme commande extérieure).


#### Étape 3

- écrivez la première migration pour `migrate up`, qui installe le schéma initial.
  Si une base précédente existe, elle devra la supprimer sans erreur et la recréer.
- écrivez la première migration pour `migrate down` qui supprime la base de données v1.



#### Étape 4

- écrivez la commande `v1 load` qui charge le fichier `data/staff.yml` dans la base installée
- écrire la commande `v1 lscompany` qui liste les entreprises et leurs effectifs
  tels qu'ils sont dans la base.
  le plus simple est les charger en mémoire et de sérialiser le résultat en YAML
  en guise d'affichage texte.
  Attention, pensez à séparer cette tâche en deux fonctions
    - la lecture depuis la base
    - l'affichage des données dans une fonction d'affichage `Show(w io.Writer, carts Carts)`,
      pour que cette dernière soit testable par un test `ExampleShow`.
- écrivez le test `ExampleShow` qui vérifie le résultat de l'affichage.
  C'est un test unitaire, donc les données sont injectées dans la fonction
  d'affichage `Show` et non lues depuis la base pendant le test.


#### Étape 5

- écrivez la commande `v1 rmcompany <siret>` qui tente de supprimer une entreprise:
   - elle y parvient si l'entreprise n'a pas d'employés si elle n'existait pas (0 ligne supprimée)
   - elle reçoit une erreur SQL qu'elle intercepte correctement si l'entreprise
     a encore des employés, et affiche une erreur orientée utilisateur
     (donc pas un message système, mais quelque chose de compréhensible)
- écrivez la commande `v1 chcompany <siret1> <siret2>` qui modifie le numéro SIRET 
     d'une entreprise
   - elle reçoit et intercepte une erreur SQL si l'identifiant cible `<siret2>` existe déjà,
     ou si `<siret1>` n'existe pas.
   - sinon, elle réussit, et le serveur SQL applique le changement aux lignes
     par la cascade de l'intégrité référentielle


#### Étape 6

- écrivez la seconde migration pour `migrate up` qui convertit la base de son schéma
  v1 à son schéma v2. Elle devra:
  - charger le fichier `data/people.yml`,
  - modifier la base en
    1. ajoutant la table `people`
    2. insérant les données lues dans la table `people` et mettant à jour les 
      noms - mais pas les salaires - des lignes de personnes employées concernées.
      Indice: faisable en une seule instruction SQL avec sous-instruction.
    3. ajoutant une clef étrangère vers la table des produits à la table des lignes
- écrivez la commande de migration `migrate down` qui redescend le schéma de v2 à v1.
- écrivez la commande `v2 lscompany` équivalente à la `v1 lscorp` pour le schéma v2.
  _Indice gain de temps_: est-elle vraiment différente ?

  
#### Étape 7

- écrivez la commande `v2 addstaff <siret> <insee>` qui:
   - échoue si le `siret` n'existe pas
   - échoue si le code `insee` n'existe pas ou s'il est déjà présent dans l'entreprise
   - réussit sinon, et ajoute la personne employée comme ligne de l'entreprise, en 
     remplissant nom et salaire de la ligne de `staff` à partir de la ligne de `people`.
- écrivez la commande `v2 rmperson <insee>` qui:
   - échoue si le code `insee` n'existe pas (SQL renvoie 0 lignes supprimées)
   - échoue si le code `insee` est utilisé dans une ligne d'une entreprise (violation d'intégrité SQL)
   - supprime le personne employée correspondante sinon
- écrivez la commande `v2 chperson <insee> "<given>" "<last>` qui:
   - échoue si le code `insee` n'existait pas déjà dans la table `people`
   - réussit sinon, en réécrivant les prénom et nom complet dans `people`, 
     sans modifier `staff`.
