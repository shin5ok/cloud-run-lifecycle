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
    with open(infile) as f:
        reader = csv.reader(f) 
        with open(outfile, "w", newline="") as w:
            writer = csv.writer(w, delimiter=',')
            for v in reader:
                print(v[-2:1])
                writer.writerow(v[-2:1])
                # location = v[0]
                # distance = _metre(v[1])
                # speed = v[2]
                # newrow = [date,location,distance,speed]
                # writer.writerow(newrow)
                # if debug_print:
                #     print(newrow)

if __name__ == '__main__':
    rewrite()