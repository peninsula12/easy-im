user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-test

social-rpc-dev:
	@make -f deploy/mk/social-rpc.mk release-test

release-test: user-rpc-dev social-rpc-dev

install-server:
	cd ./deploy/script && chmod +x release-test.sh && ./release-test.sh