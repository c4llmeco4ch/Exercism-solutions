import string, re
def abbreviate(words):
    """
    Given the proliferation of 1337speak, 
    I'm allowing for numbers and letters to both be included in acronyms
    """
    return ''.join(letter.upper() for word in re.split(r'[-_ ]',words) if len(word) > 0 and re.match('[a-z0-9A-Z]', letter := word[0]))
