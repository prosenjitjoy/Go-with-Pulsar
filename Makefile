create-pulsar:
	podman run --name dev-pulsar -p 6650:6650 -p 8080:8080 -d apachepulsar/pulsar:3.1.1 bin/pulsar standalone

delete-pulsar:
	podman rm -f dev-pulsar
	podman volume prune

generate-image:
	podman build --tag=producer --target=producer .
	podman build --tag=consumer --target=consumer .

.PHONY: create-pulsar delete-pulsar generate-image