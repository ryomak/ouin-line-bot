package ouin

import (
 // "github.com/ikawaha/kagome/tokenizer"
)

type Token struct {
	Surface  string
	Hiragana string
	Type     string
}

func GetHiraganas(sentence string) []Token {
	list := []Token{
    Token{Surface:"私",Hiragana:"わたし",Type:"主語"},
    Token{Surface:"は",Hiragana:"は",Type:"助詞"},
    Token{Surface:"隆",Hiragana:"たかし",Type:"主語"},
    Token{Surface:"です",Hiragana:"です",Type:"主語"},
  }
//	t := tokenizer.New()
//	tokens := t.Tokenize(sentence)
//	for _, token := range tokens {
//		if token.Class == tokenizer.DUMMY {
//			// BOS: Begin Of Sentence, EOS: End Of Sentence.
//			continue
//		}
//		list = append(list, Token{Surface: token.Surface, Type: token.Features()[0], Hiragana: token.Features()[7]})
//	}
	return list
}
