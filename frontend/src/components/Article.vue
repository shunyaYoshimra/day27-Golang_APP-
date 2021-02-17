<template>
  <div id="article">
    <div class="wrapper">
      <div class="article-top">
        <div>
          <router-link class="link-to-profile" :to="'/profile/' + user.id">{{user.name}}</router-link>
          <template v-if="checkedUserIds.includes(me.id)">
            <span @click="deleteCheck()" class="check green-text lighten-2">
              <i class="material-icons">check</i>checked!
            </span>
          </template>
          <template v-else>
            <span @click="checkArticle()" class="check grey-text lighten-2">
              <i class="material-icons">check</i>
            </span>
          </template>
          <br>
          <small>{{article.created_at | moment("MMMM Do YYYY")}}</small>
        </div>
        <template v-if="me.id === article.user_id">
          <div>
            <template v-if="editBool">
              <button @click="editArticle()" class="btn pink lighten-3">編集終了(Finish Editing)</button>
            </template>
            <template v-else>
              <button @click="editBool = !editBool" class="btn">編集(EDIT)</button>
            </template>
            <button @click="deleteArticle()" class="btn grey">削除(DELETE)</button>
          </div>
        </template>
      </div>
      <div class="title-wrapper">
        <template v-if="editBool">
          <input class="title edit-title" type="text" v-model="article.title">
        </template>
        <template v-else>
          <h1 class="title">{{article.title}}</h1>
        </template>
      </div>
      <div class="lines" v-for="(line, id) in lines" :key="id">
        <template v-if="link(id)">
          <template v-if="editBool">
            <textarea v-model="lines[id]"></textarea>
          </template>
          <template v-else>
            <a :href="line.line">{{line}}</a>
          </template>
        </template>
        <template v-else>
          <template v-if="editBool">
            <textarea v-model="lines[id]"></textarea>
          </template>
          <template v-else>
            <p :style="bold(id)">{{line}}</p>
          </template>
        </template>
      </div>
    </div>
   
    <ul class="pagination">
      <li class="disabled"><a @click="goBack()"><i class="material-icons">chevron_left</i></a></li>
      <li v-for="n in pageIDs" :key="n" class="waves-effect" :class="thisPage(n)"><router-link :to="'/article/' + n">{{pageIndex(n)}}</router-link></li>
      <li class="waves-effect"><a @click="goAhead()"><i class="material-icons">chevron_right</i></a></li>
    </ul>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment";
export default {
  data() {
    return {
      user: {},
      me: {},
      articleIds: [],
      article: {},
      lines: [],
      boldNumbers: [],
      linkNumbers: [],
      checkedUserIds: [],
      editBool: false,  
    }
  },
  filters: {
    moment(value, format) {
     return moment(value).format(format); 
    }
  },
  watch: {
    $route(to, from) {
      this.initData()
      this.fetchItems(to.params.id)
    }
  },
  computed: {
    bold: function() {
      self = this;
      return function (id) {
        for (let i = 0; i < self.boldNumbers.length; i ++) {
          if (self.boldNumbers[i].number === (id + 1)) {
            return {
              "font-weight": "bold",
            }
          }
        }
      }
    },
    link: function() {
      self = this;
      return function (id) {
        for (let i = 0; i < self.linkNumbers.length; i ++) {
          if (self.linkNumbers[i].number === (id + 1)) {
            return true
          }
        }
      }
    },
    pageIDs: function() {
      self = this;
      if (self.articleIds.length < 6) {
        return self.articleIds
      } else {
        for (let i = 0; i < self.articleIds.length; i ++) {
          if (self.articleIds[i] === self.article.id) {
            if (i > 1 && (self.articleIds.length - i) > 2) {
              return [(i-1), i, (i+1), (i+2), (i+3)]
            } else if (i < 2){
              return [self.articleIds[0], self.articleIds[1], self.articleIds[2], self.articleIds[3], self.articleIds[4]]
            } else if ((self.articleIds.length - i) <= 2) {
              return [self.articleIds[(self.articleIds.length - 5)], self.articleIds[(self.articleIds.length - 4)], self.articleIds[(self.articleIds.length - 3)], self.articleIds[(self.articleIds.length - 2)], self.articleIds[(self.articleIds.length - 1)]]
            }
          }
        }
      }
    },
    thisPage: function() {
      self = this;
      return function(id) {
        if (id === self.article.id) {
          return "active"
        }
      }
    },
    pageIndex: function() {
      self = this;
      return function(id) {
        for (let i = 0; i < self.articleIds.length; i ++) {
          if (self.articleIds[i] === id) {
            return i
          }
        }
      }
    }
  },
  mounted() {
   axios.get(`/api/article/${this.$route.params.id}`).then((res) => {
      this.article = res.data.article;
      this.lines = [];
      console.log(this.lines)
      for (let i = 0; i < res.data.articleLines.length; i++) {
        this.lines.push(res.data.articleLines[i].line);
      }
      for (let i = 0; i < res.data.boldNumbers.length; i++) {
        this.boldNumbers.push(res.data.boldNumbers[i]);
      }
      for (let i = 0; i < res.data.linkNumbers.length; i++) {
        this.linkNumbers.push(res.data.linkNumbers[i]);
      }
      // get user
      axios.get(`/api/user/${this.article.user_id}`).then((res) => {
        this.user = res.data.data;
      })
      // get user's article IDs
      axios.get(`/api/articles/${this.article.user_id}`).then((res) => {
        for (let i = 0; i < res.data.length; i ++) {
          this.articleIds.unshift(res.data[i].id)
        }
      })
      // get check IDs
      axios.get(`/api/checks/${this.article.id}`).then((res) => {
        for (let i = 0; i < res.data.length; i ++) {
          this.checkedUserIds.push(res.data[i].user_id)
        }
      })
    })
    // get me
    axios.get("/api/get_me").then((res) => {
      this.me = res.data.data;
    })
  },
  methods: {
    fetchItems(id) {
      axios.get(`/api/article/${id}`).then((res) => {
      this.article = res.data.article;
      for (let i = 0; i < res.data.articleLines.length; i++) {
        this.lines.push(res.data.articleLines[i].line);
      }
      for (let i = 0; i < res.data.boldNumbers.length; i++) {
        this.boldNumbers.push(res.data.boldNumbers[i]);
      }
      for (let i = 0; i < res.data.linkNumbers.length; i++) {
        this.linkNumbers.push(res.data.linkNumbers[i]);
      }
      // get user
      axios.get(`/api/user/${this.article.user_id}`).then((res) => {
        this.user = res.data.data;
      })
      // get user's article IDs
      axios.get(`/api/articles/${this.article.user_id}`).then((res) => {
        for (let i = 0; i < res.data.length; i ++) {
          this.articleIds.unshift(res.data[i].id)
        }
      })
      // // get checked user's IDs
      axios.get(`/api/checks/${this.article.id}`).then((res) => {
        for (let i = 0; i < res.data.length; i ++) {
          this.checkedUserIds.push(res.data[i].user_id)
        }
      })
    })
  },
    initData() {
      this.articleIds = [];
      this.lines = [];
      this.boldNumbers = [];
      this.linkNumbers = [];
      this.checkedUserIds = [];
    },
    changePage(id) {
      this.$router.go(`/article/${id}`)
    },
    goBack() {
      if (this.article.id !== this.articleIds[0]) {
        for(let i = 0; i < this.articleIds.length; i ++) {
          if (this.articleIds[i] === this.article.id) {
            this.$router.push(`/article/${(this.articleIds[i - 1])}`)
          }
        }
      } 
    },
    goAhead() {
      if (this.article.id !== this.articleIds[this.articleIds.length - 1]) {
        for(let i = 0; i < this.articleIds.length; i ++) {
          if (this.articleIds[i] === this.article.id) {
            this.$router.push(`/article/${(this.articleIds[i + 1])}`)
          }
        }
      }
    },
    checkArticle() {
      axios.post(`/api/checks/${this.article.id}`).then((res) => {
        this.checkedUserIds.push(this.me.id)
      })
    },
    deleteCheck() {
      axios.delete(`/api/checks/${this.article.id}`).then((res) => {
        this.checkedUserIds = this.checkedUserIds.filter(n => n !== this.me.id)
      })
    },
    deleteArticle() {
      if(confirm("本当にこの記事を削除しますか?(Are you sure to delete this article?)")) {
         axios.delete(`/api/articles/${this.article.id}`).then((res) => {
          this.$router.push(`/profile/${this.me.id}`)
        })
      }
    },
    editArticle() {
      for (let i = 0; i < (this.lines.length - 1); i ++) {
        this.lines[i] += "&divide;"
      }
      const params = new URLSearchParams();
      console.log(this.article.title)
      params.append("title", this.article.title);
      params.append("lines", this.lines);
      axios.put(`/api/articles/${this.article.id}`, params).then((res) => {
        this.lines = [];
        for (let i = 0; i < res.data.data.length; i ++) {
          this.lines.push(res.data.data[i]);
        }
      }).catch((err) => {
        console.log(err.response);
      })
      this.editBool = false;
    }
  }
}
</script>

<style lang="scss" scoped>
#article {
  // common
  .edit-icon {
    cursor: pointer;
    color: #999;
  }
  .edit-icon:hover {
    color: #333;
  }
  // part
  .wrapper {
    width: 600px;
    margin: auto;
    margin-bottom: 100px;

    .article-top {
      display: flex;
      justify-content: space-between;
      align-items: center;
      .link-to-profile{
        color: #333;
      }
      .check {
        margin-left: 2rem;
        transform: translateY(5px);
        cursor: pointer;
      }
      .link-to-profile:hover {
        color: #999;
      }
    }

    .title-wrapper {
       margin-bottom: 60px;
       .edit-title {
         margin-top: 40px;
       }
      .title {
        font-size: 1.5rem;
        font-weight: bold;
      }
    }

    .lines {
      margin-bottom: 20px;
      textarea {
        border: none;
        height: 80px;
        resize: none;
      }
      textarea:focus {
        outline: none;
      }
    }
  }

  .pagination {
    position: fixed;
    bottom: 30px;
    width: 100%;
    z-index: 100;
    text-align: center;
  }
}
@media (max-width: 480px) {
  #article {
    .wrapper {
      width: 98%;
      margin: auto;
      margin-bottom: 120px;
    }
    .pagination {
      bottom: 50px;
    }
  }
}
</style>