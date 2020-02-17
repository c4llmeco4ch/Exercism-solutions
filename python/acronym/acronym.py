import string #, re
def abbreviate(words):
    ans = ''
    for word in words.split():
        cut = word.strip(string.punctuation)
        if len(cut) == 0:
            continue
        if '-' in cut:
            for s in cut.split('-'):
                ans += s[0].upper()
        else:
            ans += cut[0].upper()
    return ans

    # So close: return ''.join(letter.upper() for word in words.split() if re.match('[a-zA-Z]', letter := word[0]))
