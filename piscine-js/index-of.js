const indexOf = (array, value, fromIndex) => {
    if (!fromIndex) {
        fromIndex = 0
    }

    for (let i = fromIndex; i < array.length; i++) {
        if (array[i] === value) {
            return i
        }
    }
    return -1
}


const lastIndexOf = (array, value, fromIndex) => {
    if (!fromIndex) {
        fromIndex = array.length - 1
    }

    for (let i = fromIndex; i >= 0; i--) {
        if (array[i] === value) {
            return i
        }
    }
    return -1
}

const includes = (array, value, fromIndex) => {
    if (!fromIndex) {
        fromIndex = 0
    }

    for (let i = fromIndex; i < array.length; i++) {
        if (array[i] === value) {
            return true
        }
    }
    return false
}