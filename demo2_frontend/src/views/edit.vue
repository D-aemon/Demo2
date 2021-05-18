<template>
  <div class="container">
    <el-button type="primary" @click="update()">退出编辑模式</el-button>
    <el-button type="success" @click="saveArticle()">保存</el-button>
    <div>
      <input type="text" v-model="title" placeholder="请输入标题" />
      title: {{ title }}
    </div>
    <div>
      <span>标签: </span>
      <el-tag
        :key="index"
        v-for="(tag, index) in tags"
        closable
        :disable-transitions="false"
        @close="handleClose(tag)"
      >
        {{ $store.state.article.id > 0 ? tag.TagName : tag }}
      </el-tag>
      <el-input
        class="input-new-tag"
        v-if="inputVisible"
        v-model="inputValue.TagName"
        ref="saveTagInput"
        size="small"
        @blur="handleInputConfirm"
      >
      </el-input>
      <el-button v-else class="button-new-tag" size="small" @click="showInput"
        >+ New Tag</el-button
      >
      <span>{{ tags }}</span>
    </div>
    <div>
      <input type="text" v-model="describe" placeholder="请输入描述" />
      describe {{ describe }}
    </div>
    <div>
      <input type="text" v-model="writeTime" placeholder="请输入书写时间" />
      writeTime {{ writeTime }}
    </div>
    <div>
      <input
        type="text"
        v-model="readingTime"
        placeholder="请输入预计阅读时间"
      />
      readingTime {{ readingTime }}
    </div>
    <div>
      <input
        type="text"
        v-model="backgroundImage"
        placeholder="请上传背景图片"
      />
      backgroundImage: {{ backgroundImage }}
    </div>
    <mavon-editor
      ref="md"
      @save="$save"
      v-model="context"
      :toolbars="toolbars"
      @imgAdd="$imgAdd"
    ></mavon-editor>
  </div>
</template>

<script>
export default {
  name: "test",
  data() {
    return {
      context: "", //输入的数据
      toolbars: {
        bold: true, // 粗体
        italic: true, // 斜体
        header: true, // 标题
        underline: true, // 下划线
        mark: true, // 标记
        superscript: true, // 上角标
        quote: true, // 引用
        ol: true, // 有序列表
        link: true, // 链接
        imagelink: true, // 图片链接
        help: true, // 帮助
        code: true, // code
        subfield: true, // 是否需要分栏
        fullscreen: true, // 全屏编辑
        readmodel: true, // 沉浸式阅读
        /* 1.3.5 */
        undo: true, // 上一步
        trash: true, // 清空
        save: true, // 保存（触发events中的save事件）
        /* 1.4.2 */
        navigation: true, // 导航目录
      },
      articleId: 0,
      title: "",
      describe: "",
      writeTime: "",
      readingTime: "",
      backgroundImage: "",
      tags: [],
      inputVisible: false,
      inputValue: {
        id: 0,
        TagName: "",
      },
    };
  },
  mounted() {
    this.init();
  },
  methods: {
    init() {
      if (this.$store.state.article.id > 0) {
        this.articleId = this.$store.state.article.id;
        this.title = this.$store.state.article.title;
        this.describe = this.$store.state.article.describe;
        this.writeTime = this.$store.state.article.writeTime;
        this.readingTime = this.$store.state.article.readingTime;
        this.backgroundImage = this.$store.state.article.backgroundImage;
        this.context = this.$store.state.article.body;
        this.tags = this.$store.state.relationA;
      }
    },
    $save(value, render) {
      if (this.title != null && this.tags != null && value != null) {
        this.$axios({
          url: "http://localhost:8082/api/article/uploadArticle",
          method: "POST",
          data: {
            title: this.title,
            tagName: this.tags,
            describe: this.describe,
            writeTime: this.writeTime,
            readingTime: this.readingTime,
            backgroundImage: this.backgroundImage,
            body: value,
          },
        }).then((res) => {
          if (res.data.code == 200 && res.data.msg == "创建成功") {
            alert("提交成功");
          } else {
            alert("提交失败");
          }
        });
      } else {
        alert("有空值");
      }
    },
    $imgAdd(pos, $file) {
      // 第一步.将图片上传到服务器.
      var formdata = new FormData();
      // 提取当前this。
      var that = this;
      formdata.append("file", $file);
      this.$axios({
        url: "http://8.136.3.67:8911/addPhoto",
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
        console.log(res.data.url);
        that.$refs.md.$img2Url(pos, res.data.url);
      });
    },
    handleClose(tag) {
      let that = this;
      this.$confirm("是否删除该标签?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          if (this.$store.state.article.id > 0) {
            this.$axios({
              method: "post",
              url: "http://localhost:8082/api/article/delTagName",
              data: {
                articleId: this.articleId,
                tagId: tag.id,
                tagName: tag.TagName,
              },
            }).then((res) => {
              console.log(res.data.code, res.data.message);
              if (res.data.code == 200 && res.data.msg == "删除成功") {
                that.$message({
                  type: "success",
                  message: "删除成功!",
                });
              } else {
                that.$message({
                  type: "success",
                  message: "删除失败!" + res.data.err,
                });
              }
            });
          }
          this.tags.splice(this.tags.indexOf(tag), 1);
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick((_) => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm() {
      if (this.inputValue.TagName) {
        let tagName = this.inputValue.TagName
        if (this.$store.state.article.id > 0) {
          this.$axios({
            method: "post",
            url: "http://localhost:8082/api/article/createTagName",
            data: {
              tagName: tagName,
              articleId: this.$store.state.article.id,
            },
          }).then(res => {
            this.inputValue.id = res.data.data.tagId
            this.inputValue.TagName = tagName
            this.tags.push(this.inputValue)
          })
        } else {
          this.tags.push(tagName);
        }
      }
      this.inputVisible = false;
      this.inputValue.TagName = "";
    },
    saveArticle() {
      this.$axios({
        method: "post",
        url: "http://localhost:8082/api/article/updateArticle",
        data: {
          articleId: this.articleId,
          title: this.title,
          describe: this.describe,
          writeTime: this.writeTime,
          readingTime: this.readingTime,
          backgroundImage: this.backgroundImage,
          body: this.context,
        },
      })
        .then((res) => {
          if (res.data.code == 200 && res.data.msg == "更新成功") {
            this.$message({
              type: "success",
              message: "更新成功!",
            });
          }
        })
        .catch((err) => {
          this.$message({
            type: "info",
            message: "更新失败!" + err,
          });
        });
    },
    update() {
      this.$router.push("/check/" + this.title)
    }
  },
};
</script>

<style scoped>
.container {
  padding: 20px 40px;
}
.el-tag + .el-tag {
  margin-left: 10px;
}
.button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding-top: 0;
  padding-bottom: 0;
}
.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}
</style>