export const getArchitects = () => {
    let col = document.getElementsByTagName("a");
    let arr1 = [];
    let arr2 = [];
    for (let elem of col) {
        arr1.push(elem);
    }
    col = document.getElementsByTagName("span");
    for (let elem of col) {
        arr2.push(elem);
    }
    console.log([arr1, arr2]);
    return [arr1, arr2];
};

export const getClassical = () => {
    let arr1 = [];
    let arr2 = [];
    let col = document.querySelectorAll("a.classical");
    for (let elem of col) {
        arr1.push(elem);
    }
    col = document.querySelectorAll("a:not(.classical)");
    for (let elem of col) {
        arr2.push(elem);
    }
    return [arr1, arr2];
};
export const getActive = () => {
    let arr1 = [];
    let arr2 = [];
    let col = document.querySelectorAll("a.classical.active");
    for (let elem of col) {
        arr1.push(elem);
    }
    col = document.querySelectorAll("a.classical:not(.active)");
    for (let elem of col) {
        arr2.push(elem);
    }
    return [arr1, arr2];
};
export const getBonannoPisano = () => {
    let arr1 = [];
    let arr2 = [];
    let col = document.querySelectorAll("a#BonannoPisano.classical.active");
    for (let elem of col) {
        arr1.push(elem);
    }
    col = document.querySelectorAll("a:not(#BonannoPisano).classical.active");
    for (let elem of col) {
        arr2.push(elem);
    }
    return [arr1[0], arr2];
};