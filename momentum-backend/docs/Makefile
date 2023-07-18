check-swagger:
	swagger version

swagger: check-swagger
	swagger generate spec -o ./swagger.yaml --scan-models

serve-swagger: check-swagger
	swagger serve -F=swagger swagger.yaml

install-swagger:
	apt update
	apt install -y apt-transport-https gnupg curl
	curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | apt-key add -
	curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/config.deb.txt?distro=debian&codename=any-version' > /etc/apt/sources.list.d/go-swagger-go-swagger.list
	apt update 
	apt install swagger