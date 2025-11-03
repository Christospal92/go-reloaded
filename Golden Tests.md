# Go-Reloaded: Golden Test Set — Full List & Content

Includes basic and combined input/output examples. Used as the source of truth.

## Numbers

**Example 1**
```
Input: 1E (hex) and 10 (bin)
Expected: 30 and 2
```

**Example 2**
```
Input: FF (hex) and 1111 (bin)
Expected: 255 and 15
```

**Example 3**
```
Input: 111000 (bin)
Expected: 56
```

## Casing

**Example 4**
```
Input: This is so exciting (up, 2)
Expected: This is SO EXCITING
```

**Example 5**
```
Input: lowercase that (low, 2)
Expected: lowercase that
```

**Example 6**
```
Input: mixed CASE (up, 2)
Expected: MIXED CASE
```

## Punctuation & Quotes

**Example 7**
```
Input: ' I am a optimist , '
Expected: 'I am an optimist,'
```

**Example 8**
```
Input: ' mix of , punctuation . and quotes '
Expected: 'mix of, punctuation. and quotes'
```

**Example 9**
```
Input: ' spaces   around , commas '
Expected: 'spaces around, commas'
```

## Articles (a → an)

**Example 10**
```
Input: a ice-cream
Expected: an ice-cream
```

**Example 11**
```
Input: a owl
Expected: an owl
```

**Example 12**
```
Input: a elephant
Expected: an elephant
```

## Tricky Combined Examples

**Example 13**
```
Input: a eagle , with 1F (hex) eggs (up,3) , she said .
Expected: an eagle, WITH 31 EGGS, she said.
```

**Example 14**
```
Input: a owl (up,3) , with 1A (hex) eggs .
Expected: AN OWL, with 26 eggs.
```

**Example 15**
```
Input: 10 (bin) birds sat on a branch (up,5) !
Expected: 2 BIRDS SAT ON A BRANCH!
```

**Example 16**
```
Input: ' a eagle and a apple ' (up,4)
Expected: 'AN EAGLE AND AN APPLE'
```

## Auditors Examples

**Example 17**
```
Input: If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?.
Expected: If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?
```

**Example 18**
```
Input: I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure.
Expected: I have to pack 5 outfits. Packed 26 just to be sure.
```

**Example 19**
```
Input: Don not be sad ,because sad backwards is das . And das not good.
Expected: Don not be sad, because sad backwards is das. And das not good.
```

**Example 20**
```
Input: harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
Expected: Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
```

## Edge cases Examples

**Example 21**
```
Input: This stays the same (up, 0)
Expected: This stays the same
```

**Example 22**
```
Input: He said: '  '
Expected: He said: ''
```

**Example 23**
```
Input: Go BIG Or Go Home (low, 10) !
Expected: go big or go home!
```
