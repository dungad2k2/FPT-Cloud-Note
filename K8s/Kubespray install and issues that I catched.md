# Objectives:
This article will guide beginners who want to set up a Kubernetes cluster using Kubespray. They also cover various cases and errors encountered during the installation.
# Outline
1. Introduction 
2. Installation step by step
3. Some errors and issues occurs 

## 1. Introduction
K8s is a container orchestration platform. It has some ways to install and setup a K8s cluster but in this article will only mention about kubespray. 
[Kubespray](https://kubespray.io/#/) is a method that can automate deploy a cluster k8s (multinode or allinone) by Ansible. I choose kubespray because it uses Ansible :D. 
## 2. Installation step by step
Install K8s by kubespray is very simple and easy to follow. All you need to do is setup an inventory file about your cluster, and config some components of K8s(DNS, CNI,....)-if you need and one admin node that installed docker.
- First, git clone the example config file from kubespray and move it to you set-up folder.
```
git clone https://github.com/kubernetes-sigs/kubespray.git -b <branch-name>
```
- Second, cd to the `/inventory/sample` after git clone:
- Third, change ansible's variable to set info about nodes like hostname/IP, ssh-user, ssh-key/ssh-password. 
- After that run kubespray container to set up K8s cluster: 
```
docker run --rm -it --mount type=bind,source="$(pwd)"/inventory/sample,dst=/inventory \ --mount type=bind,source="${HOME}"/.ssh/id_rsa,dst=/root/.ssh/id_rsa \ quay.io/kubespray/kubespray:<version> bash
#After exec to container run this command:
ansible-playbook -i /inventory/inventory.ini --private-key /root/.ssh/id_rsa cluster.yml
```
**Note**: Kubespray container version need to match with the version of kubespray that cloned before.
## 3. Optional
With this install, K8s cluster will use calico as the default CNI (Container Network Interfaces). If you want to change edit some files in this directory `~/kubespray/inventory/sample/group_vars/k8s_cluster`.
## References:
- https://kubespray.io/?ref=kodekloud.com#/
- https://github.com/kubernetes-sigs/kubespray
