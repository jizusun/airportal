.PHONY: cfssl-cert-gen run

run: 
	air

install-tools:
	go get github.com/cloudflare/cfssl/cmd/cfssl; \
	go get github.com/cloudflare/cfssl/cmd/cfssljson
	yarn global add @aws-amplify/cli

cfssl-cert-gen:
	mkdir -p certs; \
	cd certs ;\
	rm *.pem ;\
	sh ../scripts/cfssl.sh

rsync: 
	./scripts/rsync.sh