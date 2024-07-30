const generatePermutations = (s) => {
    const chars = s.split('');
    const result = [];
    backtrack(chars, 0, result);
    return result;
};

const backtrack = (chars, start, result) => {
    if (start === chars.length) {
        result.push(chars.join(''));
        return;
    }

    const seen = new Set();
    for (let i = start; i < chars.length; i++) {
        if (seen.has(chars[i])) {
            continue;
        }
        seen.add(chars[i]);

        [chars[start], chars[i]] = [chars[i], chars[start]];
        backtrack(chars, start + 1, result);
        [chars[start], chars[i]] = [chars[i], chars[start]];
    }
};

console.log("Permutations..");
const word = 'aabb';
console.log("With input : ", word);
const perm = generatePermutations(word);
perm.sort();

console.log("should return:", perm);

module.exports = generatePermutations;