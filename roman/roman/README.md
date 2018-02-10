# LetsGo! Roman Calculator Lab

> Provides calculators for converting an arabic number
to a roman glyph and vice versa

## Installation

```shell
go get github.com/imhotepio/letsgo_labs/roman
```

## Usage

### Converts Arabic To Roman

```go
g, err := roman.ToRoman(10) // => X
```

### Converts Roman To Arabic

```go
a, err := roman.ToArabic("X") // => 10
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)