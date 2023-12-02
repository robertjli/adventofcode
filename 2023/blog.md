## Advent of Code 2023

This is a log of my thoughts while solving [Advent of Code 2023](https://adventofcode.com/2023).

### Day 1: Trebuchet?!

I initially spent a good chunk of time trying to improve my project "infra", based on what I found
inconvenient from last year (or what I could remember):

- I added a script to grab the day's puzzle input
- I added a utility function to format the path for the day

I also switched from Goland to VSCode, although as I use VSCode more and more, I feel like
JetBrains's autocompletion and code navigation are better. I'll need to ask around about it.

I still have some annoyances with the project structure, namely there's still a bunch of boilerplate
that I would rather not have to write. Things like printing the output for each day, or needing to
run the project specifically from the repo root so that the filepaths work out (which I'm doing
mostly because last year's was structured that way, but if I didn't I think I don't even need the
day path formatting).

Anyway, on to the actual problem. Part 2 was trickier than I expected or remember (from last year),
and I have this tendency to want a clean efficient solution. In this case, I didn't particularly
like this algorithm of checking the string prefix ten times for each character, but hey: I'm never
going to need to look at this code again, and for an input of this size, it still completed in less
than a second. My initial thought was to use some sort of trie, but I'm glad I just banged it out.

This was a long entry; I'm sure as the month progresses I'll write much less 😅

### Day 2: Cube Conundrum

Relatively simple one today, but a good reintroduction to the string parsing aspect of AoC. I dunno
if there are better ways to do it, but I typically use a lot of `strings.Split()`,
`str[len("xyz"):]`, and `strings.HasPrefix/Suffix()`, along with a bunch of branching to handle
specific string terms in the input (in this case, "red", "green", and "blue"). Understandably, these
string operations aren't the algorithmically optimal ways to do it, but hey, it's worked so far.

One improvement to the project infra that I'm thinking of adding, is generating the files for each
day. Since I can't think of a good way to get rid of all the boilerplate, I should at least make it
easier for me to create it. And now is the best time to do it, while the actual problem solving
isn't too difficult yet, haha.
