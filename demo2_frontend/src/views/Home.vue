<template>
  <div class="container">
      <section id="services" class="dark">
        <h1 class="section-title">全部文章</h1>
        <button @click="edit()">edit</button>
        <div class="services">
          <div class="service" v-for="(item, index) in articles" :key="index" @click="check(item.title)">
            <h4>{{item.title}}</h4>
            <p>{{item.tag}}</p>
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
    };
  },
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.getAllArticles();
    },
    getAllArticles() {
      this.$axios({
        url: "http://localhost:8082/api/article/findAll",
        method: "get"
      }).then((res) => {
        this.articles = res.data.data.articles;
      });
    },
    check(item) {
      this.$router.push('/check/' + item)
    },
    edit() {
      this.$router.push('/edit')
    }
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
</style>