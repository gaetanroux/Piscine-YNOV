export const build = (num) => {
    let c = 0;
    let b = 1;
    const last = document.querySelector("body");
    var timerInt = setInterval(() => {
        const newBrick = document.createElement("div");
        newBrick.id = "brick-" + (c + 1);
        newBrick.innerHTML = c + 1;
        if (b == c) {
            newBrick.dataset.foundation = true;
            b = c + 3;
        }
        last.append(newBrick);

        c++;
        if (c == num) {
            clearInterval(timerInt);
        }
    }, 100);
};
export const repair = (...ids) => {
    ids.forEach((val) => {
        let brick = document.querySelector("div#" + val);
        if (brick) {
            if (brick.hasAttribute("data-foundation")) {
                brick.dataset.repaired = "in progress";
            } else {
                brick.dataset.repaired = true;
            }
        }
    });
};
export const destroy = () => {
    let bricks = document.querySelectorAll("div");
    let ar0 = [];
    for (let each of bricks) {
        ar0.push(each);
    }
    bricks = ar0.slice(-1)[0];
    if (bricks) {
        bricks.remove();
    }
};