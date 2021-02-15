<template>
  <div id="posts">
    <div class="search-form">
      <div>
        <input type="text" placeholder="検索(search)" v-model="keyWord">
        <button @click="searchPosts()" class="btn"><i class="material-icons">search</i></button>
      </div>
    </div>
    <div v-for="(post, id) in posts" :key="id">
      <span class="hide">{{user = userOfPost(post.post.user_id)}}</span>
      <span class="hide">{{country = countryOfPost(post.post.user_id)}}</span>
      <span class="hide">{{college = collegeOfPost(post.post.user_id)}}</span>
      <span class="hide">{{description = descriptionOfPost(post.post.user_id)}}</span>
      <div class="card" v-if="!deletedID.includes(post.post.id)">
        <template v-if="post.images.length >= 1">
          <div class="card-image waves-effect waves-block waves-light">
            <template v-if="post.images.length >= 2">
              <span class="more-images" @click="modalNum = id">+{{post.images.length - 1}}</span>
            </template>
            <img class="activator" :src="imagePath + post.images[0].file_name">
          </div>
        </template>
        <div class="card-content">
          <div class="icons">
            <i class="material-icons grey-text lighten-2 activator">person</i>
            <template v-if="post.post.user_id === me.id">
              <i @click="deletePost(post.post.id)" class="material-icons grey-text lighten-2">delete</i>
            </template>
            <template v-if="myFavorites.includes(post.post.id)">
              <span @click="DeleteFavorite(post.post.id)" class="green-text lighten-2">
                <i class="material-icons">check</i>checked!
              </span>
            </template>
            <template v-else>
              <i @click="favoritePost(post.post.id)" class="material-icons grey-text lighten-2">check</i>
            </template>
          </div>
          <p>{{post.post.content}}</p>
          <small>
            <i class="material-icons tiny">local_offer</i>
            {{post.post.tags}}
          </small>
        </div>
        <div class="card-reveal">
          <span class="card-title grey-text text-darken-4"><router-link class="grey-text text-darken-4 bold" :to="'/profile/' + user.id">{{user.name}}</router-link><i class="material-icons right">close</i></span>
          <div class="abroads-wrapper">
            <template v-if="country == null && college == null && description == null">
              <p>{{user.name}}さんの留学情報はまだ登録されていません</p>
            </template>
            <template v-else>
              <p class="country"><i class="material-icons tiny">place</i><span>{{country}}</span></p>
              <p><i class="material-icons tiny">school</i><span>{{college}}</span></p>
              <p class="description">{{description}}</p>
            </template>
          </div>
        </div>
        <!-- images modal -->
        <tempalte v-if="id === modalNum">
          <div class="images-modal">
            <i @click="closeModal()" class="material-icons close-modal">close</i>
            <template v-if="imageNum !== 0">
              <i @click="imageNum --" class="material-icons left">keyboard_arrow_left</i>
            </template>
            <template v-if="imageNum !== post.images.length - 1">
              <i @click="imageNum ++" class="material-icons right">keyboard_arrow_right</i>
            </template>
            <template v-if="post.images.length >= imageNum + 1">
              <img :src="imagePath + post.images[imageNum].file_name">
            </template>
          </div>
        </tempalte>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      me: {},
      posts: [],
      myFavorites: [],
      abroads: [],
      users: [],
      imagePath: "/dist/img/",
      imageNum: 0,
      modalNum: -1,
      deletedID: [],
      keyWord: "",
    }
  },
  computed: {
    userOfPost: function() {
      self = this;
      return function(id) {
        for (let i = 0; i < self.users.length; i ++) {
          if (self.users[i].id === id) {
            return self.users[i];
          }
        }
      }
    },
    countryOfPost: function() {
      self = this;
      return function(id) {
        for (let i = 0; i< self.abroads.length; i ++) {
          if (self.abroads[i].user_id === id) {
            return self.abroads[i].country;
          }
        }
      }
    },
    collegeOfPost: function() {
      self = this;
      return function(id) {
        for (let i = 0; i< self.abroads.length; i ++) {
          if (self.abroads[i].user_id === id) {
            return self.abroads[i].college;
          }
        }
      }
    },
    descriptionOfPost: function() {
      self = this;
      return function(id) {
        for (let i = 0; i< self.abroads.length; i ++) {
          if (self.abroads[i].user_id === id) {
            return self.abroads[i].description;
          }
        }
      }
    },
  },
  mounted() {
    // get me
    axios.get("/api/get_me").then((res) => {
      this.me = res.data.data;
    })
    // get posts and images
    axios.get("/api/posts").then((res) => {
      for (let i = 0; i < res.data.length; i ++) {
        this.posts.push(res.data[i]);
      }
    })
    // get users
    axios.get("/api/users").then((res) => {
      for (let i = 0; i < res.data.length; i ++) {
        this.users.push(res.data[i]);
      }
    })
    // get abroads
    axios.get("/api/abroads").then((res) => {
      for (let i = 0;  i < res.data.length; i ++) {
        this.abroads.push(res.data[i]);
      }
    })
    // get favorites
    axios.get("/api/favorites").then((res) => {
      for (let i = 0; i < res.data.length; i ++) {
        this.myFavorites.push(res.data[i].post_id)
      }
    })
  },
  updated() {
    $(document).ready(function(){
      $('.modal').modal();
    });
  },
  methods: {
    closeModal() {
      this.modalNum = -1;
      this.imageNum = 0;
    },
    deletePost(id) {
      if(confirm("本当にこの投稿を削除しますか?(Are you sure to delete this post?)")) {
        axios.delete(`/api/posts/${id}`)
        this.deletedID.push(id);
      }
    },
    favoritePost(id) {
      axios.post(`/api/favorites/${id}`).then((res) => {
        this.myFavorites.push(res.data.data.post_id)
      });
    },
    DeleteFavorite(id) {
      axios.delete(`/api/favorites/${id}`).then((res) => {
        this.myFavorites = this.myFavorites.filter(n => n !== id);
      })
    },
    searchPosts() {
      if (this.keyWord !== "") {
        this.posts = [];
          axios.get(`/api/searched_posts/${this.keyWord}`).then((res) => {
            console.log(res.data);
            for (let i = 0; i < res.data.length; i ++) {
              this.posts.push(res.data[i]);
            }
          })
        }
      }
  },
}
</script>

<style lang="scss" scoped>
#posts {
  i {
    cursor: pointer;
  }
  width: 600px;
  margin: auto;

  .search-form {
    div {
      width:200px;
      display: flex;
      align-items: center;
      input {
        background-color: #fafafa;
        margin-top: 10px;
        border: 1px solid #333;
        border-radius: 2px;
        font-size: 1rem;
        width: 100%;
        height:25px;
        box-sizing: border-box;
        padding: 0 1rem;
      }

      input:focus {
        outline: none;
      }
      .btn {
        height: 25px;
        background-color: #333;
        
        i {
          transform: translateY(-4px);
        }
      }
    }
    
  }
  .card {
    .card-image {
      position: relative;
      img {
        height: 500px;
        object-fit: cover;
        object-position: 50% 50%;
      }
      .more-images {
        position: absolute;
        bottom: 20px;
        right: 20px;
        width: 60px;
        height: 60px;
        background-color: #000;
        opacity: 0.7;
        color: #fff;
        font-size: 2rem;
        border-radius: 50%;
        text-align: center;
        padding-top: 10px;
      }
      .more-images:hover {
        opacity: 0.4;
      }
    }

    .card-content {
      p {
        margin: 10px 0;
      }
    }

    .card-reveal {
      .abroads-wrapper{
        .country {
          padding-top: 30px;
        }
        i {
          margin-right: 1rem;
        }
        .description {
          padding-top: 20px;
        }
      }
    }

    .images-modal {
      position: fixed;
      background-color: rgba(0, 0, 0, 0.3);
      width: 100%;
      height: 100%;
      top: 0;
      left: 0;
      z-index: 200;
      text-align: center;
      .close-modal {
        position: absolute;
        margin-left: 400px;
        top: 40px;
        width: 40px;
        height: 40px;
        background-color: #fff;
        opacity: 0.7;
        border-radius: 50%;
        text-align: center;
        padding-top: 7px;
      }
      .close-modal:hover{
        opacity:1;
      }
      .left, .right {
        position: absolute;
        top: 350px;
        width: 30px;
        height: 30px;
        background-color: #fff;
        opacity: 0.7;
        border-radius: 50%;
        text-align: center;
        padding-top: 3px;
      }
      .left {
        margin-left: 10px;
      }
      .right {
        margin-left: 360px;
      }
      img {
        margin-top: 150px;
        width: 400px;
        height: 400px;
        object-fit: cover;
        object-position: 50% 50%;
      }
    }
  }
}
@media (max-width: 480px) {
  #posts {
    width: 98%;
    margin: auto;
    margin-bottom: 60px;
    .card {
      .card-image {
        img {
          height: 300px;
        }
      }
      .images-modal {
        .close-modal {
          margin-left: 270px;
        }
        .left, .right {
          top: 300px;
        }
        .left {
          margin-left: 10px;
        }
        .right {
          margin-left: 265px;
        }
        img {
          width: 300px;
          height: 300px;
        }
      }
    }
  }
}
</style>