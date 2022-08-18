# Split Text
Implement a text splitting function in different languages

# API Design
Create a function called `split text` (casing should be language dependent)

## Arguments
The function should accept two arguments:
- text `string` (text to be split)
- breaks `array of numbers` (breakpoints to split the text by)

## Return
The function should return a structure containing two fields:
- lines `array of strings` (resulting lines)
- has extra `boolean` (signifies if the text has remaining content that was not returned)

## Implementation Details
Given a text and breaks of `N` breakpoints, 
the function should return `N` lines where the number of words of each line is the equal to it's respective breakpoint.

Below is intended use of the function in `typescript`

```typescript
  const text = "Hello World! How are ya!"
  
  const { lines, hasExtra } = splitText(text, [2, 1])
  
  console.log("lines:", lines, "extra:", hasExtra)
  // expected output:
  // lines: [ "Hello World!", "How" ] extra: true
  
  // here we specified two breakpoints 2 and 1
  // the function returns `lines` of exactly 2 strings, the first line has exactly 2 words and the second exactly 1 word
  // the function also `hasExtra` a boolean which is true because part of the input text was not returned ("are ya!")
```

## Considerations
- The function must not throw and all possible error cases must be handled gracefully by the implementation. 
Instead of errors, empty strings should be returned for lines
```typescript
  const text = "Try out of bounds"
  
  const { lines, hasExtra } = splitText(text, [2, 10, 30])
  
  console.log("lines:", lines, "extra:", hasExtra)
  // expected output:
  // lines: [ "Try out", "of bounds", "" ] extra: false
  
  // we know 10 and 30 are clearly out of bounds of the input text because the input text only has 4 words
  // but instead of the function to fail it should return as much text that is possible or an empty string.
```

# Contribution
You may check some of the already implemented examples to have a better understanding of the api design

Make a fork and send in PR for your favourite programming language 😊
Pull requests are always welcome 🤗
