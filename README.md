# VMManager

Consider a datacenter which have 2 clusters (one for each os type linux and windows and environment dev/qa). Each cluster manages the hosts machines, datastores and networks that are part of the clusters. Virtual machines are deployed to a cluster based on whether it is linux or windows VMs and datastores, networks are picked up based on available free space and VLAN address respectively. Hosts machines are picked up randomly to run the VMs on it.
We would like to pick a cluster,  datastores, hosts and networks and assign it to the following VMs according to the data associated with the VMs.

VM1 - {Name: vm5533dlv, OSType: linux, Ip: 5.4.1.12, Environment: dev}      
VM2 - {Name: vm7783dlv, OSType: linux, Ip: 5.4.9.4, Environment: dev}      
VM3 - {Name: vm2233dlv, OSType: linux, Ip: 5.4.18.7, Environment: qa}     
VM4 - {Name: vm2355dwv, OSType: windows, Ip: 2.3.8.9, Environment: qa}      
VM5 - {Name: vm2511dwv, OSType: windows, Ip: 2.3.7.2, Environment: qa}      
VM6 - {Name: vm2519dwv, OSType: windows, Ip: 2.3.0.6, Environment: qa}  
Please refer the supporting data provided in yml file (dc.yml) and use the same data as a source for your code. The data is provided in yaml representation for better understanding of the object hierarchy and relationship and need not be handled using yml.    


Write a program in go language, which takes the provided VMs data as input and assign the cluster, datastore, host and networks based on the given criteria. Print the VMs data at the end of the program that shows the cluster,datastore,host and networks assigned for the VM.
Cluster - should be picked up based on environment and OS type
Hosts   - should be picked up randomly. 
                Hint: Use some pseudo random generator function       
Datastore - should be picked up based on most available free space on a datastore. 
                  - skip the datastore which has maintenancemode flag is set to 'yes'
                   Hint: sort the datastores based on available free space in descending order. Use of any sorting  algorithm is desirable but not mandatory.
            
Networks - should be picked up based on IP address of the VM. Hint: compare IP address of the VM with VLAN address. Please ignore the 4th octet of IP address for comparison.

For example, For VM1, the program should print the following information.
VM1:
Name: vm5533dlv
OSType: linux
Ip: 5.4.1.12
Environment: dev
Cluster: cluster-dcw-lnx-dev
Host: host-dcw-lnx-dev-2
Datastore: ds-lnx-dev-3
Network: 153_dev_pg_1

For VM3, the program should print "No linux cluster found for VM3 in qa environment."
Similarly, if the program cannot find network for the give IP address, it should print "No VLAN network found for given IP 2.3.0.6 for VM6."

The programmer must demonstrate their ability to define data types using struct, define data structure (such as arrays, slices, maps, pointers as applicable. Use of nested maps and pointers in an array are preferable.), use control loops (such as for, if..else, switch). Also please handle the errors or runtime exceptions using recover and defer in functions. 
Good Luck!!!
