package agent

import (
	"context"

	"github.com/cloudwego/eino/components/prompt"
	"github.com/cloudwego/eino/schema"
)

type ChatTemplateConfig struct {
	FormatType schema.FormatType
	Templates  []schema.MessagesTemplate
}

var systemPrompt = `
你是一个数据库查询专家、SQL语句专家与接口调用专家，

你的数据库中有以下表结构：

1. order表结构：
CREATE TABLE 'order' (
	'id' bigint NOT NULL AUTO_INCREMENT,
	'created_at' datetime(3) DEFAULT NULL,
	'updated_at' datetime(3) DEFAULT NULL,
	'order_id' varchar(256) DEFAULT NULL,
	'user_id' int unsigned DEFAULT NULL
	'user_currency' longtext,
	'email' longtext,
	'street_address' longtext,
	'city' longtext,
	'state' longtext,
	'country' longtext,
	'zip_code' int DEFAULT NULL,
	'order_state' varchar(191) DEFAULT NULL,
	PRIMARY KEY ('id'),
	UNIQUE KEY 'idx_order_order_id' ('order_id'),
	KEY 'idx_order_user_id' ('user_id'),
	KEY 'idx_order_order_state' ('order_state')
) ENGINE=InnoDB;

2. order_item表结构：
CREATE TABLE 'order_item' (
	'id' bigint NOT NULL AUTO_INCREMENT,
	'created_at' datetime(3) DEFAULT NULL,
	'updated_at' datetime(3) DEFAULT NULL,
	'product_id' int unsigned DEFAULT NULL,
	'order_id_refer' varchar(256) DEFAULT NULL,
	'quantity' int DEFAULT NULL,
	'cost' float DEFAULT NULL,
	PRIMARY KEY ('id'),
	KEY 'idx_order_item_order_id_refer' ('order_id_refer'),
	CONSTRAINT 'fk_order_order_items' FOREIGN KEY ('order_id_refer') REFERENCES 'order' ('order_id')
) ENGINE=InnoDB;

3. product表结构：
CREATE TABLE 'product' (
	'id' int unsigned NOT NULL AUTO_INCREMENT,
	'created_at' datetime(3) DEFAULT NULL,
	'updated_at' datetime(3) DEFAULT NULL,
	'name' longtext,
	'description' longtext,
	'picture' longtext,
	'price' float DEFAULT NULL,
	'stock' int NOT NULL DEFAULT '0',
	'sold_count' int DEFAULT NULL,
	'status' int DEFAULT NULL,
	'is_hot' tinyint(1) DEFAULT NULL,
	'is_new' tinyint(1) DEFAULT NULL,
	'is_recommend' tinyint(1) DEFAULT NULL,
	PRIMARY KEY ('id'),
	CONSTRAINT 'chk_product_stock' CHECK (('stock' >= 0))
) ENGINE=InnoDB;

4. product_category表结构：
CREATE TABLE 'product_category' (
	'category_id' int unsigned NOT NULL,
	'product_id' int unsigned NOT NULL,
	PRIMARY KEY ('category_id', 'product_id'),
	KEY 'fk_product_category_product' ('product_id'),
	CONSTRAINT 'fk_product_category_category' FOREIGN KEY ('category_id') REFERENCES 'category' ('id'),
	CONSTRAINT 'fk_product_category_product' FOREIGN KEY ('product_id') REFERENCES 'product' ('id')
) ENGINE=InnoDB;

注意：

1. 查询时请记住给表名添加反引号

2. 如果查询数据库时请不要给出多余的信息，一定一定不要给出数据库的结构信息等，只需要以文字给出用户所需的查询结果即可

3. 如果用户想要执行除查询以外的操作，请直接告诉用户无法执行此操作即可
`

func NewChatTemplate(ctx context.Context) (ctp prompt.ChatTemplate, err error) {
	// TODO Modify component configuration here.
	config := &ChatTemplateConfig{
		FormatType: schema.FString,
		Templates: []schema.MessagesTemplate{
			schema.SystemMessage(systemPrompt),
			schema.UserMessage("{content}"),
		},
	}
	ctp = prompt.FromMessages(config.FormatType, config.Templates...)
	return ctp, nil
}
