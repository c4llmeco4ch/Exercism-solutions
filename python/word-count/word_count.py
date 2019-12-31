from string import punctuation
def count_words(sentence):
    dict = {}
    punc = punctuation.replace("'",'')
    for val in punc:
        if val in sentence:
            sentence = sentence.replace(val, ' ')
    for word in sentence.split():
        if (w := word.lower())[0] == "'":
            w = w[1:]
        if w[-1] == "'":
            w = w[:-1]
        if w not in dict:
            dict[w] = 0
        dict[w] += 1
    return dict
