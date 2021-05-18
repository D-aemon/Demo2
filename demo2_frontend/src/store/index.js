import Vue from 'vue'
import Vuex from 'vuex'
import axios from 'axios'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    relationA: [],
    article:{
      "id": 0,
      "title":"",
      "body":"",
      "describe":"",
      "writeTime":"",
      "readingTime":"",
      "backgroundImage":""
    }
  },
  //同步操作
  mutations: {
    //方法
    setArticle(state, payload) {
      state.article.id = payload.id
      state.article.title = payload.title
      state.article.body = payload.body
      state.article.describe = payload.describe
      state.article.readingTime = payload.readingTime
      state.article.writeTime = payload.writeTime
      state.article.backgroundImage = payload.backgroundImage
    },
    setTag(state, payload) {
      state.relationA = payload.relation
    }
  },
  //异步操作
  actions: {
    updateArticle(context, payload) {
      axios({
        method:"post",
        url:"http://localhost:8082/api/article/getArticleByTitle",
        data:{
          title: payload.title
        }
      }).then(res => {
        context.commit("setArticle", res.data.data.article)
        payload.success()
      }).catch(err => {
        console.log(err)
      })
    }
  },
  modules: {
  }
})
