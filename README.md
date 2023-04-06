# test-udp-bridge

This test is using an external udp brigde to expose services in Skupper. 

You can find the udp bridge code in this repository: https://github.com/grs/skupper-udp-bridge

## environment setup

For this test, we will need to set up two skupper sites and create a link between them.

In order to do that, you can follow [this tutorial](https://skupper.io/start/index.html).


## steps

### 1. udp-server setup
Deploy the udp-server in one of the skupper sites:

```
kubectl apply -f deployment-server.yaml
```

Export the deployment as a service using the external udp bridge:
```
skupper expose deployment udp-server --port 8000 --protocol udp --bridge-image=quay.io/skupper/bridge-server
```


### 2. run udp-client

In the other cluster, run the udp-client: 

```
kubectl run udp-client --image=quay.io/skupper/udp-client --restart="OnFailure"
```

The client will send a message using the UDP protocol to the address `udp-server:8000`

### 3. check the logs

The udp-server will print in its log the message received from the client:

`udp-client logs`
```
$ kubectl logs -f udp-client

Sent message to server: Hello, server!
```

`udp-server logs`
```
 $ kubectl logs -f udp-server-7dc558794c-clxts                  
 Listening for UDP packets on :8000
 Received 14 bytes from 10.244.0.7:33016: Hello, server!
```