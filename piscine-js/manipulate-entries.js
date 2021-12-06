const filterEntries = (obj, func) => {
    let newObj = {};
    Object.entries(obj)
        .filter(func)
        .map((each) => {
            newObj[each[0]] = each[1];
        });
    return newObj;
};

const mapEntries = (obj, func) => {
    let newObj = {};
    Object.entries(obj)
        .map(func)
        .map((each) => {
            newObj[each[0]] = each[1];
        });
    return newObj;
};

const reduceEntries = (obj, func, acc) => {
    const newObj = {
        ...obj,
    };
    if (acc || acc == 0) {
        return Object.entries(newObj).reduce(func, acc);
    } else {
        acc = "";
        return acc + Object.entries(newObj).reduce(func);
    }
};

const totalCalories = (obj) => {
    return reduceEntries(
        mapEntries(obj, ([k, v]) => [
            `${k}`,
            (Math.round(nutritionDB[k]["calories"] * v) / 1000) * 10,
        ]),
        (acc, [k, v]) => acc + v,
        0
    );
};

const lowCarbs = (obj) => {
    return filterEntries(
        obj,
        ([k, v]) => (v / 100) * nutritionDB[k]["carbs"] < 50
    );
};
const cartTotal = (obj) => {
    return mapEntries(obj, ([k, v]) => {
        let newObj = {};
        for (let [key, val] of Object.entries(nutritionDB[k]))
            newObj[key] = parseFloat(((val * v) / 100).toFixed(3));
        return [k, newObj];
    });
};