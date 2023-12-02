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
