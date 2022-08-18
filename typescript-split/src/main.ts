type SplitResult = {
    lines: string[],
    hasExtra: boolean,
}

function splitText(text: string, breaks: number[]): SplitResult {
    // split words by whitespace
    const words = text.split(" ")

    const lines: string[] = []
    let prevBreak = 0

    // populate lines and update previous break for each iteration
    for (const br of breaks) {
        const line = words.slice(prevBreak, prevBreak + br)

        lines.push(line.join(" "))
        prevBreak += br
    }

    // check if extra text exists
    let hasExtra = false

    if (words.slice(prevBreak).length) {
        hasExtra = true
    }

    return { lines, hasExtra }
}

function main() {
    const sentence = "hello what is the name of this app I want to get it online for my tests";

    const result = splitText(sentence, [2, 3, 4, 1]);

    for (const text of result.lines) {
        console.log(text)
    }
    console.log("has extra:", result.hasExtra);
}

main()