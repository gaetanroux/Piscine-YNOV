const slice = (n, s, m) => {
    let slice = [];
    if (!s) {
        s = 0;
    }
    if (s < 0) {
        s = n.length + s
    }
    if (!m) {
        m = n.length
    }
    if (m < 0) {
        m = n.length + m
    }
    for (let i = s; i < m; i++) {
        slice.push(n[i])
    }
    if (typeof n === 'string') {
        return slice.join('')
    }
    return slice
}
console.log(slice('abcdef', 2))