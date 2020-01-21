import re
def count_words(sentence):
    dict = {}
    sentence = re.sub("[^a-zA-Z0-9']", " ", sentence)
    for word in sentence.split():
        if (w := word.lower())[0] == "'":
            w = w[1:]
        if w[-1] == "'":
            w = w[:-1]
        if w not in dict:
            dict[w] = 0
        dict[w] += 1
    return dict
