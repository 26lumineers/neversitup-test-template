
const detectSmiley = (smileys) => {
    let count = 0;
    smileys.forEach(val => {
        if (val.length > 3 || val.length < 2) {
            console.log("invalid smile");
            return;
        }
        let hasEyes = false;
        let hasMouth = false;
        let hasNose = false;
        for (const char of val) {
            switch (char) {
                case ';':
                case ':':
                    hasEyes = true;
                    break;
                case '-':
                case '~':
                    hasNose = true;
                    break;
                case ')':
                case 'D':
                    hasMouth = true;
                    break;
            }
        }
        if (hasEyes && hasMouth && (val.length === 2 || (val.length === 3 && hasNose))) {
            count++;
        }
    });
    return count;
};

const main = () => {
    console.log("Count Smiley face :D");
    const smileys = [":)", ";(", ";}", ":-D", ":_D"];

    const smileyfaces = detectSmiley(smileys);
    console.log("smileyfaces : ", smileyfaces);
};

main();

module.exports = detectSmiley;
