## Advent of Code 2022

This is a log of my thoughts while solving [Advent of Code 2022](https://adventofcode.com/2022).
This log was started on 28 Dec 2022, and I've completed days 1–15, except for day 13. While I was
able to find plenty of time at the beginning of the month, since I was on vacation, it's been more
difficult since my return (and with the increasing difficulty of the problems).

### Day 16: Proboscidea Volcanium

28 Dec: I got part 1 working with a basic DFS, but it takes ~9 minutes to run. I think I have a
better solution by reducing the map to only link nodes that have a positive flow rate, with each
path's cost being the distance from the node. Having trouble getting the code to work, but it's 1 am
and I've spent two hours working on this already, so I'm leaving it for tomorrow.

29 Dec: I rewrote the DFS from scratch with my new reduced graph, and it's super fast—only took 236
ms for part 1. I wrote a pretty hacky part 2 with the same graph and similar logic. It took ~5
minutes to run, but gave me the right answer. Sadly I probably won't go back and improve it. I spent
about an hour on this problem today.

### Day 17: Pyroclastic Flow

3 Jan: I've been spending a bit of time here and there for this problem. It seems like it's just
begging for an object-oriented approach, so I set up different a different class for each rock and
ran the simulation for part 1 just fine. It's not happy with part 2, though.

I wonder if it's a memory allocation problem—I could look into using a circular buffer for the
chamber, because we don't care about lower sections of the chamber once they're completely covered.