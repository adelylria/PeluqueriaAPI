buildimage: 
	docker build -t peluqueria_api .

runimage:
	docker run --name peluqueria_api -p 8080:8080 peluqueria_api

.PHONY: buildimage runimage