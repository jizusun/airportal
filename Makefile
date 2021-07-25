.PHONY: cfssl-cert-gen

cfssl-cert-gen:
	mkdir -p certs; \
	cd certs ;\
	rm *.pem ;\
	sh ../scripts/cfssl.sh

rsync: 
	rsync -azP certs/{server,server-key,ca}.pem config/server.json \
		virmach:~/trojan-go-linux-amd64
