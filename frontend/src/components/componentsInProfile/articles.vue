<template>
  <div id="users-articles">
    <router-link v-for="(article, id) in articles" :key="id" :to="'/article/' + article.id">
      <div class="card-panel">
        <span>{{article.title | moment("MMMM Do YYYY")}}</span>
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
      articles: [],
    }
  },
  filters: {
    moment(value, format) {
     return moment(value).format(format); 
    }
  },
  mounted() {
    axios.get(`/api/articles/${this.$route.params.id}`).then((res) => {
      for (let i = 0; i < res.data.length; i ++) {
        this.articles.push(res.data[i]);
      }
    })
  },
}
</script>

<style lang="scss" scoped>
#users-articles {
  .card-panel {
    span {
      color: #333;
    }
  }
  .card-panel:hover {
    border: #999 1px solid;
  }
}
</style>