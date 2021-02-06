<template>
  <div id="start">
    <div class="confirm card-panel">
      <p class="welcome">Hi, {{user.name}}! Welecome to Gaidai Chat Online!!</p>
      <small>*アカウントの情報が正しいか確認してください(Please confirm your accuount registared correctly)</small>
      <div class="info">
        <p>
          名前(Name):
          <span>{{user.name}}</span>
        </p>
        <p>
          メールアドレス(E-mail):
          <span>{{user.email}}</span>
        </p>
        <p>
          学科(Subject):
          <span>{{profile.subject}}</span>
        </p>
        <template v-if="contact.media == 'twitter' ">
          <a :href="'https://www.twitter.com/' + contact.account">
            <i class="material-icons">attach_file</i>
            Twitterアカウント
          </a>
        </template>
        <template v-if="contact.media == 'instagram' ">
          <a :href="'https://www.instagram.com/' + contact.account ">
            <i class="material-icons">attach_file</i>
            Instagramアカウント
          </a>
        </template>
      </div>
    </div>
    <div class="start-btn">
      <button @click="start()" class="btn blue lighten-3">Get Started</button>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      user: {},
      profile: {},
      contact: {}
    };
  },
  mounted() {
    // get current user
    axios.get("/api/get_me").then(res => {
      console.log(res.data);
      this.user = res.data.data;
    });
    //  get current user's profile
    axios.get("/api/my_profile").then(res => {
      console.log(res.data);
      this.profile = res.data.data;
    });
    // get current user's contact
    axios.get("/api/my_contact").then(res => {
      console.log(res.data);
      this.contact = res.data.data;
    });
  },
  methods: {
    start() {
      this.$router.push("/");
    }
  },
};
</script>

<style lang="scss" scoped>
#start {
  .confirm {
    width: 600px;
    margin: 0 auto;

    .welcome {
      font-size: 1.25rem;
      font-weight: bold;
    }

    .info {
      margin-top: 20px;

      span {
        margin-left: 1rem;
        font-weight: bold;
      }
    }
  }

  .start-btn {
    width: 600px;
    margin: auto;
    margin-top: 20px;
    .btn {
      width: 600px;
    }
  }
}
</style>