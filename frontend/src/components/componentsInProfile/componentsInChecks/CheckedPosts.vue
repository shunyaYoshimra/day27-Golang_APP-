<template>
  <div id="checked-posts">
    <div class="images">
      <div class="each-image" v-for="(post, id) in posts" :key="id">
        <a class="modal-trigger" :href="'#modal' + id">
          <img :src="imagePath + post.images[0].file_name">
        </a>
        <!-- modal -->
        <div :id="'modal' + id" class="modal card">
          <div class="card-image">
            <img :src="imagePath + post.images[0].file_name">
          </div>
          <div class="card-content">
            <p>{{post.post.content}}</p>
            <small><i class="material-icons tiny">local_offer</i>{{post.post.tags}}</small>
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
      checkedNumber: [],
      posts: [],
      imagePath: "/src/img/",
    }
  },
  mounted() {
    axios.get(`/api/favorites/${this.$route.params.id}`).then((res) => {
      for (let i = 0; i < res.data.length; i ++) {
        this.checkedNumber.push(res.data[i].post_id);
      }
      for (let i = 0; i < this.checkedNumber.length; i ++) {
        axios.get(`/api/post/${this.checkedNumber[i]}`).then((res) => {
          this.posts.push(res.data.data);
        }).catch((err) => {
          if (err.response.status === 404) {
            console.log("deleted post")
          }
        })
      }
    })
  },
  updated() {
    // jquery for modal
    $(document).ready(function(){
      $('.modal').modal();
    });
  },
}
</script>

<style lang="scss" scoped>
#checked-posts {
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
  #checked-posts {
    .images {
      grid-template-columns: 1fr 1fr 1fr 1fr;
      img {
        height: 91.88px;
      }
    }
  }
}
</style>