# Evaluation SQL

**Début du TP final**: lundi 15 novembre

**Fin du projet**: dimanche 28 novembre - 23h59

## Partie 1 : Conception d'une base de données

➡️ **Imaginez la base de données d'un hôtel et écrivez un script SQL permettant de mettre en place cette base de données.**

Votre base de données doit pouvoir par exemple gérer différentes données telles que :

- les informations des clients
- les chambres
- les réservations de l'hôtel
- le menu du restaurant et des plats proposés
- ... (libre à votre imagination)

Expliquez vos choix.

## Partie 2 : Créer des fixtures (données bouchons)

➡️ **Créez un script permettant de remplir la base de données avec des données bouchons.**

Vous pouvez réaliser ce script dans le langage de votre choix.

### Ressources

- Librairie pour générer des données bouchons (ex: Faker)

## Partie 3 : Analyser une base de données

➡️ **A partir d'une base de données fourni, vous devrez réaliser des requêtes permettant d'obtenir des informations utiles (statistiques, informations sur les commandes, les employés, les villes, les clients...).**

**1** Le prix total des commandes contenant plus de 5 articles différents

**2** La liste de tous les territoires de Peacock Margaret

**3** La liste des clients vivant à "London"

**4** La liste des clients ayant commandé pour une livraison à "London" avant 2013

**5** Afficher le client qui a fait le plus de commande vers le "Brazil"

**6** Afficher la valeur total de la commande avec l'id 10260

**7** Afficher la valeur de toutes les commandes infèrieur à la moyenne de toutes les commandes

**8** Lister tous les pays de la base de données (dans les tables employee, order, customer), indiquer pour chaque pays le nombre de lignes correspondantes dans toute la base de données. (Chaque pays ne doit pas apparaître en doublon)

**9** Le nom et prénom des employés en ajoutant:

    le nombre total de commande
    le total des commandes par employé
    le pays dans lequel ils ont fait le plus de commande
    le poids total de toutes les commandes par employé
    le pourcentage de vente par rapport au meilleur vendeur


## Rendu

- Un dossier nommé `NOM_Prenom-sql` contenant :
  - un dossier `p1` avec votre script SQL dedans ET un fichier explicatif (.md ou .pdf)
  - un dossier `p2` avec votre script SQL dedans
  - un dossier `p3` avec vos requêtes et le résultat démontré (screen à l'appui)