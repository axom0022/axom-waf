build:
	go build -o axomwaf .

run: build
	sudo ./axomwaf

clean:
	rm -f axomwaf

install:
	cp axomwaf /usr/local/bin/
	cp config.json /etc/axomwaf/
