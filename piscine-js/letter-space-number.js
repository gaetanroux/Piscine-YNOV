function letterSpaceNumber(str) {
    let res = str.match(/. \d((?=\W)|$)/g);
    if (res) {
        return res
    }
    return []
}
console.log(letterSpaceNumber("Ceci est un vieux test me permettant de valider ma veille théorie."))