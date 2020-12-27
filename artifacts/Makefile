
dcrds:
	kubectl delete -f ../pkg/provision/cluster-asm/crds/
dasm:
	kubectl delete asm asm-sample

checkcrds:
	kubectl get crd | grep asm.alauda.io|grep -v NAME|wc -l
checkcrdsall:
	kubectl get crd | grep -v NAME|wc -l

c:
	kubectl apply -f ../config/samples/operator_v1alpha1_asm.yaml