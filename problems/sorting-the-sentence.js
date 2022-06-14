/**
 * @param {string} s
 * @return {string}
 */
 var sortSentence = function(s) {
    const words = s.split(' ')
    const newS = []

    const sorted = words.sort((a,b) => {
        if (a[a.length - 1] > b[b.length - 1]) return 1

        else if (a[a.length - 1] < b[b.length - 1]) return -1

        return 0
    });

    sorted.forEach(word => {
        const cleanWord = word.replace(/\d+/, '')
        newS.push(cleanWord)
    })
    return newS.join(' ')
};