import re


def parse(markdown):
    res = ''
    in_list = False
    in_list_append = False
    # 1) we never use lines
    # 2) good names save lives
    for line in markdown.split('\n'):
        # no need for 'is not None', redundant
        # can use a range for number of possible headers
        if re.match(r'#{1,6} .*', line):
            line = f'<h{(c := line.index(" "))}>{line[c + 1:]}</h{c}>'
        if (list_item := re.match(r'\* (.*)', line)):  # walrus for life
            # turns out we never '</ul>'use is_bold or is_italics
            at_start = ''
            if not in_list:
                in_list = True
                at_start = '<ul>'
            # let's keep things DRY, not WET
            curr = list_item.group(1)
            line = f'{at_start}<li>{curr}</li>'
        elif in_list:
            in_list_append = True
            in_list = False
        if not re.search('<h|<ul|<li', line):  # there won't be a p tag yet
            line = f'<p>{line}</p>'
        if (strong := re.match('(.*)__(.*)__(.*)', line)):
            line = f'{strong.group(1)}<strong>{strong.group(2)}</strong>{strong.group(3)}'
        if (em := re.match('(.*)_(.*)_(.*)', line)):
            line = f'{em.group(1)}<em>{em.group(2)}</em>{em.group(3)}'
        if in_list_append:
            line = f'</ul>{line}'
            in_list_append = False
        res += line
    if in_list:
        res += '</ul>'
    return res
