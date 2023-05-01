package main

import (
	"math/rand"
	"net/mail"
	"net/url"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type stringv2 string

func (s stringv2) TrimWhitespace() stringv2 {
	return stringv2(strings.TrimSpace(string(s)))
}

func (s stringv2) ToSnakeCase() stringv2 {
	var result strings.Builder
	for i, c := range s {
		if unicode.IsUpper(c) {
			if i > 0 {
				result.WriteRune('_')
			}
			result.WriteRune(unicode.ToLower(c))
		} else {
			result.WriteRune(c)
		}
	}
	return stringv2(result.String())
}

func (s stringv2) Reverse() stringv2 {
	var result strings.Builder
	runes := []rune(string(s))
	for i := len(runes) - 1; i >= 0; i-- {
		result.WriteRune(runes[i])
	}
	return stringv2(result.String())
}

func (s stringv2) CountOccurrences(substring stringv2) int {
	return strings.Count(string(s), string(substring))
}

func (s stringv2) ReplaceAllOccurrences(substring, replacement stringv2) stringv2 {
	return stringv2(strings.ReplaceAll(string(s), string(substring), string(replacement)))
}

func (s stringv2) IsPalindrome() bool {
	runes := []rune(string(s))
	for i := 0; i < len(runes)/2; i++ {
		if runes[i] != runes[len(runes)-i-1] {
			return false
		}
	}
	return true
}

func (s stringv2) Split(separator stringv2) []stringv2 {
	parts := strings.Split(string(s), string(separator))
	result := make([]stringv2, len(parts))
	for i, part := range parts {
		result[i] = stringv2(part)
	}
	return result
}

func (s stringv2) PadLeft(length int, padChar rune) stringv2 {
	result := strings.Builder{}
	if len(s) < length {
		for i := 0; i < length-len(s); i++ {
			result.WriteRune(padChar)
		}
	}
	result.WriteString(string(s))
	return stringv2(result.String())
}

func (s stringv2) PadRight(length int, padChar rune) stringv2 {
	result := strings.Builder{}
	result.WriteString(string(s))
	if len(s) < length {
		for i := 0; i < length-len(s); i++ {
			result.WriteRune(padChar)
		}
	}
	return stringv2(result.String())
}

func (s stringv2) ToCamelCase() stringv2 {
	parts := s.Split(" ")
	result := make([]string, len(parts))
	for i, part := range parts {
		if i == 0 {
			result[i] = strings.ToLower(string(part))
		} else {
			result[i] = strings.Title(strings.ToLower(string(part)))
		}
	}
	return stringv2(strings.Join(result, ""))
}

func (s stringv2) Contains(substr stringv2) bool {
	return strings.Contains(string(s), string(substr))
}

func (s stringv2) ContainsRune(r rune) bool {
	return strings.ContainsRune(string(s), r)
}

func (s stringv2) Cut(sep stringv2) (before, after stringv2, found bool) {
	before1, after1, found1 := strings.Cut(string(s), string(sep))
	return stringv2(before1), stringv2(after1), found1
}

func (s stringv2) CutPrefix(prefix stringv2) (stringv2, bool) {
	after, found := strings.CutPrefix(string(s), string(prefix))
	return stringv2(after), found
}

func (s stringv2) CutSuffix(suffix stringv2) (stringv2, bool) {
	before, found := strings.CutSuffix(string(s), string(suffix))
	return stringv2(before), found
}

func (s stringv2) IsEqualFold(t stringv2) bool {
	return strings.EqualFold(string(s), string(t))
}

func (sep stringv2) Join(elems []stringv2) stringv2 {
	var result stringv2
	for i, elem := range elems {
		if i == 0 {
			result += elem
		} else {
			result += sep + elem
		}
	}
	return result
}

func (s stringv2) ToUpper() stringv2 {
	return stringv2(strings.ToUpper(string(s)))
}

func (s stringv2) ToLower() stringv2 {
	return stringv2(strings.ToLower(string(s)))
}

func (s stringv2) ToUTF8() []byte {
	rs := []rune(s)
	n := 0
	for _, r := range rs {
		n += utf8.RuneLen(r)
	}
	n, bs := 0, make([]byte, n)
	for _, r := range rs {
		n += utf8.EncodeRune(bs[n:], r)
	}
	return bs
}

func (s stringv2) RandomASCIIString(length int) stringv2 {
	ascii := stringv2("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")
	var result stringv2
	for i := 1; i <= length; i++ {
		inta := rand.Intn(len(ascii))
		result += stringv2(rune(ascii[inta]))
	}
	return result
}

func (s stringv2) GetRandomRuneFromString() rune {
	return rune(s[rand.Intn(len(s))])
}

func (s stringv2) IsDigit() bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func (s stringv2) IsEmail() bool {
	_, err := mail.ParseAddress(string(s))
	return err == nil
}

func (s stringv2) IsURL() bool {
	_, err := url.Parse(string(s))
	return err == nil
}

func (s stringv2) IsValidFloat() bool {
	_, err := strconv.ParseFloat(string(s), 64)
	return err == nil
}

func (s stringv2) Shuffle() stringv2 {
	if s == "" {
		return s
	}
	runes := []rune(s)
	index := 0
	for i := len(runes) - 1; i > 0; i-- {
		index = rand.Intn(i + 1)
		runes[i], runes[index] = runes[index], runes[i]
	}
	return stringv2(string(runes))
}
