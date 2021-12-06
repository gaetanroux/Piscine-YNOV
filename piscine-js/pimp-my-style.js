import { styles } from "./pimp-my-style.data.js";
export const pimp = () => {
    let butt = document.querySelector("button");
    let clicks = 0;
    if (butt.hasAttribute("data-clicks")) {
        clicks = parseInt(butt.dataset.clicks);
        console.log(clicks, butt.dataset.clicks);
    }
    butt.dataset.clicks = clicks;
    if (butt.classList.contains("unpimp")) {
        if (butt.hasAttribute("data-clicks")) {
            clicks = parseInt(butt.dataset.clicks);
            console.log(clicks, butt.dataset.clicks);
        }
        butt.classList.remove(styles[clicks]);
        clicks--;
        if (clicks == -1) {
            butt.classList.remove("unpimp");
            clicks++;
        }
    } else {
        if (butt.hasAttribute("data-clicks")) {
            clicks = parseInt(butt.dataset.clicks);
            console.log(clicks, butt.dataset.clicks);
        }

        butt.classList.add(styles[clicks]);
        clicks++;
        if (clicks == 15) {
            butt.classList.add("unpimp");
            clicks--;
        }
    }
    butt.dataset.clicks = clicks;
    console.log("set click", clicks);
};