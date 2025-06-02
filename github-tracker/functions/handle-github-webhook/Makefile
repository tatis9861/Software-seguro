# Nombre del bucket S3 donde se subirá el archivo de la Lambda
BUCKET_NAME=ciberseguridad-lambda-platzi

# Nombre del archivo .zip que se generará con el binario de la Lambda
BUILD_NAME=handle-github-webhook.zip

# Objetivo "build": compila el binario y lo empaqueta para Lambda
build:
	# Compila el código Go para arquitectura Linux AMD64 (requerido por Lambda)
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

	# Empaqueta el binario "bootstrap" en un archivo ZIP llamado "handle-github-webhook.zip"
	# build-lambda-zip -o $(BUILD_NAME) bootstrap # No funciona el comando build-lambda-zip, así que usamos zip directamente
	zip $(BUILD_NAME) bootstrap

# Objetivo "publish": ejecuta primero "build", luego sube el archivo .zip al bucket S3
publish: build
	# Usa AWS CLI para subir el ZIP al bucket de S3
	aws s3 cp $(BUILD_NAME) s3://$(BUCKET_NAME)/$(BUILD_NAME)
