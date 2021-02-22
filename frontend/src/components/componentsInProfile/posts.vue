<template>
  <div id="users-posts">
    <div class="images">
      <div class="each-image" v-for="(post, id) in posts" :key="id">
        <a class="modal-trigger" :href="'#modal' + id">
          <img :src="imagePath + post.images[0].file_name" />
        </a>
        <!-- modal -->
        <div :id="'modal' + id" class="modal card">
          <div class="card-image">
            <img :src="imagePath + post.images[0].file_name" />
          </div>
          <div class="card-content">
            <p>{{post.post.content}}</p>
            <small>
              <i class="material-icons tiny">local_offer</i>
              {{post.post.tags}}
            </small>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      posts: [],
      imagePath: "/dist/img/"
    };
  },
  watch: {
    $route(to, from) {
      this.fetchItems(to.params.id);
      this.initData();
    }
  },
  mounted() {
    // get user's posts
    axios.get(`/api/posts/${this.$route.params.id}`).then(res => {
      for (let i = 0; i < res.data.length; i++) {
        this.posts.push(res.data[i]);
      }
    });
  },
  updated() {
    // jquery for modal
    $(document).ready(function() {
      $(".modal").modal();
    });
  },
  methods: {
    fetchItems(id) {
      // get user's posts
      axios.get(`/api/posts/${id}`).then(res => {
        for (let i = 0; i < res.data.length; i++) {
          this.posts.push(res.data[i]);
        }
      });
    },
    initData() {
      this.posts = [];
    }
  }
};
</script>

<style lang="scss" scoped>
#users-posts {
  .images {
    display: grid;
    grid-template-columns: 1fr 1fr 1fr 1fr 1fr;
    img {
      width: 100%;
      height: 120px;
      object-fit: cover;
      object-position: 50% 50%;
    }
    img:hover {
      opacity: 0.7;
    }
    .modal {
      height: 800px;
      max-width: 600px;
      .card-image {
        height: 400px;
        img {
          height: 100%;
          object-fit: cover;
          object-position: 50% 50%;
        }
      }
    }
  }
}
@media (max-width: 480px) {
  #users-posts {
    .images {
      grid-template-columns: 1fr 1fr 1fr 1fr;
      img {
        height: 91.88px;
      }
    }
  }
}
</style>