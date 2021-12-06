const reverse = (n) => {
    let isString;
    let res = [];

    if (typeof n === 'string') {
        isString = true
    }

    for (let i = n.length - 1; i >= 0; i--) {
        res.push(n[i])
    }

    if (isString) {
        return res.join('')
    }
    return res
}
console.log(reverse("Je m'appelle Gaëtan Roux"))