FROM ubuntu:latest
RUN apt-get update && apt-get install -y curl 
# 手动安装Node.js和npm
RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && apt-get install -y nodejs
RUN npm config set registry https://registry.npmmirror.com \
    && npm install -g pnpm
# 手动创建node_modules并安装依赖
RUN mkdir -p /app/node_modules && cd /app \
    && echo "console.log('OK')" > index.js \
    && PNPM_HOME=/root/.local/share/pnpm pnpm add lodash --dir /app
RUN apt-get install -y git

# 克隆项目并安装依赖
WORKDIR /app
RUN git clone https://gitclone.com/github.com/hfut-soft-ware/hfut-api.git

# 设置启动命令
WORKDIR /app/hfut-api
