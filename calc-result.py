import csv
import sys
import os
import click
import re

def _metre(v):
    k = re.findall(r"(\d*)K",v)
    m = re.findall(r"([\d\.]+)M",v)
    totalm = 0.0
    try:
        if len(m) > 0:
            totalm = float(m[0])
        if len(k) > 0:
            totalm += float(k[0]) * 1000
    except Exception as e:
        print(e)
    return totalm

@click.command()
@click.argument("infile")
@click.argument("outfile")
# @click.option("--date", required=True)
# @click.option("--debug_print", is_flag=True)
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