## **Ascii-Art**

- **Groupe :**
Clément Manant 
Gaëtan  Roux

- Procédure d'installation de votre projet en local:

Il vous faudra installer "gitbash" sur votre pc pour pouvoir exécuter le code avec toutes ses fonctionnalités.

- Dans un premier temps, il faudra aller sur notre repo : ``https://git.ytrack.learn.ynov.com/GROUX/ASCII-ART.git``.

Sur le repo, il faudra télécharger notre référentiel au format : ``zip``.

Une fois ce fichier téléchargé, il faudra extraire le dossier nommé ``ASCII-ART`` sur votre bureau.

Maintenant, vous allez pouvoir lancer ``GitBash`` qui va vous permettre d'exécuter le code.

Sur ``GitBash``, il faudra aller dans le répertoire où l'on a extrait le dossier.
Pour ce faire il faudra se déplacer en utilisant ``cd``: 
``cd Desktop`` puis ``cd ASCII-ART``.

Vous êtes maintenant dans le bon répertoire pour pouvoir exécuter notre code ;) !!

 
- Exemple d'utilisation:

Commande à suivre Ascii-art.go : ``go run Ascii-art.go`` "ce que vous souhaitez faire apparaitre" (dois toujours apparaitre en premier argument !)

![](https://i.imgur.com/twCXvP6.png)


Commande à suivre fs : ``go run Ascii-art.go`` "ce que vous souhaitez faire apparaitre" la police que vous souhaitez (standard est la police par défault, la police change uniquement si l'argument est "shadow" ou "thinkertoy")

![](https://i.imgur.com/Yp8zCsB.png)


Commande à suivre output : 

``go build``

``go run Ascii-art.go`` "ce que vous souhaitez faire apparaitre" la police que vous souhaitez --output=lenomdevotrefichier.txt (si le fichier existe déjà, il l'efface et réécrit par dessus)

cat lenomdevotrefichier.txt

![](https://i.imgur.com/wWdSfbw.png)