const sign = (n) => {
    if (n > 0) { return 1 } else if (n < 0) { return -1 } else { return 0 }

}
const sameSign = (n, m) => (sign(n) === sign(m) ? true : false)