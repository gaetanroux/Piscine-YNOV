import { colors } from "./fifty-shades-of-cold.data.js";
export const generateClasses = () => {
    let style = document.createElement("style");
    colors.forEach((elem) => {
        style.innerHTML =
            style.innerHTML +
            "." +
            elem +
            "  \n{\n   background: " +
            elem +
            ";\n }\n";
    });
    let head = document.head;
    document.head.append(style);
};
export const generateColdShades = () => {
    let body = document.querySelector("body");
    colors.forEach((elem) => {
        let ColdShade = document.createElement("div");
        ColdShade.classList.add(elem);
        ColdShade.textContent = elem;
        let myArr = [
            "aqua",
            "blue",
            "turquoise",
            "green",
            "cyan",
            "navy",
            "purple",
        ];
        let str = ColdShade.className;
        if (myArr.some((el) => str.includes(el))) body.append(ColdShade);
    });
};
export const choseShade = (shade) => {
    let divs = document.querySelectorAll("div");
    for (let div of divs) div.className = shade;
};