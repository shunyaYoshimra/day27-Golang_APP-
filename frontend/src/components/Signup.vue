<template>
  <div id="signup">
    <div class="error">{{err}}</div>
    <div class="signup-form card-panel">
      <p>SIGN UP</p>
      <div class="inputs">
        <input type="text" v-model="email" placeholder="メール(Email)" />
        <input type="text" v-model="name" placeholder="名前(Name)" />
        <input type="password" v-model="password" placeholder="パスワード(Password)" />
      </div>
      <button @click="signup()" class="btn blue lighten-3">CLICK</button>
    </div>
    <div class="to-login card-panel">
      <p>
        <router-link to="/login">登録済みの方はこちらへ(To Login)</router-link>
      </p>
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
      email: "",
      name: "",
      password: ""
    };
  },
  destroyed() {
    // reload
    this.$router.go({ path: this.$router.currentRoute.path, force: true });
  },
  methods: {
    signup() {
      const params = new URLSearchParams();
      params.append("email", this.email);
      params.append("name", this.name);
      params.append("password", this.password);
      axios
        .post("/api/signup", params)
        .then(res => {
          console.log(res.data);
          this.$router.push("/profiles/new");
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
#signup {
  .error {
    text-align: center;
    color: tomato;
  }

  .signup-form {
    width: 350px;
    height: 400px;
    margin: 0 auto;

    p {
      margin-top: 40px;
      color: #333;
      font-weight: bold;
      text-align: center;
    }

    .inputs {
      margin-top: 40px;

      input {
        border: 1px solid #666;
        border-radius: 10px;
        background-color: #fafafa;
        font-size: 0.875rem;
        color: #333;
        text-align: center;
      }

      input:focus {
        outline: none;
      }
    }

    .btn {
      margin-top: 20px;
      width: 300px;
    }
  }

  .to-login {
    width: 350px;
    margin: 20px auto 0;
    height: 70px;

    p {
      transform: translateY(-15px);
      font-weight: bold;
    }
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