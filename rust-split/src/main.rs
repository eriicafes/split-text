struct SplitResult {
    lines: Vec<String>,
    has_extra: bool,
}

fn split_text(text: &str, breaks: &[usize]) -> SplitResult {
    // split words by whitespace
    let words = text.split_whitespace();

    let mut lines: Vec<String> = Vec::with_capacity(breaks.len());
    let mut prev_break: usize = 0;

    // populate lines and update previous break for each iteration
    for br in breaks {
        let line: Vec<&str> = words.clone().skip(prev_break).take(*br).collect();

        lines.push(line.join(" "));
        prev_break += br;
    }

    // check if extra text exists
    let has_extra = !matches!(words.clone().skip(prev_break).count(), 0);

    SplitResult { lines, has_extra }
}

fn main() {
    let sentence = "hello what is the name of this app I want to get it online for my tests";

    let result = split_text(sentence, &[2, 3, 4, 1]);

    for text in &result.lines {
        println!("{text}")
    }
    println!("has extra: {}", result.has_extra);
}
