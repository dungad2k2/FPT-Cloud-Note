# OpenStack Multinode Installation (Octavia)
## Set up environment
| Node | Spec |
| ----- | -----|
| Compute Node | Ubuntu 22.04, 8GB RAM, 2 vCPUs <br> 2 network interfaces: ens33 (192.168.52.131), ens37|
| Controll Node | Ubuntu 22.04, 8GB RAM, 2 vCPUs <br> 2 network interfaces: ens33 (192.168.52.132), ens37|

**This guide is set up by VMware workstation. You can use another hypervisor to do**
Install requirements dependencies:
- ` sudo apt update ` 
- `sudo apt install git python3-dev libffi-dev gcc libssl-dev`

Install virtual environment:
- `sudo apt install python3-venv`
- `python3 -m venv /path/to/venv`: create an environment 
- `source /path/to/env/bin/activate`: active an environment that created 
- `pip install -U pip`: install pip 
- `pip install 'ansible-core>=2.16,<2.17.99'`: install ansible-core 

Install kolla-ansible: 
1. `pip install git+https://opendev.org/openstack/kolla-ansible@master`: install kolla-ansible and its dependencies
2. `sudo mkdir -p /etc/kolla`: create a kolla's directory 
3. `sudo chown $USER:$USER /etc/kolla`: change own for user 
4. `cp -r /path/to/venv/share/kolla-ansible/etc_examples/kolla/* /etc/kolla`: Copy **globals.yml** and **passwords.yml** to **/etc/kolla** directory.
5. `cp /path/to/venv/share/kolla-ansible/ansible/inventory/multinode .`: Copy multinode inventory file to the current directory
   
Install Ansible Galaxy requirements: 
`kolla-ansible install-deps`

Generate Kolla password: 
`kolla-genpwd`

Edit **globals.yml** file:
```yaml
kolla_base_distro: "ubuntu"
# VIP address is an unused IP address
Kolla_internal_vip_address: "192.168.52.140"
network_interface: "ens33"
neutron_external_interfaces: "ens37"
enable_neutron_provider_networks: "yes"
enable_octavia: "yes"
octavia_certs_country: US
octavia_certs_state: Oregon
octavia_certs_organization: Openstack
octavia_certs_organization: Octavia
enable_redis: "yes"
```

Generate Octavia certificate:
`kolla-ansible octavia-certificates` : this command is used for auto generate Octavia certificates but if you want to use other certificates follow [this link](https://docs.openstack.org/kolla-ansible/latest/reference/networking/octavia.html).

## Set up multinode file
```
[control]
# These hostname must be resolvable from your deployment host
localhost      ansible_connection=local    ansible_sudo_pass='1'

# The above can also be specified as follows:
#control[01:03]     ansible_user=kolla

# The network nodes are where your l3-agent and loadbalancers will run
# This can be the same as a host in the control group
[network]
localhost      ansible_connection=local
[compute]
192.168.52.132    ansible_ssh_user=root  ansible_become=True  ansible_private_key_file=/home/buidung/.ssh/id_rsa
localhost      ansible_connection=local
[monitoring]
localhost      ansible_connection=local

# When compute nodes and control nodes use different interfaces,
# you need to comment out "api_interface" and other interfaces from the globals.yml
# and specify like below:
#compute01 neutron_external_interface=eth0 api_interface=em1 tunnel_interface=em1

[storage]
192.168.52.132    ansible_ssh_user=root  ansible_become=True  ansible_private_key_file=/home/buidung/.ssh/id_rsa
localhost      ansible_connection=local
[deployment]
localhost      ansible_connection=local

```
After set up multinode file, we have to create key-pair for running ansible. 
## Deployment 
1. Bootstrap servers with kolla deploy dependencies:
     ```
     kolla-ansible -i ./multinode bootstrap-servers 
     ```
2. Do pre-deployment checks for hosts:
     ```
     kolla-ansible -i ./multinode prechecks
     ``` 
3. Finally proceed to actual OpenStack deployment:
     ```
     kolla-ansible -i ./multinode deploy
     ``` 

## Install other package:
     
     pip install python-openstackclient -c https://releases.openstack.org/constraints/upper/master
    
     kolla-ansible post-deploy

## Run bash-scripts for testing deployment: 
`source ./openstack/share/kolla-ansibe/init-runonce` : this script for install cirros image and create a network, flavor,... 




## Some error occurred when using this guide: 

1. Do not using host-only network interface when setup by VMware -> port-binding error caused  (host-only is private network).
## References: 
[Openstack install kolla-ansible](https://docs.openstack.org/kolla-ansible/latest/user/quickstart.html)

[Octavia](https://docs.openstack.org/kolla-ansible/latest/reference/networking/octavia.html)