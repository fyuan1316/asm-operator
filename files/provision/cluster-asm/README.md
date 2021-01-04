# ASM Helm Chart

ASM Chart based on istio chart 1.0.5. 

## Requirements

ASM chart required diablo , portal and auth-controller.So before you start ,you should install following charts:
##### portal
##### dashboard
##### appcore


##### Istio 1.0.5
##### Elasticeach 1.17.0  
##### jaeger-operator 2.1.0


## Installing ASM Chart
kubectl create ns istio-system

apiVersion: v1
data:
  .dockerconfigjson: eyJhdXRocyI6eyJpbmRleC5hbGF1ZGEuY24iOnsidXNlcm5hbWUiOiJhbGF1ZGFrOHMvcHVsbCIsInBhc3N3b3JkIjoiZVY2V3V3TkQyOTRlIiwiZW1haWwiOiJwdWxsQGFsYXVkYS5pbyIsImF1dGgiOiJZV3hoZFdSaGF6aHpMM0IxYkd3NlpWWTJWM1YzVGtReU9UUmwifX19
kind: Secret
metadata:
  name: alaudak8s
type: kubernetes.io/dockerconfigjson

kubectl create -f alaudak8s.yaml -n istio-system

helm install . -n istio --namespace istio-system
--set global.proxy.includeIPRanges="10.96.0.0/12, 10.222.0.0/16"  # Service和Pod的IP range
--set global.hub="harbor.alauda.cn/alaudak8s"      # 镜像仓库地址

helm install . -n jaeger-operator --namespace istio-system


helm install . -n asm  --namespace istio-system
--set jaeger.elasticsearch.serverurl = "http://es-elasticsearch-client.istio-system:9200"   #### es 服务地址
--set jaeger.elasticsearch.username = elastic                                                ### es用户名
--set jaeger.elasticsearch.password = changeme                                               ### es密码   
--set  global.host = "asf-dev.alauda.cn"                                      ### 服务访问方式


## Uninstalling ASM  Chart
helm del --purge asm


kubectl delete deployment asm-controller -n alauda-system ;\   
kubectl delete configmap asm-config -n alauda-system ;\
kubectl delete sa servicemesh-controller  -n alauda-system ;\
kubectl delete clusterrolebindings asm:alauda-system:cluster-admin -n alauda-system; \
kubectl delete clusterrolebindings asm:system:auth-delegator -n alauda-system; \
kubectl delete rolebindings asm-auth-reader -n alauda-system ;\
kubectl delete crd microspan.asm.alauda.io ;

#####    helm del --purge istio

    kubectl delete -f ./templates/crds.yaml ;



helm del --purge jaeger-operator

kubectl delete crd jaegers.io.jaegertracing ;\
kubectl delete deployment alauda-asm-jaeger-operator -n istio-system ;

add by lihuang
