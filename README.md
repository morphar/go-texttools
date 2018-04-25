# go-texttools
A collection of tools for manipulating texts in Go

Shorten tries to create the most sensible (to a human) shortened text.  
If possible, it will try to cut at a non-word char.  
It will strip newlines and carriage returns.  
```go
Shorten(str string, length int, appendStr string) (shorter string)
```

SpecialCharsToStandard replaces all kinds of non-ascii chars with transliterations.
```go
SpecialCharsToStandard(str string) string
```

Slug will convert a string to a [slug](https://en.wikipedia.org/wiki/Clean_URL#Slug).  
It will also do  transliteration of non-ascii chars.
```go
Slug(str string) string
```

UnCase takes a string in any "case" (kebab-case, snake_case, etc.) and creates a "normal" string.  
E.g. my-slug-string -> "My slug string"
```go
UnCase(str string) string
```

SnakeCase will convert a string to [snake_case](https://en.wikipedia.org/wiki/Letter_case#Special_case_styles).
```go
SnakeCase(str string) string
```

KebabCase will convert a string to [kebab-case](https://en.wikipedia.org/wiki/Letter_case#Special_case_styles).
```go
KebabCase(str string) string
```

CamelCase will convert a string to [camelCase](https://en.wikipedia.org/wiki/Camel_case).
```go
CamelCase(str string) string
```

PascalCase will convert a string to PascalCase.  
This is the same as [camelCase](https://en.wikipedia.org/wiki/Camel_case), but with the first letter capitalized.
```go
PascalCase(str string) string
```

StringInSlice will check if a string is in a slice and return true if it is.
```go
StringInSlice(searchStr string, strs []string) bool
```

HTMLToText converts HTML to standard text.
```go
HTMLToText(html string) (text string)
```

TextSanitizer converts HTML to standard text, but also replaces some special chars and escapings.
```go
SanitizeText(txt string) (newTxt string)
```

CP1258ToUTF8 converts a CP1258 byte array to a UTF-8 string.
```go
CP1258ToUTF8(txt []byte) (utf8Txt string)
```

RandomString creates a secure pseudorandom string using the crypto rand package.
```go
RandomString(n int) (str string)
```
