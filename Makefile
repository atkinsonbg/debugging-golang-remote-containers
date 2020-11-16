docker:
	docker build -t github.com/atkinsonbg/debugging-golang-remote-containers:latest .

dockerrun:
	docker run -it github.com/atkinsonbg/debugging-golang-remote-containers:latest