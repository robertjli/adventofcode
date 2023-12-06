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

### Day 3: Gear Ratios

I'm thinking about part 1 (and so this blog is now also a notepad), and two solutions come to mind:

- Create a `[][]rune` and store the entire schematic as is
- Create a map of each entry in the schematic, with the key being its location

Each of these storage solutions has two algorithmic solutions to calculate the final answer:

- Go through each number and check if it's adjacent to a symbol
- Go through each symbol and add all adjacent numbers (keeping in mind duplicates)

I'm going back and forth between the four solutions, as I don't think any is much easier or harder.
I'll create a map and iterate the numbers. Let's go.

...

I've just finished the problem (on Dec 4, I was too tired last night to work on it), and I think my
chosen solution worked out well. Although for part 2, I did initially try to iterate on symbols
instead of numbers, and I realized it was a bit tricky to determine if a number was a neighbor of a
symbol, while the other way around was very easy. So I ended up iterating on numbers and storing a
map of potential gears and their "pivots". O(n) space (where n is the number of numbers), but hey,
it worked. That's been the theme of this year so far 😂

### Day 4: Scratchcards

LOL, let's go O(N²) algorithm! Took a peek at the input, it seems small enough 🤣. I'm guessing
there will be some complication with part 2 that will mess it up, though, but I'll figure that out
when I get there.

...

Wow, part 2 was _not_ what I expected. I think it can be done with an array of number of copies.
When a card wins, it adds its copy count to the following cards. Okay, mentally this checks out,
let's go.

...

Well that worked out easily, I'm glad I took the risk of doing the quick, inefficient solution in
part 1! That's a good lesson for me in the future.

### Day 5: If You Give A Seed A Fertilizer

The input values are ten digits, so it's not feasible to actually create maps in memory. I'll need
to translate each seed value as I parse the input.

...

That worked for part 1, but part 2 looks like over a billion seed values, so I won't be able to
calculate each one. I think the key is that the mappings are all continuous ranges, so I only need
to test the lowest value. That does mean that one input range could map to multiple output ranges,
if it spans multiple ranges in the mapping.

...

There was one complication, that if a mapping is not defined then the destination equals the source.
Since an input range could be split, with one portion within a defined mapping and one portion with
no defined mapping, I had to figure out a way to account for that remainder. I finally realized that
I could just create mapping for these "gaps", then I wouldn't need to keep track of what portions of
ranges weren't accounted for.
