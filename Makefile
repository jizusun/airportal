.PHONY: cfssl-cert-gen run

run: 
	air

install-cfssl:
	go get github.com/cloudflare/cfssl/cmd/cfssl; \
	go get github.com/cloudflare/cfssl/cmd/cfssljson

cfssl-cert-gen:
	mkdir -p certs; \
	cd certs ;\
	rm *.pem ;\
	sh ../scripts/cfssl.sh

rsync: 
	./scripts/rsync.sh