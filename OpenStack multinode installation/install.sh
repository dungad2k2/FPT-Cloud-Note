apt update
apt install git python3-dev libffi-dev gcc libssl-dev -y
apt install python3-venv -y
mkdir openstack
python3 -m venv ./openstack
. ./openstack/bin/activate 
pip install -U pip
pip install ansible-core==2.16.10
pip install git+https://opendev.org/openstack/kolla-ansible@stable/2024.1
mkdir -p /etc/kolla
chown $USER:$USER /etc/kolla
cp -r ./openstack/share/kolla-ansible/etc_examples/kolla/* /etc/kolla
cp ./openstack/share/kolla-ansible/ansible/inventory/multinode .
kolla-ansible install-deps
kolla-genpwd