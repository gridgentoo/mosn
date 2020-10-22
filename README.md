<p align="center">
<img src="https://raw.githubusercontent.com/mosn/community/master/icons/png/mosn-labeled-horizontal.png" width="350" title="MOSN Logo" alt="MOSN logo">
</p>

[![Build Status](https://travis-ci.com/mosn/mosn.svg?branch=master)](https://travis-ci.com/mosn/mosn)
[![codecov](https://codecov.io/gh/mosn/mosn/branch/master/graph/badge.svg)](https://codecov.io/gh/mosn/mosn)
[![Go Report Card](https://goreportcard.com/badge/github.com/mosn/mosn)](https://goreportcard.com/report/github.com/mosn/mosn)
![license](https://img.shields.io/badge/license-Apache--2.0-green.svg)

### Анализ исходного кода модуля MOSN XDS.

##### XDS используется с pilot-discovery для выполнения функции обнаружения служб.

** MOSN можно получить динамически через XDS API.
- Listener，
- Route，
- Cluster， 
- Endpoint,
- Secret,
- Конфигурация.

Файл конфигурации MOSN mosn_config.json в режиме XDS:

```json
{
  "dynamic_resources": {
    "lds_config": {
      "ads": {}
    },
    "cds_config": {
      "ads": {}
    },
    "ads_config": {
      "api_type": "GRPC",
      "cluster_names": ["xxx"],
      "grpc_services": [
        {
          "envoy_grpc": {
            "cluster_name": "xds-grpc"
          }
        }
      ]
    }
  },
  "static_resources": {
    "clusters": [
      {
        "name": "xds-grpc",
        "type": "STRICT_DNS",
        "lb_policy": "RANDOM",
        "hosts": [
          {
            "socket_address": {"address": "istio-pilot.istio-system.svc.boss.twl", "port_value": 15010}
          }
        ],
        "http2_protocol_options": { }
      }
    ]
  }
}
```

Проанализируйте файл конфигурации, чтобы создать XDSConfig (конфигурация клиента XDS).

Создайте adsClient (клиент XDS).


```go
func (c *Client) Start(config *config.MOSNConfig) error {
	log.DefaultLogger.Infof("xds client start")
    //Разобрать файл конфигурации
	dynamicResources, staticResources, err := UnmarshalResources(config)
	if err != nil {
		log.DefaultLogger.Warnf("fail to unmarshal xds resources, skip xds: %v", err)
		return errors.New("fail to unmarshal xds resources")
	}
    
    //Создайте xdsConfig
	xdsConfig := v2.XDSConfig{}
	err = xdsConfig.Init(dynamicResources, staticResources)
	if err != nil {
		log.DefaultLogger.Warnf("fail to init xds config, skip xds: %v", err)
		return errors.New("fail to init xds config")
	}
    //Создать adsCLient
	stopChan := make(chan int)
	sendControlChan := make(chan int)
	recvControlChan := make(chan int)
	adsClient := &v2.ADSClient{
		AdsConfig:         xdsConfig.ADSConfig,
		StreamClientMutex: sync.RWMutex{},
		StreamClient:      nil,
		MosnConfig:        config,
		SendControlChan:   sendControlChan,
		RecvControlChan:   recvControlChan,
		StopChan:          stopChan,
	}
	adsClient.Start()
	c.adsClient = adsClient
	return nil
}
```

## Инициализировать и запустить соединение xds
adsClient.start() 
```go
func (adsClient *ADSClient) Start() {
        // Постройте двустороннее соединение потока для grpc.
	adsClient.StreamClient = adsClient.AdsConfig.GetStreamClient()
	utils.GoWithRecover(func() {
        // Аутентифицировать и начать передачу xds, а также настроить регулярную повторную передачу
		adsClient.sendThread()
	}, nil)
	utils.GoWithRecover(func() {
        // Примите выданные данные и выберите разные обработчики в зависимости от типа
		adsClient.receiveThread()
	}, nil)
}
```

Детали функции:
[https://github.com/mosn/mosn/blob/master/pkg/xds/v2/adssubscriber.go](https://github.com/mosn/mosn/blob/master/pkg/xds/v2/adssubscriber.go)

## Обработка и отправка сообщений XDS

Четыре типа регистрации процессора.

```go
func init() {
	RegisterTypeURLHandleFunc(EnvoyListener, HandleEnvoyListener)
	RegisterTypeURLHandleFunc(EnvoyCluster, HandleEnvoyCluster)
	RegisterTypeURLHandleFunc(EnvoyClusterLoadAssignment, HandleEnvoyClusterLoadAssignment)
	RegisterTypeURLHandleFunc(EnvoyRouteConfiguration, HandleEnvoyRouteConfiguration)
}
```
Примите тип данных, преобразуйте тип XDS в тип данных MOSN и добавьте соответствующий менеджер.

В качестве примера возьмем HandlerListener:
```go
func HandleEnvoyListener(client *ADSClient, resp *envoy_api_v2.DiscoveryResponse) {
	log.DefaultLogger.Tracef("get lds resp,handle it")
    //Разберите возвращенное сообщение и сгенерируйте envoy_listener
	listeners := client.handleListenersResp(resp)
	log.DefaultLogger.Infof("get %d listeners from LDS", len(listeners))
    //Преобразуйте envoy_listener в mosn_listener и добавьте ListenerAdapter
	conv.ConvertAddOrUpdateListeners(listeners)
	if err := client.reqRoutes(client.StreamClient); err != nil {
		log.DefaultLogger.Warnf("send thread request rds fail!auto retry next period")
	}
}
```


MOSN is a network proxy written in Golang. It can be used as a cloud-native network data plane, providing services with the following proxy functions:  multi-protocol, modular, intelligent, and secure. MOSN is the short name of Modular Open Smart Network-proxy. MOSN can be integrated with any Service Mesh which support xDS API. It also can be used as an independent Layer 4 or Layer 7 load balancer, API Gateway, cloud-native Ingress, etc.

## Features

As an open source network proxy, MOSN has the following core functions:

+ Support full dynamic resource configuration through xDS API integrated with Service Mesh.
+ Support proxy with TCP, HTTP, and RPC protocols.
+ Support rich routing features.
+ Support reliable upstream management and load balancing capabilities.
+ Support network and protocol layer observability.
+ Support mTLS and protocols on TLS.
+ Support rich extension mechanism to provide highly customizable expansion capabilities.
+ Support process smooth upgrade.
  
## Download&Install

Use `go get -u mosn.io/mosn`, or you can git clone the repository to `$GOPATH/src/mosn.io/mosn`.

**Notice**

- If you need to use code before 0.8.1, you may needs to run the script `transfer_path.sh` to fix the import path.
- If you are in Linux, you should modify the `SED_CMD` in `transfer_path.sh`, see the comment in the script file.

## Documentation

- [Website](https://mosn.io)
- [Changelog](CHANGELOG.md)

## Contributing

See our [contributor guide](CONTRIBUTING.md).

## Partners

Partners participate in MOSN co-development to make MOSN better.

<div>
<table>
  <tbody>
  <tr></tr>
    <tr>
      <td align="center"  valign="middle">
        <a href="https://www.antfin.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/ant.png">
        </a>
      </td>
      <td align="center"  valign="middle">
        <a href="https://www.aliyun.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/aliyun.png">
        </a>
      </td>
      <td align="center" valign="middle">
        <a href="https://www.zhipin.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/bosszhipin.png">
        </a>
      </td>
      <td align="center" valign="middle">
        <a href="https://www.dmall.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/duodian.png">
        </a>
      </td>
      </tr><tr></tr>
      <tr>
      <td align="center" valign="middle">
        <a href="https://www.kanzhun.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/kanzhun.png">
        </a>
      </td>
    </tr>
    <tr></tr>
  </tbody>
</table>
</div>

## End Users

The MOSN users. Please [leave a comment here](https://github.com/mosn/community/issues/8) to tell us your scenario to make MOSN better!

<div>
<table>
  <tbody>
  <tr></tr>
    <tr>
      <td align="center"  valign="middle">
        <a href="https://www.tenxcloud.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/tenxcloud.png">
        </a>
      </td>
      <td align="center" valign="middle">
        <a href="https://www.zhipin.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/linkedcare.png">
        </a>
      </td>
      <td align="center" valign="middle">
        <a href="https://www.xiaobaoonline.com/" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/xiaobao.png">
        </a>
      </td>
      <td align="center" valign="middle">
        <a href="https://www.wm-motor.com/" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/weima.png">
        </a>
      </td>
    </tr>
    <tr></tr>
    <tr>
      <td align="center" valign="middle">
        <a href="https://www.iqiyi.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/iqiyi.png">
        </a>
      </td>
      <td align="center" valign="middle">
        <a href="https://www.gaiaworks.cn" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/gaiya.png">
        </a>
      </td>
      <td align="center" valign="middle">
        <a href="https://www.tydic.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/tianyuandike.png">
        </a>
      </td>
      <td align="center" valign="middle">
        <a href="https://www.terminus.io" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/terminus.png">
        </a>
      </td>
    </tr>
    <tr>
      <td align="center" valign="middle">
        <a href="https://www.tuya.com" target="_blank">
          <img width="222px"  src="https://mosn.io/images/community/tuya.png">
        </a>
      </td>
    </tr>
  </tbody>
</table>
</div>

## Community

See our community materials on <https://github.com/mosn/community>.

Visit the [MOSN website](https://mosn.io/docs/community/) for more information on working groups, roadmap, community meetings, MOSN tutorials, and more.

Scan the QR code below with [DingTalk(钉钉)](https://www.dingtalk.com) to join the MOSN user group.

<p align="center">
<img src="https://gw.alipayobjects.com/mdn/rms_91f3e6/afts/img/A*NyEzRp3Xq28AAAAAAAAAAABkARQnAQ" width="150">
</p>

## Community meeting

MOSN community holds regular meetings.

- [Wednesday 8:00 PM CST(Beijing)](https://ebay.zoom.com.cn/j/96285622161) every other week
- [Meeting notes](https://docs.google.com/document/d/12lgyCW-GmlErr_ihvAO7tMmRe87i70bv2xqe4h2LUz4/edit?usp=sharing)

## Landscapes

<p align="center">
<img src="https://landscape.cncf.io/images/left-logo.svg" width="150"/>&nbsp;&nbsp;<img src="https://landscape.cncf.io/images/right-logo.svg" width="200"/>
<br/><br/>
MOSN enriches the <a href="https://landscape.cncf.io/landscape=observability-and-analysis&license=apache-license-2-0">CNCF CLOUD NATIVE Landscape.</a>
</p>

