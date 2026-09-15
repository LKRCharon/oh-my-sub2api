<div align="center">
  <img src="assets/ohmy-logo.svg" alt="Oh My Sub2API" width="96" />
  <h1>Oh My Sub2API</h1>
  <p>基于 Sub2API 的非官方分支，维护 Codex 自动任务兼容修复、可拖动侧边栏和 Apple Blue 主题。</p>

[English](README.md) · [版本下载](https://github.com/LKRCharon/oh-my-sub2api/releases) · [修改记录](CHANGES.md) · [上游项目](https://github.com/Wei-Shaw/sub2api)
</div>

## 当前版本

`0.2.5-ohmy.1` 以 **Sub2API v0.2.5** 为基线，使用独立的源码仓库、发版流程和更新渠道。

| 功能 | 具体行为 |
| --- | --- |
| Codex 自动任务兼容 | 保留带名称、缺少 `call_id` 的独立工具输出，避免请求归一化丢弃定时任务指令。 |
| 侧边栏拖动 | 桌面端可拖动边缘调整到 180–400 px，并记住宽度；聚焦拖动条后可用方向键、Home、End 调整。 |
| 主题配色 | Apple Blue 主色，配套图表颜色，支持浅色和深色；Codex 图表系列使用青绿色。 |
| 独立更新 | 更新检查与安装脚本指向本分支，正确识别 `ohmy.2`、`ohmy.10` 等修订号。 |

已配置的站点名称和 Logo 优先于默认品牌。可执行文件继续使用 `sub2api` 名称，以兼容现有部署方式。

## Docker 部署

复制部署文件，在本地 `.env` 中设置凭据后启动：

```sh
git clone https://github.com/LKRCharon/oh-my-sub2api.git
cd oh-my-sub2api/deploy
cp .env.example .env
# 先编辑 .env，设置自己的密码；不要提交此文件。
docker compose -f docker-compose.local.yml up -d
```

镜像为 `ghcr.io/lkrcharon/oh-my-sub2api:0.2.5-ohmy.1`，支持 Linux amd64/arm64。继承的 Compose 文件使用 `latest`；需要固定版本时，替换为上面的版本标签。配置、备份和二进制安装见[部署说明](deploy/README.md)。

## 从源码构建

需要 **Go 1.27.0+**、**Node.js 20+** 和 **pnpm 9.15.9**。运行网关还需 PostgreSQL 与 Redis。

```sh
pnpm --dir frontend install --frozen-lockfile
make build
./backend/bin/server -version
```

`make build` 先编译前端，再将其嵌入 Go 二进制。部署配置参考 `deploy/.env.example` 或 `deploy/config.example.yaml`。也可以构建完整镜像：

```sh
docker build -t oh-my-sub2api:local .
```

版本标签触发发布流程，生成平台二进制、GHCR 多架构镜像、SHA-256 校验文件及对应源码包。构建定义见[发布工作流](.github/workflows/release.yml)、[.goreleaser.yaml](.goreleaser.yaml) 和 [Dockerfile](Dockerfile)。

## 后续维护

测试和同步上游的方法见 [CONTRIBUTING.md](CONTRIBUTING.md)。API 密钥、SMTP/IMAP 授权码、生产配置与运行数据库放在 Git 仓库之外。示例和回归测试使用合成数据。

## 许可证与来源

沿用上游 README 声明的 [LGPL-3.0-or-later](LICENSE)。[COPYING](COPYING) 包含 LGPLv3 所纳入的 GPLv3 条款。再分发时保留版权、许可证及声明；发布修改后的二进制或镜像时，提供对应版本的完整源码与构建脚本。

本项目来自 [Wei-Shaw/sub2api](https://github.com/Wei-Shaw/sub2api)，属于非官方分支，不代表上游或 API 提供方。归属见 [NOTICE](NOTICE)，上游免责声明保留于[原始中文 README](README_UPSTREAM_CN.md)；其中赞助及服务介绍属于上游。上游 README 的商业使用表述与标准许可证存在差异，本分支不声称取得额外的上游商业授权；API 提供方的服务条款仍然适用。
