const flow = (arr) => (...args) => {
    let temp = arr[0](...args);
    arr = arr.slice(1);
    for (let each of arr) {
        temp = each.apply(this, [temp]);
    }
    return temp;
};