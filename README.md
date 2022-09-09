# go_ddns_namesilo
动态更新 https://www.namesilo.com/ 购买的域名绑定变化的 ddns

## 1、替换conf.yaml 自己具体配置 
    
    apikey：配置自己的apikey 可自己生成 https://www.namesilo.com/account/api-manager
    domain：用namesilo购买的域名
    ddns_host：需要动态绑定ddns 的域名前缀 如 home.domain.com ，应填写 "home"

## 2、运行 

```goland
go env -w GO111MODULE=on

go env -w GOPROXY=https://goproxy.cn,direct

//初始化环境
go mod tidy


```

## 或 docker启动

### 已发布到docker中央仓库

    

    docker tag 14358162d6a8 jinghan94/go_ddns_namesilo:1.0
    docker push jinghan94/go_ddns_namesilo:1.0

    docker tag 14358162d6a8 jinghan94/go_ddns_namesilo:latest
    docker push jinghan94/go_ddns_namesilo:latest
    


    docker run -d \
    --name ddns_namesilo \
    -e apikey=输入自己的 \
    -e domain=自己的域名 \
    -e ddns_host=动态绑定 ddns 的 host \
    jinghan94/go_ddns_namesilo

### 启动查看日志

     docker logs -f --tail=100 ddns_namesilo
