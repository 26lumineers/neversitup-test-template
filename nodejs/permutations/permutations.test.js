const generatePermutations = require('./index.js')

describe('generatePermutations', () => {
    const testCases = [
        { name: 'Test #1', input: 'a', expected: ['a'] },
        { name: 'Test #2', input: 'ab', expected: ['ab', 'ba'] },
        { name: 'Test #3', input: 'abc', expected: ['abc', 'acb', 'bac', 'bca', 'cab', 'cba'] },
        { name: 'Test #4', input: 'aabb', expected: ['aabb', 'abab', 'abba', 'baab', 'baba', 'bbaa'] },
    ];

    testCases.forEach(({ name, input, expected }) => {
        test(name, () => {
            const result = generatePermutations(input).sort();
            expect(result).toEqual(expected);
        });
    });
});
