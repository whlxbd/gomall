# *** Project

## introduce

- Use the [Kitex](https://github.com/cloudwego/kitex/) framework
- Generating the base code for unit tests.
- Provides basic config functions
- Provides the most basic MVC code hierarchy.

## Directory structure

|  catalog   | introduce  |
|  ----  | ----  |
| conf  | Configuration files |
| main.go  | Startup file |
| handler.go  | Used for request processing return of response. |
| kitex_gen  | kitex generated code |
| biz/service  | The actual business logic. |
| biz/dal  | Logic for operating the storage layer |

## How to run

```shell
sh build.sh
sh output/bootstrap.sh
```

## 压测指南

### 安装 ghz

```bash
go install github.com/bojand/ghz/cmd/ghz@latest
```

### 运行压测

压测文件位于 `ghz_test` 目录下, 默认生成html, 可根据实际情况修改压测文件。

```bash
ghz --config ./ghz_test/GetProduct.json
```