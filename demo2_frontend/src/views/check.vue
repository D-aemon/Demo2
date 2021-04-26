<template>
  <div class="container">
    <div>{{article.title}}</div>
    <div>{{article.tag}}</div>
              <mavon-editor
          :value="context"
          :subfield="prop.subfield"
          :defaultOpen="prop.defaultOpen"
          :toolbarsFlag="prop.toolbarsFlag"
          :editable="prop.editable"
          :scrollStyle="prop.scrollStyle"
        ></mavon-editor>
  </div>
</template>

<script>
export default {
  name: 'check',
  data(){
      return{
          article:{},
          context:""
      }
  },
  computed: {
      title() {
          return this.$route.params.title
      },
      prop() {
      let data = {
        subfield: false, // 单双栏模式
        defaultOpen: "preview", //edit： 默认展示编辑区域 ， preview： 默认展示预览区域
        editable: false,
        toolbarsFlag: false,
        scrollStyle: true,
      };
      return data;
    }
  },
  mounted() {
    this.init()  
  },
  methods: {
      init() {
          this.getArticle(this.title)
      },
      getArticle(title) {
          this.$axios({
              url: "http://localhost:8082/api/article/getArticleByTitle",
              method: "post",
              data: {
                  title: title
              }
          }).then(res => {
              this.article = res.data.data.article
              this.context = this.article.body
          })
      }
  }
}
</script>