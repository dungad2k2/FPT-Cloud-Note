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
Install K8s by kubespray is very simple and easy to follow. All you need to do is setup an inventory file about your cluster, and config some components of K8s(DNS, CNI,....)-if you need. 
- First, git clone the example config file from kubespray and move it to you set-up folder.
```
git clone https://github.com/kubernetes-sigs/kubespray.git -b <branch-name>
```