import re
from collections import Counter
def count_words(sentence):
    dict = Counter()
    sentence = re.sub("[^a-zA-Z0-9']", " ", sentence)
    for word in sentence.split():
        if (w := word.lower())[0] == "'":
            w = w[1:]
        if w[-1] == "'":
            w = w[:-1]
        dict[w] += 1
    return dict
