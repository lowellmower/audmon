## audmon (Audit Monitor)

### Description
audmon is a daemon design to run on linux machine and listen to IPC messages
related to the auditd daemon. These messages, or at least those of interest will
then be pushed back to a machine configured to recieve them and any other meta
data deemed important.

### Purpose
Mostly to scratch an itch and play with websockets and Go syscalls, but also
because my personal website is on a virtual machine with few to no firewalls and
is essentially a honey pot. Seeing what kind of shenanigans are happening such as
port scans, ssh attempts, etc. and where they are coming from will be neat.

### What This Project is Not
This project is not meant to be a tool used for production purposes or for the
purpose of gaining true insight into what is happening on your linux machine(s).
If you are in need of services like that, there are plenty of other free tools and
purchasable frameworks available. Use those.

### Project Goals
- [ ] Establish Client Server Websocket Communication
- [ ] Add logrus logging and setup
- [ ] Construct a daemon which can listen on NETLINK socket for IPC messages
- [ ] Push messages over the wire in JSON format to client machine
- [ ] Provide an interface for viewing messages (client side)
- [ ] Collect and graph metrics using something like prometheus

### Quick Links/Info
- [EC2 Dash](https://us-west-2.console.aws.amazon.com/ec2/v2/home?region=us-west-2#Instances:)
- [Digital Ocean Dash](https://cloud.digitalocean.com/droplets/177816291/graphs?i=6e05ec&period=hour)
- Client DNS  | ec2-18-237-120-181.us-west-2.compute.amazonaws.com
- Client IPv4 | 18.237.120.181
- Daemon DNS  | lowellmower.com
- Daemon IPv4 | 174.138.38.146