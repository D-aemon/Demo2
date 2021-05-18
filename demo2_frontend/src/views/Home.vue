<template>
  <div class="container">
    <section id="services" class="dark">
      <h1 class="section-title">全部文章</h1>
      <button @click="edit()">edit</button>
      <div class="tags">
        <div
          class="tag"
          v-for="(item, index) in tags"
          :key="index"
          @click="selectArtical(item.id)"
        >
          <span>{{ item.TagName }}</span>
        </div>
      </div>
      <div class="services">
        <div
          class="service"
          v-for="(item, index) in articles"
          :key="index"
        >
          <i class="el-icon-close" @click="delArticle(item, index+1)"></i>
          <h4 @click="check(item)">{{ item.title }}</h4>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
export default {
  name: "Home",
  data() {
    return {
      articles: [],
      alltags: {},
      relationT: {},
      relationA: {},
      tagRes:[]
    };
  },
  computed: {
    tags() {
        for (var i = 0; i < this.tagRes.length; i++) {
          this.alltags[this.tagRes[i].id] = this.tagRes[i]
        }
      return this.alltags
    }
  },
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.getAllArticles();
      this.getAllTags();
      this.getRelationT();
      this.getRelationA();
    },
    getAllArticles() {
      this.$axios({
        url: "http://localhost:8082/api/article/getAllArticles",
        method: "post",
      }).then((res) => {
        this.articles = res.data.data.articles;
      });
    },
    getAllTags() {
      this.$axios({
        url: "http://localhost:8082/api/article/getAllTags",
        method: "post",
      }).then((res) => {
        this.alltags[0] = {"id": 0, "TagName": "All"}
        this.tagRes = res.data.data.tagNames
      });
    },
    getRelationT() {
      this.$axios({
        url: "http://localhost:8082/api/article/getRelationT",
        method: "post",
      }).then((res) => {
        this.relationT = JSON.parse(res.data.data.res);
      });
    },
    getRelationA() {
      this.$axios({
        url: "http://localhost:8082/api/article/getRelationA",
        method: "post",
      }).then((res) => {
        this.relationA = JSON.parse(res.data.data.res);
      });
    },
    selectArtical(index) {
      if (index == 0) {
        this.getAllArticles();
      } else {
        this.$axios({
          method: "post",
          url: "http://localhost:8082/api/article/getArticlesByRelation",
          data: {
            relation: this.relationT["" + index + ""],
          },
        }).then((res) => {
          this.articles = res.data.data.articles;
        });
      }
    },
    check(item) {
      let arr = []
      let brr = this.relationA[""+ item.id +""]
      for (var i = 0; i < brr.length; i++) {
        arr.push(this.tags[""+brr[i]+""])
      }
      this.$store.commit('setTag',{
        relation: arr
      })
      this.$router.push("/check/" + item.title);
    },
    edit() {
      this.$router.push("/edit");
    },
    delArticle(item, index) {
      let that = this
      console.log(this.relationA["" + index + ""])
      this.$confirm("是否删除该文章?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.$axios({
            url: "http://localhost:8082/api/article/delArticle",
            method: "post",
            data: {
              articleId: item.id,
              relation: this.relationA["" + index + ""],
            },
          }).then((res) => {
            console.log(res.data);
            that.$message({
            type: "success",
            message: "删除成功!",
          });
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消删除",
          });
        });
    },
  },
};
</script>

<style scoped>
.container {
  background-color: #353b48;
}

.services {
  display: flex;
  justify-content: space-between;
  flex-wrap: wrap;
}

.service {
  width: calc(33% - 20px);
  text-align: center;
  border: 1px solid #48dbfb;
  border-radius: 6px;
  margin: 20px 0;
  padding: 40px 20px;
  color: #fff;
  cursor: pointer;
  transition: 0.3s linear;
  position: relative;
}

.service h4 {
  font-size: 16px;
  margin-bottom: 6px;
}

.service:hover {
  background-color: #48dbfb;
}

.section-title {
  color: #f1f1f1;
}

.tags {
  display: flex;
  justify-content: flex-start;
}

.tag {
  padding: 5px;
  line-height: 25px;
  font-size: 18px;
  color: #fff;
  border: 2px solid #48dbfb;
  border-radius: 8px;
  margin: 0 10px;
  cursor: pointer;
}

.tag:hover {
  background-color: #48dbfb;
}

i.el-icon-close {
  position: absolute;
  top: 10px;
  right: 10px;
  font-size: 20px;
}
</style>