# Demo2 :blush:
基于vue+go-gin+mysql的markdown文本存储和图片上传小Demo
## describe(描述)
Demo2采用前后端分离的方式开发，借助[ mavonEditor](https://github.com/hinesboy/mavonEditor/blob/master/doc/cn/upload-images.md)只需关注后端所需接口的实现即可。通过保存接口，用户书写的markdown文本将被存入数据库中。通过图片上传接口，用户的图片将被存入服务器指定的位置并返回响应的url。此外demo2还实现了查找文章等相关接口。

Demo2对数据库、路由、中间件、返回请求、工具类、数据模型和配置文件均进行了不同程度的封装。简单修改即可开始搭建其他业务逻辑。

Demo2借鉴了[Pipe](https://github.com/88250/pipe)“自底向上”的分类方式：
> 1. 定义分类，并配置该分类包含的标签
> 2. 查询某个分类文章列表时通过分类-->标签集 --> 标签关联的文章进行聚合
>
> 也就是说一篇文章在编排时只需打标签，访问分类时会根据该分类包含的标签将文章关联出来。这是一个自底向上的信息架构，在使用时更灵活一些，可以随时调整分类而不必从新更新文章。


## USE（如何使用）

### backend部分
#### step1 
在拉取demo1后，需要修改位于 `config/application.yml` 的配置信息
```
server:
  port: 8080
datasource:
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: your database
  username: your username
  password: your password
  charset: utf8
```
#### step 2
修改位于`controller/UserCommit`下的UploadImg函数
```
func UploadImg(ctx *gin.Context) {
	file, _ := ctx.FormFile("image")
	name := util.RandomString(10)
	filename := util.GetMd5String(name) + ".png"
	if err2 := ctx.SaveUploadedFile(file, "你想存储的位置+/"+filename); err2 != nil {
		response.Fail(ctx, gin.H{"err": err2}, "请求失败")
		return
	}
	response.Success(ctx, gin.H{"url": "服务器地址+存储位置+/" + filename},"上传成功")
}

```

#### step3
启动！
```
go run main.go
```

### frontend部分
#### step1
安装依赖
```
npm install
```
#### step2
启动！
```
npm run serve
```


## API 文档

| 方法   | url | 请求参数 | 返回参数 | 
| -------- | :----------: | :---------:|---------------|
|Get     |  /api/article/findAll | 空 | code, data, msg |
|POST     |  /api/article/commit | title, tag, body | code, data, msg |
|POST     | /api/article/getArticleByTitle | title | code, data, msg |
|POST     |  /api/article/getArticleByTag | tag | code, data, msg |
|POST     |  /api/article/uploadImg | image(file) | code, data, msg |

### api详解
----
> **/api/article/findAll** --所有文章

成功示例
```
{
  "code": 200,
  "data": {
	"articles":[
		{"title":"title1","tag":"tag1","body":"# dada\n# dada\n- dada\n- dada\n    -dadad\n```\ndadada\n```\ndada\n"},
		{"title":"title2","tag":"tag2","body":"# title2\n## dada\n```\ndadada\n```"}]
	},
  "msg": "获取成功"
}
```
---
> **/api/article/commit** --提交文章

请求示例
```
{
	"title":"title1",
	"tag":"tag1",
	"body":"# dada\n# dada\n- dada\n- dada\n    -dadad\n```\ndadada\n```\ndada\n"
}
```

成功示例
```
{
  "code": 200,
  "data": "",
  "msg": "获取成功"
}
```
-----
> **/api/article/uploadImg** --上传图片

结合[ mavonEditor](https://github.com/hinesboy/mavonEditor/blob/master/doc/cn/upload-images.md)$imgAdd的请求示例
```
$imgAdd(pos, $file) {
      // 第一步.将图片上传到服务器.
      var formdata = new FormData();
      // 提取当前this。
      var that = this
      formdata.append("image", $file);
      this.$axios({
        url: "http://localhost:8082/api/article/uploadImg",
        method: "post",
        data: formdata,
        headers: { "Content-Type": "multipart/form-data" },
      }).then((res) => {
        // 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
        /**
         * $vm 指为mavonEditor实例，可以通过如下两种方式获取
         * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
         * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
         */
        //实际使用时发现直接$vm获取不到，所以将this提取。
        that.$refs.md.$img2Url(pos, res.data.data.url);
      })
```

成功示例
```
{
  "code": 200,
  "data": {
	"url":"服务器上的存储位置"
	},
  "msg": "获取成功"
}
```
-----
> **/api/article/getArticleByTitle** --通过标题获取文章

请求示例
```
{
	"title":"title1"
}
```

成功示例
```
{
  "code": 200,
  "data": {
	"article":{
		"title":"title1",
		"tag":"tag1",
		"body":"# dada\n# dada\n- dada\n- dada\n    -dadad\n```\ndadada\n```\ndada\n"
		}
	},
  "msg": "获取成功"
}
```
-----
> **/api/article/getArticleByTag** --通过标签获取文章

请求示例
```
{
	"tag":"tag1"
}
```

成功示例
```
{
  "code": 200,
  "data": {
	"article":[{
		"title":"title2",
		"tag":"tag2",
		"body":"# dada\n# dada\n- dada\n- dada\n    -dadad\n```\ndadada\n```\ndada\n"
		},
		"title":"title3",
		"tag":"tag2",
		"body":"# dada\n# dada\n- dada\n- dada\n    -dadad"
		}]
	},
  "msg": "获取成功"
}
```
-----
