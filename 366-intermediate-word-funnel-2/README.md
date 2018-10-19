# [2018-08-22] Challenge #366 [Intermediate] Word funnel 2

Challenge
A word funnel is a series of words formed by removing one letter at a time from a starting word, keeping the remaining letters in order. For the purpose of this challenge, a word is defined as an entry in the enable1 word list. An example of a word funnel is:

gnash => gash => ash => ah
This word funnel has length 4, because there are 4 words in it.

Given a word, determine the length of the longest word funnel that it starts. You may optionally also return the funnel itself (or any funnel tied for the longest, in the case of a tie).

Examples
funnel2("gnash") => 4
funnel2("princesses") => 9
funnel2("turntables") => 5
funnel2("implosive") => 1
funnel2("programmer") => 2
Optional bonus 1
Find the one word in the word list that starts a funnel of length 10.

Optional bonus 2
For this bonus, you are allowed to remove more than one letter in a single step of the word funnel. For instance, you may step from sideboard to sidebar by removing the o and the final d in a single step. With this modified rule, it's possible to get a funnel of length 12:

preformationists =>
preformationist =>
preformations =>
reformations =>
reformation =>
formation =>
oration =>
ration =>
ratio =>
rato =>
rat =>
at
preformationists is one of six words that begin a modified funnel of length 12. Find the other five words.

---

## Output

```

$ go run solution.go
vim-go
&{gnash map[gash:0xc0043d97a0] map[] 0}
 gnash
  gash
   ash
    sh
    ah
    as
   gas
    as
4 4
9 9
5 5
1 1
2 2
twitchiest 9
streambeds 9
princesses 9
complecting 10
insolating 9
discussers 9
complected 9
stampeders 9
completing 9
staunchest 9
strappings 9
stranglers 9
composited 9
scratchers 9
consigning 9
carrousels 9
restarting 9
colonizers 9
estrangers 9
glassiness 9
 complecting
  completing
   competing
    compting
     comping
      coping
       oping
        ping
         pig
          pi
         pin
          in
          pi
      coming
```
