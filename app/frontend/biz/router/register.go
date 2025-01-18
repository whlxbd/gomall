// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	about "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/about"
	auth "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/auth"
	cart "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/cart"
	category "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/category"
	checkout "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/checkout"
	home "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/home"
	order "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/order"
	product "github.com/cloudwego/biz-demo/gomall/app/frontend/biz/router/product"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	about.Register(r)

	order.Register(r)

	checkout.Register(r)

	auth.Register(r)

	cart.Register(r)

	category.Register(r)

	product.Register(r)

	home.Register(r)
}
