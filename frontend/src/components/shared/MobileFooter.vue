<template>
  <footer>
    <div class="footer-icons">
      <div>
        <router-link to="/posts">
          <i :class="postsColor">local_airport</i>
        </router-link>
      </div>
      <div>
        <router-link to="/articles">
          <i :class="articlesColor">work</i>
        </router-link>
      </div>
      <div>
        <router-link to="/questions">
          <i :class="questionsColor">live_help</i>
        </router-link>
      </div>
      <div>
        <router-link :to="'/profile/' + me.id ">
          <i :class="profileColor">person</i>
        </router-link>
      </div>
    </div>
  </footer>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      me: {},
      postsColor: "material-icons grey-text",
      articlesColor: "material-icons grey-text",
      questionsColor: "material-icons grey-text",
      profileColor: "material-icons grey-text",
    }
  },
  mounted() {
    // get me
    axios.get("/api/get_me").then((res) => {
      this.me = res.data.data;
    })
  },
  watch: {
    '$route'(to, from) {
      this.changeIconColor()
    }
  },
  methods: {
    changeIconColor() {
      if (this.$route.path === "/posts") {
        this.postsColor = "material-icons black-text";
        this.articlesColor = "material-icons grey-text";
        this.questionsColor = "material-icons grey-text";
        this.profileColor = "material-icons grey-text";
      } else if (this.$route.path === "/questions") {
        this.questionsColor = "material-icons black-text";
        this.postsColor = "material-icons grey-text";
        this.articlesColor = "material-icons grey-text";
        this.profileColor = "material-icons grey-text";
      } else if (this.$route.path === "/articles") {
        this.articlesColor = "material-icons black-text";
        this.postsColor = "material-icons grey-text";
        this.questionsColor = "material-icons grey-text";
        this.profileColor = "material-icons grey-text";
      } else if (this.$route.path === `/profile/${this.me.id}`) {
        this.profileColor = "material-icons black-text";
        this.postsColor = "material-icons grey-text";
        this.articlesColor = "material-icons grey-text";
        this.questionsColor = "material-icons grey-text";
      } else {
        this.postsColor = "material-icons grey-text";
        this.articlesColor = "material-icons grey-text";
        this.questionsColor = "material-icons grey-text";
        this.profileColor = "material-icons grey-text";
      }
    }
  },
}
</script>

<style lang="scss" scoped>
footer {
  position: fixed;
  bottom: 0;
  width: 100%;
  z-index: 100;
  height: 50px;
  background-color: #fff;
  border-top: 1px #bdbdbd solid;
  .footer-icons {
    display: flex;
    div {
      width: 25%;
      height: 50px;
      text-align: center;
      padding-top: 10px;
    }
  }
}
</style>