const round = (n) => {
    if (n > 0) {
        if (n % 1 >= 0.5) {
            return n - (n % 1) + 1
        }
        return n - (n % 1)
    } else {
        if (n % 1 <= -0.5) {
            return n - (n % 1) - 1
        }
        return n - (n % 1)
    }
}
console.log(round(4.75))


const ceil = (n) => {
    if (round(n) >= n) {
        return round(n)
    } else {
        return round(n) + 1
    }
}

console.log(ceil(0.025))


const floor = (n) => {
    if (round(n) <= n) {
        return round(n)
    } else {
        return round(n) - 1
    }
}

console.log(floor(-7.475))

const trunc = (n) => {
    return n - n % 1
}

console.log(trunc(42.45675))