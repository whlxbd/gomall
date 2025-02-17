package agent

import (
	"context"
	"encoding/json"
	"strings"
	"fmt"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
	"github.com/whlxbd/gomall/app/aiorder/biz/dal/mysql"
)

func NewDuckDuckGoTool(ctx context.Context) (bt tool.BaseTool, err error) {
	// TODO Modify component configuration here.
	config := &duckduckgo.Config{}
	bt, err = duckduckgo.NewTool(ctx, config)
	if err != nil {
		return nil, err
	}
	return bt, nil
}

type OrdreQueryToolImpl struct {
	config *OrdreQueryToolConfig
}

type OrdreQueryToolConfig struct {
	QuerySQL string `json:"query_sql"`
	UserId   uint32 `json:"user_id"`
	OrderId  string `json:"order_id"`
}

func NewOrdreQueryTool(ctx context.Context) (bt tool.BaseTool, err error) {
	// TODO Modify component configuration here.
	config := &OrdreQueryToolConfig{}
	bt = &OrdreQueryToolImpl{config: config}
	return bt, nil
}

func (impl *OrdreQueryToolImpl) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "QueryOrder",
		Desc: "查询订单信息，可查询订单详情、订单商品等信息，同时要确保查询的订单中用户id是。参数: order_id(string)订单ID, user_id(int)用户ID, query_sql(string)查询SQL",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"user_id": {
				Type:     "number",
				Desc:     "用户ID",
				Required: true,
			},
			"order_id": {
				Type:     "string",
				Desc:     "订单ID",
				Required: true,
			},
			"query_sql": {
				Type:     "string",
				Desc:     "查询SQL",
				Required: true,
			},
		}),
	}, nil
}

func (impl *OrdreQueryToolImpl) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	var err error

	p := &OrdreQueryToolConfig{}
	if err := json.Unmarshal([]byte(argumentsInJSON), p); err != nil {
		return "", err
	}

	klog.Infof("查询订单信息: %v\n", p)

// 处理SQL语句中的引号并分割多条SQL
queries := strings.Split(p.QuerySQL, ";")
var allResults []map[string]interface{}

for _, q := range queries {
	if strings.TrimSpace(q) == "" {
		continue
	}
	
	// 不替换引号，保持原SQL语句
	query := q
	klog.Infof("执行SQL: %s\n", query)

	var resp []map[string]interface{}
	// 使用 Session 避免全局配置影响
	if err := mysql.DB.Raw(query).Scan(&resp).Error; err != nil {
		klog.Warnf("执行SQL出错: %v", err)
		return "", fmt.Errorf("SQL执行错误: %v", err)
	}

	if len(resp) > 0 {
		allResults = append(allResults, resp...)
	}
}

if len(allResults) == 0 {
	return "", fmt.Errorf("未找到相关订单信息")
}

jsonResp, err := json.Marshal(allResults)
if err != nil {
	return "", err
}
return string(jsonResp), nil
}

type SIMOrderToolImpl struct {
	config *SIMOrderToolConfig
}

type SIMOrderToolConfig struct {
	UserId   uint32 `json:"user_id"`
	Products []struct {
		ProductId uint32 `json:"product_id"`
		Quantity  int    `json:"quantity"`
	} `json:"products"`
}

func NewSIMOrderTool(ctx context.Context) (bt tool.BaseTool, err error) {
	// TODO Modify component configuration here.
	config := &SIMOrderToolConfig{}
	bt = &SIMOrderToolImpl{config: config}
	return bt, nil
}

func (impl *SIMOrderToolImpl) Info(ctx context.Context) (*schema.ToolInfo, error) {
	return &schema.ToolInfo{
		Name: "SIMOrder",
		Desc: "模拟下单，可模拟用户下单操作。参数: user_id(int)用户ID, products(array)商品列表",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"user_id": {
				Type:     "number",
				Desc:     "用户ID",
				Required: true,
			},
			"products": {
				Type:     "array",
				Desc:     "商品列表",
				Required: true,
			},
		}),
	}, nil
}

func (impl *SIMOrderToolImpl) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	return "success", nil
}
