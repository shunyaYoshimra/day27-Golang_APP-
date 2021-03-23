<template>
  <div id="articles">
    <div v-for="(article, id) in articles" :key="id">
      <router-link :to="'/article/' + article.id">
        <div class="card-panel">
          <span class="hide">{{name = userNameOfArticle(article.user_id)}}</span>
          <span class="hide">{{kind = kindOfArticle(article.user_id)}}</span>
          <span class="hide">{{description = descriptionOfArticle(article.user_id)}}</span>
          <div class="top-wrapper">
            <span class="title">{{article.title}}</span>
            <br />
            <small>{{article.created_at | moment("MMMM Do YYYY")}}</small>
          </div>
          <hr />
          <span>{{name}}</span>
          <br />
          <template v-if="kind == null && description == null">
            <small>職業情報はまだ設定されていません</small>
          </template>
          <template v-else>
            <div>
              <small>{{kind}}</small>
              <br />
              <small>{{description}}</small>
            </div>
          </template>
          <!-- <template v-if="occupation !== null">
            <small>{{occupation.kind}}</small>
            <br />
            <small>{{occupation.description}}</small>
          </template>
          <template v-else>
            <small>職業情報はまだ設定されていません</small>
          </template>-->
        </div>
      </router-link>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment";
export default {
  data() {
    return {
      me: {},
      articles: [],
      users: [],
      occupations: []
    };
  },
  filters: {
    moment(value, format) {
      return moment(value).format(format);
    }
  },
  mounted() {
    //  get articles
    axios.get("/api/articles").then(res => {
      for (let i = 0; i < res.data.length; i++) {
        this.articles.push(res.data[i]);
      }
    });
    // get me
    axios.get("/api/get_me").then(res => {
      this.me = res.data.data;
    });
    //  get users
    axios.get("/api/users").then(res => {
      // console.log(res.data);
      for (let i = 0; i < res.data.length; i++) {
        this.users.push(res.data[i]);
      }
    });
    // get occupations
    axios.get("/api/occupations").then(res => {
      console.log(res.data);
      for (let i = 0; i < res.data.length; i++) {
        this.occupations.push(res.data[i]);
      }
    });
  },
  computed: {
    userNameOfArticle: function() {
      self = this;
      return function(id) {
        for (let i = 0; i < self.users.length; i++) {
          if (self.users[i].id === id) {
            return self.users[i].name;
          }
        }
      };
    },
    kindOfArticle: function() {
      self = this;
      return function(id) {
        for (let i = 0; i < self.occupations.length; i++) {
          if (self.occupations[i].user_id === id) {
            return self.occupations[i].kind;
          }
        }
      };
    },
    descriptionOfArticle: function() {
      self = this;
      return function(id) {
        for (let i = 0; i < self.occupations.length; i++) {
          if (self.occupations[i].user_id === id) {
            return self.occupations[i].description;
          }
        }
      };
    }
  }
};
</script>

<style lang="scss" scoped>
#articles {
  width: 600px;
  margin: auto;

  .card-panel {
    color: #333;
    .top-wrapper {
      margin: 20px 0;
      .title {
        font-size: 1.5rem;
        font-weight: bold;
      }
    }
    hr {
      color: #999;
      font-weight: lighter;
    }
  }
  .card-panel:hover {
    border: 1px #999 solid;
  }
}
@media (max-width: 480px) {
  #articles {
    width: 98%;
    margin: auto;
  }
}
</style>