<template>
  <div class="container">
    <mavon-editor
      ref=md
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
      title: "titlex1",
      tag: null,
    };
  },
  methods: {
    // getContext() {
    //   this.$axios({
    //     url: "http://localhost:8082/api/article/getArticle",
    //     method: "Get",
    //     data:{
    //       "title":this.title
    //     }
    //   }).then((result) => {
    //     this.context = result.data.data.articles.Value.Body;
    //   });
    // },
    $save(value, render) {
      if (this.title != null && this.tag != null && value != null) {
        this.$axios({
          url: "http://localhost:8082/api/article/commit",
          method: "POST",
          data: {
            title: this.title,
            tag: this.tag,
            body: value,
          },
        }).then((res) => {
          if (res.data.code == 200 && res.data.msg == "添加成功") {
            alert("提交成功");
          } else {
            alert("提交失败");
          }
        });
      } else {
        alert("有空值")
      }
    },
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
      });
    },
  },
};
</script>

<style scoped>
.container {
  padding: 20px 40px;
}
</style>