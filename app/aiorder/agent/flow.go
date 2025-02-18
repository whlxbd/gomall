package agent

import (
	"context"

	"github.com/cloudwego/eino/components/tool"
	"github.com/cloudwego/eino/compose"
	"github.com/cloudwego/eino/flow/agent/react"
)

func newLambda(ctx context.Context) (lba *compose.Lambda, err error) {
	// TODO Modify component configuration here.
	config := &react.AgentConfig{
		MaxStep:            25,
		ToolReturnDirectly: map[string]struct{}{}}
	chatModelIns11, err := NewArkChatModel(ctx)
	if err != nil {
		return nil, err
	}
	config.Model = chatModelIns11
	// toolIns21, err := NewDuckDuckGoTool(ctx)
	// if err != nil {
	// 	return nil, err
	// }
	toolIns22, err := NewOrdreQueryTool(ctx)
	if err != nil {
		return nil, err
	}
	toolIns23, err := NewSIMOrderTool(ctx)
	if err != nil {
		return nil, err
	}
	config.ToolsConfig.Tools = []tool.BaseTool{ toolIns22, toolIns23}
	ins, err := react.NewAgent(ctx, config)
	if err != nil {
		return nil, err
	}
	lba, err = compose.AnyLambda(ins.Generate, ins.Stream, nil, nil)
	if err != nil {
		return nil, err
	}
	return lba, nil
}
