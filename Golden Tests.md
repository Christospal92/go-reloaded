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
Expected: AN EAGLE, WITH 31 EGGS, SHE said.
```

**Example 14**
```
Input: a owl (up,3) , with 1A (hex) eggs .
Expected: AN OWL, WITH 26 EGGS.
```

**Example 15**
```
Input: 10 (bin) birds sat on a branch (up,5) !
Expected: 2 BIRDS SAT ON AN BRANCH!
```

**Example 16**
```
Input: ' a eagle and a apple ' (up,4)
Expected: 'AN EAGLE AND AN APPLE'
```

