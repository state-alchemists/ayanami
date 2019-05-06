/**
 * Get statistics
 * @param {string[]} cityNames - Names of cities
 * @returns {population: number}
 */
function getPopulations(cityNames) {
    const data = {
        "king's landing": {
            population: 60000
        },
        braavos: {
            population: 4000
        },
        volantis: {
            population: 3000
        },
        asshai: {
            population: 70000
        },
        "old valyria": {
            population: 4000
        },
        "free cities": {
            population: 62000
        },
        qarth: {
            population: 40000
        },
        meereen: {
            population: 10000
        }
    };
    const result = cityNames.map((name) => {
        return {[name]: data[name]};
    });
    return result;
}

module.exports = {
    getPopulations,
};
