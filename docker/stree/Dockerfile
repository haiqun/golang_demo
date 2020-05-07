FROM ubuntu
# 压力测试 demo
RUN apt-get update && apt-get -y install stress
## stress --vm 1 --vm-bytes  500000M --verbose

ENTRYPOINT ["/usr/bin/stress"]