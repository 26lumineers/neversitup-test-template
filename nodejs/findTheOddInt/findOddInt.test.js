// findOddInt.test.js
const findOddInt = require('./index');

describe('findOddInt', () => {
    const testCases = [
        { name: 'Test #1', input: [7], expected: { 7: 1 } },
        { name: 'Test #2', input: [0], expected: { 0: 1 } },
        { name: 'Test #3', input: [1, 1, 2], expected: { 2: 1 } },
        { name: 'Test #4', input: [0, 1, 0, 1, 0], expected: { 0: 3 } },
        { name: 'Test #5', input: [1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1], expected: { 4: 1 } },
    ];

    testCases.forEach(({ name, input, expected }) => {
        test(name, () => {
            const result = findOddInt(input);
            const expectedKeys = Object.keys(expected);
            expectedKeys.forEach(key => {
                expect(result).toHaveProperty(key, expected[key]);
                console.log(`[${input}] should return ${key}, because it appears ${expected[key]} time(s) (which is odd).`);
            });
        });
    });
});
