<template>
  <div id="profile">
    <div class="user-info">
      <div class="top">
        <p class="title">{{user.name}}</p>
        <template v-if="me.id === user.id">
          <button @click="linkToEdit">プロフィールの編集(Edit Profile)</button>
        </template>
      </div>
      <span>{{profile.subject}}</span>
      <div class="contact">
        <span>{{user.email}}</span>
        <template v-if="contactBool">
          <template v-if="contact.media == 'twitter' ">
            <a :href="'https://www.twitter.com/' + contact.account">
              <i class="material-icons tiny">attach_file</i>
              Twitterアカウント
            </a>
          </template>
          <template v-if="contact.media == 'instagram' ">
            <a :href="'https://www.instagram.com/' + contact.account ">
              <i class="material-icons tiny">attach_file</i>
              Instagramアカウント
            </a>
          </template>
        </template>
        <template v-else>
          <router-link to="/contacts/new">Contact</router-link>
        </template>
      </div>
      <div class="description">
        <p>{{profile.description}}</p>
      </div>
    </div>
    <template v-if="abroadBool">
      <div class="user-abroad">
        <p class="title">留学(Study Abroad)</p>
        <span><i class="material-icons tiny">place</i></span>{{abroad.country}}<br>
        <span><i class="material-icons tiny">school</i></span>{{abroad.college}}
        <p class="description">{{abroad.description}}</p>
      </div>
    </template>
    <template v-if="occupationBool">
      <div class="user-occupation">
        <p class="title">就職(Job Seeking)</p>
        <span><i class="material-icons tiny">label</i></span>{{occupation.kind}}<br>
        <span><i class="material-icons tiny">work</i></span>{{occupation.company}}
        <p class="description">{{occupation.description}}</p>
      </div>
    </template>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      me: {},
      user: {},
      profile: {},
      profileBool: true,
      contact: {},
      contactBool: true,
      editBool: false,
      abroad: {},
      abroadBool: true,
      occupation: {},
      occupationBool: true,
      componentID: 1,
    };
  },
  watch: {
    $route(to, from) {
      this.fetchItems(to.params.id);
    }
  },
  // computed: {
  //   whichComponent: function() {
  //     self = this;
  //     if (this.componentID === 1) {
  //       return "user-posts"
  //     } else if (this.componentID === 2) {
  //       return "user-articles"
  //     } else if (this.componentID === 3) {
  //       return "user-questions"
  //     } else if (this.componentID === 4) {
  //       return "user-checks"
  //     }
  //   }
  // },
  mounted() {
    // get me
    axios.get("/api/get_me").then(res => {
      this.me = res.data.data;
    });
    // get user
    axios.get(`/api/user/${this.$route.params.id}`).then(res => {
      this.user = res.data.data;
    });
    // get profile
    axios.get(`/api/profile/${this.$route.params.id}`).then(res => {
      this.profile = res.data.data;
    }).catch((err) => {
      console.log(err)
      if (err.response.status === 404) {
        this.profileBool = false
      }
    });
    // get contact
    axios.get(`/api/contact/${this.$route.params.id}`).then(res => {
      this.contact = res.data.data;
    }).catch((err) => {
      if (err.response.status === 404) {
        this.contactBool = false
      }
    });
    // get abroad
    axios.get(`/api/abroad/${this.$route.params.id}`).then((res) => {
      this.abroad = res.data.data;
    }).catch((err) => {
      if (err.response.status === 404) {
        this.abroadBool = false
      }
    })
    // get occupation
    axios.get(`/api/occupation/${this.$route.params.id}`).then((res) => {
      this.occupation = res.data.data;
    }).catch((err) => {
      if (err.response.status === 404) {
        this.occupationBool = false
      }
    })
  },
  methods: {
    fetchItems(id) {
        // get me
      axios.get("/api/get_me").then(res => {
        this.me = res.data.data;
      });
      // get user
      axios.get(`/api/user/${id}`).then(res => {
        this.user = res.data.data;
      });
      // get profile
      axios.get(`/api/profile/${id}`).then(res => {
        this.profile = res.data.data;
      }).catch((err) => {
        if (err.response.status === 404) {
          this.profileBool = false
        }
      });
      // get contact
      axios.get(`/api/contact/${id}`).then(res => {
        this.contact = res.data;
      }).catch((err) => {
        if (err.response.status === 404) {
          this.contactBool = false
        }
      });
      // get abroad
      axios.get(`/api/abroad/${id}`).then((res) => {
        this.abroad = res.data;
      }).catch((err) => {
        if (err.response.status === 404) {
          this.abroadBool = false
        }
      })
      // get occupation
      axios.get(`/api/occupation/${id}`).then((res) => {
        this.occupation = res.data;
      }).catch((err) => {
        if (err.response.status === 404) {
          this.occupationBool = false
        }
      })
    },
    linkToEdit() {
      this.$router.push("/profiles/edit");
    }
  }
};
</script>

<style lang="scss" scoped>
#profile {
  width: 600px;
  margin: auto;
  margin-bottom: 200px;
  // common
  .title {
      font-size: 1.25rem;
      font-weight: bold;
    }

  // user-info
  .user-info {
    .top {
      align-items: center;
      display: flex;
      justify-content: space-between;
      button {
        background-color: #fafafa;
        border: 1px #999 solid;
        border-radius: 5px;
        padding: 5px;
        cursor: pointer;
      }
      button:hover {
        background-color: #fff;
      }
    }
    .contact {
      span {
        margin-right: 2rem;
      }
    }
    .description {
      margin-top: 15px;
    }
  }

  // user-abroad
  .user-abroad {
    margin-top: 60px;
    span {
      margin-right: 1rem;
    }
    .description {
      margin-top: 20px;
    }
  }

  // user-occupation
  .user-occupation {
    margin-top: 60px;
    span {
      margin-right: 1rem;
    }
    .description {
      margin-top: 20px;
    }
  }

  .select-tags {
    margin-top: 60px;
    color: #666;
    .above-line {
      margin-bottom: 0;
    }
    .bottom-line {
      margin-top: 0;
    }
    .tags-wrapper {
      height: 50px;
      align-items: center;
      display: flex;
      .each-tag {
        cursor: pointer;
        height: 50px;
        width: 150px;
        text-align: center;
        padding-top: 10px;
      }
      .each-tag:hover {
        background-color: #eee;
      }
    }
  }
}
@media (max-width: 480px) {
  #profile {
    width: 100%;
  }
}
</style>