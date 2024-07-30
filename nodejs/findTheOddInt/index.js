const findOddInt = (arr) => {
    const countMap = {};
    arr.forEach(num => {
        countMap[num] = (countMap[num] || 0) + 1;
    });
    const result = {};
    for (const key in countMap) {
        if (countMap[key] % 2 !== 0) {
            result[key] = countMap[key];
        }
    }
    return result;
};
const arrayInt = [1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1]
const result = findOddInt(arrayInt)
console.log(result);

module.exports = findOddInt;