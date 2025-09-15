import subprocess

with open(".env", "r") as env_file:
    for line in env_file.readlines():
        key, val = line.strip().split("=")
        subprocess.run(
            f"heroku config:set {key}={val} --app primal-bloodline", shell=True
        )
