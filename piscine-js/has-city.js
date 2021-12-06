let hasCity = (country, arr) => {
    return (city) =>
        arr.some((elem, ind, arr) => city === elem) ?
        city + " is a city from " + country :
        city + " is not a city from " + country;
};