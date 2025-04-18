### Manual Scheduling
Mỗi pod sẽ có một trường được đặt tên là NodeName, mặc định khi tạo Pod thì trường này sẽ không được set, k8s sẽ add vào mainfest file một cách tự động. 
Khi mà muốn schedule POD vào một node thì ta cần phải tiền hành thêm một binding object vào manifest file.
Ví dụ:
```
apiVersion: v1
kind: Pod
metadata:
 name: nginx
 labels:
  name: nginx
spec:
 containers:
 - name: nginx
   image: nginx
   ports:
   - containerPort: 8080
 nodeName: node02
```
### Taints và Tolerations
Taints và Tolerations được sử dụng để set restrictions trên pod có thể được schedule đến những node nào. 
Taints định nghĩa là một thuộc tính được áp dụng cho một node để có thể từ chối việc schedule pod trên node đó, nếu không pod sẽ phải matching toleration cho taint node đó. 
Taints được sử dụng cho việc reserving nodes. Một taints sẽ bao gồm:
- Key: unique indentifier cho taints
- Value: optional value được gắn với key
- Effect: Định nghĩa hành vi cho taints này. Trong effect có 3 behavior là: **Noschedule**(Pod mà không matching được toleration thì không được đặt trên node), **PreferNoSchedule**(kubernetes cố gắng tránh scheduling pod vào node nếu không matching toleration, nhưng nếu không còn node nào available thì node này sẽ được schedule), **NoExecute**(Pod được tạo mà không matching toleration thì sẽ bị evicted và pod mới sẽ không được schedule)
Tạo taints cho một node 
```
kubectl taint nodes <nodename> <key>=<value>:<effect>
##remove taint
kubectl taint nodes <nodename> <key>=<value>:<effect>-
```
### Tolerations
Toleration là một thuộc tính áp dụng cho một pod cho phép nó có thể schedule trên các node mà matching taint. Tolerations được định nghĩa trong pod spec và phải match về taint key, value và effect để có thể được schedule trên tainted node. Tolerations không bảo đảm việc pod sẽ được schedule trên một node cụ thể nào đó, nó chỉ cho phải schedule trên những tainted node. Node selectors, affinities, hoặc các phương thưc scheduling khác sẽ quyết định cuối cùng
Ví dụ:
```
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  tolerations:
  - key: "key"
    operator: "Equal"
    value: "value"
    effect: "NoSchedule"
```
### Node Selector và NodeAffinity
Để sử dụng Node Selector ta chỉ cần thêm một thuộc tính gọi là Node Selector và specify label trong file triển khai. Để sử dụng Node selector cần phải gán label cho node trước:
```
kubectl label nodes <node-name> <label-key>=<label-value>
```
Ví dụ
```
apiVersion: v1
kind: Pod
metadata:
 name: myapp-pod
spec:
 containers:
 - name: data-processor
   image: data-processor
 nodeSelector:
  size: Large
```
Ở đây size là label key còn Large là label value.
Tuy nhiên để tùy chỉnh và có những exception ở mức chi tiết hơn trong việc chọn node để schedule thì ta cần phải sử dụng NodeAffinity. 
Trong NodeAffinity có các dạng operator như sau:
- In : node phải có label key và giá trị của nó phải match với một trong các value được lists
- NotIn: node phải có label key và giá trị của nó phải **không** match với value nào được list.
- Exists: node phải có label key (value không quan trọng)
- DoesNotExist: Node phải không có label key
- Gt: Label value phải lớn hơn một giá trị nào cho trước
- Lt: Label value phải nhỏ hơn một giá trị cho trước 
Show label in node 
```
kubectl get node <node-name> --show-labels
```
Show which nodes are the pods placed on 
```
kubectl get pods -o wide
```
### Resource Limits
Mặc định mỗi container trong pod hoặc một pod đều yêu cầu 0.5 CPU và 256 Mb memory. Nếu application trong pod yêu cầu nhiều hơn default resources ta cần phải set thêm ở file definition, resource này được gọi là **resource request**.
Ví dụ:
```
apiVersion: v1
kind: Pod
metadata:
  name: simple-webapp-color
  labels:
    name: simple-webapp-color
spec:
 containers:
 - name: simple-webapp-color
   image: simple-webapp-color
   ports:
    - containerPort:  8080
   resources:
     requests:
      memory: "1Gi"
      cpu: "1"
```
Đối với resource limit default k8s sẽ set là 1CPU và 512 Mb RAM. 
Khi resource dành cho pod exceed limit:
- Đối với CPU thì sẽ xảy ra hiện tượng CPU throttling. Pod sẽ không bị killed nhưng CPU usage sẽ bị giới hạn container sẽ chạy lâu hơn.
- Đối với memory thì container sẽ OOMKilled. Nó có thể sẽ restart liên tục nếu option `RestartPolicy: Always` được bật.
### DaemonSet
DaemonSet đảm bảo một instance đơn lẻ của một pod nào đó được chạy trên tất cả các node trong cluster. DaemonSet thường được dùng cho các task chạy background hoặc system daemon trên mỗi node như là log collector, exporter, network proxies,...
Khác với ReplicaSet, thì các pod khi sử dụng ReplicaSet sẽ được phân ra các node theo sự chỉ định của scheduler. 
![[Pasted image 20250418142436.png]]
