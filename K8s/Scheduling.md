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
