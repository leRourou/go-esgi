Vous travaillez dans le service IT d'une entreprise. Vous devez écrire un programme Go qui analyse une liste d'appareils et produit un rapport de statut. Chaque appareil est représenté par trois valeurs : son nom, son type et son année d'achat.
À implémenter

Déclarez une slice de strings contenant les données brutes de la flotte, sous la forme "NOM,TYPE,ANNEE" (ex : "LAPTOP-01,laptop,2019"). Incluez au moins 5 appareils variés (laptop, tablet, server, phone) avec des années entre 2017 et 2022.

En utilisant une seule boucle for, parcourez la slice et construisez deux nouvelles slices : recents (année > 2019) et obsoletes (année ≤ 2019). Affichez le nombre d'éléments dans chacune.

Pour chaque appareil de recents, affichez son nom et son niveau de priorité de maintenance via un switch sur le type. Utilisez fallthrough pour que "server" affiche d'abord "Critique — SLA 24h" puis tombe également dans le cas "laptop" qui affiche "Standard — ticket J+1". Les types "tablet" et "phone" affichent "Faible priorité".
