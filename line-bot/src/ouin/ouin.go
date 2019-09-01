package ouin

import(
  "github.com/ikawaha/kagome/tokenizer"
)

type Token struct{
  Surface string
  Hiragana string
  Type string
}

func GetHiraganas(sentence string)[]Token {
  list := []Token{}
  t := tokenizer.New()
	tokens := t.Tokenize(sentence)
  for _, token := range tokens {
		if token.Class == tokenizer.DUMMY {
			// BOS: Begin Of Sentence, EOS: End Of Sentence.
			continue
		}
    list = append(list,Token{Surface:token.Surface,Type:token.Features()[0],Hiragana:token.Features()[7]})
	}
  return list
}
