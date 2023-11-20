ASSET_BINNARY=app.exe

build:
	@echo start build
	set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ./dist/${ASSET_BINNARY} ./cmd/mitra
	@echo done building ..

start:
	@echo starting app
	@env  ./dist/${ASSET_BINNARY} &
	@echo "running ..."

stop:
	@echo Stopping ...
	@taskkill /IM ${ASSET_BINNARY} /F
	@echo Stopped