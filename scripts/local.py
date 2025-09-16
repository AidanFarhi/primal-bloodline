import os
import signal
import subprocess

env = os.environ.copy()

with open("local.env", "r") as env_file:
    for line in env_file.readlines():
        key, val = line.strip().split("=")
        env[key] = val

proc = subprocess.Popen(["air"], env=env)

try:
    proc.wait()
except KeyboardInterrupt:
    proc.send_signal(signal.SIGINT)
    proc.wait()
