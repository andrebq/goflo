.PHONY: start-noflo ssh

start-noflo:
	docker run --rm -it -p 9999:9999 flowhub/noflo-ui

ssh:
	cd ssh && \
	docker build -t ssh-tunnel .
