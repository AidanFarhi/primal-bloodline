import os
import subprocess

env = os.environ.copy()

with open("local.env", "r") as env_file:
    for line in env_file.readlines():
        key, val = line.strip().split("=")
        env[key] = val

subprocess.run("air", env=env)
