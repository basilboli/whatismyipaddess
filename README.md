<<<<<<< HEAD
=======
# Description

Stupid little server returning your real ip address using X-Real-IP header provided by proxy

# Note

THIS IS NOT PRODUCTION-READY, USE IT WITH CAUTION.

# Build 

`make build`

# Deploy with nginx 

Use sample nginx configuration for your proxy.

# Deploy with kubernetes

`make deploy`

Oh damn, it's should be an overkill for most cases to deploy it with kubernetes but maybe useful in some cases like testing that your outgoing traffic always resolves to a single ip address (if you master GA exclude internal traffic magic that's for you).

# ISSUES 

I have not succeeded to get it working with kubernetes. Apparently it's possible using iptables-based kube-proxy (see http://stackoverflow.com/questions/32112922/how-to-read-client-ip-addresses-from-http-requests-behind-kubernetes-services)
>>>>>>> 23d57947ff8cb7f41f4dadd2a43002210a623dbd
