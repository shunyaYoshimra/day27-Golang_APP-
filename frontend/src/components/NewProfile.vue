<template>
  <div id="new-profile">
    <div class="error">{{err}}</div>
    <div class="profile-form card-panel">
      <p>PROFILE</p>
      <label for="profile">プロフィール(Profile)</label>
      <textarea v-model="description" id="profile"></textarea>
      <div class="subject">
        <select v-model="subject">
          <option value disabled selected>学科(Subject)</option>
          <option value="英米(Engilish)">英米(Engilish)</option>
          <option value="イスパニア(Spanish)">イスパニア(Spanish)</option>
          <option value="中国(Chinese)">中国(Chinese)</option>
          <option value="ロシア(Russian)">ロシア(Russian)</option>
          <option value="国際関係(IR)">国際関係(IR)</option>
          <option value="2部英米(Engilish')">2部英米(Engilish')</option>
          <option value="院生(Graduate)">院生(Graduate Students)</option>
          <option value="留学生(Exchange Students)">留学生(Exchange Students)</option>
          <option value="その他(Other)">その他(Other)</option>
        </select>
      </div>
      <label>*必須(Essential)</label>
      <button @click="createProfile()" class="btn blue lighten-3">Next</button>
    </div>
    <div class="copy-right">
      <small>&copy; Shunya Yoshimura@2020</small>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data() {
    return {
      err: "",
      description: "",
      subject: "",
    };
  },
  mounted() {
    // jquery for select
    $(document).ready(function() {
      $("select").formSelect();
    });
  
  },
  methods: {
    createProfile() {
      const params = new URLSearchParams();
      params.append("description", this.description);
      params.append("subject", this.subject);
      axios
        .post("/api/profiles", params)
        .then(res => {
          this.$router.push("/contacts/new");
        })
        .catch(err => {
          console.log(err.response);
          this.err = err.response.data.message;
        });
    }
  }
};
</script>

<style lang="scss" scoped>
#new-profile {
  .error {
    text-align: center;
    color: tomato;
  }

  .profile-form {
    width: 350px;
    height: 400px;
    margin: 0 auto;

    p {
      margin-top: 40px;
      color: #333;
      font-weight: bold;
      text-align: center;
    }

    textarea {
      border: 1px solid #666;
      border-radius: 10px;
      background-color: #fafafa;
      color: #333;
      height: 80px;
    }

    textarea:focus {
      outline: none;
    }
  }

  .btn {
    margin-top: 20px;
    width: 300px;
  }

  .copy-right {
    margin-top: 20px;
    text-align: center;

    small {
      color: #666;
    }
  }
}
</style>