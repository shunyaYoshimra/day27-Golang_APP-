<template>
  <div id="users-checks">
    <div class="choose-which">
      <div class="checked-posts" @click="componentID = 1" :style="postBorder">
        <span :class="postColor">
          <i class="material-icons">check</i>Posts
        </span>
      </div>
      <div class="checked-articles" @click="componentID = 2" :style="articleBorder">
        <span :class="articleColor">
          <i class="material-icons">check</i>Articles
        </span>
      </div>
    </div>
    <hr class="bottom-line">
    <component :is="whichComponent"></component>
  </div>
</template>

<script>
import CheckedPosts from "./componentsInChecks/CheckedPosts.vue"
import CheckedArticles from "./componentsInChecks/CheckedArticles.vue"
export default {
  components: {
    "checked-posts": CheckedPosts,
    "checked-articles": CheckedArticles,
  },
  data() {
    return {
      componentID: 1,
    }
  },
  watch: {

  },
  computed: {
    whichComponent: function() {
      self = this;
      if (this.componentID === 1) {
        return "checked-posts"
      } else if (this.componentID === 2) {
        return "checked-articles"
      } 
    },
    postBorder: function() {
      self = this;
      if (this.componentID === 1) {
        return {
          "border-bottom": "#333 2px solid"
        }
      }
    },
    articleBorder: function() {
      self = this;
      if (this.componentID === 2) {
        return {
          "border-bottom": "#333 2px solid"
        }
      }
    },
    postColor: function() {
      self = this;
      if (this.componentID === 1) {
        return "green-text lighten-2"
      }
    },
    articleColor: function() {
      self = this;
      if (this.componentID === 2) {
        return "green-text lighten-2"
      }
    }
  }
}
</script>

<style lang="scss" scoped>
#users-checks {
  .choose-which {
    margin-top: 30px;
    display: flex;
    .checked-posts, .checked-articles {
      cursor: pointer;
      width: 300px;
      height: 50px;
      text-align: center;
      margin-bottom: 0;
      padding-top: 10px;
    }
    .checked-posts:hover, .checked-articles:hover {
       background-color: #eee;
    }
  }
  hr {
    margin-top: 0;
  }
}
</style>