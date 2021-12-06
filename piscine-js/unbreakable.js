function split(str, m) {
    let lenStr = str.length;
    let lenSepatator = m.length;
    let words = [];
    let mCount = 0;
    let word = ''
    if (m != '') {
        for (let i = 0; i < lenStr; i++) {
            word += String(str[i]);
            let lenWord = word.length;
            if (lenWord => lenSepatator) {
                if (word.slice(lenWord - lenSepatator, lenWord) == m) {
                    mCount++
                    words.push(word.slice(0, lenWord - lenSepatator))
                    word = ''
                }
            }
            if (i == lenStr - 1 && word != '') {
                words.push(word)
            }
        }
    }

    if (mCount == words.length) {
        words.push('')
        return words
    }
    return words
}


function join(n, m) {
    if (m === '') {
        m = ','
    }
    let res = ''
    for (let i = 0; i < n.length; i++) {
        res += n[i]
        if (i != n.length - 1) {
            res += m
        }
    }
    return res
}