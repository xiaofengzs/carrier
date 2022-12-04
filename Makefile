
run: crd dcokerbuild dockerpush apply_controller apply_backend

crd:
	kubectl apply -f artifacts/crd.yaml

build:
	CGO_ENABLE=0 go build -o controller ./

dcokerbuild:
	docker build . -t tr1stanzhi/carrier:latest

dockerpush:
	docker push tr1stanzhi/carrier:latest

apply_backend:
	kubectl apply -f artifacts/backend.yaml

apply_controller:
	kubectl apply -f artifacts/controller-deployment.yaml

.PHONY: crd build dcokerbuild dockerpush apply_controller apply_backend