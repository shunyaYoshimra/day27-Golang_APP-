<template>
  <div id="checked-articles">
    <router-link v-for="(article, id) in articles" :key="id" :to="'/article/' + article.id">
      <div class="card-panel">
        <span>{{article.title}}</span>
        <br />
        <small>{{article.created_at | moment("MMMM Do YYYY")}}</small>
      </div>
    </router-link>
  </div>
</template>

<script>
import axios from "axios";
import moment from "moment";
export default {
  data() {
    return {
      checkedNumber: [],
      articles: []
    };
  },
  filters: {
    moment(value, format) {
      return moment(value).format(format);
    }
  },
  mounted() {
    axios.get(`/api/user_checks/${this.$route.params.id}`).then(res => {
      for (let i = 0; i < res.data.length; i++) {
        this.checkedNumber.push(res.data[i].article_id);
      }
      for (let i = 0; i < this.checkedNumber.length; i++) {
        axios
          .get(`/api/article/${this.checkedNumber[i]}`)
          .then(res => {
            this.articles.push(res.data.article);
          })
          .catch(err => {
            if (err.response.status === 404) {
              console.log("deleted article");
            }
          });
      }
    });
  }
};
</script>

<style lang="scss" scoped>
#checked-articles {
  .card-panel {
    span,
    small {
      color: #333;
    }
  }
  .card-panel:hover {
    border: #999 1px solid;
  }
}
</style>