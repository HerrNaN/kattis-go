#!/bin/python
import os
import sys
import subprocess
import time

if len(sys.argv) != 2:
    raise ValueError("Wrong number of arguments")

path = sys.argv[1]
if not path.endswith("/"):
    path += "/"

sample_path = path+"samples/"

files = os.listdir(sample_path)
files.sort()

ins = []
outs = []

for name in files:
    if name.endswith(".in"):
        ins.append(name)
    if name.endswith(".ans"):
        outs.append(name)

if len(ins) != len(outs):
    raise ValueError("Mismatching number of inputs and outputs")

for n in range(len(ins)):
    with open(sample_path+ins[n]) as i:
        with open(sample_path+outs[n]) as o:
            expected = o.read()
            try:
                t1 = time.time()
                completed = subprocess.run(["go", "run", path+"main.go"], stdin=i, capture_output=True, check=True, text=True)
                t2 = time.time()
                if completed.stdout != expected:
                    print(f"# {n+1} - WRONG ANSWER [{round((t2-t1) * 1000)}ms]")
                    print("EXPECTED:")
                    print(expected)
                    print("ACTUAL")
                    print(completed.stdout)
                else:
                    print(f"# {n+1} - CORRECT [{round((t2-t1) * 1000)}ms]")

            except subprocess.SubprocessError:
                print(f"# {n+1} - ERROR")

