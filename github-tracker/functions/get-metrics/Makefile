# Nombre del bucket de S3 donde se subirá el archivo Lambda
BUCKET_NAME=ciberseguridad-lambda-platzi

# Nombre del archivo ZIP que se generará con el binario de la Lambda
BUILD_NAME=get-metrics.zip

# Objetivo "build" — compila el código Go y genera el archivo ZIP para Lambda
build:
	# Compila el archivo main.go con sistema operativo Linux y arquitectura amd64,
	# y lo guarda con el nombre "bootstrap" (nombre requerido para runtimes personalizados de Lambda)
	GOOS=linux GOARCH=amd64 go build -o bootstrap main.go

	# Empaqueta el binario "bootstrap" en un archivo ZIP llamado "get-metrics.zip"
	#build-lambda-zip -o $(BUILD_NAME) bootstrap #no funciona el comando build-lambda-zip
	zip $(BUILD_NAME) bootstrap

# Objetivo "publish" — primero ejecuta "build", luego sube el ZIP al bucket S3
publish: build
	# Sube el archivo ZIP al bucket S3 especificado
	aws s3 cp $(BUILD_NAME) s3://$(BUCKET_NAME)/$(BUILD_NAME)
