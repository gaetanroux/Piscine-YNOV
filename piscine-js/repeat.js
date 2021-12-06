const repeat = (c, r) => {
    let compteur = 0
    let resultat = ""
    while (compteur < r) {
        compteur += 1
        resultat += c.toString()
    }
    return resultat
}