# Ansible playbook
An ansible playbook to install a simple k8s cluster with a master node and two worker nodes. The CNI plugin is pcn-k8s. Vagrant is used to provide the vms

## Requirements
 * Vagrant 2.2.5 or higher
 * Ansible 2.8.5 or higher

## Adding additional worker-nodes
If more worker-nodes are needed than:
* add the dns name or the IP address of the new node in inventory.ini under the "k8s_worker_nodes" section;
* add a new file in the "host_vars" folder named as the new node and inside write the ip address of the node. it is used during the k8s installation;

## Remote user 
The remote user which runs the ansible commands is defined in two places:
* ansible.cfg
* group_vars/all/vars.yml

## Variables
All the variables are stored in group_vars/all/vars.yml.

## Ansible authentication
The playbook is configured to login to the machines using ssh-keys. Change the ansible.cfg if the password authentication is needed.
Before playing the playbook make sure that you change the path to the ssh-key in "vagrant-ssh-key.yml". 