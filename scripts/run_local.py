import os
import subprocess

# create a copy of the environment variables
env = os.environ.copy()

# open the local environment vars file and set each one
with open('local.env', 'r') as env_file:
    for line in env_file.readlines():
        key, val = line.strip().split('=')
        env[key] = val

# run the golang air command with the environment variables
subprocess.run('air', env=env)
