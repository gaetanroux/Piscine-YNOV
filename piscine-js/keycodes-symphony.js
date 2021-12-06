export const compose = () => {
    let del = (str) => {
        let arr = document.querySelectorAll("div");
        if (str === "Escape") {
            for (let each of arr) each.remove();
        }
        if (str === "Backspace") {
            let last = arr[arr.length - 1];
            if (last) last.remove();
        }
    };

    document.addEventListener(
        "keydown",
        (event) => {
            const keyName = event.key;
            if (keyName === "Backspace" || keyName === "Escape") {
                return del(keyName);
            }
            if (keyName.length === 1 && /[a-z]/.test(keyName)) {
                let note = document.createElement("div");
                note.className = "note";
                let color = "#" + (((1 << 24) * Math.random()) | 0).toString(16);
                note.style.background = color;
                note.textContent = keyName;
                document.querySelector("body").append(note);
            }
        },
        false
    );
};