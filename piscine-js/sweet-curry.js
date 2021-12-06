let mult2 = (x) => {
    return (y) => x * y;
};
let add3 = (x) => {
    return (y) => (z) => x + y + z;
};
let sub4 = (x) => {
    return (y) => (z) => (a) => x - y - z - a;
};