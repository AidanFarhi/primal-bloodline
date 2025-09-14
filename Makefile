# run app locally
local:
	python ./scripts/local.py

# set app config and deploy to Heroku
deploy:
	python ./scripts/deploy.py
	git push heroku main