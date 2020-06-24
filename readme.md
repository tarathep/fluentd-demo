# Fluentd DEMO

![alt text](https://raw.githubusercontent.com/tarathep/fluentd-demo/master/architecture.jpg)
1. build custom-fluentd
> sudo docker build -t custom-fluentd .
2. start mongodb ,zookeeper, kafka
3. run appzone1<br>
    3.1. cd fluentd : config fluentd<br>
    3.2. run fluend<br>
    > sudo docker container run --name fluentd-zone1 --rm -it -v $(pwd)/config.conf:/fluentd/etc/config.conf -v $(pwd)/log_NAS:/log -e FLUENTD_CONF=config.conf -p 24224:24224 fluentd-custom<br>
    
    3.3. Build applog<br>
    > sudo docker build -t applog-zone1 .<br>
    
    3.4. RUN applog<br>
    > sudo docker run -it --rm --name applog-zone1 -e HOST={fluentd-zone1} -v /home/devops/go/src/local/workspace/appzone1/fluentd/log_NAS:/log applog-zone1

4. run appzone2<br>
    4.1. cd fluentd : config fluentd<br>
    4.2. run fluend<br>
    > sudo docker container run --name fluentd-zone2 --rm -it -v $(pwd)/config.conf:/fluentd/etc/config.conf -e FLUENTD_CONF=config.conf -p 24224:24224 fluentd-custom<br>
    
    4.3. Build applog<br>
    > sudo docker build -t applog-zone2 .<br>
    
    4.4. RUN applog<br>
    > sudo docker run -it --rm --name applog-zone2 -e HOST={fluentd-zone2} -e PORT=24224 -v /home/devops/go/src/local/workspace/appzone1/fluentd/log_NAS:/log applog-zone2<br>

5. run monitor<br>
    5.1 BUILD<br>
    > sudo docker build -t monitor .<br>
    
    5.2 RUN<br>
    > sudo docker run -it --rm --name monitor -e BROKER=dev.tarathep.com -e TOPIC=log-messages monitor .<br>
    
