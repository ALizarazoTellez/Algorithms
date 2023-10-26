#!/usr/bin/env python3

import csv, sys

MULTIPLE = 20

EXCLUDES = [
    0,
    1,
    2,
    10,
    11,
]

def main():
    tables: list[tuple[3]] = []

    for a in range(MULTIPLE+1):
        if a in EXCLUDES: continue

        for b in range(MULTIPLE+1):
            if b > a: continue
            if b in EXCLUDES: continue

            tables.append((a, b, a*b))

    writer = csv.writer(sys.stdout, delimiter=',',
                        quotechar='"', quoting=csv.QUOTE_MINIMAL)
    for table in tables:
        writer.writerow(table)


if __name__ == "__main__":
    main()
