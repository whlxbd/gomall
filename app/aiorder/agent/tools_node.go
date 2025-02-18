package agent

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/cloudwego/eino-ext/components/tool/duckduckgo"
	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/schema"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/whlxbd/gomall/app/aiorder/biz/dal/mysql"

	"github.com/whlxbd/gomall/app/aiorder/infra/rpc"
	cart "github.com/whlxbd/gomall/rpc_gen/kitex_gen/cart"
	order "github.com/whlxbd/gomall/rpc_gen/kitex_gen/order"
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
	Currency string `json:"user_currency"`
	Address  struct {
		StreetAddress string `json:"street_address"`
		State         string `json:"state"`
		City          string `json:"city"`
		Country       string `json:"country"`
		ZipCode       string `json:"zipCode"`
	} `json:"address"`
	Email       string `json:"email"`
	Order_items []struct {
		Item struct {
			ProductId uint32 `json:"productId"`
			Quantity  int32  `json:"quantity"`
		} `json:"item"`
		Cost float32 `json:"cost"`
	} `json:"order_items"`
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
		Desc: "模拟下单，可模拟用户下单操作。参数: user_id(int)用户ID, user_currency(string)用户货币, address(string)用户地址, email(string)用户邮箱, order_items(array)订单商品",
		ParamsOneOf: schema.NewParamsOneOfByParams(map[string]*schema.ParameterInfo{
			"user_id": {
				Type:     "number",
				Desc:     "用户ID",
				Required: true,
			},
			"user_currency": {
				Type:     "string",
				Desc:     "用户货币",
				Required: true,
			},
			"address": {
				Type:     "object",
				Desc:     "用户地址",
				Required: true,
				SubParams: map[string]*schema.ParameterInfo{
					"street_address": {
						Type:     "string",
						Desc:     "街道地址",
						Required: true,
					},
					"state": {
						Type:     "string",
						Desc:     "省",
						Required: true,
					},
					"city": {
						Type:     "string",
						Desc:     "城市",
						Required: true,
					},
					"country": {
						Type:     "string",
						Desc:     "国家",
						Required: true,
					},
					"zipCode": {
						Type:     "string",
						Desc:     "邮编",
						Required: false,
					},
				},
			},
			"email": {
				Type:     "string",
				Desc:     "用户邮箱",
				Required: true,
			},
			"order_items": {
				Type:     "array",
				Desc:     "订单商品，包含每种商品的价格以及商品相关信息",
				Required: true,
				ElemInfo: &schema.ParameterInfo{
					Type: "object",
					SubParams: map[string]*schema.ParameterInfo{
						"item": {
							Type:     "object",
							Desc:     "商品，包含商品ID以及商品数量",
							Required: true,
							SubParams: map[string]*schema.ParameterInfo{
								"product_id": {
									Type:     "number",
									Desc:     "商品ID",
									Required: true,
								},
								"quantity": {
									Type:     "integer",
									Desc:     "该商品数量",
									Required: true,
								},
							},
						},
						"cost": {
							Type:     "number",
							Desc:     "购买的这种商品总价",
							Required: true,
						},
					},
				},
			},
		}),
	}, nil
}

func (impl *SIMOrderToolImpl) InvokableRun(ctx context.Context, argumentsInJSON string, opts ...tool.Option) (string, error) {
	fmt.Println(argumentsInJSON)
	orderReq := &order.PlaceOrderReq{}
	if err := json.Unmarshal([]byte(argumentsInJSON), orderReq); err != nil {
		return "", err
	}

	fmt.Printf("模拟下单: %+v\n", orderReq)

	add_cartReq := &cart.AddItemReq{}
	add_cartReq.UserId = orderReq.UserId
	for _, item := range orderReq.OrderItems {
		add_cartReq.Item = item.Item
		if _, err := rpc.CartClient.AddItem(ctx, add_cartReq); err != nil {
			fmt.Println(err)
			return "", err
		}
	}

	// get_cartReq := &cart.GetCartReq{UserId: orderReq.UserId}
	// get_cartResp, err := rpc.CartClient.GetCart(ctx, get_cartReq)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return "", err
	// }

	// for _, item := range get_cartResp.Cart.Items {
	// 	orderReq.OrderItems = append(orderReq.OrderItems, &order.OrderItem{
	// 		Item: &cart.CartItem{
	// 			ProductId: item.ProductId,
	// 			Quantity:  item.Quantity,
	// 		},
	// 	})
	// }

	orderResp, err := rpc.OrderClient.PlaceOrder(ctx, orderReq)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	// _, err := rpc.OrderClient.PlaceOrder(ctx, req)
	// if err != nil {
	// 	klog.Errorf("PlaceOrder failed: %v", err)
	// 	return "", err
	// }

	return fmt.Sprintf("成功，订单号为%s", orderResp.Order.OrderId), nil
}
