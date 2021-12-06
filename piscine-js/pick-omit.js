const pick = (obj, keys) => {
    const newObj = {};
    keys instanceof Array ? (keys = keys) : (keys = new Array(keys));
    for (let each of keys) {
        if (obj.hasOwnProperty(each)) {
            newObj[each] = obj[each];
        }
    }
    return newObj;
};
const omit = (obj, keys) => {
    const newObj = {
        ...obj,
    };
    keys instanceof Array ? (keys = keys) : (keys = new Array(keys));
    for (let each of keys) {
        if (obj.hasOwnProperty(each)) {
            delete newObj[each];
        }
    }
    return newObj;
};