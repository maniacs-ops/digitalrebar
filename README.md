# Digital Rebar Deploy Scripts

> Maintained by RackN, Inc.

Items in this repo are used to install and configure [Digital Rebar](https://digitalrebar.githib.io) and associated workloads.

## Documentation

They are intended as a companion to the [Digital Rebar documentation](https://github.com/digitalrebar/doc).

See ansible group vars:

* group_vars/all.yml for values, defaults. ansible group vars:
* group_vars/mac.yml for a mac with docker tools ansible group vars:
* group_vars/vagrant.yml for a vagrant install

## Deploy Paths

The following deploy approaches are available:

* [Ansible Playbook(s)](install_ansible.md)  <==== STEP-BY-STEP INSTALL
* [Vagrant Install](install_vagrant.md) runs on your local system
* [Run In Packet.net](run_in_packet.sh) automatically runs on a hosted metal server (account needed)

> Note: All installs use the same Ansible playbooks.
