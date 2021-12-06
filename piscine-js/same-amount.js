function sameAmount(str, r1, r2) {
    r1 = new RegExp(r1, 'g');
    r2 = new RegExp(r2, 'g')

    let res1 = str.match(r1);
    let res2 = str.match(r2);

    if (res1 == null) {
        res1 = []
    }
    if (res2 == null) {
        res2 = []
    }

    return res1.length === res2.length
}

console.log(sameAmount(`abcdefghijklmnopqrstuvwxyz azaom ffl ùevcj fflmhqd haa âh psfl hvz ooo gzao o qq io oqpzhvhp`, /fs[^q]/, /q /))