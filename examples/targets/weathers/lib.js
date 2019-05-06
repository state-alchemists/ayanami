const fs = require('fs');

/**
 * Get weathers
 * @param {string[]} cityNames - Names of cities
 * @returns {main: string, temp: number}
 */
function getWeathers(cityNames) {
    const content = fs.readFileSync(`${__dirname}/data.json`);
    const data = JSON.parse(content);
    const data ={
        "king's landing": {
            main: "clear",
            temp: 281.34
        },
        braavos: {
            main: "clouds",
            temp: 269.35
        },
        volantis: {
            main: "clouds",
            temp: 294.21
        },
        asshai: {
            main: "rain",
            temp: 235.43
        },
        "old valyria": {
            main: "rain",
            temp: 252.46
        },
        "free cities": {
            main: "rain",
            temp: 281.25
        },
        qarth: {
            main: "clouds",
            temp: 243.32
        },
        meereen: {
            main: "clear",
            temp: 284.32
        }
    };
    const result = cityNames.map((name) => {
        return {[name]: data[name]};
    });
    return result;
}

module.exports = {
    getWeathers,
};
