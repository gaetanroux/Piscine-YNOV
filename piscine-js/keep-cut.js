const cutFirst = (n, v) => n.slice(2)
console.log(cutFirst('abcdef', 2))

const cutLast = (n, v) => n.slice(0, -2)
console.log(cutLast('abcdef', 2))

const cutFirstLast = (n, v) => n.slice(2, -2)
console.log(cutFirstLast('abcdef', 2))

const keepFirst = (n, v) => n.slice(0, 2)
console.log(keepFirst('abcdef', 2))

const keepLast = (n, v) => n.slice(-2)
console.log(keepLast('abcdef', 2))

const keepFirstLast = (n, v) => n.length > 4 ? keepFirst(n, v) + keepLast(n, v) : n;
console.log(keepFirstLast('abcdef', 2))