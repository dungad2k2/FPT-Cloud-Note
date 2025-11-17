# This is some note about container challenge in https://labs.iximiuz.com/ 

1. Execute Host Commands inside a running container:
   - In this challenge, we must query the endpoint that bound to a specific port inside the container but do not exec into this container.
   - Using `nsenter -t <PID_of_container> -n bash` to access network namspace of container and query this endpoint
2. Run a sidecar container in the namespace of another container:
   - In this challenge, we'll need to start a new container in the namespaces of another. The sidecar container should share the PID, IPC, network namespaces of the target container
   - Using 'docker run -d -i --name sidecar --network container:<name-of-the-target-container> --pid container:<name-of-the-target-container> --ipc container:<name-of-the-target-container> busybox' to run a sidecar container that share namespaces with target container.
3. 