const escapeStr = "\`\\\/\"\'"
const arr = Object.freeze([4, '2'])

const obj = {
    str: "string",
    num: 1,
    bool: true,
    undef: undefined
}
Object.freeze(obj)

const nested = {
    arr: ([4, undefined, '2']),
    obj: {
        str: "string",
        num: 1,
        bool: true,

    }


}
Object.freeze(nested)
Object.freeze(nested.arr)
Object.freeze(nested.obj)