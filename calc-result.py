import csv
import sys
import os
import click
import re

@click.command()
@click.argument("infile")
@click.argument("outfile")
def rewrite(infile, outfile):
    results = []
    with open(infile) as f:
        reader = csv.reader(f) 
        with open(outfile, "w", newline="") as w:
            writer = csv.writer(w, delimiter=',')
            for v in reader:
                try:
                    item = v[-2:]
                    # print(item)
                    m = re.findall(r"new-min-instance:\s\((.+)\)\sterminated:", item[0])
                    if len(m) == 0:
                        continue
                    print(item[1], m[0])
                except Exception as e:
                    print(e)

if __name__ == '__main__':
    rewrite()