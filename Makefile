# run app locally
local:
	python ./scripts/local.py

# set app config and deploy to Heroku
deploy-with-config:
	python ./scripts/set_heroku_config.py
	git push heroku main

# deploy to Heroku
deploy:
	git push heroku main