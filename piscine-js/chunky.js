function chunk(array, len) {
    let counter = 0;
    let res = [];
    let tempArray = [];

    for (let i = 0; i < array.length; i++) {
        if (tempArray.length == len) {
            res.push(tempArray);
            tempArray = [];
        }
        tempArray.push(array[i])
    }
    if (tempArray != []) {
        res.push(tempArray)
    }
    return res
}
console.log(chunk("ceci est un test et me permet de validé mon code !"))