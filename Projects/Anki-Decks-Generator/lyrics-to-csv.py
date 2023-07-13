#!/usr/bin/env python

import re, sys, csv
REGEX = re.compile(r'^(?P<es>.*)$\s*^(?P<en>.*)$\s*', re.MULTILINE)


def get_verses(s: str) -> [list]:
    regexs = []
    double = []

    for re in REGEX.finditer(s):
        group = re.groupdict()

        if group["en"] == "" or group["es"] == "":
            continue
        elif group["en"] in double:
            continue

        regexs.append([group["en"], group["es"]])
        double.append(group["en"])

    return regexs


def main():
    txt: str
    with open(sys.argv[1]) as file:
                          txt = file.read()

    verses = get_verses(txt)
    # Write Rows.
    with open(sys.argv[1]+".csv", "w", newline="") as file:
            writer = csv.writer(file, delimiter=',',
                            quotechar='"', quoting=csv.QUOTE_MINIMAL)
            for verse in verses:
                writer.writerow(verse)


main()
