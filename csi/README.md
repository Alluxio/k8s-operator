## Alluxio CSI Driver

To see how to run Alluxio CSI Driver on kubernetes, please check out:
https://docs.alluxio.io/os/user/stable/en/kubernetes/Running-Alluxio-On-Kubernetes.html#csi

To build Alluxio CSI, enter the root directory of `Alluxio/k8s-operator` repository, and run
```console
$ docker build -t <image-name> -f ./build/csi/Dockerfile .
```
