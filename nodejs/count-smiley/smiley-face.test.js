const detectSmiley = require('./index');

describe('detectSmiley', () => {
    const testCases = [
        {
            name: 'Test #1',
            input: [":)", ";(", ";}", ":-D"],
            expected: 2,
        },
        {
            name: 'Test #2',
            input: [";D", ":-(", ":-)", ";~)"],
            expected: 3,
        },
        {
            name: 'Test #3',
            input: [";]", ":[", ";*", ":$", ";-D"],
            expected: 1,
        },
    ];

    testCases.forEach(({ name, input, expected }) => {
        test(name, () => {
            const result = detectSmiley(input);
            expect(result).toBe(expected);
            console.log(`countSmileys [${input}] = ${expected}`);
        });
    });
});
