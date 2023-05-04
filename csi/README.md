## Alluxio CSI Driver

AlluxioCSI implements Container Storage Interface(https://github.com/container-storage-interface/spec)
for Alluxio on Kubernetes. By using AlluxioCSI, AlluxioFuse process can be spawned on demand
instead of run as a long-running process.

To see how to run Alluxio CSI Driver, please check out:
https://docs.alluxio.io/os/user/stable/en/kubernetes/Running-Alluxio-On-Kubernetes.html#csi

To build Alluxio CSI, enter the root directory of `Alluxio/k8s-operator` repository, and run
```console
$ docker build -t <image-name> -f ./build/csi/Dockerfile .
```
